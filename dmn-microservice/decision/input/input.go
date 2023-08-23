package input

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/marianalavoura/dmn-microservice/model"
	"github.com/marianalavoura/dmn-microservice/parser/types"
)

func NewInput(jsonObj []byte) (*model.DecisionInput, error) {
	var input model.DecisionInput

	err := json.Unmarshal(jsonObj, &input)

	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	return &input, nil
}

func GetPositionInInput(input model.DecisionInput, inputDefinition []types.InputDefinition, numberOfInputs int) *[]int {
	// Defines the position of each XML input in the object received
	val := reflect.Indirect(reflect.ValueOf(input))

	var positionInInput = make([]int, numberOfInputs)

	for i := 0; i < numberOfInputs; i++ {
		exists := false
		for j := 0; j < val.NumField(); j++ {
			if strings.EqualFold(inputDefinition[i].InputName, val.Type().Field(j).Name) {
				exists = true
				positionInInput[i] = j
				break
			}
		}

		if !exists {
			panic("Input not found in object received from DB")
		}
	}

	return &positionInInput
}
