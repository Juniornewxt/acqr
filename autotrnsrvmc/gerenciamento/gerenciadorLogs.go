package gerenciamento

import (
	"log"
	"net"
	"runtime"
	"time"
)

// Função para logar detalhes da mensagem, thread e IPs
func LogMensagens(direcao, message string, entIP, dstIP net.Addr) {
	// Captura o ID da goroutine (thread)
	goroutineID := runtime.NumGoroutine()

	// Gera a data e hora atual
	horaAtual := time.Now().Format("2006-01-02 15:04:05")

	// Log detalhado da mensagem, incluindo os IPs de origem e destino
	log.Printf("[%s] [Goroutine: %d] [%s] [Origem: %s] [Destino: %s] - %s\n",
		horaAtual, goroutineID, direcao, entIP.String(), dstIP.String(), message)
}
