package msnkafka

import (
	"autoadmsrvmc/modelos"
	"autoadmsrvmc/msnkafka/conn"
	"encoding/json"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func EnviaFila0800(mti, pan, dmhhmmss, stan, fiic, codresp, addinfo, netdata, codrede, privatedata string) error {

	kafka_topic_adm := os.Getenv("TOPIC_ADM")
	// Criar produtor usando a função externa, sem passar brokers
	p, err := conn.ConectorProdutor()
	if err != nil {
		log.Println("Erro ao criar o produtor:", err)

	}
	defer p.Close()

	// Criando a mensagem
	//mtirecebido, pan, dmhhmmss, stan, fiic, codresp, addinfo, netdata, codrede, privatedata
	msg := modelos.MensagemAdm{
		Mti:         mti,
		Pan:         pan,
		Dmhhmmss:    dmhhmmss,
		Stan:        stan,
		Fiic:        fiic,
		Codresp:     codresp,
		Addinfo:     addinfo,
		Netdata:     netdata,
		Codrede:     codrede,
		Privatedata: privatedata,
	}

	// Convertendo a mensagem para JSON
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	topic := kafka_topic_adm

	// Enviar a mensagem para o Kafka
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msgJSON,
	}, nil)

	if err != nil {
		log.Fatal(err)
	}

	// Confirmação do envio
	e := <-p.Events()
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		log.Printf("Kafka Erro ao produzir: %v\n", m.TopicPartition.Error)
		return err
	} else {
		log.Printf("Kafka Mensagem enviada para %v [%v] em %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	return nil
}
