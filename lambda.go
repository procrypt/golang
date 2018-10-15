package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

type BotRequest struct {
	Messages []Message `json:"messages,omitempty"`
}

type BotResponce struct {
	Messages []Message `json:"messages,omitempty"`
}

type Message struct {
	Type         string        `json:"type,omitempty"`
	Unstructured *UnstructuredMessage `json:"unstructured,omitempty"`
}

type UnstructuredMessage struct {
	ID        string `json:"id,omitempty"`
	Text      string `json:"text,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

func handelRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	botRequest := BotRequest{}
	req := request.Body
	json.Unmarshal([]byte(req),&botRequest)

	responseSlice := []Message{}
	for _,v := range botRequest.Messages {
		if v.Unstructured.Text == "Hi" {
			body := Message{
				Type: request.HTTPMethod,
				Unstructured: &UnstructuredMessage{
					ID: request.RequestContext.RequestID,
					Text: "Hi! How are you doing today?",
					Timestamp: time.Now().String(),
				},
			}
			responseSlice = append(responseSlice, body)
			bs, _ := json.Marshal(BotResponce{Messages:responseSlice})
			return events.APIGatewayProxyResponse{Body: string(bs), StatusCode:200} , nil
		} else if v.Unstructured.Text == "Hey" {
			body := Message{
				Type: request.HTTPMethod,
				Unstructured: &UnstructuredMessage{
					ID: request.RequestContext.RequestID,
					Text: "Hey There!!",
					Timestamp: time.Now().String(),
				},
			}
			responseSlice = append(responseSlice, body)
			bs, _ := json.Marshal(BotResponce{Messages:responseSlice})
			return events.APIGatewayProxyResponse{Body: string(bs), StatusCode:200} , nil
		} else if v.Unstructured.Text == "Supp" {
			body := Message{
				Type: request.HTTPMethod,
				Unstructured: &UnstructuredMessage{
					ID: request.RequestContext.RequestID,
					Text: "What up, Homie!!",
					Timestamp: time.Now().String(),
				},
			}
			responseSlice = append(responseSlice, body)
			bs, _ := json.Marshal(BotResponce{Messages:responseSlice})
			return events.APIGatewayProxyResponse{Body: string(bs), StatusCode:200} , nil
		}
	}
	body := Message{
		Type: request.HTTPMethod,
		Unstructured: &UnstructuredMessage{
			ID: request.RequestContext.RequestID,
			Text: `Hi, I'm dumb. I can not read your mind. Although, my creator is tyring hard to make me intelligent. In the mean time, type something instead, like 'Hi', 'Hello', etc..!!`,
			Timestamp: time.Now().String(),
		},
	}
	responseSlice = append(responseSlice, body)
	bs, _ := json.Marshal(BotResponce{Messages:responseSlice})
	return events.APIGatewayProxyResponse{Body: string(bs), StatusCode:200} , nil
}

func main() {
	lambda.Start(handelRequest)
}
