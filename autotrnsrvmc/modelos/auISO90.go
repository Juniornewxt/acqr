package modelos

type IsoDe90 struct {
	MtiOrin    string `iso8583:"01"`
	StanOrin   string `iso8583:"02"`
	HhmmssOrin string `iso8583:"03"`
	AcqrOrin   string `iso8583:"04"`
	InstFOrin  string `iso8583:"05"`
}
