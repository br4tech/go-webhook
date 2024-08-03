# Webhook de Notificação de Pagamentos

Este projeto implementa um webhook robusto e escalável para notificações de status de pagamento, utilizando:

- **Arquitetura Hexagonal**: Separação clara de responsabilidades, promovendo testabilidade e flexibilidade.
- **MongoDB**: Banco de dados NoSQL flexível e escalável para armazenar inscrições e eventos.
- **MongoDB Generic Driver**: Interface genérica para interagir com o MongoDB, facilitando a manutenção e a troca de drivers, se necessário.
- **Docker**: Contêinerização para facilitar o desenvolvimento, teste e implantação.
- **Docker Compose**: Orquestração de múltiplos contêineres (webhook, MongoDB) para simplificar o ambiente de desenvolvimento.

## Componentes Principais

### Domínio
- **Entidades**: PaymentStatus (enum) e PaymentEvent (struct).
- **Ports**: Interfaces para abstrair a persistência de dados e o envio de notificações.

### Aplicação
- **Webhook Handler**: Lida com as requisições do webhook, valida, processa e notifica os inscritos.
- **Rotas**:
  - `/webhook`: Endpoint para receber notificações de status de pagamento.
  - `/subscriptions`: Endpoint para gerenciar inscrições de usuários em pagamentos.

### Adapters
- **MongoDB Adapter**: Implementa o port de persistência, interagindo com o MongoDB para armazenar e recuperar dados.
- **Notification Adapter**: Implementa o port de notificação, enviando notificações para os usuários inscritos (e-mail, push, etc.).

### Infraestrutura
- **Docker**: Contêineriza a aplicação webhook e o MongoDB.
- **Docker Compose**: Define e orquestra os contêineres.

## Benefícios

- **Escalabilidade**: O uso do MongoDB e a arquitetura hexagonal permitem que o sistema lide com um grande volume de notificações e inscrições.
- **Flexibilidade**: A arquitetura hexagonal facilita a troca de componentes (banco de dados, provedor de notificações) sem impactar o domínio.
- **Testabilidade**: A separação de responsabilidades permite escrever testes unitários e de integração de forma eficaz.
- **Facilidade de Desenvolvimento**: Docker e Docker Compose simplificam a configuração do ambiente de desenvolvimento.

## Como Executar
   - Em desenvolvimento....

## Acesse a Aplicação

- **Webhook**: [http://localhost:8080/webhook](http://localhost:8080/webhook)
- **Inscrições**: [http://localhost:8080/subscriptions](http://localhost:8080/subscriptions)


## Estrutura:
```bash
  |____README.md
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
  | | |____webhook_handler.go
  | | |____payment_handler.go
```
## Observações

- Certifique-se de ter o Docker e o Docker Compose instalados em sua máquina.
- Adapte as configurações do Docker Compose (portas, nomes de contêineres, etc.) para o seu ambiente.

