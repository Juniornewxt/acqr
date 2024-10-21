# Projeto ACQR

![Projeto ACQR](assets/slogan_projeto_acqr.png)

**Projeto ACQR** é um sistema em desenvolvimento para gerenciamento de adquirentes de cartão de crédito, totalmente implementado em Golang. O objetivo é criar uma solução completa para o processamento de transações, com suporte a múltiplas bandeiras de cartões como Visa, Mastercard, ELO, e outras.

## 🚧 Status do Projeto

O Projeto ACQR ainda está em desenvolvimento e, até o momento, conta com as seguintes funcionalidades:

- **Autorizador de Compras para Mastercard**: Implementação inicial para o processamento de transações Mastercard.
  - **Aplicativo de Tratamento de Mensagens**: Recebe e processa mensagens de autorização.
  - **Aplicativo de Conexão com Bandeiras**: Mantém a comunicação com a bandeira para encaminhar as mensagens de transações e responder as mensagens administrativas.

## 🔧 Estrutura do Sistema de Autorização

O sistema de autorização será composto por três serviços principais:

1. **Recepcionista**: Recebe transações de POS, entre outros canais e direciona para o autorizador da bandeira correspondente via gRPC, também é o responsavel por enviar as mensagens para outros satelites, alem do do HSM.
2. **Mensageiro**: Processa as mensagens recebidas antes de enviá-las para o aplicativo de conexão com a bandeira. Suporta protocolos ISO e gRPC, faz algumas validaçoes e pode negar a transaçao sem que seja encaminhado a bandeira.
3. **Administrador**: Responsável por comunicar-se diretamente com a bandeira, ele quem trata as mensagens 0800 de rede, e troca de chave, gravando a chave no redis para uso das aplicaçoes acima, e gerando fila no kafka para cada movimento.

## 🥷 Colocando a mão na massa

1. Você deve configurar o arquivo .env corresponde do serviço, por padrao kafka e redis, estao desabilitados, se deseja usar, deve configurar os parametros no .env
   Os parametros KAFKA_ENVIO e REDIS_ENVIO devem estar = true, caso contrario deixe false se só quer testar, nao esqueça de configurar as portas de entrada e saida.
   
    Exemplo: "autoadmsrvmc"
```{#define configurações de ambiente 

#ip e porta de conexão de entrada e saída
TCPMCADMEIN=localhost:8080 #porta de entrada
TCPMCADMBAN=mcserver:1002 #porta de saida

#timeout comunicacao entre servidores.
MCTIMEOUT=3 #timeout bandeira Mastercard
CLIENTTIMOUT=3 #timeout cliente

#envio de mensagem para kafka
KAFKA_ENVIO=true
KAFKA_SERV=localhost:9092
KAFKA_SECP=SASL_PLAINTEXT
KAFKA_SASL=PLAIN
KAFKA_USER=usuario
KAFKA_PASS=usuario-secret
TOPIC_TRN=MENSAGENS-TRN-AUTO
TOPIC_ADM=MENSAGENS-ADM-AUTO

#envio de mensagem para redis
REDIS_ENVIO=true
REDIS_SERV=localhost:6379
REDIS_USER=usuario
REDIS_PASS=usuario-secret
REDIS_DB=0

```

2. Você precisa executar os serviços autoadmsrvmc e autotrnsrvmc
  ``` autotrnsrvmc                                                
2024/10/19 13:41:09 Iniciando Servidor MC v1...
2024/10/19 13:41:09 Iniciando servidor TCP...
2024/10/19 13:41:09 Servidor aguardando conexões localhost:8082...

  autoadmsrvmc
2024/10/19 17:55:50 Iniciando Servidor MC Admninstrativo v1...
2024/10/19 17:55:50 Conectado ao servidor da bandeira: 192.168.1.214:1002
2024/10/19 17:55:50 Escutando na porta localhost:8080
```
autotrnsrvmc deve ser conectado a porta, o qual subiu o autoadmsrvmc, que por sua vez deve estar conectado a bandeira.
você tambem pode conectar o serviço autotrnsrvmc direto ao servidor da bandeira, porem nao terá o tratamento de mensagens 0800, nem ficará conectado o tempo todo, mas é util se você quiser testar alguma mensagem, visto que tirando a 0800, ele quem tem a logica de tratamento das mensagens MC.

**NOTA: o servidor da Mastercard deve estar preparado para receber mensagens em EBCDIC**

