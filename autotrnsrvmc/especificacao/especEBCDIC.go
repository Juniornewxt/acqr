package especificacao

import (
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/encoding"
	"github.com/moov-io/iso8583/field"
	"github.com/moov-io/iso8583/padding"
	"github.com/moov-io/iso8583/prefix"
	"github.com/moov-io/iso8583/sort"
)

func NewSpecEBCDIC() *iso8583.MessageSpec {
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
				Pad:         padding.Left('0'),
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
				Description: "CODIGO RESPOSTA",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
				Pad:         padding.Left('0'),
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
			477: field.NewComposite(&field.Spec{
				Length:      999,
				Description: "Informacoes adicionais de solic resp",
				Pref:        prefix.EBCDIC.LLL,
				Tag: &field.TagSpec{
					//Length: 2,
					//Enc:  encoding.EBCDIC,
					Sort: sort.StringsByInt,
				},
				Subfields: map[string]field.Field{
					"00": field.NewString(&field.Spec{
						Length:      1,
						Description: "TCC",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"14": field.NewString(&field.Spec{
						Length:      5,
						Description: "ATI",
						Pref:        prefix.EBCDIC.LL,
						Enc:         encoding.EBCDIC,
					}),
					"74": field.NewString(&field.Spec{
						Length:      3,
						Description: "API",
						Pref:        prefix.EBCDIC.LL,
						Enc:         encoding.EBCDIC,
					}),
					"80": field.NewString(&field.Spec{
						Length:      2,
						Description: "PSC",
						Pref:        prefix.EBCDIC.LL,
						Enc:         encoding.EBCDIC,
					}),
				},
			}),

			48: field.NewComposite(&field.Spec{
				Length:      999,
				Description: "INFO ADD",
				Pref:        prefix.EBCDIC.LLL,
				Tag: &field.TagSpec{
					//Length: 3,
					//Enc:  encoding.EBCDIC,
					Sort: sort.StringsByInt,
				},
				Subfields: map[string]field.Field{
					"00": field.NewString(&field.Spec{
						Length:      1,
						Description: "TCC",
						Enc:         encoding.EBCDIC,
						Pref:        prefix.EBCDIC.Fixed,
					}),
					"01": field.NewComposite(&field.Spec{
						Length:      999,
						Description: "Composite",
						Pref:        prefix.EBCDIC.Fixed,
						Tag: &field.TagSpec{
							Enc:  encoding.EBCDIC,
							Sort: sort.StringsByInt,
						},
						Subfields: map[string]field.Field{
							"14": field.NewString(&field.Spec{
								Length:      5,
								Description: "ATI",
								Enc:         encoding.EBCDIC,
								Pref:        prefix.EBCDIC.LL,
							}),
							"20": field.NewString(&field.Spec{
								Length:      1,
								Description: "Cardholder Verification Method",
								Enc:         encoding.EBCDIC,
								Pref:        prefix.EBCDIC.LL,
							}),
							"22": field.NewString(&field.Spec{
								Length:      34,
								Description: "INSTALL",
								Enc:         encoding.EBCDIC,
								Pref:        prefix.EBCDIC.LL,
							}),
							"26": field.NewString(&field.Spec{
								Length:      3,
								Description: "INDENTIFICADOR DE CARTEIRA",
								Enc:         encoding.EBCDIC,
								Pref:        prefix.EBCDIC.LL,
							}),
							"29": field.NewString(&field.Spec{
								Length:      1,
								Description: "Info Add Localizacao do POS",
								Enc:         encoding.EBCDIC,
								Pref:        prefix.EBCDIC.LL,
							}),
							"74": field.NewString(&field.Spec{
								Description: "API",
								Length:      3,
								Enc:         encoding.EBCDIC,
								Pref:        prefix.EBCDIC.LL,
							}),
							"77": field.NewString(&field.Spec{
								Description: "PAYMENT TRANSACION INDICATOR",
								Length:      3,
								Enc:         encoding.EBCDIC,
								Pref:        prefix.EBCDIC.LL,
							}),
							"80": field.NewString(&field.Spec{
								Length:      2,
								Description: "PSC",
								Enc:         encoding.EBCDIC,
								Pref:        prefix.EBCDIC.LL,
							}),
							"91": field.NewString(&field.Spec{
								Length:      15,
								Description: "ARD",
								Enc:         encoding.EBCDIC,
								Pref:        prefix.EBCDIC.LL,
							}),
							"95": field.NewString(&field.Spec{
								Length:      6,
								Description: "Código de Promoção da Mastercard",
								Enc:         encoding.EBCDIC,
								Pref:        prefix.EBCDIC.LL,
							}),
						},
					}),
					//	"95": field.NewString(&field.Spec{
					//	Length:      6,
					//	Description: "Código de Promoção da Mastercard",
					//	Enc:         encoding.EBCDIC,
					//	Pref:        prefix.EBCDIC.LLL,
					//}),
				},
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
				//Pref: prefix.Hex.Fixed,
			}),
			551: field.NewString(&field.Spec{
				Length:      999,
				Description: "ICC Data – EMV Having Multiple Tags",
				Enc:         encoding.Binary,
				Pref:        prefix.EBCDIC.LLL,
			}),
			55: field.NewComposite(&field.Spec{
				Length:      255,
				Description: "ICC Data – EMV Having Multiple Tags",
				Pref:        prefix.EBCDIC.LLL,
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
				Description: "ADVICE REASON CODE",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
			}),
			61: field.NewComposite(&field.Spec{
				Length:      26,
				Description: "Point-of-Service [POS] Data",
				Pref:        prefix.EBCDIC.LLL,
				Tag: &field.TagSpec{
					//Length: 3,
					//Enc:  encoding.EBCDIC, //comentado nao informa o campo
					Sort: sort.StringsByInt,
				},
				Subfields: map[string]field.Field{
					"01": field.NewString(&field.Spec{
						Length:      1,
						Description: "Point of Service Terminal Attendance",
						Pref:        prefix.EBCDIC.Fixed, //se informado tlv colocar com LL sem fixo nao vai o tamanho do campo
						Enc:         encoding.EBCDIC,
					}),
					"02": field.NewString(&field.Spec{
						Length:      1,
						Description: "Reservado",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"03": field.NewString(&field.Spec{
						Length:      1,
						Description: "Point of Service Terminal Location",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"04": field.NewString(&field.Spec{
						Length:      1,
						Description: "Point of Service Cardholder Presence",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"05": field.NewString(&field.Spec{
						Length:      1,
						Description: "Point of Service Card Presence",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"06": field.NewString(&field.Spec{
						Length:      1,
						Description: "Point of Service Card Capture Capabilities",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"07": field.NewString(&field.Spec{
						Length:      1,
						Description: "Point of Service Transaction Status",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"08": field.NewString(&field.Spec{
						Length:      1,
						Description: "POS Transaction security",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"09": field.NewString(&field.Spec{
						Length:      1,
						Description: "Reversado",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"10": field.NewString(&field.Spec{
						Length:      1,
						Description: "Cardholder-Activated Terminal (CAT) Level",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"11": field.NewString(&field.Spec{
						Length:      1,
						Description: "Point of Service Card Data Terminal Input Capability Indicator",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"12": field.NewString(&field.Spec{
						Length:      2,
						Description: "Point of Service Authorization Life Cycle",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"13": field.NewString(&field.Spec{
						Length:      3,
						Description: "Point of Service Country Code",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"14": field.NewString(&field.Spec{
						Length:      10,
						Description: "Point of Service Postal Code",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
				},
			}),
			611: field.NewString(&field.Spec{
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
				Length:      999,
				Description: "DADOS DE REDE",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.LLL,
			}),
			//	124: field.NewString(&field.Spec{
			//		Length:      124,
			//		Description: "DE124",
			//		Enc:         encoding.ASCII,
			//		Pref:        prefix.ASCII.LLL,
			//	}),
			901: field.NewString(&field.Spec{
				Length:      42,
				Description: "DADOS CANCELAMENTO",
				Enc:         encoding.EBCDIC,
				Pref:        prefix.EBCDIC.Fixed,
			}),
			90: field.NewComposite(&field.Spec{
				Length:      42,
				Description: "DADOS CANCELAMENTO",
				Pref:        prefix.EBCDIC.Fixed,
				Tag: &field.TagSpec{
					//Length: 3,
					//Enc:  encoding.EBCDIC, //comentado nao informa o campo
					Sort: sort.StringsByInt,
				},
				Subfields: map[string]field.Field{
					"01": field.NewString(&field.Spec{
						Length:      4,
						Description: "MTI TRN ORIGINAL",
						Pref:        prefix.EBCDIC.Fixed, //se informado tlv colocar com LL sem fixo nao vai o tamanho do campo
						Enc:         encoding.EBCDIC,
					}),
					"02": field.NewString(&field.Spec{
						Length:      6,
						Description: "STAN TRN ORIGINAL",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"03": field.NewString(&field.Spec{
						Length:      10,
						Description: "MMDDhhmms TRN ORIGINAL",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
					}),
					"04": field.NewString(&field.Spec{
						Length:      11,
						Description: "ACQR TRN ORIGINAL",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
						Pad:         padding.Left('0'),
					}),
					"05": field.NewString(&field.Spec{
						Length:      11,
						Description: "COD. INSTIT. ENC. TRN ORIGINAL",
						Pref:        prefix.EBCDIC.Fixed,
						Enc:         encoding.EBCDIC,
						Pad:         padding.Left('0'),
					}),
				},
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
