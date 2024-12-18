package main

import (
	"log"
	"net"
	"os"

	"autotrnsrvmc/trabalhaconexoes"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "autotrnsrvmc/proto/proto"
)

func main() {
	log.Println("Iniciando Servidor MC v1...")

	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Lê o tipo de servidor (gRPC ou TCP) do .env
	tipoServer := os.Getenv("TIPO_SERVER")
	if tipoServer == "" {
		log.Fatalf("TIPO_SERVER não definido no arquivo .env")
	}

	// Inicia o servidor correto com base no parâmetro lido
	switch tipoServer {
	case "grpc":
		log.Println("Iniciando servidor gRPC...")
		grpcPortaEnt := os.Getenv("GRPC_PORTA_ENT")

		lis, err := net.Listen("tcp", grpcPortaEnt)
		if err != nil {
			log.Fatalf("Falha ao iniciar o servidor gRPC: %v", err)
		}

		grpcServer := grpc.NewServer()
		pb.RegisterTransacaoServiceServer(grpcServer, &trabalhaconexoes.EstruturaServer{})

		log.Printf("Servidor gRPC aguardando conexões %s...", grpcPortaEnt)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Falha ao rodar o servidor gRPC: %v", err)
		}
		//runGRPCServer()
	case "tcp":
		log.Println("Iniciando servidor TCP...")
		//runTCPServer()
		// Porta de rede de entrada
		tcpPortaEnt := os.Getenv("TCP_PORTA_ENT")

		// Inicia a escuta na porta definida no .env
		concliente, err := net.Listen("tcp", tcpPortaEnt)
		if err != nil {
			log.Fatalf("Erro ao abrir a porta TCP: %v", err)
		}
		defer concliente.Close()
		log.Printf("Servidor aguardando conexões %s...", tcpPortaEnt)

		// Loop contínuo para aceitar múltiplas conexões
		// O "conexao define a conexao entre cliente aplicacao e servidor
		for {
			conexao, err := concliente.Accept()
			if err != nil {
				log.Printf("Erro ao aceitar conexão: %v", err)
				continue
			}

			// Trata a conexão de forma assíncrona (multithreading)
			//go lidaComConexao(conexao)
			//go trabalhaconexoes.TrabalhaConexaoTCP(conexao)
			go trabalhaconexoes.TrabalhaConexaoTCP(conexao)
		}
	default:
		log.Fatalf("SERVER_TYPE inválido: %s. Use 'grpc' ou 'tcp'.", tipoServer)
	}

}
