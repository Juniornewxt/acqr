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

	"github.com/moov-io/iso8583/network"
)

func Mciso100(iso_padrao []byte) ([]byte, error) {

	// Remove os primeiros 2 bytes
	iso_padrao_mod := iso_padrao[2:]
	// Trabalha com string sem os 2 primeiros bytes
	retorno := string(iso_padrao_mod)

	r_spec := especificacao.NewSpecASCII()

	r_message := iso8583.NewMessage(r_spec)

	r_message.Unpack([]byte((retorno)))

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

	iso55, err := r_message.GetString(55)
	if err != nil {
		panic(err)
	}
	//fmt.Println("PRINT DO 55", iso55)

	var objetosISO55 modelos.ISO55dados
	manipulaISO55, err := especificacao.LeBerTLV55(iso55, &objetosISO55)
	if err != nil {
		fmt.Println("Erro:", err)

	}

	// Criar uma instância da estrutura
	var ISO55Manipulada modelos.ISO55dados

	// Preenche a estrutura com os dados TLV
	err = especificacao.EstruturaTLV(manipulaISO55, &ISO55Manipulada)
	if err != nil {
		fmt.Println("Erro:", err)
		//return
	}
	//fmt.Println("PRINT DO 55 CAPTURADO", ISO55Manipulada.AC9F26)

	iso61, err := r_message.GetString(61)
	if err != nil {
		panic(err)
	}

	// Envia para funçao que lida com dados no formato TLV para manipulacao
	tlvManipulado, err := especificacao.LeTLVASCII(iso61)
	if err != nil {
		fmt.Println("Erro:", err)
		//return
	}

	// Criar uma instância da estrutura
	var isoData61 modelos.IsoDe61ret
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
	//recebido := &modelos.SolicitacaoAutorizacao2{}
	//err = r_message.Unmarshal(recebido)
	//fmt.Println("TESTE DO QUE RECEBEU PARA BANDEIRA 1 ", recebido.DadosChip.AC9F26)
	//fmt.Println("TESTE DO QUE RECEBEU PARA BANDEIRA TUDO", recebido)

	b, err := r_message.Pack()
	if err != nil {
		panic(err)

	}

	log.Printf("MENSAGEM RECEBIDA:\n")
	log.Printf("% x\n", b)
	log.Printf("MENSAGEM ABERTA ASCII:\n")
	iso8583.Describe(r_message, os.Stdout)

	//Inicia o tratamento da mensagem em EBCDIC para Mastercard

	new_spec := especificacao.NewSpecEBCDIC()

	new_message := iso8583.NewMessage(new_spec)

	r_message.Unpack(b) //Abre a ISO formatada no inicio para ser usada somente de alguns campos

	cartao, err := r_message.GetString(35)
	if err != nil {
		panic(err)
	}
	pan, err := r_message.GetString(2)
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
	sequencecard, err := r_message.GetString(23)
	if err != nil {
		panic(err)
	}
	//new_message.Field(22, entrymode) //entry mode
	trilha2, err := r_message.GetString(35)
	if err != nil {
		panic(err)
	}
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
	//new_message.Field(42, mid) //cod comercio
	//new_message.Field(43, "POSTO DM JR            DIADEMA  EVANG076")
	moeda, err := r_message.GetString(49)
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
	//Logica para definir se é debito single ou dual
	var singledualmsn string
	var ParcelISO112 string //vai informacao de cpf cnpj é o mesmo campo usado para pacelado
	if Iso48EntrManip.DocumentoSigleMSN != "" && (proccod == "0" || proccod == "112000" || proccod == "2000") {
		singledualmsn = "0200"
		ParcelISO112 = "029014" + Iso48EntrManip.DocumentoSigleMSN
		//DE14 nao é enviado em sigle msn
		validade = ""
	} else {
		singledualmsn = "0100"
	}
	pedidoautoriza := &modelos.SolicitacaoAutorizacao{} //Absorve a estrutura do SolicitacaoAutorizacao para o pedidoautoriza
	//pedidoautoriza.MTI = "0100"
	pedidoautoriza.MTI = singledualmsn
	pedidoautoriza.PAN = num_cartao
	pedidoautoriza.ProcCod = proccod
	pedidoautoriza.Valor = valor
	pedidoautoriza.DataHoraTransacao = mmddhhmmss
	pedidoautoriza.STAN = sy_trace
	pedidoautoriza.Hhmmss = hhmmss
	pedidoautoriza.Mmdd = mmdd
	pedidoautoriza.Validade = validade
	pedidoautoriza.Mcc = mcc
	pedidoautoriza.EntryMode = entrymode
	if entrymode == "50" || entrymode == "70" || entrymode == "51" || entrymode == "71" || entrymode == "52" || entrymode == "72" || entrymode == "58" || entrymode == "78" || entrymode == "810" || entrymode == "820" || entrymode == "811" || entrymode == "821" || entrymode == "812" || entrymode == "822" || entrymode == "818" || entrymode == "828" { //ou 51 ou 71, fora esses nao vai o DE23
		pedidoautoriza.SeqCardNum = sequencecard
	}
	//pedidoautoriza.Adquirente = "00000000025"
	pedidoautoriza.Adquirente = acqr
	//pedidoautoriza.CodInstiRemetente = "0026"
	pedidoautoriza.CodInstiRemetente = acqrremente
	if entrymode != "10" && entrymode != "11" && entrymode != "12" && entrymode != "18" && entrymode != "810" && entrymode != "811" && entrymode != "812" && entrymode != "818" && entrymode != "101" && entrymode != "100" && entrymode != "102" && entrymode != "108" { //10 compra digitada nao envia trilha2 35 sub0 Sem inf de senha 1 terminal permite senha,
		// Se nao tiver valor de transacao é consulta crediario.
		if valor != "0" {
			// contiacao... 2 terminal nao tem capacidade de pin, 8 terminal tem capacidade mas nao esta ativo
			pedidoautoriza.Trilha2 = trilha2
		}
	}
	pedidoautoriza.Nsu = (strings.Replace((data.Format(("150405.000000"))), ".", "", -1))
	pedidoautoriza.Terminal = tid
	pedidoautoriza.CodComercio = mid
	//pedidoautoriza.NomeEndereco = "POSTO DM JR            DIADEMA  EVANG076"
	pedidoautoriza.NomeEndereco = endereco
	// Verificar se é cash in
	var InfomeCarteira string
	// Verificar se tem dados para transaçao parcelada recebido no DE48
	var TemParcelamento string
	var Parcelamento string
	var IndicadorPgto string
	//var ParcelISO112 string
	var TipoParcelado48 string
	//fmt.Println("QUE PEGOU DO 48", Iso48EntrManip.TipoParceleado)
	//panic(err)
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

	pedidoautoriza.InfoAdd = &modelos.IsoDe48{
		//TCC: "F",
		TCC: Iso48EntrManip.Tcc,
		CompsicaoTCC: &modelos.IsoDe48sub{
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
	if entrymode == "51" || entrymode == "71" || entrymode == "901" || entrymode == "911" || entrymode == "801" || entrymode == "011" || entrymode == "101" { //ou 51 ou 71
		pedidoautoriza.Pin = "FEE8CA6A"
		//senha, err := r_message.GetString(52)
		//if err != nil {
		//	panic(err)
		//}
		//pedidoautoriza.Pin = senha
	}

	//pedidoautoriza.Pin = recebido.Pin
	if entrymode == "50" || entrymode == "70" || entrymode == "51" || entrymode == "71" || entrymode == "52" || entrymode == "72" || entrymode == "58" || entrymode == "78" || entrymode == "810" || entrymode == "820" || entrymode == "811" || entrymode == "821" || entrymode == "812" || entrymode == "822" || entrymode == "818" || entrymode == "828" {

		// 9F03 Usado para cashback, se vier nulo preencher com zero.
		var ModOutroValor9F03 string
		if ISO55Manipulada.OutroValor9F03 == "" {
			ModOutroValor9F03 = "000000000000"
		} else {
			ModOutroValor9F03 = ISO55Manipulada.OutroValor9F03
		}
		// Se nao tiver valor de transacao é consulta crediario.
		if valor != "0" {
			pedidoautoriza.DadosChip = &modelos.IsoDe55{

				AC9F26:                ISO55Manipulada.AC9F26,                //9F26
				InfoCript9F27:         ISO55Manipulada.InfoCript9F27,         //9F27
				IAD9F10:               ISO55Manipulada.IAD9F10,               //9F10
				NumImprevi9F37:        ISO55Manipulada.NumImprevi9F37,        //9F37
				ContadorTrnApp9F36:    ISO55Manipulada.ContadorTrnApp9F36,    //9F36
				TRV95:                 ISO55Manipulada.TRV95,                 //95
				DataTrn9A:             ISO55Manipulada.DataTrn9A,             //9A
				TipoTrn9C:             ISO55Manipulada.TipoTrn9C,             //9C
				ValorAutoriz9F02:      ISO55Manipulada.ValorAutoriz9F02,      //9F02
				CodMoedaTrn5F2A:       ISO55Manipulada.CodMoedaTrn5F2A,       //5F2A
				PerfilIntercApp82:     ISO55Manipulada.PerfilIntercApp82,     //82
				CodPaisTerm9F1A:       ISO55Manipulada.CodPaisTerm9F1A,       //9F1A
				ResultCVM9F34:         ISO55Manipulada.ResultCVM9F34,         //9F34
				CapacidadeTermEUA9F33: ISO55Manipulada.CapacidadeTermEUA9F33, //9F33
				NumDedicado84:         ISO55Manipulada.NumDedicado84,         //84
				//OutroValor9F03:        ISO55Manipulada.OutroValor9F03,        //9F03
				OutroValor9F03: ModOutroValor9F03,
				////opcionais abaixo  :ISO55Manipulada.////opcionais abaixo , //
				NumSeqPan5F34:         ISO55Manipulada.NumSeqPan5F34,         //5F34
				DadoProprCadSelAp9F0A: ISO55Manipulada.DadoProprCadSelAp9F0A, //9F0A
				//TipoTerminal9F35:      ISO55Manipulada.TipoTerminal9F35,      //9F35
				IDF9F1E:            ISO55Manipulada.IDF9F1E,            //9F1E
				CodCatgTrn9F53:     ISO55Manipulada.CodCatgTrn9F53,     //9F53
				NumVersaoApp9F09:   ISO55Manipulada.NumVersaoApp9F09,   //9F09
				ContadorSeqTrn9F41: ISO55Manipulada.ContadorSeqTrn9F41, //9F41
				//TermCapForaEUA9F33  :ISO55Manipulada.//TermCapForaEUA9F33 , //9F33
				DadosTerceiro9F6E: ISO55Manipulada.DadosTerceiro9F6E, //9F6E
			}
		}
	}
	//Verificar se é elegivel cash in
	if mcc == "6540" && entrymode == "102" || entrymode == "108" {
		//Gera DE60
		InfomeCarteira = "101"
		pedidoautoriza.AdvReasonCod = "1210000cash in"
	} else {
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
			POSliveCycle12:   isoDatasub61.POSliveCycle12,
			POScodPais13:     isoDatasub61.POScodPais13,
			POScodPostal14:   isoDatasub61.POScodPostal14,
		}
	}

	//pedidoautoriza.DadosPOS = posdata

	pedidoautoriza.DadosADDParcelado = ParcelISO112
	//pedidoautoriza.DadosADDParcelado = "0010042112"
	//pedidoautoriza.DadosRegistro = De120
	//pedidoautoriza.DadosPrivados = De126

	err = new_message.Marshal(pedidoautoriza)

	d, err := new_message.Pack()
	if err != nil {
		panic(err)

	}

	// renomeia a iso criada acima com pacote b
	packed := d

	// criar o cabeçalho binario 2 byts
	header := network.NewBinary2BytesHeader()
	//header := network.NewASCII4BytesHeader()
	header.SetLength(len(packed))

	var buf bytes.Buffer

	header.WriteTo(&buf)
	//_, err = buf.Write(tpdu)
	_, err = buf.Write(packed)
	//_, err = Write(buf.Bytes())
	log.Printf("MENSAGEM ABERTA EBCDIC:\n")
	log.Printf("% x\n", packed)
	iso8583.Describe(new_message, os.Stdout)
	fmt.Println()
	log.Printf("FIM DO PROCESSAMENTO :) \n")

	log.Println("MENSAGEM EBCDIC FORMATADA PARA SER ENVIADA AO SERVIDOR:", buf.Bytes())
	return buf.Bytes(), err
}
