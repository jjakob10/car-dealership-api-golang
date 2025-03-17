# Car Dealership API

Esta é uma API simples para gerenciar um inventário de carros em uma concessionária. A API permite adicionar, atualizar, deletar e obter informações sobre carros.

## Endpoints

### Obter todos os carros

- **URL:** `/`
- **Método:** `GET`
- **Resposta de Sucesso:**
  - **Código:** `200 OK`
  - **Conteúdo:** Lista de carros em formato JSON

### Obter carro pelo ID

- **URL:** `/car`
- **Método:** `GET`
- **Parâmetros de URL:** `id` (string) - ID do carro
- **Resposta de Sucesso:**
  - **Código:** `200 OK`
  - **Conteúdo:** Carro em formato JSON

### Adicionar carros

- **URL:** `/add`
- **Método:** `POST`
- **Corpo da Requisição:** Lista de carros em formato JSON
- **Resposta de Sucesso:**
  - **Código:** `201 Created`
  - **Conteúdo:** Mensagem de sucesso

### Atualizar carro

- **URL:** `/update`
- **Método:** `PUT`
- **Corpo da Requisição:** Carro em formato JSON
- **Resposta de Sucesso:**
  - **Código:** `200 OK`
  - **Conteúdo:** Mensagem de sucesso

### Deletar carro

- **URL:** `/delete`
- **Método:** `DELETE`
- **Parâmetros de URL:** `id` (string) - ID do carro
- **Resposta de Sucesso:**
  - **Código:** `200 OK`
  - **Conteúdo:** Mensagem de sucesso

## Como Executar

1. Certifique-se de ter o [Go](https://golang.org/dl/) instalado.
2. Clone este repositório.
3. Navegue até o diretório do projeto.
4. Execute `go run main.go` para iniciar o servidor.

O servidor estará ouvindo na porta `8080`.

## Estrutura do Projeto

- `main.go`: Contém os handlers para os endpoints da API.
- `cars.go`: Contém as funções para manipulação dos dados dos carros.
- `cars.json`: Arquivo JSON que armazena os dados dos carros.

## Dependências

- Nenhuma dependência externa além da biblioteca padrão do Go.

## Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
