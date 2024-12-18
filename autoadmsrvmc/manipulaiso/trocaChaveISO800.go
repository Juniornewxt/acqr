package manipulaiso

import (

	//"encoding/csv"
	"bytes"
	"log"
	"os"

	"github.com/moov-io/iso8583"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"

	//"github.com/moov-io/iso8583/specs"
	especificacao "autoadmsrvmc/especificacao"
	"autoadmsrvmc/modelos"
	"autoadmsrvmc/redis"

	"github.com/moov-io/iso8583/network"
)

var msnKAFKA bool
var msnREDIS bool

func TrocaChave(iso []byte) ([]byte, error) {

	// Remove os primeiros 2 bytes
	iso_mod := iso[2:]
	// Configura se deve enviar a mensagem para kafka
	msnKAFKA = os.Getenv("KAFKA_ENVIO") == "true"
	if msnKAFKA {
		go EnviaKafka(iso_mod)
	} else {
		log.Println("MENSAGEM PARA KAFKA NAO ENVIADA ESTA DESATIVADA")
	}

	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_mod)

	p_spec := especificacao.NewSpecEBCDIC()

	p_message := iso8583.NewMessage(p_spec)

	p_message.Unpack([]byte((retorno)))

	msnREDIS = os.Getenv("REDIS_ENVIO") == "true"
	if msnREDIS {
		// Criar uma instância da estrutura
		var isoData48 modelos.IsoDe48ret
		// Criar uma instância da estrutura
		var isoDatasub48 modelos.IsoDe48sub11ret
		// Preenche a estrutura com os dados TLV

		iso48, err := p_message.GetString(48)
		if err != nil {
			panic(err)
		}
		// Envia para funçao que lida com dados no formato TLV para manipulacao
		tlvManipulado, err := especificacao.LeTLVASCIIt2(iso48)
		if err != nil {
			log.Println("Erro tlvManipulado:", err)
			//return
		}

		err = especificacao.EstruturaTLV(tlvManipulado, &isoData48)
		if err != nil {
			log.Println("Erro EstruturaTLV:", err)
			//return
		}
		// Preenche a estrutura os dados enviado somemte o campo que queremos
		err = especificacao.EstruturaDados(isoData48.De48sub11, &isoDatasub48)
		if err != nil {
			log.Println("Erro EstruturaDados:", err)
			//return
		}
		//Envia chaves para o redis
		err = redis.CarregaRedis("DM", isoDatasub48.IdClasseChave, isoDatasub48.NumIndiceChave, isoDatasub48.NumCicloChave, isoDatasub48.ChavePEKcripSenha, isoDatasub48.ValorVerifChave)
		if err != nil {
			log.Println("Erro Ao carregar dados no Reids:", err)
			var errotrocachava []byte
			errotrocachava, err = TrocaChaveErroDM([]byte(retorno))
			if err != nil {
				log.Println("Erro TrocaChaveErroDM:", err)

			}
			return errotrocachava, err
		}
	} else {
		log.Println("MENSAGEM PARA REDIS NAO ENVIADA ESTA DESATIVADA")
	}
	//Fim de envio de msn para redis
	b, err := p_message.Pack()
	if err != nil {
		panic(err)
	}
	//fmt.Println("CHAVE CRIP", isoDatasub48.ChavePEKcripSenha)
	log.Printf("MENSAGEM RECEBIDA DA BANDEIRA:\n")
	log.Printf("% x\n", b)
	iso8583.Describe(p_message, os.Stdout)

	new_spec := especificacao.NewSpecEBCDIC()

	pnew_message := iso8583.NewMessage(new_spec)

	p_message.Unpack(b) //Abre a ISO formatada no inicio para ser usada somente de alguns campos

	pnew_message.MTI("0810")
	pan, err := p_message.GetString(2)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(2, pan)

	mmddhhmmss, err := p_message.GetString(7)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(7, mmddhhmmss) //data e hora mmddhhmmss
	sy_trace, err := p_message.GetString(11)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(11, sy_trace)

	fwd_inst, err := p_message.GetString(33)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(33, fwd_inst)

	//fwd_inst, err := p_message.GetString(39)
	////if err != nil {
	//	panic(err)
	//}
	pnew_message.Field(39, "00") // Troca aceita
	///
	de48, err := p_message.GetString(48)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(48, de48)
	rede, err := p_message.GetString(63)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(63, rede)
	////////////////////////////////////////////////////////////////
	rede_inf, err := p_message.GetString(70)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(70, rede_inf)

	d, err := pnew_message.Pack()
	if err != nil {
		panic(err)

	}

	packed := d

	header := network.NewBinary2BytesHeader()

	header.SetLength(len(packed))

	var buf bytes.Buffer

	header.WriteTo(&buf)

	_, err = buf.Write(packed)

	log.Printf("MENSAGEM A SER ENVIADA GERENCIAMENTO DE REDE DUAL MSG:\n")
	iso8583.Describe(pnew_message, os.Stdout)
	log.Printf("FIM DO PROCESSAMENTO :) \n")

	// Configura se deve enviar a mensagem para kafka
	msnKAFKA = os.Getenv("KAFKA_ENVIO") == "true"
	if msnKAFKA {
		go EnviaKafka(buf.Bytes()[2:])
	} else {
		log.Println("MENSAGEM PARA KAFKA NAO ENVIADA ESTA DESATIVADA")
	}

	return buf.Bytes(), err
}

