package separaiso

import (

	//"encoding/csv"
	especificacao "autoadmsrvmc/especificacao"
	manipulaiso "autoadmsrvmc/manipulaiso"
	"log"

	"github.com/moov-io/iso8583"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"
	//"github.com/moov-io/iso8583/specs"
)

func Admiso(iso_mti []byte) ([]byte, error) {

	// Remove os primeiros 2 bytes
	iso_mti_mod := iso_mti[2:]

	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_mti_mod)

	r_spec := especificacao.NewSpecEBCDIC()

	r_message := iso8583.NewMessage(r_spec)

	r_message.Unpack([]byte((retorno)))

	valida_mti, err := r_message.GetMTI()
	if err != nil {
		panic(err)
	}
	if valida_mti < "0800" {
		log.Println("MENSAGEM TARDIA DESCARTADA, SERÁ ENVIADO DESFAZIMENTO PELO SERVTRN MTI:", valida_mti)
		return []byte(err.Error()), nil
	}

	valida_codrede, err := r_message.GetString(70)
	if err != nil {
		panic(err)
	}

	valida_dados_rede, err := r_message.GetString(63)
	if err != nil {
		panic(err)
	}

	// Valida se um pedido de single ou dual msg
	valida_dados_band := valida_dados_rede[:2]

	//Tratamento de msg credito
	if valida_mti == "0800" && valida_codrede == "161" && valida_dados_band == "MC" {
		enviado, err := manipulaiso.TrocaChave([]byte(iso_mti))
		if err != nil {
			log.Println("Erro ao encaminhar mensagem 0800 para troca chave dual msg:", err)
			return enviado, err
		}
		//Tratamento de msg debito
		return enviado, err
	} else if valida_mti == "0800" && valida_codrede == "161" && valida_dados_band == "MS" {
		enviado, err := manipulaiso.TrocaChaveMaestro([]byte(iso_mti))
		if err != nil {
			log.Println("Erro ao encaminhar mensagem 0800 para troca chave single msg:", err)
			return enviado, err
		}

		return enviado, err
		// Tratamento de confirmacao de msg
	} else if valida_mti == "0820" {
		enviado, err := manipulaiso.RetornoADM([]byte(iso_mti))
		if err != nil {
			log.Println("Erro ao encaminhar mensagem 0800 iso Padrao:", err)
			return enviado, err
		}

		return enviado, err
	} else {
		//err := fmt.Errorf("MTI inválido: %s", valida_mti)
		log.Println("Erro: Nenhuma condicao no admISO foi encontrada para essa mensagem:", iso_mti_mod, err)
		return nil, err
	}
}
func Rastreio(iso_mti []byte) ([]byte, error) {

	// Remove os primeiros 2 bytes
	iso_mti_mod := iso_mti[2:]
	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_mti_mod)

	r_spec := especificacao.NewSpecEBCDIC()

	r_message := iso8583.NewMessage(r_spec)

	r_message.Unpack([]byte((retorno)))

	ddmmhhmmss, err := r_message.GetString(07)
	if err != nil {
		panic(err)
	}

	stan, err := r_message.GetString(11)
	if err != nil {
		panic(err)
	}
	uneosdois := stan + ddmmhhmmss
	//log.Println("STAN LOCALIZADO", stan)
	return []byte(uneosdois), err
}
