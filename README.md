# Webhook de Notificação de Pagamentos com Arquitetura Hexagonal, MongoDB e Docker

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

## Observações

- Certifique-se de ter o Docker e o Docker Compose instalados em sua máquina.
- Adapte as configurações do Docker Compose (portas, nomes de contêineres, etc.) para o seu ambiente.

