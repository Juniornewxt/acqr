package manipulaiso

import (

	//"encoding/csv"
	"bytes"
	"log"
	"os"
	"time"

	"github.com/moov-io/iso8583"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"

	//"github.com/moov-io/iso8583/specs"
	especificacao "autoadmsrvmc/especificacao"

	"github.com/moov-io/iso8583/network"
)

func Ativacao0800() ([]byte, error) {

	data := time.Now()

	spec := especificacao.NewSpecEBCDIC()

	message := iso8583.NewMessage(spec)

	//message.SetHeader("0587")
	//network.NewBinary2BytesHeader()
	message.MTI("0800")                           //587
	message.Field(2, "00026")                     //pan
	message.Field(7, data.Format(("0102150405"))) //data e hora mmddhhmmss
	message.Field(11, data.Format(("150405")))    //trace
	message.Field(33, "0026")                     //Forwarding Institution ID Code
	message.Field(70, "081")                      //270 status conexao  //081 ativacao de host //082 desativacao de host
	b, err := message.Pack()
	if err != nil {
		panic(err)

	}
	log.Printf("MENSAGEM ENVIADA \n")
	iso8583.Describe(message, os.Stdin)
	log.Printf("% x\n", b)

	log.Printf("MENSAGEM FORMATADA \n")
	iso8583.Describe(message, os.Stdin)

	packed := b

	header := network.NewBinary2BytesHeader()

	header.SetLength(len(packed))

	var buf bytes.Buffer

	header.WriteTo(&buf)

	_, err = buf.Write(packed)

	return buf.Bytes(), err
}
