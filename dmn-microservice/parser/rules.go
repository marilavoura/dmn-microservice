package parser

import (
	"github.com/marianalavoura/dmn-microservice/parser/types"
)

func NewRules(definitions types.Definitions, numberOfInputs int, numberOfOutputs int, numberOfRules int) *[]types.Rules {
	var rules = make([]types.Rules, numberOfRules)

	for i := 0; i < numberOfRules; i++ {
		rules[i] = types.Rules{
			Inputs:  make([]string, numberOfInputs),
			Outputs: make([]string, numberOfOutputs),
		}
	}

	for i := 0; i < numberOfRules; i++ {
		//fmt.Println("Regra", i)
		for j := 0; j < numberOfInputs; j++ {
			rules[i].Inputs[j] = definitions.Decision.DecisionTable.Rule[i].InputEntry[j].Text
			//fmt.Print(rules[i].Inputs[j])
		}
		//fmt.Println()
		for k := 0; k < numberOfOutputs; k++ {
			rules[i].Outputs[k] = definitions.Decision.DecisionTable.Rule[i].OutputEntry[k].Text
			//fmt.Print(rules[i].Outputs[k])
		}

		//fmt.Println()
		//fmt.Println()
	}

	return &rules
}
