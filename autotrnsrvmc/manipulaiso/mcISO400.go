package manipulaiso

import (

	//"encoding/csv"

	especificacao "autotrnsrvmc/especificacao"
	"autotrnsrvmc/modelos"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/moov-io/iso8583"
	//"github.com/moov-io/iso8583/cmd/iso8583/describe"

	//"github.com/moov-io/iso8583/specs"
	"github.com/moov-io/iso8583/network"
)

func Mciso400(iso_padrao []byte) ([]byte, error) {

	// 1 se renomeia a iso criada acima com pacote iso_padrao

	fmt.Println("PRINT DA FUNCAO MC", iso_padrao)

	// Remove os primeiros 2 bytes
	iso_padrao_mod := iso_padrao[2:]
	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_padrao_mod)

	r_spec := especificacao.NewSpecASCII()

	r_message := iso8583.NewMessage(r_spec)

	r_message.Unpack([]byte((retorno)))

	b, err := r_message.Pack()
	if err != nil {
		panic(err)

	}
	fmt.Printf("\n MENSAGEM RECEBIDA MONTAGEM DE CANCELAMENTO:\n")
	fmt.Printf("% x\n", b)
	fmt.Printf("\n MENSAGEM ABERTA ASCII:\n")
	iso8583.Describe(r_message, os.Stdout)

	new_spec := especificacao.NewSpecEBCDIC()

	new_message := iso8583.NewMessage(new_spec)

	r_message.Unpack(b) //Abre a ISO formatada no inicio para ser usada somente de alguns campos

	valida_mti, err := r_message.GetMTI()
	if err != nil {
		panic(err)
	}
	pan, err := r_message.GetString(2)
	if err != nil {
		panic(err)
	}

	cartao, err := r_message.GetString(35)
	if err != nil {
		panic(err)
	}
	//Abre o DE 35, le ate = para encontrar o numero do cartão
	var num_cartao string
	if pan == "" {
		if cartao != "" {
			for _, procura_cartao := range cartao {
				if procura_cartao == '=' || procura_cartao == 'D' || procura_cartao == 'C' {
					break
				}
				num_cartao += string(procura_cartao)
			}
		} else {
			log.Println("DE02 OU DE35 NAO PODE ESTAR VAZIOS")
		}
	} else {

		num_cartao = pan

	}

	proccod, err := r_message.GetString(3)
	if err != nil {
		panic(err)
	}

	valor, err := r_message.GetString(4)
	if err != nil {
		panic(err)
	}
	mmddhhmmss, err := r_message.GetString(7)
	if err != nil {
		panic(err)
	}
	//new_message.Field(7, mmddhhmmss) //data e hora mmddhhmmss
	sy_trace, err := r_message.GetString(11)
	if err != nil {
		panic(err)
	}
	//new_message.Field(11, sy_trace)
	hhmmss, err := r_message.GetString(12)
	if err != nil {
		panic(err)
	}
	//new_message.Field(12, hhmmss) //hhmmss
	mmdd, err := r_message.GetString(13)
	if err != nil {
		panic(err)
	}
	//new_message.Field(13, mmdd) //mmdd
	validade, err := r_message.GetString(14)
	if err != nil {
		panic(err)
	}
	//new_message.Field(14, validade) //validade do cartao
	mcc, err := r_message.GetString(18)
	if err != nil {
		panic(err)
	}
	//new_message.Field(18, mcc) //mcc
	entrymode, err := r_message.GetString(22)
	if err != nil {
		panic(err)
	}
	acqr, err := r_message.GetString(32)
	if err != nil {
		panic(err)
	}
	acqrremente, err := r_message.GetString(33)
	if err != nil {
		panic(err)
	}
	endereco, err := r_message.GetString(43)
	if err != nil {
		panic(err)
	}
	//new_message.Field(22, entrymode) //entry mode
	//trilha2, err := r_message.GetString(35)
	//if err != nil {
	//	panic(err)
	//}
	//new_message.Field(35, trilha2) // trilha 2 cartao/ identificador/ validade/ cvv
	data := time.Now()
	//new_message.Field(37, strings.Replace((data.Format(("150405.000000"))), ".", "", -1))
	tid, err := r_message.GetString(41)
	if err != nil {
		panic(err)
	}
	//new_message.Field(41, tid) //terminal
	mid, err := r_message.GetString(42)
	if err != nil {
		panic(err)
	}
	iso48ent, err := r_message.GetString(48)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	// Envia para funçao que lida com dados no formato TLV para manipulacao
	tlvManipu48ent, err := especificacao.LeTLVASCII(iso48ent)
	if err != nil {
		fmt.Println("Erro:", err)
		//return
	}
	// Criar uma instância da estrutura
	var Iso48EntrManip modelos.Iso48Entrada

	// Preenche a estrutura com os dados TLV
	err = especificacao.EstruturaTLV(tlvManipu48ent, &Iso48EntrManip)
	if err != nil {
		fmt.Println("Erro:", err)
		//return
	}
	//new_message.Field(42, mid) //cod comercio
	//new_message.Field(43, "POSTO DM JR            DIADEMA  EVANG076")
	moeda, err := r_message.GetString(49)
	if err != nil {
		panic(err)
	}
	senha, err := r_message.GetString(52)
	if err != nil {
		panic(err)
	}
	//new_message.Field(49, moeda) //moeda
	//de55, err := r_message.GetString(55)
	//if err != nil {
	//	panic(err)
	//}
	//new_message.Field(55, de55) //dados do chip
	posdata, err := r_message.GetString(61)
	if err != nil {
		panic(err)
	}
	// Envia para funçao que lida com dados no formato TLV para manipulacao
	tlvManipulado, err := especificacao.LeTLVASCII(posdata)
	if err != nil {
		fmt.Println("Erro:", err)
		//return
	}

	// Criar uma instância da estrutura
	var isoData61 modelos.IsoDe61ret
	var isoData90 modelos.IsoDe90Entrada
	// Criar uma instância da estrutura
	var isoDatasub61 modelos.IsoDe61sub03ret

	// Preenche a estrutura com os dados TLV
	err = especificacao.EstruturaTLV(tlvManipulado, &isoData61)
	if err != nil {
		fmt.Println("Erro:", err)
		//return
	}
	// Preenche a estrutura os dados enviado somemte o campo que queremos
	err = especificacao.EstruturaDados(isoData61.DadosBandeira03, &isoDatasub61)
	if err != nil {
		fmt.Println("Erro:", err)
		//return
	}
	dadoscanc, err := r_message.GetString(90)
	if err != nil {
		panic(err)
	}
	// Envia para funçao que lida com dados no formato TLV para manipulacao
	//	tlvManipulado90, err := especificacao.LeTLVASCII(dadoscanc)
	//	if err != nil {
	//		fmt.Println("Erro:", err)
	//return
	//	}
	// Preenche a estrutura com os dados TLV
	//err = especificacao.EstruturaTLV(tlvManipulado90, &isoData90)
	//if err != nil {
	//	fmt.Println("Erro:", err)
	//return
	//}
	// Preenche a estrutura os dados enviado somemte o campo que queremos
	err = especificacao.EstruturaDados(dadoscanc, &isoData90)
	if err != nil {
		fmt.Println("Erro:", err)
		//return
	}

	//Logica para definir se é debito single ou dual
	var ReplecAmountDeb string
	var DadosRedeDeb string
	var AdvReasonCodDeb string
	var SetlmMmdd2 string
	//var singledualmsncanc string
	//var singledualmsn string
	var ParcelISO112 string //vai informacao de cpf cnpj é o mesmo campo usado para pacelado
	//if Iso48EntrManip.DocumentoSigleMSN != "" && (proccod == "20" || proccod == "202000" || proccod == "202000") {
	if valida_mti == "0420" {
		//singledualmsn = "0420"
		//singledualmsncanc = "0200"
		AdvReasonCodDeb = "000000000000000000000000000000000000000000000000000000000000"
		DadosRedeDeb = "MS0000000000"
		ReplecAmountDeb = "000000000000000000000000000000000000000000"
		ParcelISO112 = "029014" + Iso48EntrManip.DocumentoSigleMSN
		//Add DE15 para siglemsn
		SetlmMmdd2 = validade
		//DE14 nao é enviado em sigle msn
		validade = ""
		//Campos abaixo nao devem ser enviados siglemsn
		mcc = ""
		entrymode = ""
		//trilha2 = ""
		mid = ""
		//endereco = ""

	}
	pedidoautoriza := &modelos.SolicitacaoAutorizacao{} //Absorve a estrutura do SolicitacaoAutorizacao para o pedidoautoriza
	pedidoautoriza.MTI = valida_mti
	pedidoautoriza.PAN = num_cartao
	pedidoautoriza.ProcCod = proccod
	pedidoautoriza.Valor = valor
	pedidoautoriza.DataHoraTransacao = mmddhhmmss
	pedidoautoriza.STAN = sy_trace
	pedidoautoriza.Hhmmss = hhmmss
	pedidoautoriza.Mmdd = mmdd
	pedidoautoriza.Validade = validade
	pedidoautoriza.SetlmMmdd = SetlmMmdd2
	pedidoautoriza.Mcc = mcc
	pedidoautoriza.EntryMode = entrymode
	if entrymode == "51" || entrymode == "71" { //ou 51 ou 71, fora esses nao vai o DE23
		pedidoautoriza.SeqCardNum = "001"
	}
	pedidoautoriza.Adquirente = acqr
	pedidoautoriza.CodInstiRemetente = acqrremente
	//pedidoautoriza.Trilha2 = trilha2
	pedidoautoriza.Nsu = (strings.Replace((data.Format(("150405.000000"))), ".", "", -1))
	pedidoautoriza.Cod_resp = "30"
	pedidoautoriza.Terminal = tid
	pedidoautoriza.CodComercio = mid
	pedidoautoriza.NomeEndereco = endereco
	// Verificar se é cash in
	var InfomeCarteira string
	// Verificar se tem dados para transaçao parcelada recebido no DE48
	var TemParcelamento string
	var Parcelamento string
	var IndicadorPgto string
	// var ParcelISO112 string
	var TipoParcelado48 string
	// fmt.Println("QUE PEGOU DO 48", Iso48EntrManip.TipoParceleado)
	// panic(err)
	if Iso48EntrManip.Parcelas == "" || Iso48EntrManip.Parcelas == "00" || Iso48EntrManip.Parcelas == "01" {
		TemParcelamento = ""
		Parcelamento = ""
	} else {
		TipoParcelado48 = Iso48EntrManip.TipoParceleado
		TemParcelamento = Iso48EntrManip.Parcelas
		Parcelamento = "PARCEL"
		IndicadorPgto = "P10"
		//ParcelISO112 = "021003" + TipoParcelado48 + "B0220030" + TemParcelamento
		if TipoParcelado48 == "25" && valor < "1" {
			// Consulta de crediario sub19 é o valor simulador, colocar depois no csv
			//ParcelISO112 = "001004" + TipoParcelado48 + TemParcelamento + "019012000000100000"
			ParcelISO112 = "001004" + TipoParcelado48 + TemParcelamento + "019012" + Iso48EntrManip.VlrConsulCrediario
		} else {
			ParcelISO112 = "001004" + TipoParcelado48 + TemParcelamento
		}
	}
	// Logica para informar se foi usado senha online ou nao
	var senhaOnline string
	if valida_mti == "0420" {
		senhaOnline = ""
	} else if senha != "" {
		senhaOnline = "S" //Pode significar assinatura, “Verificação da senha off-line” (para transaçõescom chip), “M-PIN” (para Dispositivo Móvel com capacidade para entrada senha) ou “Nenhum CVM foi utilizado”
	} else {
		senhaOnline = "P" //Verificação da senha on-line
	}
	pedidoautoriza.InfoAdd = &modelos.IsoDe48{
		//TCC: "F",
		TCC: Iso48EntrManip.Tcc,
		CompsicaoTCC: &modelos.IsoDe48sub{
			VerifSenha: senhaOnline,
			//MerchantCategoryCode=" ",
			//Sub48_22i=&De48sub22{
			//Sub48_22="0000",
			//Sub48_22i="0000000000000000000000000000000000",
			//},
			//MerchantPostalCode="50C",
			//MerchantWebsite=   "TV",
			//Ard=               "111111111111111",
			InfoCarteira: InfomeCarteira, //necessario para cash in
			//InfoPOSLocal:     "B",
			PaymentIndicator: IndicadorPgto,
			InfoParcelado:    Parcelamento,
		},
		//InfoParcelado: "PARCEL",
	}
	pedidoautoriza.Moeda = moeda
	if entrymode == "51" || entrymode == "21" { //ou 51 ou 21
		pedidoautoriza.Pin = "FEE8CA6A"
	}
	//if entrymode == "51" || entrymode == "71" {
	//	pedidoautoriza.DadosChip = de55
	//}
	if valida_mti != "0420" {
		//Gera DE60
		//	InfomeCarteira = "101"
		//	pedidoautoriza.AdvReasonCod = "1210000cash in"
		//} else {
		//Gera DE61 obrigatorios para outras transacoes
		pedidoautoriza.DadosPOS = &modelos.IsoDe61{
			POSta01:          isoDatasub61.POSta01,
			Reserv02:         isoDatasub61.Reserv02,
			POStl03:          isoDatasub61.POStl03,
			POSchPresent04:   isoDatasub61.POSchPresent04,
			POScardPresent05: isoDatasub61.POScardPresent05,
			POScardCapCap06:  isoDatasub61.POScardCapCap06,
			POStrnStatus07:   isoDatasub61.POStrnStatus07,
			POStrnSecur08:    isoDatasub61.POStrnSecur08,
			Reserv09:         isoDatasub61.Reserv09,
			POSCat10:         isoDatasub61.POSCat10,
			POSIputCap11:     isoDatasub61.POSIputCap11,
			//POSliveCycle12:   isoDatasub61.POSliveCycle12,
			POSliveCycle12: "00",
			POScodPais13:   isoDatasub61.POScodPais13,
			POScodPostal14: isoDatasub61.POScodPostal14,
		}
	}
	pedidoautoriza.AdvReasonCod = AdvReasonCodDeb
	pedidoautoriza.DadosRede = DadosRedeDeb
	//pedidoautoriza.DadosPOS = posdata
	pedidoautoriza.DadosCanc = &modelos.IsoDe90{
		MtiOrin:    isoData90.MtiOrin,
		StanOrin:   isoData90.StanOrin,
		HhmmssOrin: isoData90.HhmmssOrin,
		AcqrOrin:   isoData90.AcqrOrin,
		InstFOrin:  isoData90.InstFOrin,
	}
	pedidoautoriza.ReplecAmount = ReplecAmountDeb
	//pedidoautoriza.DadosPOS = posdata

	pedidoautoriza.DadosADDParcelado = ParcelISO112
	//pedidoautoriza.DadosRegistro = De120
	//pedidoautoriza.DadosPrivados = De126

	err = new_message.Marshal(pedidoautoriza)

	d, err := new_message.Pack()
	if err != nil {
		panic(err)

	}

	// 1 se renomeia a iso criada acima com pacote b
	packed := d

	// 2 se criar o cabeçalho binario 2 byts
	header := network.NewBinary2BytesHeader()
	//header := network.NewASCII4BytesHeader()
	header.SetLength(len(packed))

	// 3 se criar o tpdu ELO
	//tpdu := []byte("60 00 06 00 00")

	// 4 - combinar tudo o que temos em um buf

	var buf bytes.Buffer

	header.WriteTo(&buf)
	//_, err = buf.Write(tpdu)
	_, err = buf.Write(packed)
	//_, err = Write(buf.Bytes())
	fmt.Printf("\n MENSAGEM ABERTA EBCDIC:\n")
	fmt.Printf("% x\n", packed)
	iso8583.Describe(new_message, os.Stdout)
	fmt.Printf("\n FIM DO PROCESSAMENTO :) \n")
	//fmt.Println("Pressione 'Enter' para sair...")
	//fmt.Scanln()
	//conn.Close()
	fmt.Println("\n MENSAGEM EBCDIC ENVIADA AO SERVIDOR:\n", buf.Bytes())
	return buf.Bytes(), err
}
