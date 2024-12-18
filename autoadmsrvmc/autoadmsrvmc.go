package main

import (
	//	"autoadmsrvmc/controlemsg"
	"autoadmsrvmc/trabalhaconexoes"
	"log"
	"os"

	//	"autoadmsrvmc/controlemsg"

	//manipulaiso "command-line-arguments/Users/junior/Documents/go/github.com/juniornewxt/projetoacqr/auto/servidor/autoadmsrvmc/manupulaiso/ativaISO0800.go"

	"github.com/joho/godotenv"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
	// Portas do .env
	portaEntrada := os.Getenv("TCPMCADMEIN")
	if portaEntrada == "" {
		log.Fatalf("TCPMCADMEIN não definido no arquivo .env")
	}
	portaBandeira := os.Getenv("TCPMCADMBAN")
	if portaBandeira == "" {
		log.Fatalf("TCPMCADMBAN não definido no arquivo .env")
	}
	log.Println("Iniciando Servidor MC Admninstrativo v1...")
	// Conectar ao servidor da bandeira assim que o programa iniciar e manter a conexão aberta
	connBandeira, err := trabalhaconexoes.ConexaoSrvBandeira(portaBandeira)
	if err != nil {
		log.Fatalf("Erro ao conectar ao servidor externo %s: %v", portaBandeira, err)
	}
	defer connBandeira.Close()

	//log.Printf("Conectado ao servidor externo %s\n", portaBandeira)
	//Envia Sonda solicitando ativaçao
	//msn3, err := controlemsg.MsgAtivacao(connBandeira)
	//Verifica se houve erro
	//if err != nil {
	//	log.Println("Erro:", err)
	//} else {
	//	log.Println("Valor retornado:", msn3)
	//	log.Println("Servidor será parado")
	//	log.Fatalln("Valor retornado:", msn3)
	//}
	// Canal para coordenar mensagens originadas da porta de entrada
	msgPortaEntrada := make(chan []byte)

	// Goroutine para escutar mensagens oriundas da bandeira
	go trabalhaconexoes.MsgOrigemBandeira(connBandeira, msgPortaEntrada)

	// Se conectado a bandeira, inicia o servidor na porta de entrada, aguardando conexoes para encaminhar a porta da bandeira
	trabalhaconexoes.AdmTrnEntpBand(portaEntrada, connBandeira, msgPortaEntrada)
}
