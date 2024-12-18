package trabalhaconexoes

import (
	"log"
	"net"
)

// encaminhaMensagem envia a mensagem para outro servidor e retorna a resposta
func EncaminhaMensagem(mensagem string, connBandeira net.Conn) (string, error) {

	// Tamanho do buffer para leitura de dados
	//const tamanhoBuffer = 1024
	const tamanhoBuffer = 4096

	//buf := make([]byte, 4096)

	// Envia a mensagem para o servidor de destino
	if _, err := connBandeira.Write([]byte(mensagem)); err != nil {
		return "", err
	}

	// Buffer para armazenar a resposta
	recvBuf := make([]byte, tamanhoBuffer)

	// Lê a resposta do servidor de destino
	recebidaFormatada, err := connBandeira.Read(recvBuf)
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

	return retorno, nil
}

// encaminhaMensagem envia a mensagem para outro servidor e retorna a resposta
func EncaminhaMensagem2(mensagem string, connBandeira net.Conn) (string, error) {

	// Tamanho do buffer para leitura de dados
	//const tamanhoBuffer = 1024
	const tamanhoBuffer = 4096

	//buf := make([]byte, 4096)

	// Envia a mensagem para o servidor de destino
	if _, err := connBandeira.Write([]byte(mensagem)); err != nil {
		return "", err
	}

	// Buffer para armazenar a resposta
	recvBuf := make([]byte, tamanhoBuffer)

	// Lê a resposta do servidor de destino
	recebidaFormatada, err := connBandeira.Read(recvBuf)
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

	return retorno, nil
}
