package modelos

type IsoDe55 struct {
	AC9F26                string `iso8583:"9F26"`
	InfoCript9F27         string `iso8583:"9F27"`
	IAD9F10               string `iso8583:"9F10"`
	NumImprevi9F37        string `iso8583:"9F37"`
	ContadorTrnApp9F36    string `iso8583:"9F36"`
	TRV95                 string `iso8583:"95"`
	DataTrn9A             string `iso8583:"9A"`
	TipoTrn9C             string `iso8583:"9C"`
	ValorAutoriz9F02      string `iso8583:"9F02"`
	CodMoedaTrn5F2A       string `iso8583:"5F2A"`
	PerfilIntercApp82     string `iso8583:"82"`
	CodPaisTerm9F1A       string `iso8583:"9F1A"`
	ResultCVM9F34         string `iso8583:"9F34"`
	CapacidadeTermEUA9F33 string `iso8583:"9F33"`
	NumDedicado84         string `iso8583:"84"`
	OutroValor9F03        string `iso8583:"9F03"`
	////opcionais abaixo
	NumSeqPan5F34         string `iso8583:"5F34"`
	DadoProprCadSelAp9F0A string `iso8583:"9F0A"`
	TipoTerminal9F35      string `iso8583:"9F35"`
	IDF9F1E               string `iso8583:"9F1E"`
	CodCatgTrn9F53        string `iso8583:"9F53"`
	NumVersaoApp9F09      string `iso8583:"9F09"`
	ContadorSeqTrn9F41    string `iso8583:"9F41"`
	//TermCapForaEUA9F33    string `iso8583:"9F33"`
	DadosTerceiro9F6E string `iso8583:"9F6E"`
}
