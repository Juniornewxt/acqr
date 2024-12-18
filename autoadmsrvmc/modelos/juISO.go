package modelos

// Estrutura do DE 48
type IsoDe48ret struct {
	De48sub01 string `isoJUNIOR:"01"`
	De48sub02 string `isoJUNIOR:"02"`
	De48sub03 string `isoJUNIOR:"03"`
	De48sub04 string `isoJUNIOR:"04"`
	De48sub05 string `isoJUNIOR:"05"`
	De48sub06 string `isoJUNIOR:"06"`
	De48sub07 string `isoJUNIOR:"07"`
	De48sub08 string `isoJUNIOR:"08"`
	De48sub09 string `isoJUNIOR:"09"`
	De48sub10 string `isoJUNIOR:"10"`
	De48sub11 string `isoJUNIOR:"11"`
	De48sub12 string `isoJUNIOR:"12"`
}

// Estrutura do subcampos DE 48 11
type IsoDe48sub11ret struct {
	IdClasseChave     string `dado:"1" tamanho:"2"`
	NumIndiceChave    string `dado:"2" tamanho:"2"`
	NumCicloChave     string `dado:"3" tamanho:"2"`
	ChavePEKcripSenha string `dado:"4" tamanho:"32"`
	ValorVerifChave   string `dado:"5" tamanho:"16"`
}
