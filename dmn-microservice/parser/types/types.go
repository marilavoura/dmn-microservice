package types

import "encoding/xml"

// XML types

type Definitions struct {
	XMLName 		xml.Name		`xml:"definitions"`
	Decision		Decision		`xml:"decision"`
}

type Decision struct {
	XMLName 		xml.Name		`xml:"decision"`
	DecisionTable	DecisionTable	`xml:"decisionTable"`
}

type DecisionTable struct {
	XMLName 		xml.Name		`xml:"decisionTable"`
	HitPolicy		string			`xml:"hitPolicy,attr"`
	Input			[]Input			`xml:"input"`
	Output			[]Output		`xml:"output"`
	Rule			[]Rule			`xml:"rule"`
}

type Input struct {
	XMLName 		xml.Name		`xml:"input"`
	Label			string			`xml:"label,attr"`
	InputExpression	InputExpression	`xml:"inputExpression"`
	InputValues		InputValues 	`xml:"inputValues"`
}

type InputExpression struct {
	XMLName 		xml.Name		`xml:"inputExpression"`
	TypeRef			string			`xml:"typeRef,attr"`
	Text			string			`xml:"text"`
}

type InputValues struct {
	XMLName 		xml.Name		`xml:"inputValues"`
	Text			string			`xml:"text"`
}

type Output struct {
	XMLName 		xml.Name		`xml:"output"`
}

type Rule struct {
	XMLName 		xml.Name		`xml:"rule"`
	InputEntry 		[]InputEntry	`xml:"inputEntry"`
	OutputEntry 	[]OutputEntry	`xml:"outputEntry"`
}

type InputEntry struct {
	XMLName 		xml.Name		`xml:"inputEntry"`
	Text			string			`xml:"text"`
}

type OutputEntry struct {
	XMLName 		xml.Name		`xml:"outputEntry"`
	Text			string			`xml:"text"`
}

// Definition

type InputDefinition struct {
	InputName		string
	AllowedType		string
	AllowedValues 	[]string
}

type Rules struct {
	Inputs			[]string
	Outputs			[]string
}



