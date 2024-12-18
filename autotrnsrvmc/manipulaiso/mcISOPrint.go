package manipulaiso

import (

	//"encoding/csv"

	especificacao "autotrnsrvmc/especificacao"
	"log"
	"os"

	"github.com/moov-io/iso8583"
)

func McisoPrint(iso_padrao []byte) ([]byte, error) {
	// Remove os primeiros 2 bytes
	iso_mod := iso_padrao[2:]

	retorno := string(iso_mod)

	p_spec := especificacao.NewSpecEBCDIC()

	p_message := iso8583.NewMessage(p_spec)

	p_message.Unpack([]byte((retorno)))

	b, err := p_message.Pack()
	if err != nil {
		panic(err)
	}
	log.Printf("MENSAGEM RECEBIDA DA BANDEIRA:\n")
	log.Printf("% x\n", b)
	iso8583.Describe(p_message, os.Stdout)
	log.Printf("FIM DO PROCESSAMENTO :) \n")

	return nil, err
}
