package parser

import (
	"strings"

	"github.com/marianalavoura/dmn-microservice/parser/types"
)

func NewInputDefinitions(definitions types.Definitions, numberOfInputs int) *[]types.InputDefinition {
	var inputDefinition = make([]types.InputDefinition, numberOfInputs)

	//fmt.Println("Definição das entradas")
	for i := 0; i < numberOfInputs; i++ {
		inputDefinition[i] = types.InputDefinition{
			InputName:     definitions.Decision.DecisionTable.Input[i].Label,
			AllowedType:   definitions.Decision.DecisionTable.Input[i].InputExpression.TypeRef,
			AllowedValues: strings.Split(strings.ReplaceAll(definitions.Decision.DecisionTable.Input[i].InputValues.Text, "\"", ""), ","),
		}

		//fmt.Printf("%s - %s - %s\n", inputDefinition[i].InputName, inputDefinition[i].AllowedType, inputDefinition[i].AllowedValues)
	}

	return &inputDefinition
}
