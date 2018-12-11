package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Result struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total    int     `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				ObjectKey        string   `json:"objectKey"`
				Bucket           string   `json:"bucket"`
				CreatedTimestamp string   `json:"createdTimestamp"`
				Label            []string `json:"label"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func elasticSearch(query []string) []string {
	data := []string{}
	r := Result{}
	for _, v := range query {
		req, err := http.NewRequest("GET", "<elasticSearchURL>"+v, nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		bs, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal([]byte(bs), &r)
		fmt.Println(string(bs))
		if len(r.Hits.Hits) != 0 {
			for _, v := range r.Hits.Hits {
				data = append(data, v.Source.ObjectKey)
			}
		}
		defer resp.Body.Close()
	}
	return data
}

func lexSearch(req events.LexEvent) (events.LexEvent, error) {
	result := []string{}
	attachments := []events.Attachment{}
	if req.CurrentIntent.Name == "ProcryptPics" {
		if req.InputTranscript == "show me birds" || req.InputTranscript == "birds" {
			result = elasticSearch([]string{"Bird"})
		} else if req.InputTranscript == "show me cats" || req.InputTranscript == "cats" {
			result = elasticSearch([]string{"Cat"})
		} else if req.InputTranscript == "show me cats and dogs" || req.InputTranscript == "show me dogs and cats" {
			result = elasticSearch([]string{"Cat", "Dog"})
		}
		for _, v := range result {
			attachments = append(attachments, events.Attachment{ImageURL: "<s3 bucket url>" + v})
		}
		dailogAction := events.LexEvent{
			DialogAction: &events.LexDialogAction{
				Type:             "Close",
				FulfillmentState: "Fulfilled",
				Message: map[string]string{
					"contentType": "PlainText",
					"content":     "Result",
				},
				ResponseCard: &events.LexResponseCard{
					Version:            1,
					ContentType:        "application/vnd.amazonaws.card.generic",
					GenericAttachments: attachments,
				},
			},
		}
		return dailogAction, nil
	}
	dailogAction := events.LexEvent{
		DialogAction: &events.LexDialogAction{
			Type:             "Close",
			FulfillmentState: "Fulfilled",
			Message: map[string]string{
				"contentType": "PlainText",
				"content":     "No search results",
			},
		},
	}
	return dailogAction, nil
}

func main() {
	lambda.Start(lexSearch)
}
