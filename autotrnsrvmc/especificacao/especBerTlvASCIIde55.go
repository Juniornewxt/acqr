package especificacao

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Essa funcao tem como objetivo receber uma string com os dados bertlf em ascii e separara os dados conforme dados e estruturava passados na execucao.

// Função que retorna um map de campos e valores, lidando com tags de struct e ignorando campos ausentes
func LeBerTLV55(dados string, estrutura interface{}) (map[string]string, error) {
	valorEstrutura := reflect.ValueOf(estrutura).Elem()
	tipoEstrutura := valorEstrutura.Type()

	result := make(map[string]string)

	for i := 0; i < valorEstrutura.NumField(); i++ {
		// Obter a tag "campo" da struct
		campoTag := tipoEstrutura.Field(i).Tag.Get("isoJUNIOR")
		if campoTag == "" {
			// Se não houver tag "campo", usar o nome do campo como fallback
			campoTag = tipoEstrutura.Field(i).Name
		}

		// Procurar o nome do campo na string
		posCampo := strings.Index(dados, campoTag)
		if posCampo == -1 {
			// Campo não encontrado, ir para o próximo campo
			continue
		}

		// Ajustar a posição de início com base na posição encontrada
		inicio := posCampo + len(campoTag)

		// Encontrar o tamanho do valor do campo
		if inicio+2 > len(dados) {
			return nil, fmt.Errorf("dados insuficientes para ler o tamanho do campo %s", campoTag)
		}

		tamanhoCampoStr := dados[inicio : inicio+2]
		tamanhoCampo, err := strconv.Atoi(tamanhoCampoStr)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter o tamanho do campo %s: %v", campoTag, err)
		}

		// Multiplicar o tamanho por 2 (pois o tamanho é em bytes e cada byte é representado por 2 caracteres hexadecimais)
		tamanhoCampo *= 2

		inicio += 2

		// Extrair o valor do campo
		if inicio+tamanhoCampo > len(dados) {
			return nil, fmt.Errorf("dados insuficientes para ler o valor do campo %s", campoTag)
		}

		valor := dados[inicio : inicio+tamanhoCampo]
		result[campoTag] = valor

		// Atualizar o início para o próximo campo
		inicio += tamanhoCampo
	}

	return result, nil
}
