package trabalhaconexoes

import (
	"autotrnsrvmc/grpciso"
	"autotrnsrvmc/modelos"
	pb "autotrnsrvmc/proto/proto"
	"context"
	"log"
)

// Criar uma instância da estrutura
type EstruturaServer modelos.Server

func (s *EstruturaServer) ProcessTransacao(ctx context.Context, req *pb.TransacaoRequest) (*pb.TransacaoResponse, error) {
	if req.MTI == "" {
		log.Println("o campo 'MTI' é obrigatório")
		//return nil, errors.New("o campo 'MTI' é obrigatório")
		return &pb.TransacaoResponse{
			ResponseCode: "91", // Código de sucesso
			Message:      "O campo 'MTI' é obrigatório"}, nil
	}
	if req.DE11 == "" {
		log.Println("o campo 'Trace' é obrigatório")
		//return nil, errors.New("o campo 'Trace' é obrigatório")
		return &pb.TransacaoResponse{
			ResponseCode: "91", // Código de sucesso
			Message:      "O campo 'Trace' é obrigatório"}, nil
	}

	//Envia dados recebidos no gRPC para a funcao responsavel por converter em ISO
	grpcpiso, err := grpciso.GrpcISO(req.MTI, req.DE02, req.DE03, req.DE04, req.DE07, req.DE11, req.DE12, req.DE13, req.DE14, req.DE18, req.DE19, req.DE22, req.DE23, req.DE32, req.DE33, req.DE35, req.DE38, req.DE39, req.DE41, req.DE42, req.DE43, req.DE48, req.DE49, req.DE52, req.DE55, req.DE61, req.DE120, req.DE126)
	if err != nil {
		log.Printf("Erro ao ler mensagem: %v", err)
		return nil, err
	}
	// Envia que recebeu do construtor de ISO para o gerenciador de mensagens
	grpcpisoretorno, err := GerenciaMensagens(grpcpiso)
	if err != nil {
		log.Printf("Erro ao ler mensagem: %v", err)
		return nil, err
	}
	// Envia que recebeu do gerenciador de mensagens para o conversor de ISO em json
	isogrpc, err := grpciso.IsoGRPC(grpcpisoretorno)
	if err != nil {
		log.Printf("Erro ao ler mensagem: %v", err)
		return nil, err
	}
	// Valida se a transacao foi autorizada ou nao
	var DE39 string
	if isogrpc.DE39 == "" {
		DE39 = "00"
	} else {
		DE39 = isogrpc.DE39
	}
	var MSN_ERRO string
	if isogrpc.DE39 != "" {
		MSN_ERRO = "Transação nao autorizada, falha no sistema"
	} else {
		MSN_ERRO = "Transação processada com sucesso"
	}
	// Retorna o resultado ao gRPC
	return &pb.TransacaoResponse{
		//ResponseCode: "00", // Código de sucesso
		ResponseCode: DE39, // Código de sucesso
		///De11: isogrpc.DE11,
		//Message: "Transação processada com sucesso",
		Message: MSN_ERRO,
	}, nil
}
