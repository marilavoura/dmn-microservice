package decision

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/marianalavoura/dmn-microservice/decision/input"
	"github.com/marianalavoura/dmn-microservice/model"
	"github.com/marianalavoura/dmn-microservice/parser/types"
	"github.com/marianalavoura/dmn-microservice/utils"
)

func GetRuleOutput(receivedInput model.DecisionInput, rules []types.Rules, inputDefinition []types.InputDefinition, numberOfInputs int, numberOfOutputs int, numberOfRules int) *model.DecisionOutput {
	var invalidOutput = model.DecisionOutput{
		Id:        receivedInput.Id,
		Taxes:     -1,
		Discounts: -1,
	}

	var output model.DecisionOutput

	positionInInput := *input.GetPositionInInput(receivedInput, inputDefinition, numberOfInputs)
	invoiceObject := reflect.ValueOf(receivedInput)

	ruleMatches := true

	for i := 0; i < numberOfRules; i++ {
		ruleMatches = true
		for j := 0; j < numberOfInputs; j++ {
			invoiceField := fmt.Sprint(invoiceObject.FieldByIndex([]int{positionInInput[j]}).Interface())
			allowedType := inputDefinition[j].AllowedType

			switch allowedType {
			case "string":
				ruleFieldFormatted := strings.ReplaceAll(rules[i].Inputs[j], "\"", "")
				ruleFieldArray := strings.Split(ruleFieldFormatted, ", ")

				if !utils.Contains(ruleFieldArray, invoiceField) {
					ruleMatches = false
				}
			case "number":
				invoiceNumberField, _ := strconv.ParseFloat(invoiceField, 64)

				var firstCharacter string

				if len(rules[i].Inputs[j]) > 1 {
					firstCharacter = string(rules[i].Inputs[j][0])
				}

				switch firstCharacter {
				case "[":
					r := regexp.MustCompile(`\[(\d*\.\d*)\.\.(\d*\.\d*)\]`)

					lowerLimit, _ := strconv.ParseFloat(r.FindStringSubmatch(rules[i].Inputs[j])[1], 64)
					upperLimit, _ := strconv.ParseFloat(r.FindStringSubmatch(rules[i].Inputs[j])[2], 64)

					if invoiceNumberField < lowerLimit || invoiceNumberField > upperLimit {
						ruleMatches = false
					}
				case ">":
					secondCharacter := rules[i].Inputs[j][1]

					r := regexp.MustCompile(`>={0,1} (\d*\.\d*)`)
					lowerLimit, _ := strconv.ParseFloat(r.FindStringSubmatch(rules[i].Inputs[j])[1], 64)

					if invoiceNumberField < lowerLimit || (invoiceNumberField == lowerLimit && secondCharacter != '=') {
						ruleMatches = false
					}
				case "<":
					secondCharacter := rules[i].Inputs[j][1]

					r := regexp.MustCompile(`<={0,1} (\d*\.\d*)`)
					upperLimit, _ := strconv.ParseFloat(r.FindStringSubmatch(rules[i].Inputs[j])[1], 64)

					if invoiceNumberField > upperLimit || (invoiceNumberField == upperLimit && secondCharacter != '=') {
						ruleMatches = false
					}
				default:
					//fmt.Print("Regra não encontrada")

					return &invalidOutput
				}
			default:
				//fmt.Println("Tipo permitido:")
				//fmt.Println(inputDefinition[j].AllowedType)

				return &invalidOutput
			}

		}

		if ruleMatches {
			output.Id = receivedInput.Id
			output.Taxes, _ = strconv.ParseFloat(rules[i].Outputs[0], 64)
			output.Discounts, _ = strconv.ParseFloat(rules[i].Outputs[1], 64)

			return &output
		}
	}

	return &output
}