// }
func TrocaChaveMaestro(iso []byte) ([]byte, error) {

	// Remove os primeiros 2 bytes
	iso_mod := iso[2:]
	// Configura se deve enviar a mensagem para kafka
	msnKAFKA = os.Getenv("KAFKA_ENVIO") == "true"
	if msnKAFKA {
		go EnviaKafka(iso_mod)
	} else {
		log.Println("MENSAGEM PARA KAFKA NAO ENVIADA ESTA DESATIVADA")
	}

	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_mod)

	p_spec := especificacao.NewSpecEBCDIC()

	p_message := iso8583.NewMessage(p_spec)

	p_message.Unpack([]byte((retorno)))

	msnREDIS = os.Getenv("KAFKA_ENVIO") == "true"
	if msnREDIS {
		// Criar uma instância da estrutura
		var isoData48 modelos.IsoDe48ret
		// Criar uma instância da estrutura
		var isoDatasub48 modelos.IsoDe48sub11ret
		// Preenche a estrutura com os dados TLV

		iso48, err := p_message.GetString(48)
		if err != nil {
			panic(err)
		}
		// Envia para funçao que lida com dados no formato TLV para manipulacao
		tlvManipulado, err := especificacao.LeTLVASCIIt2(iso48)
		if err != nil {
			log.Println("Erro tlvManipulado:", err)
			//return
		}

		err = especificacao.EstruturaTLV(tlvManipulado, &isoData48)
		if err != nil {
			log.Println("Erro EstruturaTLV:", err)
			//return
		}
		// Preenche a estrutura os dados enviado somemte o campo que queremos
		err = especificacao.EstruturaDados(isoData48.De48sub11, &isoDatasub48)
		if err != nil {
			log.Println("Erro EstruturaDados:", err)
			//return
		}
		//Envia chaves para o redis
		err = redis.CarregaRedis("SM", isoDatasub48.IdClasseChave, isoDatasub48.NumIndiceChave, isoDatasub48.NumCicloChave, isoDatasub48.ChavePEKcripSenha, isoDatasub48.ValorVerifChave)
		if err != nil {
			log.Println("Erro Ao carregar dados no Redis:", err)
			var errotrocachava []byte
			errotrocachava, err = TrocaChaveErroSM([]byte(retorno))
			if err != nil {
				log.Println("Erro TrocaChaveErroSM:", err)

			}
			return errotrocachava, err
		}

	} else {
		log.Println("MENSAGEM PARA REDIS NAO ENVIADA ESTA DESATIVADA")
	}

	b, err := p_message.Pack()
	if err != nil {
		panic(err)
	}
	log.Printf("MENSAGEM RECEBIDA DA BANDEIRA:\n")
	log.Printf("% x\n", b)
	iso8583.Describe(p_message, os.Stdout)

	// Inicia a criação da spec nova para que não se repita os dados da ISO original, salve quando solicitado p_message
	new_spec := especificacao.NewSpecEBCDIC()

	pnew_message := iso8583.NewMessage(new_spec)

	p_message.Unpack(b) //Abre a ISO formatada no inicio para ser usada somente de alguns campos

	pnew_message.MTI("0810")

	mmddhhmmss, err := p_message.GetString(7)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(7, mmddhhmmss) //data e hora mmddhhmmss
	sy_trace, err := p_message.GetString(11)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(11, sy_trace)

	fwd_inst, err := p_message.GetString(33)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(33, fwd_inst)

	//fwd_inst, err := p_message.GetString(39)
	////if err != nil {
	//	panic(err)
	//}
	pnew_message.Field(39, "00") // Troca aceita
	///
	rede, err := p_message.GetString(63)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(63, rede)
	////////////////////////////////////////////////////////////////
	rede_inf, err := p_message.GetString(70)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(70, rede_inf)
	////////////////////////////////////////////////////////////////
	dadosprivados, err := p_message.GetString(126)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(126, dadosprivados)

	d, err := pnew_message.Pack()
	if err != nil {
		panic(err)

	}

	packed := d

	header := network.NewBinary2BytesHeader()

	header.SetLength(len(packed))

	var buf bytes.Buffer

	header.WriteTo(&buf)

	_, err = buf.Write(packed)

	log.Printf("MENSAGEM A SER ENVIADA GERENCIAMENTO DE REDE MAESTRO:\n")
	iso8583.Describe(pnew_message, os.Stdout)
	log.Printf("FIM DO PROCESSAMENTO :) \n")

	// Configura se deve enviar a mensagem para kafka
	msnKAFKA = os.Getenv("KAFKA_ENVIO") == "true"
	if msnKAFKA {
		go EnviaKafka(buf.Bytes()[2:])
	} else {
		log.Println("MENSAGEM PARA KAFKA NAO ENVIADA ESTA DESATIVADA")
	}

	return buf.Bytes(), err
}

