package trabalhaconexoes

import (
	//"autotrnsrvmc/gerenciamento"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
)

// Variável para controlar se a mensagem completa será logada
//var logMSNrede bool

// encaminhaMensagem envia a mensagem para outro servidor e retorna a resposta
func EncaminhaMensagem(mensagem, endereco string, srcAddr net.Addr) (string, error) {
	goroutineID := runtime.NumGoroutine()
	// Definir timeout para a conexão com servidor externo
	mcTimeout := os.Getenv("MCTIMEOUT")
	mcTimeoutf, err := strconv.Atoi(mcTimeout)
	if err != nil {
		log.Println("Falha ao enviar a mensagem, erro ao converter string para int MCTIMEOUT:", err)
	}
	timeout := time.Duration(mcTimeoutf) * time.Second
	//timeout := 1 * time.Second
	dialer := net.Dialer{
		Timeout: timeout, // Timeout de conexão
	}

	// Estabelece a conexão TCP com o servidor de destino com timeout de conexão
	conn, err := dialer.Dial("tcp", endereco)
	if err != nil {
		log.Printf("Erro ao conectar ao servidor de destino: %v", err)
		return "", err
	}
	defer conn.Close()

	// Aplica o timeout para leitura e escrita (ambos)
	conn.SetDeadline(time.Now().Add(timeout))

	// Mostra o IP do cliente que se conectou
	bandAddr := conn.RemoteAddr().String()
	log.Println("Mensagem enviada para o servidor de destino:", bandAddr, ([]byte(mensagem)), "GO Routine:", goroutineID)

	// Envia a mensagem para o servidor de destino
	if _, err := conn.Write([]byte(mensagem)); err != nil {
		return "", err
	}

	// Buffer para armazenar a resposta
	recvBuf := make([]byte, tamanhoBuffer)

	// Lê a resposta do servidor de destino
	recebidaFormatada, err := conn.Read(recvBuf)
	if err != nil {
		// Verifica se o erro foi causado por um timeout
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			log.Printf("Timeout ao ler a resposta do servidor de destino.")
			return "", err
		}
		log.Printf("Erro ao ler a resposta do servidor de destino: %v", err)
		return "", err
	}

	retorno := string(recvBuf[:recebidaFormatada])

	// Mostra o IP do cliente que se conectou
	bandRAddr := conn.RemoteAddr().String()
	log.Println("Mensagem de retorno recebida do servidor de destino:", bandRAddr, ([]byte(retorno)), "GO Routine:", goroutineID)

	// Verifica se deve logar a mensagem completa
	//logMSNredeDefinida := os.Getenv("MSN_DEFINIDA")
	//if logMSNrede {
	//	gerenciamento.LogMensagens("Recebido", retorno, conn.RemoteAddr(), srcAddr)
	//} else {
	//	gerenciamento.LogMensagens("Recebido", logMSNredeDefinida, conn.RemoteAddr(), srcAddr)
	//	}

	return retorno, nil
}

// encaminhaMsnSemRet apenas envia a mensagem, sem esperar retorno
func EncaminhaMsnSemRet(mensagem, endereco string, srcAddr net.Addr) error {
	goroutineID := runtime.NumGoroutine()
	tcpAddr, err := net.ResolveTCPAddr("tcp", endereco)
	if err != nil {
		log.Printf("Erro ao resolver endereço: %v", err)
		return err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Printf("Erro ao conectar ao servidor de destino: %v", err)
		return err
	}
	defer conn.Close()

	// Mostra o IP do cliente que se conectou
	bandAddr := conn.RemoteAddr().String()

	// Envia a mensagem
	if _, err := conn.Write([]byte(mensagem)); err != nil {
		return err
	}
	log.Println("Mensagem de confirmacao enviada para o servidor de destino:", bandAddr, ([]byte(mensagem)), "GO Routine:", goroutineID)
	// Log da mensagem enviada
	//gerenciamento.LogMensagens("Enviado", mensagem[:8], srcAddr, conn.RemoteAddr())
	// Verifica se deve logar a mensagem completa
	//logMSNredeDefinida := os.Getenv("MSN_DEFINIDA")
	//if logMSNrede {
	//	gerenciamento.LogMensagens("Recebido", mensagem, srcAddr, conn.RemoteAddr())
	//} else {
	//	gerenciamento.LogMensagens("Recebido", logMSNredeDefinida, srcAddr, conn.RemoteAddr())
	//}

	// Não aguarda resposta
	return nil
}
