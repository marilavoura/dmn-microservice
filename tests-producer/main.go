package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/marianalavoura/producer/model"
)

type DecisionOutput struct {
	Id        int     `json:"id"`
	Taxes     float64 `json:"taxes"`
	Discounts float64 `json:"discounts"`
}

func main() {
	tests := model.NewTests()

	// Producer and consumer definition
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	})
	if err != nil {
		panic(err)
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "tests-consumer",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	// Start consuming invoices
	c.SubscribeTopics([]string{"decision-output"}, nil)

	// Produce invoice messages to topic (asynchronously)

	topic := "invoices"
	for _, invoice := range tests {
		deliveryChan := make(chan kafka.Event)

		formattedInvoice, err := json.Marshal(invoice)

		fmt.Println(invoice)

		if err != nil {
			panic(err)
		}

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(formattedInvoice),
		}, deliveryChan)

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}

		close(deliveryChan)
	}

	run := true

	for run {
		passed := true

		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))

			var decisionOutput DecisionOutput

			err = json.Unmarshal(msg.Value, &decisionOutput)

			if err != nil {
				fmt.Print(err)
			}

			if decisionOutput.Discounts != tests[decisionOutput.Id].Discounts || decisionOutput.Taxes != tests[decisionOutput.Id].Taxes {
				passed = false

				fmt.Println("Test failed")
				fmt.Println(decisionOutput)
				fmt.Println("Expected output: ")
				fmt.Println(tests[decisionOutput.Id].Discounts, " - ", tests[decisionOutput.Id].Taxes)
			}

			if decisionOutput.Id == len(tests)-1 {
				if passed {
					fmt.Println("All tests passed.")
				}

				run = false

				defer p.Close()
				defer c.Close()
			}

		} else {
			fmt.Printf("Error while consuming message: %v (%v)\n", err, msg)
		}
	}
}
