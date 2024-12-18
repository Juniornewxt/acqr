package separaiso

import (

	//"encoding/csv"
	especificacao "autotrnsrvmc/especificacao"

	"github.com/moov-io/iso8583"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"
	//"github.com/moov-io/iso8583/specs"
)

func PerguntaAscii(retorno string) (string, error) {

	r_spec := especificacao.NewSpecASCII()

	r_message := iso8583.NewMessage(r_spec)

	r_message.Unpack([]byte((retorno)))

	valida_mti, err := r_message.GetMTI()
	if err != nil {
		panic(err)
	}
	return valida_mti, err
}
func PerguntaEbcdic(retorno string) (string, error) {

	r_spec := especificacao.NewSpecEBCDICret()

	r_message := iso8583.NewMessage(r_spec)

	r_message.Unpack([]byte((retorno)))

	valida_mti, err := r_message.GetMTI()
	if err != nil {
		panic(err)
	}
	return valida_mti, err
}
