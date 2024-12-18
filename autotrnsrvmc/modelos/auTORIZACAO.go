package modelos

type SolicitacaoAutorizacao struct {
	MTI               string   `iso8583:"0"`
	PAN               string   `iso8583:"2"`
	ProcCod           string   `iso8583:"3"`
	Valor             string   `iso8583:"4"`
	DataHoraTransacao string   `iso8583:"7"`
	STAN              string   `iso8583:"11"`
	Hhmmss            string   `iso8583:"12"`
	Mmdd              string   `iso8583:"13"`
	Validade          string   `iso8583:"14"`
	SetlmMmdd         string   `iso8583:"15"`
	Mcc               string   `iso8583:"18"`
	EntryMode         string   `iso8583:"22"`
	SeqCardNum        string   `iso8583:"23"`
	Adquirente        string   `iso8583:"32"`
	CodInstiRemetente string   `iso8583:"33"`
	Trilha2           string   `iso8583:"35"`
	Nsu               string   `iso8583:"37"`
	Cod_auto          string   `iso8583:"38"`
	Cod_resp          string   `iso8583:"39"`
	Terminal          string   `iso8583:"41"`
	CodComercio       string   `iso8583:"42"`
	NomeEndereco      string   `iso8583:"43"`
	InfoAdd           *IsoDe48 `iso8583:"48"`
	Moeda             string   `iso8583:"49"`
	Pin               string   `iso8583:"52"`
	DadosChip         *IsoDe55 `iso8583:"55"`
	AdvReasonCod      string   `iso8583:"60"`
	DadosPOS          *IsoDe61 `iso8583:"61"`
	DadosRede         string   `iso8583:"63"`
	DadosCanc         *IsoDe90 `iso8583:"90"`
	ReplecAmount      string   `iso8583:"95"`
	DadosADDParcelado string   `iso8583:"112"`
	DadosRegistro     string   `iso8583:"120"`
	DadosPrivados     string   `iso8583:"126"`
}
