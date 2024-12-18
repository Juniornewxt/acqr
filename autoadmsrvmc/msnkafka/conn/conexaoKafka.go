package conn

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConectorProdutor() (*kafka.Producer, error) {

	kafka_ip_porta := os.Getenv("KAFKA_SERV")
	kafka_sec_prot := os.Getenv("KAFKA_SECP")
	kafka_sasl := os.Getenv("KAFKA_SASL")
	kafka_usuario := os.Getenv("KAFKA_USER")
	kafka_senha := os.Getenv("KAFKA_PASS")
	// Broker Kafka definido diretamente aqui
	//brokers := "localhost:9092" // IP e porta do Kafka

	// Configura o produtor com o broker embutido
	//p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	//if err != nil {
	//	log.Printf("Erro ao criar o produtor Kafka: %v", err)
	//	return nil, err
	//}
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafka_ip_porta,
		"security.protocol": kafka_sec_prot,
		"sasl.mechanisms":   kafka_sasl,
		"sasl.username":     kafka_usuario,
		"sasl.password":     kafka_senha,
	})
	if err != nil {
		log.Println("Erro ao criar o produtor:", err)
		return nil, err
	}
	return p, nil
}
