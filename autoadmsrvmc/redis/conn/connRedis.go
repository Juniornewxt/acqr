package conn

import (
	"context"
	"log"

	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// Cria um contexto global
//var ctx = context.Background()

// Função para criar a conexão com o Redis
func NovoClienteRedis() (*redis.Client, context.Context) {

	redis_ip_porta := os.Getenv("REDIS_SERV")
	redis_usuario := os.Getenv("REDIS_USER")
	redis_password := os.Getenv("REDIS_PASS")
	redis_db := os.Getenv("REDIS_DB")

	db_int, err := strconv.Atoi(redis_db)
	if err != nil {
		log.Println("Erro ao converter string para int:", err)

	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_ip_porta, // Endereço do Redis
		Username: redis_usuario,  // Nome de usuário ACL
		Password: redis_password, // Senha do usuário ACL
		DB:       db_int,         // Use o banco de dados padrão
	})
	ctx := context.Background() // Cria um novo contexto
	return rdb, ctx
}
