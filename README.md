# leilao

## Boa tarde, pessoal!

Segue o desafio técnico do módulo internals, eu não fiz o clone do projeto como vocês pediram porque eu fui 
acompanhado as aulas e eu mesmo fui impelentando o projeto. 

## A função para calcular a duração do leilão está localizado na: 
* linha 35 do arquivo internal/infra/database/auction/create_auction.go

A verificação já acontecia no FindAuctionId, lá ele busca o auction e depois verificamos se ele ainda está válido, 
caso já tenha passado o tempo ele irá atualizar o auction e retornar internal_error.NewInternalServerError. Essa verificação da
duração servirá também para a criação do bid, já que o mesmo precisa e também faz a consulta do auction pelo FindAuctionId.

A duração do auction está nas variáveis de ambiente com a chave e value:
* AUCTION_DURATION=20s



## Para executar o projeto, rode o docker compose com
* docker compose up -d
ou 
* docker-compose up -d