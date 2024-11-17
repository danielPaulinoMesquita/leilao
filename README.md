# leilao

## Boa tarde, pessoal!

Segue o desafio técnico do módulo internals, eu não fiz o clone do projeto como vocês pediram porque eu fui 
acompanhado as aulas e eu mesmo fui implementando o projeto. 

## A função para calcular a duração do leilão está localizado na: 
* linha 35 do arquivo internal/infra/database/auction/create_auction.go

A verificação já acontecia no FindAuctionId, lá ele busca o auction e depois verifica se ele ainda está válido, 
caso já tenha passado o tempo ele irá atualizar o auction e retornar internal_error.NewInternalServerError. Essa verificação da
duração servirá também para a criação do bid, já que o mesmo precisa e também faz a consulta do auction pelo FindAuctionId.

A duração do auction está nas variáveis de ambiente com a chave e value:
* AUCTION_DURATION=20s

## Para executar o projeto, rode o docker compose com
* docker compose up -d
ou 
* docker-compose up -d

### Para fazer um teste, você deve criar uma auction:

* Você pode realizar os testes também pelo teste.http: cmd/test/test.http

POST: localhost:8080/auctions

and body:
{
"product_name": "TERCEIRO Objeto",
"category": "BICICLETA",
"description": "Daniel",
"condition": 0
}

### Depois fazer a consulta de todas as auctions para verificar as que foram criadas, usando GET:
GET: localhost:8080/auctions?status=0&category&productName

a resposta será parecida com essa:
[
{
"id": "473e3cbb-1bcb-4482-bd88-55cb6596b3aa",
"product_name": "TERCEIRO Objeto",
"category": "BICICLETA",
"description": "teste de brasília",
"condition": 0,
"status": 0,
"timestamp": "2024-11-11T00:53:26-03:00"
}
]

### Com id em mãos, faça a consulta para verificar e atualizar a auction caso precise, então:
GET: localhost:8080/auctions/{id}

caso a auction esteja desatualizada ela irá retornar
{
"message": "Auction has expired",
"error": "internal_server_error",
"code": 500,
"Cause": null
}

ou a auction em questão. 

### Unit Test:

* Criado um teste unitário para verificar a função criada, para isso você pode executar o comando para ir até a pasta:
- cd internal/infra/database/auction/
depois executar
- go test -v