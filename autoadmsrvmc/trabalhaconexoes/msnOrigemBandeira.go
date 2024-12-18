package trabalhaconexoes

import (
	"autoadmsrvmc/separaiso"

	"io"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
)

// Função para escutar mensagens não solicitadas que tiverem origem da bandeira
func MsgOrigemBandeira(connBandeira net.Conn, msgPortaEntrada chan []byte) {
	buf := make([]byte, 4096) // Buffer para armazenar as mensagens recebidas
	for {
		// Definir timeout para a conexão com servidor externo
		mcTimeout := os.Getenv("MCTIMEOUT")
		mcTimeoutf, err := strconv.Atoi(mcTimeout)
		timeout := time.Duration(mcTimeoutf) * time.Second
		if err != nil {
			log.Println("Falha ao enviar a mensagem, erro ao converter string para int MCTIMEOUT:", err)
		}

		goroutineID := runtime.NumGoroutine()
		// Mostra o IP conectado
		bandAddr := connBandeira.RemoteAddr().String()
		//log.Println("Mensagem enviada para o servidor de destino:", bandAddr, ([]byte(mensagem)))
		// Define um tempo limite para leitura de mensagens da porta da bandeira
		connBandeira.SetReadDeadline(time.Now().Add((timeout)))
		n, err := connBandeira.Read(buf) // Aqui Tenta ler dados da conexão com a bandeira
		if err != nil {
			if err == io.EOF {
				log.Println("Conexão com o servidor da Bandeira fechada.")
				log.Fatalf("Perda de conexao com o servidor externo")
				return
			}
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue // Ignora timeout e tenta ler novamente
			}
			log.Printf("Erro ao ler mensagem da Bandeira: %v", err)
			continue
		}

		if n > 0 {
			msg := buf[:n] // Obtém a mensagem recebida
			//controlemsg.MsgRecebidaBand(msg)

			// Verifica se há uma mensagem no canal msgPortaEntrada (Mensagens recebidas pela porta de entrada do servidor)
			select {
			case msgPortaEntrada <- msg:
				// Se houver uma mensagem originada da porta de entreda, trata como uma resposta e envia para a porta de entreda
				log.Println("Resposta recebida do servidor da Bandeira para a porta de entrada:", bandAddr, msg, "GO Routine:", goroutineID)
			default:
				// Se não houver mensagens aguardando de a porta de entrada, trata como mensagem não solicitada recebida da Bandeira
				log.Println("Mensagem não solicitada recebida do servidor da Bandeira:", bandAddr, msg, "GO Routine:", goroutineID)
				msn1, err := separaiso.Admiso([]byte(msg))
				if err != nil {
					log.Printf("Erro ao processar mensagem da bandeira")
					//return
				}
				msn2, err := EncaminhaMensagem(string(msn1), connBandeira)
				if err != nil {
					log.Printf("Erro ao formatar resposta para cliente: %v", err)
					//return
				}
				// Trata o retorno da bandeira
				_, err = separaiso.Admiso([]byte(msn2))
				if err != nil {
					log.Printf("Erro ao processar mensagem da bandeira")
					//return
				}
			}
		}
	}
}
