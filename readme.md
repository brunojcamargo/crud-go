
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

```cURL
  curl --location 'http://localhost:3001/tweets'
```

#### Adicionar um tweet

```cURL
  curl --location --request POST 'http://localhost:3001/tweet' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "description": "teste"
    }'
```
| Parâmetro | Tipo     | Descrição                    |
| :-------- | :------- | :--------------------------- |
| `description`| `string` | descrição para o tweet |

#### Excluir um tweet

```cURL
  curl --location --request DELETE 'http://localhost:3001/tweet/${id}'
```

| Parâmetro | Tipo     | Descrição                    |
| :-------- | :------- | :--------------------------- |
| `id`      | `string` | **Obrigatório**. id do tweet |

#### Editar um tweet

```cURL
  curl --location --request PUT 'http://localhost:3001/tweet/${id}' \
    --header 'Content-Type: application/json' \
    --data '{
        "description":"Texto da nova descrição aqui"
    }'
```

| Parâmetro | Tipo     | Descrição                    |
| :-------- | :------- | :--------------------------- |
| `id`      | `string` | **Obrigatório**. id do tweet |
| `description`| `string` | **Obrigatório**. texto para a nova descrição |
