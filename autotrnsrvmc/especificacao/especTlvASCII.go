package especificacao

import (
	"fmt"
	"strconv"
	//"strings"
)

// LeTLVASCII interpreta uma string TLV no formato ASCII e retorna um mapa de tags para valores.
func LeTLVASCII(tlv string) (map[string]string, error) {
	resultado := make(map[string]string)
	for len(tlv) > 0 {
		// Lê a Tag (3 primeiros caracteres)
		if len(tlv) < 3 {
			return nil, fmt.Errorf("dados insuficientes para ler a tag")
		}
		tag := tlv[:3]
		tlv = tlv[3:]

		// Lê o Comprimento (3 caracteres)
		if len(tlv) < 3 {
			return nil, fmt.Errorf("dados insuficientes para ler o comprimento")
		}
		tamanhoString := tlv[:3]
		tamanho, err := strconv.Atoi(tamanhoString)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter comprimento: %v", err)
		}
		tlv = tlv[3:]

		// Lê o Valor de acordo com o comprimento especificado
		if len(tlv) < tamanho {
			return nil, fmt.Errorf("dados insuficientes para ler o valor da tag %s", tag)
		}
		valor := tlv[:tamanho]
		tlv = tlv[tamanho:]

		// Adiciona o campo ao resultado
		resultado[tag] = valor
	}
	return resultado, nil
}
