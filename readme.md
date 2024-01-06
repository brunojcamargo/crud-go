
# CRUD com GO

CRUD em memória com GO desenvolvido em meu primeiro contato com a linguagem.


## Dependência

[Go 1.21.5](https://go.dev/dl/)


## Rodando localmente

Clone o projeto

```bash
  git clone https://github.com/brunojcamargo/crud-go
```

Entre no diretório do projeto

```bash
  cd crud-go
```

Instale as dependências

```bash
  go mod tidy
```

Inicie o servidor

```bash
  go run main.go
```


## Documentação da API

#### Retorna todos os tweets

```http
  GET http://localhost:3001/tweets
```

#### Adicionar um tweet

```http
  POST http://localhost:3001/tweet
  Content-Type: application/json
  {
    "description": "Texto da descrição aqui"
  }
```

#### Excluir um tweet

```http
  DELETE http://localhost:3001/tweet/${id}
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. id do tweet |

#### Editar um tweet

```http
  PUT http://localhost:3001/tweet/${id}
  Content-Type: application/json
  {
    "description": "Texto da nova descrição aqui"
  }
```

| Parâmetro   | Tipo       | Descrição                                   |
| :---------- | :--------- | :------------------------------------------ |
| `id`      | `string` | **Obrigatório**. id do tweet |
