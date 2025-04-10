package parser

type Flag struct {
	Label string
	Value any
}

type ParseOptions struct {
	Second string
	Minute string
	Hour   string
	Dom    string
	Month  string
	Dow    string
	Flags  []Flag
}
