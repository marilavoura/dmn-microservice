package verify

import (
	"strings"

	"github.com/marianalavoura/dmn-microservice/utils"
)

func stringMatches(ruleStr string, invStr string) bool {
	ruleFieldFormatted := strings.ReplaceAll(ruleStr, "\"", "")
	ruleFieldArray := strings.Split(ruleFieldFormatted, ", ")

	if !utils.Contains(ruleFieldArray, invStr) {
		return false
	}

	return true
}
