package separaiso

import (
	"fmt"
	//"encoding/csv"

	manipulaiso "autotrnsrvmc/manipulaiso"
	"log"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"
	//"github.com/moov-io/iso8583/specs"
)

func Mciso(iso_mti []byte) ([]byte, error) {

	// 1 se renomeia a iso criada acima com pacote iso_padrao

	//fmt.Println("PRINT DA FUNCAO MC", iso_mti)

	// Remove os primeiros 2 bytes
	iso_mti_mod := iso_mti[2:]
	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_mti_mod)

	//r_spec := especificacao.NewSpecASCII()

	//r_message := iso8583.NewMessage(r_spec)

	//r_message.Unpack([]byte((retorno)))

	//valida_mti, err := r_message.GetMTI()
	valida_mti, err := PerguntaAscii(retorno)
	if err != nil {
		panic(err)
	}
	valida_mtiEBC, err := PerguntaEbcdic(retorno)
	if err != nil {
		panic(err)
	}
	//log.Println("PRINT DO MTI PADRAO", valida_mti)
	//valida_mti = string(valida_mti)
	if valida_mti == "0200" {
		enviado, err := manipulaiso.Mciso100(iso_mti)
		if err != nil {
			log.Println("Erro ao encaminhar mensagem 0200 iso Padrao:", err)
			return enviado, err
		}
		//fmt.Println("RECEBIDO DO MC QUE RECEBEU A 0200", enviado)
		return enviado, err
	} else if valida_mti == "0400" {
		enviado, err := manipulaiso.Mciso400(iso_mti)
		if err != nil {
			log.Println("Erro ao encaminhar mensagem 0800 iso Padrao:", err)
			return enviado, err
		}
		//	fmt.Println("RECEBIDO DO MC QUE RECEBEU A 0400", enviado)
		return enviado, err
	} else if valida_mtiEBC == "0410" {
		enviado, err := manipulaiso.McisoPrint(iso_mti)
		if err != nil {
			log.Println("Erro ao encaminhar mensagem 0410 iso Print:", err)
			return enviado, err
		}
		//	fmt.Println("RECEBIDO DO MC QUE RECEBEU A 0400", enviado)
		return enviado, err
	} else if valida_mtiEBC == "0430" {
		enviado, err := manipulaiso.McisoPrint(iso_mti)
		if err != nil {
			log.Println("Erro ao encaminhar mensagem 0430 iso Print:", err)
			return enviado, err
		}
		//	fmt.Println("RECEBIDO DO MC QUE RECEBEU A 0400", enviado)
		return enviado, err
	} else {
		err := fmt.Errorf("MTI inválido: %s", valida_mti)
		log.Println("Erro:", err)
		return nil, err
	}
}
