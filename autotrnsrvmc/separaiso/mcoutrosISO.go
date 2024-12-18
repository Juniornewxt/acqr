package separaiso

import (

	//"encoding/csv"

	especificacao "autotrnsrvmc/especificacao"
	manipulaiso "autotrnsrvmc/manipulaiso"
	"fmt"

	"github.com/moov-io/iso8583"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"
	//"github.com/moov-io/iso8583/specs"
)

func MCoutrosiso(iso_mti []byte) ([]byte, error) {

	// 1 se renomeia a iso criada acima com pacote iso_padrao

	//fmt.Println("PRINT DA FUNCAO PADRAO", iso_mti)

	// Remove os primeiros 2 bytes
	iso_mti_mod := iso_mti[2:]
	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_mti_mod)

	r_spec := especificacao.NewSpecEBCDICret()

	r_message := iso8583.NewMessage(r_spec)

	r_message.Unpack([]byte((retorno)))

	valida_mti, err := r_message.GetMTI()
	if err != nil {
		panic(err)
	}
	cod_auto, err := r_message.GetString(39)
	if err != nil {
		panic(err)
	}
	//pnew_message.Field(39, cod_auto) //codigo autorizacao
	//fmt.Println("PRINT DO MTI PADRAO", valida_mti)
	//valida_mti = string(valida_mti)
	if valida_mti == "0110" && (cod_auto == "" || cod_auto == "00") {
		enviado, err := manipulaiso.McMti180(iso_mti)
		if err != nil {
			fmt.Println("Erro ao encaminhar mensagem 0110 para confirmaçao Mti180:", err)
			return enviado, err
		}
		//fmt.Println("RECEBIDO DO MC QUE RECEBEU A 0110", enviado)
		return enviado, err
		//} else if valida_mti == "0410" {
		//	enviado, err := manipulaiso.Padraoiso400(iso_mti)
		//	if err != nil {
		//		fmt.Println("Erro ao encaminhar mensagem 0410 iso Padrao:", err)
		//		return enviado, err
		//	}
		//fmt.Println("RECEBIDO DO MC QUE RECEBEU A 0410", enviado)
		//return enviado, err
	} else {
		err := fmt.Errorf("Esse Codigo Resposta Nao precisa de confirmacao: %s", string(cod_auto))
		fmt.Println("Erro:", err)
		return nil, err
	}
}
