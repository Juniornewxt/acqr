package trabalhaconexoes

import (
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
	//"autotrnsrvmc/gerenciamento"
	//"github.com/joho/godotenv"
)

// Variável para controlar se a mensagem completa será logada
//var logMSNrede bool

// Tamanho do buffer para leitura de dados
const tamanhoBuffer = 1024

// handleConnection lida com uma única conexão de cliente
func TrabalhaConexaoTCP(conexao net.Conn) {
	defer conexao.Close()
	goroutineID := runtime.NumGoroutine()
	// Carrega as variáveis de ambiente do arquivo .env
	//	if err := godotenv.Load(); err != nil {
	//		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	//}

	// Configura se deve logar a mensagem completa
	//logMSNrede = os.Getenv("LOG_MSN_REDE") == "true"

	// Buffer para receber a mensagem
	msnbuffer := make([]byte, tamanhoBuffer)

	// Lê a mensagem recebida
	recebidaFormatada, err := conexao.Read(msnbuffer)
	if err != nil {
		log.Printf("Erro ao ler mensagem: %v", err)
		return
	}

	// Mostra o IP do cliente que se conectou
	entAddr := conexao.RemoteAddr().String()
	log.Println("Mensagem recebida do servidor de entrada:", entAddr, ([]byte(msnbuffer)), "GO Routine:", goroutineID)

	// Log da mensagem recebida
	mensagemRecebida := string(msnbuffer[:recebidaFormatada])

	mensagemDevolver, err := GerenciaMensagens(mensagemRecebida)

	// Verifica se deve logar a mensagem completa
	//	logMSNredeDefinida := os.Getenv("MSN_DEFINIDA")
	//if logMSNrede {
	//		gerenciamento.LogMensagens("Recebido", mensagemRecebida, conexao.RemoteAddr(), conexao.LocalAddr())
	//	} else {
	//		gerenciamento.LogMensagens("Recebido", logMSNredeDefinida, conexao.RemoteAddr(), conexao.LocalAddr())
	//	}

	// Define timeout para leitura do servidor do cliente
	clienteTimeout := os.Getenv("CLIENTTIMOUT")
	clienteTimeoutf, err := strconv.Atoi(clienteTimeout)
	if err != nil {
		log.Println("Falha ao enviar a mensagem, erro ao converter string para int CLIENTTIMOUT:", err)
		return
	}
	// Adiciona time out na conexao com o cliente
	conexao.SetReadDeadline(time.Now().Add(time.Duration(clienteTimeoutf) * time.Second))

	// Envia a resposta ao cliente
	conexao.Write([]byte(mensagemDevolver))

	//if logMSNrede {
	//		gerenciamento.LogMensagens("Recebido", string(mensagemDevolver), conexao.LocalAddr(), conexao.RemoteAddr())
	//	} else {
	//	gerenciamento.LogMensagens("Recebido", logMSNredeDefinida, conexao.LocalAddr(), conexao.RemoteAddr())
	//	}
}
