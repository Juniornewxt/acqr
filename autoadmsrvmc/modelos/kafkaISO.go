package modelos

type MensagemAuth struct {
	Mti               string `json:"mti              "`
	Pan               string `json:"pan              "`
	ProcCod           string `json:"proccod          "`
	Valor             string `json:"valor            "`
	DataHoraTransacao string `json:"datahoratransacao"`
	STAN              string `json:"stan             "`
	Hhmmss            string `json:"hhmmss           "`
	Mcc               string `json:"mcc              "`
	EntryMode         string `json:"entrymode        "`
	Adquirente        string `json:"adquirente       "`
	CodInstiRemetente string `json:"codinstiremetente"`
	Nsu               string `json:"nsu              "`
	CodAuto           string `json:"codauto          "`
	Codresp           string `json:"codresp          "`
	Terminal          string `json:"terminal         "`
	CodComercio       string `json:"codcomercio      "`
	NomeEndereco      string `json:"nomeendereco     "`
	Moeda             string `json:"moeda            "`
	DadosADDParcelado string `json:"dadosaddparcelado"`
}
type MensagemAdm struct {

	//mtirecebido, pan, dmhhmmss, stan, fiic, codresp, addinfo, netdata, codrede, privatedata

	Mti         string `json:"mti"`
	Pan         string `json:"pan"`
	Dmhhmmss    string `json:"dmhhmmss"`
	Stan        string `json:"stan"`
	Fiic        string `json:"fiic"`
	Codresp     string `json:"codresp"`
	Addinfo     string `json:"addinfo"`
	Netdata     string `json:"netdata"`
	Codrede     string `json:"codrede"`
	Privatedata string `json:"privatedata"`
}
