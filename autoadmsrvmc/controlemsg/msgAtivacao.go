package controlemsg

import (

	//"encoding/csv"
	//"bytes"
	"log"
	"os"

	//	"strings"
	//"time"
	"net"

	"github.com/moov-io/iso8583"

	"autoadmsrvmc/especificacao"
	"autoadmsrvmc/manipulaiso"
	"autoadmsrvmc/trabalhaconexoes"
	//"github.com/moov-io/iso8583/network"
)

func MsgAtivacao(connBandeira net.Conn) (string, error) {

	//Chama a funcao de msg para criar a mensagem 0800 de ativacao
	msn0, err := manipulaiso.Ativacao0800()
	// Verifica se houve erro
	if err != nil {
		log.Println("Erro:", err)
	} else {
		log.Println("Valor retornado:", msn0)
	}
	// Envia a mensagem recebida da porta de entrada para o servidor da bandeira
	msn1, err := trabalhaconexoes.EncaminhaMensagem2(string(msn0), connBandeira)
	if err != nil {
		log.Printf("Erro ao formatar resposta para cliente: %v", err)
		return "", err
	}
	//log.Println("MSN RECEBIDA", msn1)

	//Inicia registro de horas
	//data := time.Now()
	// 1 se renomeia a iso criada acima com pacote iso_padrao

	//fmt.Println("PRINT DA FUNCAO MC", iso_padrao)

	// Remove os primeiros 2 bytes
	iso_padrao := msn1
	iso_padrao_mod := iso_padrao[2:]
	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_padrao_mod)

	spec := especificacao.NewSpecEBCDIC()

	message := iso8583.NewMessage(spec)

	message.Unpack([]byte((retorno)))

	b, err := message.Pack()
	if err != nil {
		panic(err)
	}
	log.Printf("MENSAGEM RECEBIDA DA BANDEIRA SOLICITACAO DE ATIVACAO:\n")
	log.Printf("% x\n", b)
	iso8583.Describe(message, os.Stdout)

	cod_resposta, err := message.GetString(39)
	if err != nil {
		panic(err)
	}
	if cod_resposta != "" {
		log.Fatalln("ATIVACAO NEGADA PELA BANDEIRA, COD RESPOSTA DIFERENTE DE 00")
	}
	resp_rede, err := message.GetString(70)
	if err != nil {
		panic(err)
	}
	if resp_rede != "81" {
		log.Println("ATIVACAO PODE NAO TER OCORRIDO PELA BANDEIRA, RESPOSTA DIFERENTE DE 81, COD RECEBIDO:", resp_rede)
	}

	return resp_rede, err
}
