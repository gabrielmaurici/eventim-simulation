## eventim-simulation

Esse projeto simula uma experiência de fila virtual, reserva e compra de ingressos em um sistema de vendas online inspirado no site Eventim. Foi utilizado Golang para desenvolver as APIs de gerenciamento de fila, reserva e compra de ingressos, e também para workers que atualizam as posições na fila, contabilizam os compradores ativos e expiram as reservas. A comunicação entre os serviços é facilitada pelo RabbitMQ, que promove a troca eficiente de mensagens e coordena as operações internas do sistema. O WebSocket transmite atualizações em tempo real para a interface do usuário, informando a posição na fila.

O Redis foi utilizado como cache distribuído para gerenciar as posições na fila, a contagem de compradores ativos e o controle das reservas temporárias. Para persistência dos dados dos ingressos e garantia de controle sobre seus estados, foi escolhido o MySQL.

A interface foi construída em React, possibilitando uma interação contínua com as APIs e WebSocket. O Nginx atua como balanceador de carga e proxy reverso, garantindo um desempenho adequado em cenários de alta demanda. Toda a aplicação foi containerizada com Docker e orquestrada com Docker Compose, promovendo escalabilidade e facilitando a replicação do sistema.

![Eventim Simulation](https://github.com/user-attachments/assets/e83f27c1-a84b-4fe2-877b-9bfc638103fc)

## 🛠️ Construído com

* [GoLang](https://go.dev/) - Usado para criar API's, Workers e WebSocket
* [go-chi](https://go-chi.io/#/) - Usado para fazer o roteamento com middleware das API's
* [Gorilla WebSocket](https://pkg.go.dev/github.com/gorilla/websocket) - Usado para fazer o WebSocket
* [Redis](https://redis.io/) - Usado para cache distribuído
* [MySQL](https://www.mysql.com/) - Usado para persistência de dados
* [Rabbit MQ](https://www.rabbitmq.com/) - Usado para comunicação entre sistemas
* [Nginx](https://nginx.org/) - Usado para balanceamento de carga e proxy reverso
* [Docker](https://www.docker.com/) - Usado para criação do ambiente em desenvolvimento

## 📋 Pré-requisitos

* [Docker e Docker Compose](https://www.docker.com/) - Para rodar o projeto

## 💻 Como usar

Rode docker compose up na raiz do projeto:

```
docker compose up
```

Aguarde a inicialização de todos os projetos e navegue até localhost:8080

![Eventim Simulation GIF](https://github.com/user-attachments/assets/0cafe349-c43a-4ef6-beb2-64c21e900164)