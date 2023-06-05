# Clean API Genesis

Essa é uma refatoração de um teste técnico que eu ja havia feito, segue o link deste projeto: https://github.com/NathanOnCodes/api-rest-genesis
<br>
<br>
Eu refatorei usando Inversão de Controle, Injeção de dependência e outras boas práticas do SOLID.
<br>
Apliquei alguns conceitos de clean architecture para a separação de pastas, mas não segui a risca.

## Requisitos

- Docker

## Como executar
1. Crie um pasta chamada api_rest, abra o vscode e abra essa pasta
```sh
mkdir api_rest
```
2. Clone o repositório:

```sh
git clone https://github.com/SEU_USUARIO/clean-api-genesis.git
```

3. Entre na pasta: 

```sh
cd api_rest
```

4. Construa uma imagem Docker para o mongodb: 

```sh
docker run --name database_mongodb -d -p 27017:27017 mongo
```

# Endpoints Post /exchange/:amount/:from/:to/:rate

Converte uma quantidade de uma moeda para outra, com base em uma taxa de conversão.

Parâmetros:
- amount: valor a ser convertido (número decimal)
- from: código da moeda de origem (string)
- to: código da moeda de destino (string)
- rate: taxa de conversão (número decimal)
<br>

# Outros Endpoints:
GET /exchange - Exibe mensagem de boas-vindas e instruções de como utilizar a API.
GET /exchange/logs - Exibe o histórico de conversões realizadas.
### Licença MIT
