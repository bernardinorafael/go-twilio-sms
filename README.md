# Twilio Verification Service

Este projeto implementa um serviço simples para enviar e verificar códigos de verificação via SMS usando a API da Twilio.

## Requisitos

- Go 1.16 ou superior
- Conta Twilio com serviço de verificação configurado
- Arquivo `.env` com as variáveis de ambiente necessárias

## Instalação

1. Clone o repositório:

   ```sh
   git clone https://github.com/seu-usuario/go-twilio-sms.git
   cd twilio-verification-service
   ```

2. Instale as dependências:

   ```sh
   go mod tidy
   ```

3. Crie uma conta Twilio e configure o serviço de verificação com suporte para envio de SMS.

- Acesse [Twilio](https://www.twilio.com/) e crie uma conta.
- Após criar a conta, navegue até a seção de serviços de verificação (Verify Services).
- Crie um novo serviço de verificação e habilite o canal SMS.

4. Crie um arquivo `.env` na raiz do projeto e utilize o arquivo `.env.example` como base.

5. Execute o projeto/

   ```sh
   go run cmd/main.go
   ```
