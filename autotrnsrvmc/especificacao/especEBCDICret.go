package especificacao

import (
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/padding"
	"github.com/moov-io/iso8583/prefix"
)

func NewSpecEBCDICret() *iso8583.MessageSpec {
	return &iso8583.MessageSpec{
		Fields: map[int]field.Field{
			0: field.NewString(&field.Spec{
				Length:      4,
				Description: "MTI",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				//Pad:         padding.Left('0'),
			}),
			1: field.NewBitmap(&field.Spec{
				Description: "Bitmap",
				Enc:         encoding.Binary,
				Pref:        prefix.Hex.Fixed,
			}),
			2: field.NewString(&field.Spec{
				Length:      16,
				Description: "PAN - CARTAO",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LL,
			}),
			3: field.NewNumeric(&field.Spec{
				Length:      6,
				Description: "CODIGO PROCESSAMENTO",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			4: field.NewNumeric(&field.Spec{
				Length:      12,
				Description: "VALOR DA TRANSACAO",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			5: field.NewNumeric(&field.Spec{
				Length:      12,
				Description: "VALOR DA TRANSACAO AGENDA",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			6: field.NewNumeric(&field.Spec{
				Length:      12,
				Description: "VALOR DA TRANSACAO CARDHOLDER",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			7: field.NewNumeric(&field.Spec{
				Length:      10,
				Description: "DATA HORA GMT TRANSMISSAO",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			9: field.NewNumeric(&field.Spec{
				Length:      8,
				Description: "CONVERSAO AGENDA",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			10: field.NewNumeric(&field.Spec{
				Length:      8,
				Description: "CONVERSAO CARD HOLDER",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			11: field.NewNumeric(&field.Spec{
				Length:      6,
				Description: "DE011",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			12: field.NewNumeric(&field.Spec{
				Length:      6,
				Description: "HORA LOCAL TRANSACAO",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			13: field.NewNumeric(&field.Spec{
				Length:      4,
				Description: "DATA LOCAL TRANSACAO",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			14: field.NewNumeric(&field.Spec{
				Length:      4,
				Description: "VALIDADE",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			15: field.NewNumeric(&field.Spec{
				Length:      4,
				Description: "DATA AGENDA",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			16: field.NewNumeric(&field.Spec{
				Length:      4,
				Description: "DATA CONVERSAO",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			18: field.NewNumeric(&field.Spec{
				Length:      4,
				Description: "MCC",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			19: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "DE19",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			22: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "ENTRY MODE",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			23: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "DE23",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			24: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "DE24",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			32: field.NewString(&field.Spec{
				Length:      11,
				Description: "CODIGO ADQUIRENTE",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LL,
			}),
			33: field.NewString(&field.Spec{
				Length:      6,
				Description: "Forwarding Institution ID Code",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LL,
				//Pad:         padding.Left('0'),
			}),
			35: field.NewString(&field.Spec{
				Length:      37,
				Description: "TRILHA 2",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LL,
			}),
			37: field.NewString(&field.Spec{
				Length:      12,
				Description: "NSU",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			38: field.NewNumeric(&field.Spec{
				Length:      6,
				Description: "CODIGO DE AUTORIZACAO",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			39: field.NewString(&field.Spec{
				Length:      2,
				Description: "CODIGO DE RESPOSTA",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				//Pad:         padding.Left('0'),
			}),
			41: field.NewString(&field.Spec{
				Length:      8,
				Description: "TERMINAL",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			42: field.NewString(&field.Spec{
				Length:      15,
				Description: "CODIGO DO EC",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			43: field.NewString(&field.Spec{
				Length:      40,
				Description: "DADOS EC",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			46: field.NewString(&field.Spec{
				Length:      999,
				Description: "Identificador do Cartao -CVE2",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
			}),
			48: field.NewString(&field.Spec{
				Length:      999,
				Description: "Informacoes adicionais de solic resp",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
			}),
			49: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "CODIGO DE MOEDA",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			50: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "MOEDA DE AGENDA",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			51: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "MOEDA CARDHOLDER",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
			}),
			52: field.NewString(&field.Spec{
				Length:      8,
				Description: "SENHA CRIPTOGRAFADA",
				Enc:         encoding.Binary,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			55: field.NewString(&field.Spec{
				Length:      999,
				Description: "ICC Data – EMV Having Multiple Tags",
				Enc:         encoding.Binary,
				Pref:        prefix.EBCDIC.LLL,
			}),
			56: field.NewString(&field.Spec{
				Length:      37,
				Description: "CONTA DE PAGAMENTO",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
			}),
			58: field.NewString(&field.Spec{
				Length:      60,
				Description: "DE58",
				Enc:         encoding.Binary,
				Pref:        prefix.EBCDIC.LLL,
			}),
			60: field.NewString(&field.Spec{
				Length:      999,
				Description: "DE60",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
			}),
			61: field.NewString(&field.Spec{
				Length:      999,
				Description: "Point-of-Service [POS] Data",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
			}),
			62: field.NewString(&field.Spec{
				Length:      999,
				Description: "Dados para Identificar",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
			}),
			63: field.NewString(&field.Spec{
				Length:      9,
				Description: "DADOS DE REDE",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
				Pad:         padding.Right(' '),
			}),
			90: field.NewString(&field.Spec{
				Length:      42,
				Description: "DADOS CANCELAMENTO",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				//Pad:         padding.Left('0'),
			}),
			95: field.NewString(&field.Spec{
				Length:      42,
				Description: "REPLACEMENTS AMOUNTS",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			112: field.NewString(&field.Spec{
				Length:      999,
				Description: "ADD ADC - Usado no Parcelamento",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
			}),
			120: field.NewString(&field.Spec{
				Length:      999,
				Description: "DADOS COMPLEMENTARES TRANSACAO",
				Enc:         encoding.Binary,
				Pref:        prefix.EBCDIC.LLL,
			}),
			126: field.NewString(&field.Spec{
				Length:      999,
				Description: "Identificador do Cartao -CVE2",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
			}),
			127: field.NewString(&field.Spec{
				Length:      999,
				Description: "DE127",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
			}),
		},
	}

}
