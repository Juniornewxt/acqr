package especificacao

import (
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/padding"
	"github.com/moov-io/iso8583/prefix"
	"github.com/moov-io/iso8583/sort"
)

func NewSpecASCII() *iso8583.MessageSpec {
	return &iso8583.MessageSpec{
		Fields: map[int]field.Field{
			0: field.NewString(&field.Spec{
				Length:      4,
				Description: "MTI",
				Enc:         encoding.Binary,
				Pref:        prefix.ASCII.Fixed,
				//Pad:         padding.Left('0'),
			}),
			1: field.NewBitmap(&field.Spec{
				Description: "Bitmap",
				//Enc:         encoding.Binary,
				Enc:  encoding.BytesToASCIIHex,
				Pref: prefix.Hex.Fixed,
			}),
			2: field.NewString(&field.Spec{
				Length:      16,
				Description: "PAN - CARTAO",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LL,
			}),
			3: field.NewNumeric(&field.Spec{
				Length:      6,
				Description: "CODIGO PROCESSAMENTO",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			4: field.NewNumeric(&field.Spec{
				Length:      12,
				Description: "VALOR DA TRANSACAO",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			7: field.NewNumeric(&field.Spec{
				Length:      10,
				Description: "DATA HORA GMT TRANSMISSAO",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			11: field.NewNumeric(&field.Spec{
				Length:      6,
				Description: "DE011",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			12: field.NewNumeric(&field.Spec{
				Length:      6,
				Description: "HORA LOCAL TRANSACAO",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			13: field.NewNumeric(&field.Spec{
				Length:      4,
				Description: "DATA LOCAL TRANSACAO",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			14: field.NewNumeric(&field.Spec{
				Length:      4,
				Description: "VALIDADE",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			18: field.NewNumeric(&field.Spec{
				Length:      4,
				Description: "MCC",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			19: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "DE19",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			22: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "ENTRY MODE",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			23: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "DE23",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			24: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "DE24",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			32: field.NewString(&field.Spec{
				Length:      11,
				Description: "CODIGO ADQUIRENTE",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LL,
			}),
			33: field.NewString(&field.Spec{
				Length:      6,
				Description: "Forwarding Institution ID Code",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LL,
				Pad:         padding.Left('0'),
			}),
			35: field.NewString(&field.Spec{
				Length:      37,
				Description: "TRILHA 2",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LL,
			}),
			37: field.NewString(&field.Spec{
				Length:      12,
				Description: "NSU",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			38: field.NewNumeric(&field.Spec{
				Length:      6,
				Description: "CODIGO DE AUTORIZACAO",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			39: field.NewString(&field.Spec{
				Length:      2,
				Description: "CODIGO RESPOSTA",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			41: field.NewString(&field.Spec{
				Length:      8,
				Description: "TERMINAL",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			42: field.NewString(&field.Spec{
				Length:      15,
				Description: "CODIGO DO EC",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			43: field.NewString(&field.Spec{
				Length:      40,
				Description: "DADOS EC",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			46: field.NewString(&field.Spec{
				Length:      999,
				Description: "Identificador do Cartao -CVE2",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LLL,
			}),
			48: field.NewString(&field.Spec{
				Length:      999,
				Description: "Informacoes adicionais de solic resp",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LLL,
			}),
			//48: field.NewComposite(&field.Spec{
			//	Length:      999,
			//	Description: "INFO ADD",
			//	Pref:        prefix.ASCII.LLL,
			//	Tag: &field.TagSpec{

			//		Sort: sort.StringsByInt,
			//		},
			//	Subfields: map[string]field.Field{
			//		"00": field.NewString(&field.Spec{
			//			Length:      1,
			//			Description: "TCC",
			//			Enc:         encoding.ASCII,
			//				Pref:        prefix.ASCII.Fixed,
			//			}),
			//		"01": field.NewComposite(&field.Spec{
			//			Length:      999,
			//			Description: "Composite",
			//			Pref:        prefix.ASCII.Fixed,
			//			Tag: &field.TagSpec{
			//				Enc:  encoding.ASCII,
			//				Sort: sort.StringsByInt,
			//			},
			//			Subfields: map[string]field.Field{
			//				"14": field.NewString(&field.Spec{
			//					Length:      5,
			//					Description: "ATI",
			//					Enc:         encoding.ASCII,
			//					Pref:        prefix.ASCII.LL,
			//				}),
			//				"20": field.NewString(&field.Spec{
			//					Length:      1,
			//					Description: "Cardholder Verification Method",
			//					Enc:         encoding.ASCII,
			//					Pref:        prefix.ASCII.LL,
			//				}),
			//				"22": field.NewString(&field.Spec{
			//					Length:      34,
			//					Description: "INSTALL",
			//					Enc:         encoding.ASCII,
			//					Pref:        prefix.ASCII.LL,
			//				}),
			//				"74": field.NewString(&field.Spec{
			//					Description: "API",
			//					Length:      3,
			//					Enc:         encoding.ASCII,
			//					Pref:        prefix.ASCII.LL,
			//				}),
			///				"80": field.NewString(&field.Spec{
			//				Length:      2,
			//				Description: "PSC",
			//				Enc:         encoding.ASCII,
			//				Pref:        prefix.ASCII.LL,
			//			}),
			//			"91": field.NewString(&field.Spec{
			//				Length:      15,
			//				Description: "ARD",
			//				Enc:         encoding.ASCII,
			//				Pref:        prefix.ASCII.LL,
			//			}),
			//		},
			//	}),
			//},
			//	}),
			49: field.NewNumeric(&field.Spec{
				Length:      3,
				Description: "CODIGO DE MOEDA",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
				Pad:         padding.Left('0'),
			}),
			52: field.NewString(&field.Spec{
				Length:      16,
				Description: "SENHA CRIPTOGRAFADA",
				Enc:         encoding.Binary,
				Pref:        prefix.ASCII.Fixed,
			}),
			54: field.NewString(&field.Spec{
				Length:      999,
				Description: "VALORES ADICIONAIS",
				Enc:         encoding.Binary,
				Pref:        prefix.ASCII.LLL,
			}),
			55: field.NewString(&field.Spec{
				Length:      999,
				Description: "DADOS CHIP",
				Enc:         encoding.Binary,
				Pref:        prefix.ASCII.LLL,
			}),
			51: field.NewComposite(&field.Spec{
				Length:      255,
				Description: "ICC Data – EMV Having Multiple Tags",
				Pref:        prefix.ASCII.LLL,
				Tag: &field.TagSpec{
					Enc:  encoding.BerTLVTag,
					Sort: sort.StringsByHex,
				},
				Subfields: map[string]field.Field{
					"9F27": field.NewString(&field.Spec{
						Description: "Criptograma do Aplicativo (AC)",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F26": field.NewString(&field.Spec{
						Description: "Dados de Informações do Criptograma",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F10": field.NewString(&field.Spec{
						Description: "Dados do Aplicativo do Emissor (IAD)",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F37": field.NewString(&field.Spec{
						Description: "Número Imprevisível",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F36": field.NewString(&field.Spec{
						Description: "Contador de Transações do Aplicativo ",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"95": field.NewString(&field.Spec{
						Description: "Resultado da Verificação do Terminal(TVR)",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9A": field.NewString(&field.Spec{
						Description: "Data da Transação",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9C": field.NewString(&field.Spec{
						Description: "Tipo de Transação",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F02": field.NewString(&field.Spec{
						Description: "Valor Autorizado",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"5F2A": field.NewString(&field.Spec{
						Description: "Código da Moeda de Transação",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"82": field.NewString(&field.Spec{
						Description: "Perfil de Intercâmbio do Aplicativo",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F1A": field.NewString(&field.Spec{
						Description: "Código de País do Terminal ",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F34": field.NewString(&field.Spec{
						Description: "Resultados do Método de Verificação do Titular do Cartão (CVM)",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F33": field.NewString(&field.Spec{
						Description: "Capacidade do Terminal - Somente EUA",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"84": field.NewString(&field.Spec{
						Description: "Nome do Arquivo Dedicado",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F03": field.NewString(&field.Spec{
						Description: "Outro Valor",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					///Opcionais abaixo
					"5F34": field.NewString(&field.Spec{
						Description: "Número da Sequência do Número do Cartão (PAN) do Aplicativo",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F0A": field.NewString(&field.Spec{
						Description: "Dados de Propriedade Cadastrados de Seleção do Aplicativo",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F35": field.NewString(&field.Spec{
						Description: "Tipo de Terminal ",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F1E": field.NewString(&field.Spec{
						Description: "Número de Série do Dispositivo da Interface (IFD)",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F53": field.NewString(&field.Spec{
						Description: "Código de Categoria da Transação",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F09": field.NewString(&field.Spec{
						Description: "Número da Versão do Aplicativo ",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					"9F41": field.NewString(&field.Spec{
						Description: "Contador da Sequência da Transação",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
					//	"9F33": field.NewString(&field.Spec{
					//		Description: "Capacidade do Terminal - Fora dos EUA",
					//		Enc:         encoding.Binary,
					//		Pref:        prefix.BerTLV,
					//	}),
					"9F6E": field.NewString(&field.Spec{
						Description: "Dados de Terceiros",
						Enc:         encoding.Binary,
						Pref:        prefix.BerTLV,
					}),
				},
			}),
			58: field.NewString(&field.Spec{
				Length:      60,
				Description: "DE58",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LLL,
			}),
			60: field.NewString(&field.Spec{
				Length:      999,
				Description: "DE60",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LLL,
			}),
			61: field.NewString(&field.Spec{
				Length:      999,
				Description: "Point-of-Service [POS] Data",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LLL,
			}),
			611: field.NewComposite(&field.Spec{
				Length:      999,
				Description: "Point-of-Service [POS] Data",
				Pref:        prefix.ASCII.LLL,
				Tag: &field.TagSpec{
					Length: 3,
					Enc:    encoding.ASCII,
					Sort:   sort.StringsByInt,
				},
				Subfields: map[string]field.Field{
					"01": field.NewString(&field.Spec{
						Length:      8,
						Description: "DADOS APLICACAO DA LOJA",
						Pref:        prefix.ASCII.LL,
						Enc:         encoding.ASCII,
					}),
					//	"02": field.NewString(&field.Spec{
					//		Length:      20,
					//		Description: "NUM SERIE PINPAD",
					//		Pref:        prefix.ASCII.LLL,
					//		Enc:         encoding.ASCII,
					//	}),
					//	"03": field.NewString(&field.Spec{
					//		Length:      26,
					//		Description: "DADOS PARA AS BANDEIRAS",
					//		Pref:        prefix.ASCII.Fixed,
					//		Enc:         encoding.ASCII,
					//	}),
					//	"04": field.NewString(&field.Spec{
					//		Length:      16,
					//		Description: "VERSAO DA APLICACAO",
					//		Pref:        prefix.ASCII.LLL,
					//		Enc:         encoding.ASCII,
					//	}),
					//"05": field.NewString(&field.Spec{
					//		Length:      20,
					//		Description: "FABRICANTE PINPAD",
					//		Pref:        prefix.ASCII.LLL,
					//		Enc:         encoding.ASCII,
					//	}),
					//	"06": field.NewString(&field.Spec{
					//		Length:      20,
					//	Description: "VERSAO HARDWARE",
					//	Pref:        prefix.ASCII.LLL,
					//	Enc:         encoding.ASCII,
					//}),
					//	"07": field.NewString(&field.Spec{
					//	Length:      20,
					//	Description: "FIRWARE PINPAD",
					//	Pref:        prefix.ASCII.LLL,
					//	Enc:         encoding.ASCII,
					//}),
				},
			}),
			62: field.NewString(&field.Spec{
				Length:      999,
				Description: "Dados para Identificar",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LLL,
			}),
			//	124: field.NewString(&field.Spec{
			//		Length:      124,
			//		Description: "DE124",
			//		Enc:         encoding.ASCII,
			//		Pref:        prefix.ASCII.LLL,
			//	}),
			90: field.NewString(&field.Spec{
				Length:      42,
				Description: "DADOS CANCELAMENTO",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.Fixed,
			}),
			112: field.NewString(&field.Spec{
				Length:      999,
				Description: "ADD ADC - Usado no Parcelamento",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LLL,
			}),
			120: field.NewString(&field.Spec{
				Length:      999,
				Description: "DADOS COMPLEMENTARES TRANSACAO",
				Enc:         encoding.Binary,
				Pref:        prefix.ASCII.LLL,
			}),
			126: field.NewString(&field.Spec{
				Length:      999,
				Description: "Identificador do Cartao -CVE2",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LLL,
			}),
			127: field.NewString(&field.Spec{
				Length:      999,
				Description: "DE127",
				Enc:         encoding.ASCII,
				Pref:        prefix.ASCII.LLL,
			}),
		},
	}

}
