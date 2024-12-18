package especificacao

import (
	"fmt"
	"reflect"
	"strconv"
)

// Ajustar a função para lidar com campos vazios
func EstruturaDados(dados string, resultado interface{}) error {
	valr := reflect.ValueOf(resultado).Elem()
	tipo := valr.Type()

	inicio := 0

	for i := 0; i < valr.NumField(); i++ {
		campo := valr.Field(i)
		dadoTag := tipo.Field(i).Tag.Get("dado") // Usando dadoTag para log ou depuração
		tamanhoTag := tipo.Field(i).Tag.Get("tamanho")

		// Converter o tamanho da tag
		tamanho, err := strconv.Atoi(tamanhoTag)
		if err != nil {
			return fmt.Errorf("erro ao converter tamanho do campo %s: %v", dadoTag, err)
		}

		// Verifica se há dados suficientes para preencher o campo
		if inicio+tamanho > len(dados) {
			// Se os dados forem insuficientes, pular este campo sem erro
			fmt.Printf("Dados insuficientes para preencher o campo %s. Continuando...\n", dadoTag)
			continue
		}

		valor := dados[inicio : inicio+tamanho]
		inicio += tamanho

		// Definir o valor somente se houver algo para definir
		if campo.CanSet() && valor != "" {
			campo.SetString(valor)
			fmt.Printf("Campo %s preenchido com valor: %s\n", dadoTag, valor)
		} else {
			fmt.Printf("Campo %s está vazio ou não pode ser definido\n", dadoTag)
		}
	}
	return nil
}
