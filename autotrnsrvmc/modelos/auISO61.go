package modelos

type IsoDe61 struct {
	POSta01          string `iso8583:"01"`
	Reserv02         string `iso8583:"02"`
	POStl03          string `iso8583:"03"`
	POSchPresent04   string `iso8583:"04"`
	POScardPresent05 string `iso8583:"05"`
	POScardCapCap06  string `iso8583:"06"`
	POStrnStatus07   string `iso8583:"07"`
	POStrnSecur08    string `iso8583:"08"`
	Reserv09         string `iso8583:"09"`
	POSCat10         string `iso8583:"10"`
	POSIputCap11     string `iso8583:"11"`
	POSliveCycle12   string `iso8583:"12"`
	POScodPais13     string `iso8583:"13"`
	POScodPostal14   string `iso8583:"14"`
}
