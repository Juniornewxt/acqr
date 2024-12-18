package manipulaiso

import (

	//"encoding/csv"
	"bytes"
	"log"
	"os"
	"strings"
	"time"

	"github.com/moov-io/iso8583"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"

	//"github.com/moov-io/iso8583/specs"
	especificacao "autotrnsrvmc/especificacao"

	"github.com/moov-io/iso8583/network"
)

func Padraoiso100erro(iso_padrao []byte) ([]byte, error) {

	//Inicia registro de horas
	data := time.Now()
	// 1 se renomeia a iso criada acima com pacote iso_padrao

	//fmt.Println("PRINT DA FUNCAO MC", iso_padrao)

	// Remove os primeiros 2 bytes
	iso_padrao_mod := iso_padrao[2:]
	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_padrao_mod)

	p_spec := especificacao.NewSpecASCII()

	p_message := iso8583.NewMessage(p_spec)

	p_message.Unpack([]byte((retorno)))

	b, err := p_message.Pack()
	if err != nil {
		panic(err)
	}
	log.Printf("MENSAGEM RECEBIDA TRATAMENTO ISO PADRAO ERRO:\n")
	log.Printf("% x\n", b)
	iso8583.Describe(p_message, os.Stdout)

	// Inicia a criação da spec nova para que não se repita os dados da ISO original, salve quando solicitado p_message
	new_spec := especificacao.NewSpecASCII()

	pnew_message := iso8583.NewMessage(new_spec)

	p_message.Unpack(b) //Abre a ISO formatada no inicio para ser usada somente de alguns campos

	pnew_message.MTI("0210")
	proccod, err := p_message.GetString(3)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(3, proccod) //proc code
	valor, err := p_message.GetString(4)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(4, valor) //valor
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
	//hhmmss, err := p_message.GetString(12)
	//if err != nil {
	//panic(err)
	//}
	pnew_message.Field(12, data.Format(("150405"))) //hhmmss
	//mmdd, err := p_message.GetString(13)
	////if err != nil {
	//panic(err)
	//}
	pnew_message.Field(13, data.Format(("0102"))) //mmdd
	///
	auto_resp, err := p_message.GetString(38)
	if err != nil {
		panic(err)
	}
	if auto_resp == "" {
		pnew_message.Field(38, data.Format(("150405")))
	} else {
		pnew_message.Field(38, auto_resp) //codigo de resposta
	}
	//pnew_message.Field(38, auto_resp) //codigo de resposta
	//cod_auto, err := p_message.GetString(39)
	//if err != nil {
	//	panic(err)
	//}
	pnew_message.Field(39, "96") //codigo autorizacao
	//
	tid, err := p_message.GetString(41)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(41, tid) //terminal
	mid, err := p_message.GetString(42)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(42, mid) //cod comercio
	pnew_message.Field(41, tid) //terminal
	de48, err := p_message.GetString(48)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(48, de48)
	moeda, err := p_message.GetString(49)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(49, moeda) //moeda
	////////////////////////////////////////////////////////////////
	de55, err := p_message.GetString(55)
	if err != nil {
		panic(err)
	}
	//fmt.Println("DE 55 QUE VOLTOU", de55)
	if de55 != "" {
		pnew_message.Field(55, "9100") //dados do chip
	}
	//pnew_message.Field(55, de55) //dados do chip
	//entrymode, err := p_message.GetString(22)
	//if err != nil {
	//	panic(err)
	//}
	//if entrymode == "51" || entrymode == "71" {
	//	pnew_message.Field(55, "FF210F91091122334455667788998A023030")
	//}
	//if cod_auto == "" {
	//	pnew_message.Field(60, "C@    DEBIT MASTERCARD - Via Cliente    @TESTEPAY - TESTE111111 E            @RUA XXXX XXXXX C                 SP BR@CNPJ:01000000000100                   @TID: GT000131@EC:XXXXXXXXX@                    @VENDA CREDITO A VISTA                 @************0045   @19/06/24                         09:10@VALOR APROVADO:               R$ 11,00@CV:XXXXXXXXXXXX            AUTO:000000@DOC:000000@TERM:XXXXXXXX@@               ")
	//	pnew_message.Field(62, "C@    DEBIT MASTERCARD - Via Cliente    @TESTEPAY - TESTE111111 E            @RUA XXXX XXXXX C                 SP BR@CNPJ:01000000000100                   @TID: GT000131@EC:XXXXXXXXX@                    @VENDA CREDITO A VISTA                 @************0045   @19/06/24                         09:10@VALOR APROVADO:               R$ 11,00@CV:XXXXXXXXXXXX            AUTO:000000@DOC:000000@TERM:XXXXXXXX@@ MEDIANTE A SENHA               ")
	//}
	//data := time.Now()
	pnew_message.Field(127, strings.Replace((data.Format(("150405.000000"))), ".", "", -1))
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
