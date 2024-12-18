package trabalhaconexoes

import (
	"log"
	"net"
	//	"runtime"
)

// ConexaoSrvBandeira estabelece uma conexão TCP com o servidor da bandeira e retorna a conexão e o erro, se houver
func ConexaoSrvBandeira(portaBandeira string) (net.Conn, error) {
	//goroutineID := runtime.NumGoroutine()
	connBandeira, err := net.Dial("tcp", portaBandeira)
	if err != nil {
		return nil, err // Retorna nil para a conexão e o erro encontrado
	}
	// Mostra o IP conectado
	bandAddr := connBandeira.RemoteAddr().String()

	//log.Println("Conectado ao servidor da bandeira:", bandAddr, "GO Routine:", goroutineID)
	log.Println("Conectado ao servidor da bandeira:", bandAddr)
	return connBandeira, nil // Retorna a conexão e nil indicando sucesso
}