3. Por padrao autotrnsrvmc esta configurado para receber mensagens ISO de entrada, mas você pode mudar no arquivo .env o parametro TIPO_SERVER
 ```  #define se o servidor vai receber mensagens via tcp ou grcp
TIPO_SERVER=tcp  # Pode ser "grpc" ou "tcp"
```
Se for usar o grpc aqui esta um exemplo de proto que você vai precisar carregar.
```protobuf
syntax = "proto3";

package transacao;

option go_package = "/proto;transacao";

message TransacaoRequest {
    string MTI = 1;  // "0200", que seria o MTI (Message Type Indicator)
    string DE02 = 2; // PAN (Primary Account Number)
    string DE03 = 3; // Código de processamento
    string DE04 = 4; // Valor da transação
    string DE07 = 5; // Data e hora no formato MMDDhhmmss
    string DE11 = 6; // Número de sistema (trace number)
    string DE12 = 7; // Hora local (hhmmss)
    string DE13 = 8; // Data local (MMDD)
    string DE14 = 9; // Validade do cartão
    string DE18 = 10; // MCC (Merchant Category Code)
    string DE19 = 11; // País
    string DE22 = 12; // Modo de entrada (entry mode)
    string DE23 = 13; // Código de identificação de aplicação
    string DE32 = 14; // Código de identificador do adquirente
    string DE33 = 15; // Forwarding Institution ID Code
    string DE35 = 16; // Trilha 2 do cartão (track 2)
    string DE38 = 17; // CODIGO DE AUTORIZACAO
    string DE39 = 18; // CODIGO RESPOSTA
    string DE41 = 19; // Terminal ID
    string DE42 = 20; // Merchant ID (ID do comércio)
    string DE43 = 21; // Nome e localização do comerciante
    string DE48 = 22; // Dados adicionais (campo 48)
    string DE49 = 23; // Código da moeda
    string DE52 = 24; // Senha (PIN Data)
    string DE55 = 25; // EMV Data (campo 55)
    string DE61 = 26; // Dados adicionais (campo 61)
    string DE120 = 27; // Dados adicionais (campo 120)
    string DE126 = 28; // Dados adicionais (campo 126)
}

message TransacaoResponse {
    string response_code = 1;
    string message = 2;
}

service TransacaoService {
    rpc ProcessTransacao(TransacaoRequest) returns (TransacaoResponse);
}
```
**NOTA: as mensagens via GRPC ainda nao foram 100% testadas, você pode ter problema, por exemplo, o proto nao nao tem o DE112, que pode dar erro em algumas transaçoes como parcelado emissor.**

4. Enviando mensagem, via TCP "ISO8583" formato ASCII Padrao
   Header 2 bytes Binario
     Exemplo: transaçao trila 2 sem senha
```F0   MTI.............................: 0200
F3   CODIGO PROCESSAMENTO............: 3000
F4   VALOR DA TRANSACAO..............: 1000
F7   DATA HORA GMT TRANSMISSAO.......: 1018164059
F11  DE011...........................: 783172
F12  HORA LOCAL TRANSACAO............: 164059
F13  DATA LOCAL TRANSACAO............: 1018
F14  VALIDADE........................: 3112
F18  MCC.............................: 5411
F22  ENTRY MODE......................: 902
F32  CODIGO ADQUIRENTE...............: 00000000025
F33  Forwarding Institution ID Code..: 0026
F35  TRILHA 2........................: 5666****0000=**************
F41  TERMINAL........................: JUNI0SIM
F42  CODIGO DO EC....................: 000001020116592
F43  DADOS EC........................: POSTO DM JR            DIADEMA  EVANG076
F48  DADOS ADD.......................: 001008000000010020010003001F0210022102200212
F49  CODIGO DE MOEDA.................: 986
F61  Point-of-Service [POS] Data.....: ******************************
```
 ``` Formato:
  DE03 Numerico tamanho de 6 fixo
  DE04 Numerico tamanho de 12 fixo
  DE07 Numerico tamanho de 10 fixo
  DE11 Numerico tamanho de 6 fixo
  DE12 Numerico tamanho de 6 fixo
  DE13 Numerico tamanho de 4 fixo
  DE14 Numerico tamanho de 4 fixo
  DE18 Numerico tamanho de 4 fixo
  DE22 Numerico tamanho de 3 fixo
  DE32 String tamanho de 11 LL
  DE33 String tamanho de 6 LL
  DE35 String tamanho de 37 LL
  DE41 String tamanho de 8 fixo
  DE42 String tamanho de 15 fixo
  DE43 String tamanho de 40 fixo
  DE48 String tamanho de 999 LLL
  DE49 Numerico tamanho de 3 fixo
  DE61 String tamanho de 999 LLL
  ```
  Para conseguir o envio de transaçoes parceladas, crediario, consulta, o DE48 precisa ser enviado na estrutura que desenhei, ainda a ser melhor detalhado na documentaçao que vou anexar futuramente.
  Exemplos de DE48
  ```Compra a vista, DE03=003000, DE48=001008000000010020010003001F0210022202200200 ou 001008000000010020010003001F0210022202200201 
  Compra parcelada loja, DE03=003000,    DE48=001008000000010020010003001F0210022102200212 "os 2 ultimos digitos sao a quantidade de parcelas"
  Compra parcelada emissor, DE03=003000, DE48=001008000000010020010003001F0210022002200208 "os 2 ultimos digitos sao a quantidade de parcelas"
  Compra consulta de crediario, DE48=001008000000010020010003001 0210022502200212023012000000200000
  Compra compra crediario, DE48=001008000000010020010003001 0210022502200207 "os 2 ultimos digitos sao a quantidade de parcelas"
  Compra Maestro, DE03=002000, DE48=001008000000010020010003001F021002220220020002401400394460005887
```
  Para tentar facilitar, deixei um programa que simula uma transacao de trilha 2 "com dados mocados" para executar, basta:
  
    POS_asciiTarjaMOCv1 localhost:8082 1000 001008000000010020010003001F0210022202200200
    Sendo os parametros ip:porta valor DE48

  **Se precisar de algo mais sofisticado, que testa varios tipos de transacoes com os dados via arquivo csv, me procure.**

  Se tudo ocorrer bem, você deve receber algo parecido com isso:

