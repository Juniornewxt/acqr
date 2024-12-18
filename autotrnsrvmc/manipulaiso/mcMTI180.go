package manipulaiso

import (

	//"encoding/csv"
	"bytes"
	"log"
	"os"

	"github.com/moov-io/iso8583"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"

	//"github.com/moov-io/iso8583/specs"
	especificacao "autotrnsrvmc/especificacao"

	"github.com/moov-io/iso8583/network"
)

func McMti180(iso_padrao []byte) ([]byte, error) {

	//Inicia registro de horas
	//data := time.Now()
	// 1 se renomeia a iso criada acima com pacote iso_padrao

	// Remove os primeiros 2 bytes
	iso_padrao_mod := iso_padrao[2:]
	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_padrao_mod)

	p_spec := especificacao.NewSpecEBCDICret()

	p_message := iso8583.NewMessage(p_spec)

	p_message.Unpack([]byte((retorno)))

	b, err := p_message.Pack()
	if err != nil {
		panic(err)
	}
	log.Printf("MENSAGEM RECEBIDA PARA CONFIRMACAO DE AUTORIZACAO:\n")
	log.Printf("% x\n", b)
	iso8583.Describe(p_message, os.Stdout)

	// Inicia a criação da spec nova para que não se repita os dados da ISO original, salve quando solicitado p_message
	new_spec := especificacao.NewSpecEBCDICret()

	pnew_message := iso8583.NewMessage(new_spec)

	p_message.Unpack(b) //Abre a ISO formatada no inicio para ser usada somente de alguns campos

	pnew_message.MTI("0180")
	mmddhhmmss, err := p_message.GetString(7)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(7, mmddhhmmss) //data e hora mmddhhmmss
	sy_trace, err := p_message.GetString(11)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(11, sy_trace)
	cod_auto, err := p_message.GetString(39)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(39, cod_auto) //codigo autorizacao

	dadosrede, err := p_message.GetString(63)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(63, dadosrede)
	//pnew_message.Field(127, strings.Replace((data.Format(("150405.000000"))), ".", "", -1))

	d, err := pnew_message.Pack()
	if err != nil {
		panic(err)

	}

	// 1 se renomeia a iso criada acima com pacote b
	packed := d

	// 2 se criar o cabeçalho binario 2 byts
	header := network.NewBinary2BytesHeader()
	//header := network.NewASCII4BytesHeader()
	header.SetLength(len(packed))

	// 3 se criar o tpdu ELO
	//tpdu := []byte("60 00 06 00 00")

	// 4 - combinar tudo o que temos em um buf

	var buf bytes.Buffer

	header.WriteTo(&buf)
	//_, err = buf.Write(tpdu)
	_, err = buf.Write(packed)
	//_, err = Write(buf.Bytes())
	log.Printf("MENSAGEM A SER ENVIADA:\n")
	iso8583.Describe(pnew_message, os.Stdout)
	log.Printf("FIM DO PROCESSAMENTO :) \n")
	//fmt.Println("Pressione 'Enter' para sair...")
	//fmt.Scanln()
	//conn.Close()
	return buf.Bytes(), err
}
