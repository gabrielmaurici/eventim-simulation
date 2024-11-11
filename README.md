## eventim-simulation

Esse projeto simula uma experiência de fila virtual, reserva e compra de ingressos em um sistema de vendas online inspirado no site Eventim. Foi utilizado Golang para desenvolver as APIs de gerenciamento de fila, reserva e compra de ingressos, além de workers para atualização das posições na fila, contagem de compradores ativos e expiração das reservas, com o WebSocket enviando atualizações em tempo real para a interface do usuário sobre a posição na fila. Redis foi utilizado como cache distribuído para o gerenciamento de posições na fila, contagem de compradores ativos e controle das reservas temporárias. MySQL foi escolhido para persistir os dados dos ingressos, garantindo controle sobre os estados dos mesmos.

A interface foi construída em React, oferecendo uma interação integrada com as APIs e WebSocket, enquanto o Nginx desempenhou o papel de balanceamento de carga para assegurar um desempenho adequado em cenários de alta demanda fazendo também proxy reverso. A aplicação foi containerizada com Docker e orquestrada com Docker Compose, facilitando a escalabilidade e a replicação do sistema.

## 🛠️ Construído com

* [GoLang](https://go.dev/) - Usado para criar API's, Workers e WebSocket
* [go-chi](https://go-chi.io/#/) - Usado para fazer o roteamento com middleware das API's
* [Gorilla WebSocket](https://pkg.go.dev/github.com/gorilla/websocket) - Usado para fazer o WebSocket
* [Redis](https://redis.io/) - Usado para cache distribuído
* [MySQL](https://www.mysql.com/) - Usado para persistência de dados
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