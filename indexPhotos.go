package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type Payload struct {
	Object    string   `json:"objectKey"`
	Bucket    string   `json:"bucket"`
	TimeStamp string   `json:"createdTimestamp"`
	Label     []string `json:"label"`
}

func elasticSearchIndex(p *Payload) {
	payloadBytes, err := json.Marshal(p)
	fmt.Println(string(payloadBytes))
	if err != nil {
		fmt.Println(err)
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", "<ElasticSearchVPC URL>", body)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	bs, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bs))
	defer resp.Body.Close()
}

func indexPhotos(s3Event events.S3Event) {
	p := Payload{}
	for _, record := range s3Event.Records {
		s3 := record.S3
		svc := rekognition.New(session.New())
		input := &rekognition.DetectLabelsInput{
			Image: &rekognition.Image{
				S3Object: &rekognition.S3Object{
					Bucket: aws.String(s3.Bucket.Name),
					Name:   aws.String(s3.Object.Key),
				},
			},
			MaxLabels:     aws.Int64(10),
			MinConfidence: aws.Float64(70.00),
		}
		result, err := svc.DetectLabels(input)
		if err != nil {
			fmt.Println(err)
		}

		for _, v := range result.Labels {
			p.Label = append(p.Label, *v.Name)
		}
		p.Object = *input.Image.S3Object.Name
		p.Bucket = *input.Image.S3Object.Bucket
		p.TimeStamp = time.Now().String()
		elasticSearchIndex(&p)
	}

}

func main() {
	lambda.Start(indexPhotos)
}
