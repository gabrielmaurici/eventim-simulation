## chatbot-webscraping-chatgpt

![eventim-simulation](https://github.com/user-attachments/assets/e6e94927-a182-42da-89d1-1afc3a702d78)


Bot de WhatsApp desenvolvido em Node.js que realiza chamadas gRPC para um server .NET. O server .NET integra as APIs da OpenAI, permitindo interações com DALL-E e ChatGPT. Além disso, o projeto inclui um serviço de web scraping que coleta dados sobre resultados e próximas partidas de futebol.


## 🛠️ Construído com

* [.NET 8.0](https://learn.microsoft.com/pt-br/dotnet/core/whats-new/dotnet-8/overview) - Usado para criar server gRPC, integração com Open AI e webscraping de partidas de futebol
* [gRPC](https://grpc.io/) - Usado para integração entre as aplicações em Node JS e .NET
* [OpenAI API](https://platform.openai.com/docs/overview) - API de inteligência artificial para gerar chat e imagens através dos modelos GPT e Dall-e
* [Selenium WebDriver](https://www.selenium.dev/documentation/webdriver/) - Lib usada para webscraping
* [Node JS](https://nodejs.org/pt) - Usado para criar o bot de Whats App com um client gRPC para integração um server gRPC .NET
* [whatsapp-web.js](https://wwebjs.dev/) - Lib java script para integração com Whats App

## 📋 Pré-requisitos

* [Docker e Docker Compose](https://www.docker.com/) - Para rodar o projeto
* [OpenAI API Key](https://platform.openai.com/docs/quickstart/create-and-export-an-api-key) - Criar API KEY e exportar uma variavel de ambiente para configurar o projeto

Além de criar a API KEY, é necessário criar um projeto, permitir que o projeto acesse os modelos gpt-3.5-turbo e dall-e-3
![Permitir acesso aos modelos OpenAI](assets/models.png)

Como são modelos pagos, você pode acessar a parte de cobrança para adicionar um valor e poder fazer requests
![Realizar pagamento OpenAI](assets/billing.png)


## 💻 Como usar

Se você seguiu os paços da criação da API Key na seção dos pré requisitos exportando a variável de ambiente com o nome "OPENAI_API_KEY", o projeto está pronto para ser inciado, caso tenha exportado com outro nome, basta alterar o docker compose com o nome que você exportou.
![Altera docker compose api key](assets/docker-compose-apikey.png)

Com o projeto pronto para iniciar, rode o docker compose na raiz do projeto:
```
docker compose up
```

Ao iniciar o bot, aparecerá um QR Code para você ler pelo whats app e permitir que o bot possa ler e responder as mensagens. Após ler o QR Code uma mensagem irá informar que o client foi conectado no terminal:

![QR Code Whats App](assets/qr-code-wp.png)

Alguns exemplos das funcionalidades:

![Bot](assets/bot.png)
![IA Chat](assets/ia-chat.png)
![IA Imagem](assets/ia-imagem.png)
![Ultima pártida](assets/ultima-partida.png)
![Próxima partida](assets/proxima-partida.png)