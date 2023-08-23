package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/marianalavoura/dmn-microservice/parser/types"
)

func NewDefinitions(path string) *types.Definitions {
	xmlFile, err := os.Open(path)

	if err != nil {
		panic(err)
	} 

	fmt.Println("Successfully Opened " + path)
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)
	var definitions types.Definitions
	xml.Unmarshal(byteValue, &definitions)

	return &definitions
}