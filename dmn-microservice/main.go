package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/marianalavoura/dmn-microservice/decision"
	"github.com/marianalavoura/dmn-microservice/decision/input"
	"github.com/marianalavoura/dmn-microservice/parser"
)

func main() {
	path := "./model/Decide Discounts and Taxes.xml"

	definitions := *parser.NewDefinitions(path)

	// Inputs
	numberOfInputs := len(definitions.Decision.DecisionTable.Input)
	//fmt.Println("Número de entradas: ", numberOfInputs)
	//fmt.Println()

	// Outputs
	numberOfOutputs := len(definitions.Decision.DecisionTable.Output)
	//fmt.Println("Número de saídas: ", numberOfOutputs)
	//fmt.Println()

	// Rules
	numberOfRules := len(definitions.Decision.DecisionTable.Rule)
	//fmt.Println("Número de regras: ", numberOfRules)
	//fmt.Println()

	inputDefinition := *parser.NewInputDefinitions(definitions, numberOfInputs)

	rules := *parser.NewRules(definitions, numberOfInputs, numberOfOutputs, numberOfRules)

	// Producer and consumer definition
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "decide-discouts-taxes",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Falha na entrega: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Mensagem enviada para %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	c.SubscribeTopics([]string{"invoices"}, nil)

	run := true

	for run {
		msg, err := c.ReadMessage(-1)

		if err == nil {
			fmt.Printf("Message on %s: %s\n\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v\n", err)
		}

		input, _ := input.NewInput(msg.Value)

		fmt.Println("Inputs: ", input)

		output := decision.GetRuleOutput(*input, rules, inputDefinition, numberOfInputs, numberOfOutputs, numberOfRules)

		// Produce messages to topic (asynchronously)
		topic := "decision-output"
		formattedOutput, err := json.Marshal(output)

		fmt.Print()

		if err != nil {
			panic(err)
		}

		fmt.Println(output)

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(formattedOutput),
		}, nil)
	}

	c.Close()
}
