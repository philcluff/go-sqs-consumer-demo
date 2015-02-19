package main

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/sqs"
	"time"
	"os"
)

type MessageProcessor struct {
	client sqs.SQS
}

func main() {
	creds, _ := aws.EnvCreds()
	client := sqs.New(creds, os.Getenv("AWS_REGION"), nil)
	mp := &MessageProcessor{*client}
	mp.pollQueue() // Blocks
}

func (mp *MessageProcessor) pollQueue() {
	for {

		fmt.Println("Long polling for a message... (Will wait for 10s)")

		// Fetch some messages, hide it for 10 seconds.
		rmr := &sqs.ReceiveMessageRequest{
			MaxNumberOfMessages: aws.Integer(10),
			QueueURL:            aws.String(os.Getenv("SQS_QUEUE")),
			VisibilityTimeout:   aws.Integer(10),
			WaitTimeSeconds: aws.Integer(10),
		}
		rmre, _ := mp.client.ReceiveMessage(rmr)

		// Sleep a little if we didn't find any messages in this poll.
		if len(rmre.Messages) < 1 {
			fmt.Println("No messages on queue. Will sleep for 1s, then long poll again.")
			time.Sleep(time.Second)
		}

		// Iterate over the messages we received, and dispatch a processor for each.
		for _, message := range rmre.Messages {
			go mp.processMessage(message)
		}
	}
}

func (mp *MessageProcessor) processMessage(message sqs.Message) {
	fmt.Println(*message.Body)
	dmr := &sqs.DeleteMessageRequest{
		QueueURL:      aws.String(os.Getenv("SQS_QUEUE")),
		ReceiptHandle: message.ReceiptHandle,
	}
	err := mp.client.DeleteMessage(dmr)
	fmt.Println(err)
}
