# BFF E-commerce

Este projeto é um Backend for Frontend (BFF) para um site de e-commerce, implementado em Go. Ele fornece funcionalidades de login e autorização, seguindo uma arquitetura de camadas bem definida. Abaixo, detalhamos a estrutura do projeto, as tecnologias utilizadas, a injeção de dependências e os endpoints disponíveis.

## Contexto Geral

A ideia é usar este BFF para se conectar a um front-end desenvolvido em Angular com TypeScript de um e-commerce similar à Amazon. O BFF será a aplicação responsável por intermediar as requisições do front para os microsserviços (como de produtos, compras, rastreio, etc.) que estarão em sub-redes privadas na AWS. A comunicação ocorrerá por meio de autenticação de usuários e roles do IAM, para evitar exposição na internet pública. Este BFF estará dentro de um API Gateway, sendo também responsável por login com autenticação multifator e autorização, além de aplicar observabilidade por meio de tracers, permitindo encontrar e analisar gargalos de comunicação e processamento entre todos os microserviços que serão chamados para atender a solicitação de quem está consumindo determinado endpoint.

Ao decorrer deste projeto, a ideia é utilizar o Redis para armazenar códigos randômicos para comparar com os que serão inseridos pelo usuário na hora da autenticação multifator.

O projeto como um todo, ao decorrer do tempo, utilizará DynamoDB, PostgreSQL, filas SQS, SNS, S3 buckets, Lambdas e outros serviços da AWS.

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
    "email": "mockuser@example.com",
    "cep": "12345678",
    "country": "Brazil",
    "city": "São Paulo",
    "address": "Rua Mock, 123",
    "password": "mock@password"
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
    "email": "mockuser@example.com",
    "password": "mock@password"
  }
  ```
- **Response**:
  - **Status 200**: Login efetuado com sucesso.
  - **Status 400**: Erro de validação.
  - **Status 404**: Erro de usuário não encontrado.
  - **Status 401**: Senha incorreta.

  ### Serviço Externo

- **URL**: `/otherservice`
- **Método**: GET
- **Descrição**: Acessa um serviço externo. Requer autenticação JWT com a role "user".
- **Response**:
  - **Status 200**: Serviço encontrado com sucesso.
  - **Status 403**: Permissão negada.
  - **Status 401**: Token inválido ou ausente.

## Como Executar

1. Clone o repositório.
2. Variáveis de ambiente no arquivo estão corretas apontando para um banco de dados em memória `.env`.
3. Execute o comando `go run cmd/server/main.go` para iniciar o servidor.

## Testando com Docker

É possível testar esta API BFF usando o Dockerfile para criar uma imagem e executá-la na porta 8080 com os seguintes comandos:

## Futuras Implementações

- **Observabilidade e OpenTelemetry**: Planejo adicionar rastreamento e monitoramento utilizando OpenTelemetry em uma próxima feature.

- **CI/CD Pipeline**: O repositório contará com uma pipeline CI/CD utilizando Terraform para provisionamento em nuvem, especificamente na AWS.

- **Documentação Swagger**: Posteriormente, será feita uma documentação Swagger do projeto para facilitar a visualização e teste dos endpoints.
