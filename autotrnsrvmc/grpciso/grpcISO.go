package grpciso

import (

	//"encoding/csv"

	especificacao "autotrnsrvmc/especificacao"
	"fmt"
	"os"

	//manipulaiso "projetoacqr/manipulaiso"

	"github.com/moov-io/iso8583"
)

func GrpcISO(MTI, DE02, DE03, DE04, DE07, DE11, DE12, DE13, DE14, DE18, DE19, DE22, DE23, DE32, DE33, DE35, DE38, DE39, DE41, DE42, DE43, DE48, DE49, DE52, DE55, DE61, DE120, DE126 string) (string, error) {

	spec := especificacao.NewSpecASCII()

	message := iso8583.NewMessage(spec)

	message.MTI(MTI) //587

	message.Field(3, DE03) //proc code
	message.Field(4, DE04) //valor
	message.Field(7, DE07) //data e hora mmddhhmmss
	message.Field(11, DE11)
	message.Field(12, DE12) //hhmmss
	message.Field(13, DE13) //mmdd
	message.Field(14, DE14) //validade do cartao
	message.Field(18, DE18) //mcc

	message.Field(22, DE22) //entry mode

	message.Field(23, DE23) //card sequence number

	message.Field(32, DE32) //adquirente
	message.Field(33, DE33) //Forwarding Institution ID Code
	message.Field(35, DE35) // trilha 2 cartao/ identificador/ validade/ cvv

	message.Field(41, DE41) //terminal
	message.Field(42, DE42) //cod comercio
	message.Field(43, DE43)
	message.Field(48, DE48)
	message.Field(49, DE49) //moeda

	message.Field(52, DE52) //senha

	message.Field(55, DE55)

	message.Field(61, DE61)

	message.Field(120, DE120)

	message.Field(126, DE126)

	iso_criada, err := message.Pack()
	if err != nil {
		panic(err)

	}

	fmt.Printf("GRPC MENSAGEM FORMATADA ISO\n")
	iso8583.Describe(message, os.Stdin)
	//Adiciona um header ficticio para nao quebrar os outros blocos
	hiso_criada := "00" + string(iso_criada)
	return hiso_criada, nil
}
