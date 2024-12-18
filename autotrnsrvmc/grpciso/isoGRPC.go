package grpciso

import (

	//"encoding/csv"

	"log"
	"os"

	"github.com/moov-io/iso8583"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"

	//"github.com/moov-io/iso8583/specs"
	especificacao "autotrnsrvmc/especificacao"
	"autotrnsrvmc/modelos"
)

// func IsoGRPC(iso_padrao []byte) (Campos, error) {
func IsoGRPC(iso_padrao string) (modelos.GrpcIsoM, error) {

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
	log.Printf("MENSAGEM RECEBIDA CONVERSAO ISO GRPC:\n")
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
	hhmmss, err := p_message.GetString(12)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(12, hhmmss) //hhmmss
	mmdd, err := p_message.GetString(13)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(13, mmdd) //mmdd
	///
	auto_resp, err := p_message.GetString(38)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(38, auto_resp) //codigo de resposta
	cod_resp, err := p_message.GetString(38)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(38, cod_resp) //codigo de resposta

	cod_auto, err := p_message.GetString(39) //codigo autorizacao
	if err != nil {
		panic(err)
	}
	pnew_message.Field(39, cod_auto) //terminal
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
	de55, err := p_message.GetString(55)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(55, de55) //dados do chip
	de127, err := p_message.GetString(127)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(127, de127) //dados do chip
	//	d, err := pnew_message.Pack()
	//if err != nil {
	//	panic(err)

	//}
	log.Printf("MENSAGEM A SER ENVIADA PARA GRPC:\n")
	iso8583.Describe(pnew_message, os.Stdout)
	log.Printf("FIM DO PROCESSAMENTO :) \n")
	//fmt.Println("Pressione 'Enter' para sair...")
	//fmt.Scanln()
	//conn.Close()
	//return buf.Bytes(), err
	// Retorna os campos em uma estrutura
	return modelos.GrpcIsoM{
		DE11: (sy_trace),
		DE39: cod_auto,
		DE41: (sy_trace),
		DE43: (de48),
	}, nil
}
