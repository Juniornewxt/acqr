package modelos

type IsoDe48sub struct {
	//MerchantName1        string `iso8583:"01"`
	MerchantCategoryCode string `iso8583:"14"`
	//Sub48_22i            *De48sub22 `iso8583:"22"`
	VerifSenha         string `iso8583:"20"`
	Sub48_22i          string `iso8583:"22"`
	InfoCarteira       string `iso8583:"26"`
	InfoPOSLocal       string `iso8583:"29"`
	MerchantPostalCode string `iso8583:"74"`
	PaymentIndicator   string `iso8583:"77"`
	MerchantWebsite    string `iso8583:"80"`
	Ard                string `iso8583:"91"`
	InfoParcelado      string `iso8583:"95"`
}
type IsoDe48 struct {
	TCC          string      `iso8583:"00"`
	CompsicaoTCC *IsoDe48sub `iso8583:"01"`
	//InfoParcelado string      `iso8583:"95"`
}
