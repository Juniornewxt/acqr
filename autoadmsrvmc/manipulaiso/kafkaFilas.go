package manipulaiso

import (

	//"encoding/csv"

	"github.com/moov-io/iso8583"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"

	//"github.com/moov-io/iso8583/specs"
	especificacao "autoadmsrvmc/especificacao"
	"autoadmsrvmc/msnkafka"
)

func EnviaKafka(iso []byte) ([]byte, error) {

	spec := especificacao.NewSpecEBCDIC()

	message := iso8583.NewMessage(spec)

	message.Unpack([]byte((iso)))

	mtirecebido, err := message.GetMTI()
	if err != nil {
		panic(err)
	}
	// VALIDA QUAL TIPO DE MSN
	if mtirecebido >= "0800" && mtirecebido <= "0899" {

		pan, err := message.GetString(2)
		if err != nil {
			panic(err)
		}
		if pan == "" {
			pan = "..."
		}
		dmhhmmss, err := message.GetString(7)
		if err != nil {
			panic(err)
		}
		if dmhhmmss == "" {
			dmhhmmss = "..."
		}
		stan, err := message.GetString(11)
		if err != nil {
			panic(err)
		}
		if stan == "" {
			stan = "..."
		}
		fiic, err := message.GetString(33)
		if err != nil {
			panic(err)
		}
		if fiic == "" {
			fiic = "..."
		}
		codresp, err := message.GetString(39)
		if err != nil {
			panic(err)
		}
		if codresp == "" {
			codresp = "00"
		}
		addinfo, err := message.GetString(48)
		if err != nil {
			panic(err)
		}
		if addinfo == "" {
			addinfo = "..."
		}
		netdata, err := message.GetString(63)
		if err != nil {
			panic(err)
		}
		if netdata == "" {
			netdata = "..."
		}
		codrede, err := message.GetString(70)
		if err != nil {
			panic(err)
		}
		if codrede == "" {
			codrede = "..."
		}
		privatedata, err := message.GetString(126)
		if err != nil {
			panic(err)
		}
		if privatedata == "" {
			privatedata = "..."
		}

		//fmt.Println("ISSO AQUI É PARA O KAFKA", mtirecebido, pan, dmhhmmss, stan, fiic, codresp, addinfo, netdata, codrede, privatedata)

		msnkafka.EnviaFila0800(mtirecebido, pan, dmhhmmss, stan, fiic, codresp, addinfo, netdata, codrede, privatedata)

	} else if mtirecebido >= "0100" && mtirecebido <= "0299" {

		pan, err := message.GetString(2)
		if err != nil {
			panic(err)
		}
		if pan == "" {
			pan = "..."
		} else {
			pan = MascaraCartao(pan)
		}
		proccod, err := message.GetString(3)
		if err != nil {
			panic(err)
		}
		if proccod == "" {
			proccod = "..."
		}
		valor, err := message.GetString(4)
		if err != nil {
			panic(err)
		}
		if valor == "" {
			valor = "..."
		}
		dmhhmmss, err := message.GetString(7)
		if err != nil {
			panic(err)
		}
		if dmhhmmss == "" {
			dmhhmmss = "..."
		}
		stan, err := message.GetString(11)
		if err != nil {
			panic(err)
		}
		if stan == "" {
			stan = "..."
		}
		hhmmss, err := message.GetString(12)
		if err != nil {
			panic(err)
		}
		if hhmmss == "" {
			hhmmss = "..."
		}
		mcc, err := message.GetString(18)
		if err != nil {
			panic(err)
		}
		if mcc == "" {
			mcc = "..."
		}
		EntryMode, err := message.GetString(22)
		if err != nil {
			panic(err)
		}
		if EntryMode == "" {
			EntryMode = "..."
		}
		Adquirente, err := message.GetString(32)
		if err != nil {
			panic(err)
		}
		if Adquirente == "" {
			Adquirente = "..."
		}
		CodInstiRemetente, err := message.GetString(33)
		if err != nil {
			panic(err)
		}
		if CodInstiRemetente == "" {
			CodInstiRemetente = "..."
		}
		Nsu, err := message.GetString(37)
		if err != nil {
			panic(err)
		}
		if Nsu == "" {
			Nsu = "..."
		}
		CodAuto, err := message.GetString(38)
		if err != nil {
			panic(err)
		}
		if CodAuto == "" {
			CodAuto = "..."
		}
		Codresp, err := message.GetString(39)
		if err != nil {
			panic(err)
		}
		if Codresp == "" {
			Codresp = "00"
		}
		Terminal, err := message.GetString(41)
		if err != nil {
			panic(err)
		}
		if Terminal == "" {
			Terminal = "..."
		}
		CodComercio, err := message.GetString(42)
		if err != nil {
			panic(err)
		}
		if CodComercio == "" {
			CodComercio = "..."
		}
		NomeEndereco, err := message.GetString(43)
		if err != nil {
			panic(err)
		}
		if NomeEndereco == "" {
			NomeEndereco = "..."
		}
		Moeda, err := message.GetString(49)
		if err != nil {
			panic(err)
		}
		if Moeda == "" {
			Moeda = "..."
		}
		DadosADDParcelado, err := message.GetString(112)
		if err != nil {
			panic(err)
		}
		if DadosADDParcelado == "" {
			DadosADDParcelado = "..."
		}

		msnkafka.EnviaFilaAuth(mtirecebido,
			pan,
			proccod,
			valor,
			dmhhmmss,
			stan,
			hhmmss,
			mcc,
			EntryMode,
			Adquirente,
			CodInstiRemetente,
			Nsu,
			CodAuto,
			Codresp,
			Terminal,
			CodComercio,
			NomeEndereco,
			Moeda,
			DadosADDParcelado)
	}
	return nil, err
	//return buf.Bytes(), err
}
func MascaraCartao(cartao string) string {
	// Verifica se o numero do cartao "Pan" é menor que 19 caracteres
	if len(cartao) < 8 {
		return cartao
	}

	// Pega os 4 primeiros e os 4 últimos caracteres do cartao
	inicio := cartao[:4]
	fim := cartao[len(cartao)-4:]

	// Concatena os 4 primeiros com os 4 últimos numeros do cartao
	return inicio + "****" + fim
}