```F0   MTI....................................: 0210
F2   PAN - CARTAO...........................: 5666****0000
F3   CODIGO PROCESSAMENTO...................: 3000
F4   VALOR DA TRANSACAO.....................: 1000
F7   DATA HORA GMT TRANSMISSAO..............: 1018164059
F11  DE011..................................: 783172
F12  HORA LOCAL TRANSACAO...................: 164059
F13  DATA LOCAL TRANSACAO...................: 1018
F38  CODIGO DE AUTORIZACAO..................: 428411
F39  CODIGO RESPOSTA........................: 00
F41  TERMINAL...............................: JUNI0SIM
F49  CODIGO DE MOEDA........................: 986
F60  DADOS COMPROVAMENTE DE VENDA CLINTE....: C@    DEBIT MASTERCARD - Via Cliente    @NINJASEC - TESTE111111 E            @RUA XXXX XXXXX C                 SP BR@CNPJ:01000000000100                   @TID: JUNI0SIM@EC:XXXXXXXXX@                    @VENDA CREDITO A VISTA                 @************0000   @18/10/24                         16:40@VALOR APROVADO:               R$ 10.00@CV:XXXXXXXXXXXX            AUTO:000000@DOC:000000@TERM:XXXXXXXX@@               
F62  DADOS COMPROVAMENTE DE VENDA COMERCIO..: C@    DEBIT MASTERCARD - Via Cliente    @NINJASEC - TESTE111111 E            @RUA XXXX XXXXX C                 SP BR@CNPJ:01000000000100                   @TID: JUNI0SIM@EC:XXXXXXXXX@                    @VENDA CREDITO A VISTA                 @************0000   @18/10/24                         16:40@VALOR APROVADO:               R$ 10.00@CV:XXXXXXXXXXXX            AUTO:000000@DOC:000000@TERM:XXXXXXXX@@ MEDIANTE A ASSINATURA               
F127 INDENTIFICADOR UNICO DO ADQUIRENTE.....: 164059805490

 FIM DO PROCESSAMENTO :)
```

## ✈️ Próximos Passos

Os macro próximos passos para o desenvolvimento do projeto incluem:

- Implementar suporte para outras bandeiras de cartões.
- Documentaçao de como montar a mensagem ISO Padrao de entrada.
- Melhorar o processamento de mensagens ISO/gRPC.
- Desenvolver o serviço do Recepcionista que deve consultar tabela de bins e direcionar aos satelites.
- Desenvolver interaç±ao com HSM.
- Desenvolver o serviço de clearing que será responsavel pela troca de arquivos com a bandeira.
- Desenvolver o serviço de gerenciamenteo de comercio, credenciamento, incluindo a parte de serviços como venda, de seguro, aluguel de maquina, sub do sub, entre outros...

## 💬 Contato

Se você estiver interessado em contribuir ou tiver alguma dúvida, sinta-se à vontade para entrar em contato.

---

> **Nota:** Este projeto está em fase inicial de desenvolvimento e funcionalidades estão sendo continuamente adicionadas.
