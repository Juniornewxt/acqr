package trabalhaconexoes

import (
	"log"
	"os"

	"autotrnsrvmc/manipulaiso"
	"autotrnsrvmc/separaiso"
)

// Variável para controlar se a mensagem de confirmacao é enviada
var msnMTI180 bool

// TrabalhaConexaoProcessaDados lida com a lógica de processar mensagens, independente da conexão
func GerenciaMensagens(mensagemRecebida string) (string, error) {

	// Configura se deve enviar a mensagem de confirmacao mti180
	msnMTI180 = os.Getenv("ISO_CONFIRM") == "true"

	// Processamento da mensagem
	msn, err := separaiso.Mciso([]byte(mensagemRecebida))
	if err != nil {
		log.Printf("Erro ao processar mensagem ISO MC: %v", err)
		return "", err
	}

	// Envia a mensagem para outro servidor e obtém a resposta
	tcpHPortaMC := os.Getenv("TCP_HP_MC")
	resposta, err := EncaminhaMensagem(string(msn), tcpHPortaMC, nil)
	if err != nil {
		log.Printf("Erro ao encaminhar mensagem: %v", err)
		// Prepara a resposta final para o cliente
		log.Printf("Formatando mensagem de erro para servidor de entrada ...")
		msn3, err := separaiso.Padraoiso([]byte(mensagemRecebida))
		if err != nil {
			log.Printf("Erro ao formatar resposta para cliente: %v", err)
			return "", err
		}

		// Tentar enviar mensagem de desfazimento para o outro servidor (bandeira)
		// Envia a mensagem de desfazimento de forma assíncrona
		go func() {
			log.Printf("Enviando mensagem de desfazimento para o servidor de destino ...")
			msn4, err := manipulaiso.Mciso400rev([]byte(mensagemRecebida))
			if err != nil {
				log.Printf("Erro ao processar mensagem ISO MC: %v", err)
				return
			}
			msn5, err := EncaminhaMensagem(string(msn4), tcpHPortaMC, nil)
			if err != nil {
				log.Printf("Erro ao encaminhar mensagem, deve ser enviado para fila: %v", err)
				return
			}
			_, err = separaiso.Mciso([]byte(msn5))
			if err != nil {
				log.Printf("Erro ao formatar resposta para cliente: %v", err)
				return
			}
		}()
		// fim tratamento de reversao

		return string(msn3), nil

	}
	// Configura se deve enviar a mensagem de confirmacao mti180
	msnMTI180 = os.Getenv("ISO_CONFIRM") == "true"
	// Inicia processo de envio ou nao de mensagem de confirmacao
	if msnMTI180 {
		// Processa a confirmação da mensagem de autorização
		msnConfirm, err := separaiso.MCoutrosiso([]byte(resposta))
		if err != nil {
			log.Printf("Erro ao processar confirmação MC: %v", err)
			//return
		}
		// Envia confirmação sem esperar resposta
		if err := EncaminhaMsnSemRet(string(msnConfirm), tcpHPortaMC, nil); err != nil {
			log.Printf("Erro ao enviar confirmação: %v", err)
		}
	} else {
		log.Println("MENSAGEM DE CONFIRMACAO NAO ENVIADA ESTA DESATIVADA")
	}

	// Prepara a resposta final
	msn2, err := separaiso.Padraoiso([]byte(resposta))
	if err != nil {
		log.Printf("Erro ao formatar resposta para cliente: %v", err)
		return "", err
	}
	return string(msn2), nil
}
