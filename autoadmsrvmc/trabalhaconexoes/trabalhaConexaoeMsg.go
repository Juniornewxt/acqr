package trabalhaconexoes

import (
	manipulaiso "autoadmsrvmc/manipulaiso"
	separaiso "autoadmsrvmc/separaiso"
	"io"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
)

var msnKAFKA bool

// Função AdmTrnEntpBand escuta na porta de entrada e encaminha mensagens entre entrada e bandeira
func AdmTrnEntpBand(portaEntrada string, connBandeira net.Conn, msgPortaEntrada chan []byte) {
	goroutineID := runtime.NumGoroutine()
	escutaPorta, err := net.Listen("tcp", portaEntrada)
	if err != nil {
		log.Fatalf("Erro ao iniciar listener na porta %s: %v", portaEntrada, err)
	}
	defer escutaPorta.Close()

	log.Printf("Escutando na porta %s\n", portaEntrada)

	for {
		// Aceita novas conexões na porta de entrada
		connPortaEntrada, err := escutaPorta.Accept()
		if err != nil {
			log.Printf("Erro ao aceitar conexão na porta %s: %v", portaEntrada, err)
			continue
		}
		// Mostra o IP conectado
		entAddr := connPortaEntrada.RemoteAddr().String()
		//log.Println("Mensagem enviada para o servidor de destino:", bandAddr, ([]byte(mensagem)))
		log.Println("Nova conexão recebida na porta de entrada:", entAddr, "GO Routine:", goroutineID)

		// Para cada nova conexão, cria uma goroutine para lidar com a comunicação entre a porta de entrada e a porta da bandeira
		go TratabalhaConexoes(connPortaEntrada, connBandeira, msgPortaEntrada)
		//TratabalhaConexoes(connPortaEntrada, connBandeira, msgPortaEntrada)
	}
}

// Função TratabalhaConexoes lida com a comunicação entre a conexão da porta de entrada e o servidor da bandeira
func TratabalhaConexoes(connPortaEntrada net.Conn, connBandeira net.Conn, msgPortaEntrada chan []byte) {
	goroutineID := runtime.NumGoroutine()
	defer connPortaEntrada.Close() // Garante que a conexão da porta de entrada será fechada após o processamento

	buf := make([]byte, 4096) // Buffer para leitura de dados da conexão da porta de entrada

	for {
		// Mostra o IP conectado
		entAddr := connPortaEntrada.RemoteAddr().String()
		bandAddr := connBandeira.RemoteAddr().String()
		// Definir timeout para a conexão com servidor externo
		mcTimeout := os.Getenv("MCTIMEOUT")
		mcTimeoutf, err := strconv.Atoi(mcTimeout)
		timeout := time.Duration(mcTimeoutf) * time.Second
		// Define timeout para leitura do servidor do cliente
		clienteTimeout := os.Getenv("CLIENTTIMOUT")
		clienteTimeoutf, err := strconv.Atoi(clienteTimeout)
		timeoutclient := time.Duration(clienteTimeoutf) * time.Second
		// Aqui se define um tempo limite para leitura de dados da conexão de entrada (Para nao aguardar vida toda)
		connPortaEntrada.SetReadDeadline(time.Now().Add(timeoutclient))
		n, err := connPortaEntrada.Read(buf) // Le os dados da conexão de entrada
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Println("Timeout lendo da conexão de entrada")
				break
			}
			if err != io.EOF {
				log.Printf("Erro ao ler da conexão de entrada: %v", err)
			}
			break
		}

		if n > 0 {
			msg := buf[:n] // Obtém a mensagem recebida da porta de entrada
			log.Println("Mensagem recebida da porta de entrada:", bandAddr, msg, "GO Routine:", goroutineID)

			// Envia a mensagem recebida da porta de entrada para o servidor da bandeira
			_, err = connBandeira.Write(msg)
			if err != nil {
				log.Printf("Erro ao enviar mensagem para o servidor da bandeira: %v", err, bandAddr)
				break
			}
			log.Println("Mensagem enviada para o servidor da bandeira:", bandAddr, msg, "GO Routine:", goroutineID)

			// Aguardar a resposta do servidor da bandeira (através do canal msgPortaEntrada)
			select {

			case resp := <-msgPortaEntrada:

				//Envio de msg para tratamento da ISO e validar se é o mesmo stan recebido no retorno

				envio_stan, err := separaiso.Rastreio(msg)
				//envio_stan, err := separaiso.Rastreio([]byte(<-msgPortaEntrada))
				if err != nil {
					log.Printf("Erro ao formatar resposta para cliente: %v", err)
					//return
				}
				// Envia para Kafka se habilitado
				msnKAFKA = os.Getenv("KAFKA_ENVIO") == "true"
				if msnKAFKA {
					go manipulaiso.EnviaKafka(msg[2:])
				} else {
					log.Println("MENSAGEM PARA KAFKA NAO ENVIADA ESTA DESATIVADA")
				}
				// realiza uma copia da resp para garantir que nao foi modificado
				//respCopy := make([]byte, len(resp))
				//copy(respCopy, resp)
				//log.Printf("Conteúdo de resp antes de rodar Rastreio: %v", resp)
				resp_stan, err := separaiso.Rastreio(resp)
				if err != nil {
					log.Printf("Erro ao formatar resposta para cliente: %v", err)
					//return
				}
				// Envia para Kafka se habilitado
				msnKAFKA = os.Getenv("KAFKA_ENVIO") == "true"
				if msnKAFKA {
					go manipulaiso.EnviaKafka(resp[2:])
				} else {
					log.Println("MENSAGEM PARA KAFKA NAO ENVIADA ESTA DESATIVADA")
				}
				//log.Printf("Conteúdo de resp antes de comparar stan: %v", resp)
				// Valida se stan é o mesmo do envio
				log.Println("STAN ENVIADO", string(envio_stan))
				log.Println("STAN RECEBIDO", string(resp_stan))
				if string(envio_stan) != (string(resp_stan)) {
					//log.Println("STAN RECEBIDO É DIFERENTE DO ENVIADO")
					//log.Printf("Conteúdo de resp antes admiso: %v", string(resp))
					//log.Printf("COPIA Conteúdo de resp antes admiso: %v", respCopy)
					msn1, err := separaiso.Admiso(resp)
					if err != nil {
						log.Printf("Erro ao processar mensagem da bandeira")
						return
					}
					msn2, err := EncaminhaMensagem(string(msn1), connBandeira)
					if err != nil {
						log.Printf("Erro ao formatar resposta para cliente: %v", err)
						return
					}
					log.Println("ENVIA MENSAGEM PARA TRATAMENTO DE 0800", string(msn2))

				} else {
					log.Println("STAN LOCALIZADO, SEGUE FLUXO DE MSN")
					//}
					// Se a resposta chegar, envia de volta para a conexão de entrada
					_, err = connPortaEntrada.Write(resp)
					if err != nil {
						log.Println("Erro ao enviar resposta para o servidor de entrada:", err, entAddr, "GO Routine:", goroutineID)
					} //
				}
			case <-time.After(timeout):
				// Timeout se não houver resposta do servidor da bandeira
				log.Println("Timeout esperando resposta do servidor da bandeira para o servidor de entrada", entAddr, "GO Routine:", goroutineID)
			}
		}

	}

	log.Println("Fim das mensagens desta conexao, todas as mensagens foram tratadas")
	log.Println("Aguardando novas conexoes/mensagens")
}

// Função que recebe a mensagem do canal
//func receberDoCanal(ch chan []byte) []byte {
//return <-ch
//}
