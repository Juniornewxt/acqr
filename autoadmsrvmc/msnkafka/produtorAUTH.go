package msnkafka

import (
	"autoadmsrvmc/modelos"
	"autoadmsrvmc/msnkafka/conn"
	"encoding/json"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func EnviaFilaAuth(mti, pan, proccod, valor, datahoratransacao, stan, hhmmss, mcc, entrymode, adquirente, codinstiremetente, nsu, codauto, codresp, terminal, codcomercio, nomeendereco, moeda, dadosaddparcelado string) error {

	kafka_topic_trn := os.Getenv("TOPIC_TRN")
	// Criar produtor usando a função externa, sem passar brokers
	p, err := conn.ConectorProdutor()
	if err != nil {
		log.Println("Erro ao criar o produtor:", err)

	}
	defer p.Close()

	// Criando a mensagem
	//mtirecebido, pan, dmhhmmss, stan, fiic, codresp, addinfo, netdata, codrede, privatedata
	msg := modelos.MensagemAuth{
		Mti:               mti,
		Pan:               pan,
		ProcCod:           proccod,
		Valor:             valor,
		DataHoraTransacao: datahoratransacao,
		STAN:              stan,
		Hhmmss:            hhmmss,
		Mcc:               mcc,
		EntryMode:         entrymode,
		Adquirente:        adquirente,
		CodInstiRemetente: codinstiremetente,
		Nsu:               nsu,
		CodAuto:           codauto,
		Codresp:           codresp,
		Terminal:          terminal,
		CodComercio:       codcomercio,
		NomeEndereco:      nomeendereco,
		Moeda:             moeda,
		DadosADDParcelado: dadosaddparcelado,
	}

	// Convertendo a mensagem para JSON
	msgJSON, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	topic := kafka_topic_trn

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
		//log.Printf("Mensagem enviada para %v [%v] em %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset, string(m.Value))
		log.Printf("Kafka Mensagem enviada para %v [%v] em %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
	return nil
}
