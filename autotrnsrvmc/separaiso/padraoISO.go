package separaiso

import (

	//"encoding/csv"

	manipulaiso "autotrnsrvmc/manipulaiso"
	"fmt"
	"log"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"
	//"github.com/moov-io/iso8583/specs"
)

func Padraoiso(iso_mti []byte) ([]byte, error) {

	// 1 se renomeia a iso criada acima com pacote iso_padrao

	//fmt.Println("PRINT DA FUNCAO PADRAO", iso_mti)

	// Remove os primeiros 2 bytes
	iso_mti_mod := iso_mti[2:]
	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_mti_mod)

	//r_spec := especificacao.NewSpecEBCDIC()

	//r_message := iso8583.NewMessage(r_spec)

	//r_message.Unpack([]byte((retorno)))

	//valida_mti, err := r_message.GetMTI()
	//if err != nil {
	//	panic(err)
	//}
	valida_mti, err := PerguntaAscii(retorno)
	if err != nil {
		panic(err)
	}
	valida_mtiEBC, err := PerguntaEbcdic(retorno)
	if err != nil {
		panic(err)
	}
	//fmt.Println("PRINT DO MTI PADRAO", valida_mti)
	//valida_mti = string(valida_mti)
	if valida_mti == "0200" {
		enviado, err := manipulaiso.Padraoiso100erro(iso_mti)
		if err != nil {
			fmt.Println("Erro ao encaminhar mensagem 0100 iso Padrao:", err)
			return enviado, err
		}
		//fmt.Println("RECEBIDO DO MC QUE RECEBEU A 0110", enviado)
		return enviado, err
	} else if valida_mti == "0400" {
		enviado, err := manipulaiso.Padraoiso400erro(iso_mti)
		if err != nil {
			fmt.Println("Erro ao encaminhar mensagem 0400 iso Padrao:", err)
			return enviado, err
		}
		//fmt.Println("RECEBIDO DO MC QUE RECEBEU A 0110", enviado)
		return enviado, err
	} else if valida_mtiEBC == "0110" {
		enviado, err := manipulaiso.Padraoiso100(iso_mti)
		if err != nil {
			fmt.Println("Erro ao encaminhar mensagem 0110 iso Padrao:", err)
			return enviado, err
		}
		//fmt.Println("RECEBIDO DO MC QUE RECEBEU A 0110", enviado)
		return enviado, err
	} else if valida_mtiEBC == "0410" {
		enviado, err := manipulaiso.Padraoiso400(iso_mti)
		if err != nil {
			fmt.Println("Erro ao encaminhar mensagem 0410 iso Padrao:", err)
			return enviado, err
		}
		//fmt.Println("RECEBIDO DO MC QUE RECEBEU A 0410", enviado)
		return enviado, err
	} else if valida_mtiEBC == "0210" {
		enviado, err := manipulaiso.Padraoiso100(iso_mti)
		if err != nil {
			fmt.Println("Erro ao encaminhar mensagem 0210 iso Padrao:", err)
			return enviado, err
		}
		//fmt.Println("RECEBIDO DO MC QUE RECEBEU A 0410", enviado)
		return enviado, err
	} else {
		log.Println("SAI Aqui")
		err := fmt.Errorf("MTI inválido: %s", valida_mti)
		fmt.Println("Erro:", err)
		return nil, err
	}
}
