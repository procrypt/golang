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

func elasticSearch(query string) []string {
	data := []string{}
	r := Result{}
	req, err := http.NewRequest("GET", ""+query, nil)
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
	if len(r.Hits.Hits) != 0 {
		for _, v := range r.Hits.Hits {
			data = append(data, v.Source.ObjectKey)
		}
	} else {
		return data
	}
	defer resp.Body.Close()
	return data
}

func lexSearch(req events.LexEvent) (events.LexEvent, error) {
	utterances := []string{"Tree", "Road", "Bird", "Human", "Town", "Downtown", "Asphalt", "Plant"}
	attachments := []events.Attachment{}
	if req.CurrentIntent.Name == "Training" {
		for _, v := range utterances {
			if req.InputTranscript == v {
				result := elasticSearch(req.InputTranscript)
				for _, v := range result {
					attachments = append(attachments, events.Attachment{ImageURL: "" + v})
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

				json.Marshal(dailogAction)
				return dailogAction, nil
			}
		}
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
	json.Marshal(dailogAction)
	return dailogAction, nil
}

func main() {
	lambda.Start(lexSearch)
}
