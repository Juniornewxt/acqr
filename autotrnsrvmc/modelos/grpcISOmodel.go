package modelos

import (
	pb "autotrnsrvmc/proto/proto"
)

type GrpcIsoM struct {
	DE11 string
	DE38 string
	DE39 string
	DE41 string
	DE43 string
}
type Server struct {
	pb.UnimplementedTransacaoServiceServer
}
