# Stress Test CLI em Go

Este é um projeto de CLI (Command-Line Interface) em Go que permite realizar testes de estresse em um serviço web.

## Requisitos

- Go versão 1.22
- Docker

## Executando os Testes

Para executar os testes, utilize o seguinte comando:

```
go test ./...
```

## Construindo a Imagem Docker (Opcional)

Para construir a imagem Docker da aplicação, execute o seguinte comando:

```
docker build -t taisso/stress_test .
```

## Executando a Aplicação no Docker

Para executar a aplicação no Docker, utilize o seguinte comando:

```
docker run --rm taisso/stress_test --url=http://localhost:8080 --requests=1000 --concurrency=10
```
> Imagem já publicada no docker hub

## Flags da CLI

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requests.
- `--concurrency`: Número de chamadas simultâneas.