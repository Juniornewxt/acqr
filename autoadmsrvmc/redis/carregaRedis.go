package redis

import (
	"log"

	"autoadmsrvmc/modelos"
	"autoadmsrvmc/redis/conn"
)

func CarregaRedis(IDm, IdClasseChave, NumIndiceChave, NumCicloChave, ChavePEKcripSenha, ValorVerifChave string) error {
	// Inicializando o cliente Redis
	rdb, ctx := conn.NovoClienteRedis()

	// Dados para inserir no Redis
	i48 := modelos.IsoDe48sub11ret{
		IdClasseChave:     IdClasseChave,
		NumIndiceChave:    NumIndiceChave,
		NumCicloChave:     NumCicloChave,
		ChavePEKcripSenha: ChavePEKcripSenha,
		ValorVerifChave:   ValorVerifChave,
	}

	// Inserir dados no Redis (usando um hash)
	// IDm DM = Chave Dual msn ou credito
	// IDm SM = Chave Single msn ou debito maestro
	err := rdb.HSet(ctx, "trocachave:"+IDm+"", map[string]interface{}{
		"id":     i48.IdClasseChave,
		"indice": i48.NumIndiceChave,
		"ciclo":  i48.NumCicloChave,
		"pek":    i48.ChavePEKcripSenha,
		"verif":  i48.ValorVerifChave,
	}).Err()

	if err != nil {
		log.Println("Erro ao inserir dados no Redis: %v", err)
		return err
	}

	log.Println("Dados inseridos no Redis com sucesso, ID:", IDm)

	return nil
}
