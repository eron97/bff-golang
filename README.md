# BFF E-commerce

Este projeto é um Backend for Frontend (BFF) para um site de e-commerce, implementado em Go. Ele fornece funcionalidades de login e autorização, seguindo uma arquitetura de camadas bem definida. Abaixo, detalhamos a estrutura do projeto, as tecnologias utilizadas, a injeção de dependências e os endpoints disponíveis.

## Arquitetura

O projeto segue a arquitetura de camadas, dividindo responsabilidades em diferentes pacotes:

- **Controller**: Responsável por lidar com as requisições HTTP e chamar os serviços apropriados.
- **Service**: Contém a lógica de negócios e interage com a camada de persistência.
- **Persistence**: Gerencia a comunicação com o banco de dados.
- **DTOs (Data Transfer Objects)**: Define os objetos que são transferidos entre as camadas.
- **Middlewares**: Implementa funcionalidades transversais como validação e autenticação.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação principal.
- **Fiber**: Framework web para Go, utilizado para criar os endpoints HTTP.
- **GORM**: ORM (Object-Relational Mapping) para interagir com o banco de dados.
- **JWT**: Utilizado para autenticação e autorização.
- **Zap**: Biblioteca de logging.
- **GoMock**: Para criação de mocks em testes.
- **Testify**: Framework de testes.

## Injeção de Dependências

A injeção de dependências é realizada através de construtores que recebem interfaces como parâmetros. Isso permite uma fácil substituição de implementações, facilitando testes e manutenção. Por exemplo, a camada de serviço recebe uma interface de criptografia e uma interface de persistência:

## Testes

Os testes são uma parte crucial do projeto. Utilizamos um banco de dados em memória para testes, garantindo que os testes sejam rápidos e isolados. Além disso, utilizamos a biblioteca GoMock da Uber para gerar mocks das interfaces, permitindo testes unitários eficazes.



## Endpoints

### Registro de Usuário

- **URL**: `/register`
- **Método**: POST
- **Descrição**: Cria um novo usuário no sistema.
- **Request Body**:
  ```json
  {
    "first_name": "John",
    "last_name": "Doe",
    "email": "john.doe@example.com",
    "cep": "12345678",
    "country": "Brazil",
    "city": "São Paulo",
    "address": "Rua das Flores, 123",
    "password": "password#@#@!2121"
  }
  ```
- **Response**:
  - **Status 201**: Usuário criado com sucesso.
  - **Status 400**: Erro de validação.
  - **Status 500**: Erro interno no servidor.

### Login de Usuário

- **URL**: `/login`
- **Método**: POST
- **Descrição**: Autentica um usuário e retorna um token JWT.
- **Request Body**:
  ```json
  {
    "email": "john.doe@example.com",
    "password": "password#@#@!2121"
  }
  ```
- **Response**:
  - **Status 200**: Login efetuado com sucesso.
  - **Status 400**: Erro de validação.
  - **Status 500**: Erro interno no servidor.

### Obter Usuário

- **URL**: `/api/user/:id`
- **Método**: GET
- **Descrição**: Retorna as informações de um usuário específico. Requer autenticação JWT.
- **Response**:
  - **Status 200**: Informações do usuário.
  - **Status 403**: Permissão negada.
  - **Status 401**: Token inválido ou ausente.

## Como Executar

1. Clone o repositório.
2. Configure as variáveis de ambiente no arquivo `.env`.
3. Execute o comando `go run cmd/server/main.go` para iniciar o servidor.

## Testando com Docker

É possível testar esta API BFF usando o Dockerfile para criar uma imagem e executá-la na porta 8080 com os seguintes comandos:

## Futuras Implementações

- **Observabilidade e OpenTelemetry**: Planejo adicionar rastreamento e monitoramento utilizando OpenTelemetry em uma próxima feature.

- **CI/CD Pipeline**: O repositório contará com uma pipeline CI/CD utilizando Terraform para provisionamento em nuvem, especificamente na AWS.