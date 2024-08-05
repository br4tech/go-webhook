# Notificação de Pagamentos

Este projeto implementa 3 aplicacoes no fluxo de venda e pagamento:

   1. Client responsavel por criar um pedido(order) de venda
   2. Subscriber responsavel por criar as inscricaoes dos webhook(outbound)
   3. Payment responsavel port transitar os pagemento  e informar os webhook inscritos 
  
  Todas as aplicacoes se encontra nesse repo, com uma estrutura robusto e escalável utilizando:

- **Arquitetura Hexagonal**: Separação clara de responsabilidades, promovendo testabilidade e flexibilidade.
- **MongoDB**: Banco de dados NoSQL flexível e escalável para armazenar inscrições e eventos.
- **PostgreSQL**: Banco de dados Relacional para armazenar produtos e pedidos.
- **Docker**: Contêinerização para facilitar o desenvolvimento, teste e implantação.
- **Docker Compose**: Orquestração de múltiplos contêineres (client, payment e subscriber) para simplificar o ambiente de desenvolvimento.

## Componentes Principais

### Domínio
- **Entidades**: Product, Order, OrderItem, Payment, Subscriber (struct).
- **Ports**: Interfaces para abstrair a persistência de dados e DI.

### Aplicação

- **Client**: Lida com a criacao de produto e pedido no PostgreSQL.
  - **Rotas**:
    - `/product`: Endpoint para criar produto.
    - `/order`: Endpoint para criar um pedido, com base no produto criado.

- **Payment**: Lida com o fluxo de pagamento e envia o evento para o webhook configurado.
  - **Rotas**:
    - `/payment`: Endpoint para transitar o status de pagamento do pedido.

- **Subscribe**: Lida com o fluxo de inscritos, atraves dele sabemos a rota para enviar o evento.
  - **Rotas**:
    - `/subscriber`: Endpoint para gerenciar inscrições de usuários que vao receber os eventos de pagamentos.


### Adapters

- **MongoAdapter**: Implementa o port de persistência, interagindo com o MongoDB para armazenar e recuperar dados.
- **PostgresAdapter**: Implementa o port persistência, interagindo com o PostgreSQL para armazenar e recuperar dados..

### Infraestrutura

- **Docker**: Contêineriza a aplicação webhook e o MongoDB.
- **Docker Compose**: Define e orquestra os contêineres.

## Benefícios

- **Escalabilidade**: O uso do MongoDB e a arquitetura hexagonal permitem que o sistema lide com um grande volume de notificações e inscrições.
- **Flexibilidade**: A arquitetura hexagonal facilita a troca de componentes (banco de dados, provedor de notificações) sem impactar o domínio.
- **Testabilidade**: A separação de responsabilidades permite escrever testes unitários e de integração de forma eficaz.
- **Facilidade de Desenvolvimento**: Docker e Docker Compose simplificam a configuração do ambiente de desenvolvimento.

## Como Executar

   1. Clone o repositório:
   ```bash
    git clone https://github.com/br4tech/go-webhook.git
   ```

   2. Inicie os serviços com Docker Compose:
   ```bash
    docker-compose up
   ```

## Acesse as Aplicaçoes

  ### Client

  - **Product**: [http://localhost:8080/product](http://localhost:8080/product)
  - **Order**:   [http://localhost:8080/order](http://localhost:8080/order)

  ### Payment
  
  - **Payment**: [http://localhost:8081/payment](http://localhost:8081/payment)

  ### Subscribe

  - **Subscriber**: [http://localhost:8082/subscriber](http://localhost:8082/subscriber)

## Estrutura:
```bash
|____README.md
|____config.yml
|____cmd
| |____subscriber
| | |____config.yml
| | |____main.go
| | |____Dockerfile
| | |____server
| | | |____server.go
| | | |____echo_server.go
| |____client
| | |____config.yml
| | |____main.go
| | |____Dockerfile
| | |____server
| | | |____server.go
| | | |____echo_server.go
| |____payment
| | |____config.yml
| | |____main.go
| | |____Dockerfile
| | |____server
| | | |____server.go
| | | |____echo_server.go
|____go.mod
|____.vscode
| |____launch.json
|____go.sum
|____docker-compose.yml
|____config
| |____config.go
|____internal
| |____dto
| | |____payment.go
| | |____product.go
| | |____subscribe.go
| | |____order_item.go
| | |____order.go
| |____repositories
| | |____payment_repository.go
| | |____order_repository.go
| | |____subscribe_repository.go
| | |____order_item_repository.go
| | |____product_repository.go
| |____adapter
| | |____mongo.go
| | |____postgres.go
| |____utils
| | |____filters
| | | |____greater_than.go
| | | |____equal.go
| |____model
| | |____postgres
| | | |____product.go
| | | |____order_item.go
| | | |____order.go
| | | |____model.go
| | |____mongo
| | | |____payment.go
| | | |____subscribe.go
| |____core
| | |____port
| | | |____repository.go
| | | |____database.go
| | | |____handler.go
| | | |____usecase.go
| | |____usecase
| | | |____payment_use_case.go
| | | |____order_use_case.go
| | | |____subscription_usecase.go
| | | |____product_use_case.go
| | |____domain
| | | |____payment.go
| | | |____product.go
| | | |____subscribe.go
| | | |____payment_status.go
| | | |____order_item.go
| | | |____order.go
| | | |____payment_event.go
| |____handler
| | |____product_handler.go
| | |____subscribe_handler.go
| | |____handler.go
| | |____payment_handler.go
| | |____order_handler.go
```
## Observações

- Certifique-se de ter o Docker e o Docker Compose instalados em sua máquina.
- Adapte as configurações do Docker Compose (portas, nomes de contêineres, etc.) para o seu ambiente.
- Hoje estamos compartilhando os dados de casos de uso e repository, porem com a estrutura adotada podemos levar nosso fonte para diferentes repo.

