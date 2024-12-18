package modelos

type Iso48Entrada struct {
	VersaoTB           string `isoJUNIOR:"001"`
	CargaDesatualizada string `isoJUNIOR:"002"`
	Tcc                string `isoJUNIOR:"003"`
	TipoParceleado     string `isoJUNIOR:"021"`
	Parcelas           string `isoJUNIOR:"022"`
	VlrConsulCrediario string `isoJUNIOR:"023"`
	DocumentoSigleMSN  string `isoJUNIOR:"024"`
}

// Estrutura do DE 61
type IsoDe61ret struct {
	DadosAplicao01    string `isoJUNIOR:"001"`
	NumSeriePinpad02  string `isoJUNIOR:"002"`
	DadosBandeira03   string `isoJUNIOR:"003"`
	VersaoAplicacao04 string `isoJUNIOR:"004"`
	FabricaPinpad05   string `isoJUNIOR:"005"`
	HardwareVersao06  string `isoJUNIOR:"006"`
	FirmwarePinpad07  string `isoJUNIOR:"007"`
}

// Estrutura do subcampos DE 61 03
type IsoDe61sub03ret struct {
	POSta01          string `dado:"1" tamanho:"1"`
	Reserv02         string `dado:"2" tamanho:"1"`
	POStl03          string `dado:"3" tamanho:"1"`
	POSchPresent04   string `dado:"4" tamanho:"1"`
	POScardPresent05 string `dado:"5" tamanho:"1"`
	POScardCapCap06  string `dado:"6" tamanho:"1"`
	POStrnStatus07   string `dado:"7" tamanho:"1"`
	POStrnSecur08    string `dado:"8" tamanho:"1"`
	Reserv09         string `dado:"9" tamanho:"1"`
	POSCat10         string `dado:"10" tamanho:"1"`
	POSIputCap11     string `dado:"11" tamanho:"1"`
	POSliveCycle12   string `dado:"12" tamanho:"2"`
	POScodPais13     string `dado:"12" tamanho:"3"`
	POScodPostal14   string `dado:"14" tamanho:"10"`
}
type ISO55dados struct {
	AC9F26                string `isoJUNIOR:"9F26"`
	InfoCript9F27         string `isoJUNIOR:"9F27"`
	IAD9F10               string `isoJUNIOR:"9F10"`
	NumImprevi9F37        string `isoJUNIOR:"9F37"`
	ContadorTrnApp9F36    string `isoJUNIOR:"9F36"`
	TRV95                 string `isoJUNIOR:"95"`
	DataTrn9A             string `isoJUNIOR:"9A"`
	TipoTrn9C             string `isoJUNIOR:"9C"`
	ValorAutoriz9F02      string `isoJUNIOR:"9F02"`
	CodMoedaTrn5F2A       string `isoJUNIOR:"5F2A"`
	PerfilIntercApp82     string `isoJUNIOR:"82"`
	CodPaisTerm9F1A       string `isoJUNIOR:"9F1A"`
	ResultCVM9F34         string `isoJUNIOR:"9F34"`
	CapacidadeTermEUA9F33 string `isoJUNIOR:"9F33"`
	NumDedicado84         string `isoJUNIOR:"84"`
	OutroValor9F03        string `isoJUNIOR:"9F03"`
	////opcionais abaixo
	NumSeqPan5F34         string `isoJUNIOR:"5F34"`
	DadoProprCadSelAp9F0A string `isoJUNIOR:"9F0A"`
	TipoTerminal9F35      string `isoJUNIOR:"9F35"`
	IDF9F1E               string `isoJUNIOR:"9F1E"`
	CodCatgTrn9F53        string `isoJUNIOR:"9F53"`
	NumVersaoApp9F09      string `isoJUNIOR:"9F09"`
	ContadorSeqTrn9F41    string `isoJUNIOR:"9F41"`
	//TermCapForaEUA9F33    string `isoJUNIOR:"9F33"`
	DadosTerceiro9F6E string `isoJUNIOR:"9F6E"`
}
type IsoDe90Entrada struct {
	MtiOrin    string `dado:"1" tamanho:"4"`
	StanOrin   string `dado:"2" tamanho:"6"`
	HhmmssOrin string `dado:"3" tamanho:"10"`
	AcqrOrin   string `dado:"4" tamanho:"11"`
	InstFOrin  string `dado:"5" tamanho:"11"`
}