func RetornoADM(iso []byte) ([]byte, error) {

	// Remove os primeiros 2 bytes
	iso_mod := iso[2:]

	// Configura se deve enviar a mensagem para kafka
	msnKAFKA = os.Getenv("KAFKA_ENVIO") == "true"
	if msnKAFKA {
		go EnviaKafka(iso_mod)
	} else {
		log.Println("MENSAGEM PARA KAFKA NAO ENVIADA ESTA DESATIVADA")
	}
	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_mod)

	p_spec := especificacao.NewSpecEBCDIC()

	p_message := iso8583.NewMessage(p_spec)

	p_message.Unpack([]byte((retorno)))

	b, err := p_message.Pack()
	if err != nil {
		panic(err)
	}
	log.Printf("MENSAGEM RECEBIDA DA BANDEIRA:\n")
	log.Printf("% x\n", b)
	iso8583.Describe(p_message, os.Stdout)
	log.Printf("FIM DO PROCESSAMENTO :) \n")
	log.Printf("Aguardando novas mensagens se ouver... \n")

	return nil, err
}
func TrocaChaveErroSM(iso []byte) ([]byte, error) {

	p_spec := especificacao.NewSpecEBCDIC()

	p_message := iso8583.NewMessage(p_spec)

	//p_message.Unpack([]byte((iso)))

	// Inicia a criação da spec nova para que não se repita os dados da ISO original, salve quando solicitado p_message
	new_spec := especificacao.NewSpecEBCDIC()

	pnew_message := iso8583.NewMessage(new_spec)

	p_message.Unpack(iso) //Abre a ISO formatada no inicio para ser usada somente de alguns campos

	pnew_message.MTI("0810")

	mmddhhmmss, err := p_message.GetString(7)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(7, mmddhhmmss) //data e hora mmddhhmmss
	sy_trace, err := p_message.GetString(11)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(11, sy_trace)

	fwd_inst, err := p_message.GetString(33)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(33, fwd_inst)

	//fwd_inst, err := p_message.GetString(39)
	////if err != nil {
	//	panic(err)
	//}
	pnew_message.Field(39, "96") // Troca aceita
	///
	rede, err := p_message.GetString(63)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(63, rede)
	////////////////////////////////////////////////////////////////
	rede_inf, err := p_message.GetString(70)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(70, rede_inf)
	////////////////////////////////////////////////////////////////
	dadosprivados, err := p_message.GetString(126)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(126, dadosprivados)

	d, err := pnew_message.Pack()
	if err != nil {
		panic(err)

	}

	packed := d

	header := network.NewBinary2BytesHeader()

	header.SetLength(len(packed))

	var buf bytes.Buffer

	header.WriteTo(&buf)

	_, err = buf.Write(packed)

	log.Printf("MENSAGEM A SER ENVIADA GERENCIAMENTO DE REDE MAESTRO:\n")
	iso8583.Describe(pnew_message, os.Stdout)
	log.Printf("FIM DO PROCESSAMENTO :) \n")

	// Configura se deve enviar a mensagem para kafka
	msnKAFKA = os.Getenv("KAFKA_ENVIO") == "true"
	if msnKAFKA {
		go EnviaKafka(buf.Bytes()[2:])
	} else {
		log.Println("MENSAGEM PARA KAFKA NAO ENVIADA ESTA DESATIVADA")
	}

	return buf.Bytes(), err
}
func TrocaChaveErroDM(iso []byte) ([]byte, error) {

	p_spec := especificacao.NewSpecEBCDIC()

	p_message := iso8583.NewMessage(p_spec)

	//p_message.Unpack([]byte((iso)))

	// Inicia a criação da spec nova para que não se repita os dados da ISO original, salve quando solicitado p_message
	new_spec := especificacao.NewSpecEBCDIC()

	pnew_message := iso8583.NewMessage(new_spec)

	p_message.Unpack(iso) //Abre a ISO formatada no inicio para ser usada somente de alguns campos

	pnew_message.MTI("0810")
	pan, err := p_message.GetString(2)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(2, pan)

	mmddhhmmss, err := p_message.GetString(7)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(7, mmddhhmmss) //data e hora mmddhhmmss
	sy_trace, err := p_message.GetString(11)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(11, sy_trace)

	fwd_inst, err := p_message.GetString(33)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(33, fwd_inst)

	//fwd_inst, err := p_message.GetString(39)
	////if err != nil {
	//	panic(err)
	//}
	pnew_message.Field(39, "96") // Troca aceita
	///
	de48, err := p_message.GetString(48)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(48, de48)
	rede, err := p_message.GetString(63)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(63, rede)
	////////////////////////////////////////////////////////////////
	rede_inf, err := p_message.GetString(70)
	if err != nil {
		panic(err)
	}
	pnew_message.Field(70, rede_inf)

	d, err := pnew_message.Pack()
	if err != nil {
		panic(err)

	}

	packed := d

	header := network.NewBinary2BytesHeader()

	header.SetLength(len(packed))

	var buf bytes.Buffer

	header.WriteTo(&buf)

	_, err = buf.Write(packed)

	log.Printf("MENSAGEM A SER ENVIADA GERENCIAMENTO DE REDE DUAL MSG:\n")
	iso8583.Describe(pnew_message, os.Stdout)
	log.Printf("FIM DO PROCESSAMENTO :) \n")

	// Configura se deve enviar a mensagem para kafka
	msnKAFKA = os.Getenv("KAFKA_ENVIO") == "true"
	if msnKAFKA {
		go EnviaKafka(buf.Bytes()[2:])
	} else {
		log.Println("MENSAGEM PARA KAFKA NAO ENVIADA ESTA DESATIVADA")
	}

	//go EnviaKafka(buf.Bytes()[2:])

	return buf.Bytes(), err
}
