<p align="center">
  <img src="https://www.lifeonline.com.br/site/site_v4/assets/img/Integra/Eulabs.png" alt="Eulabs">
</p>
<h1 align="center">Eulabs Hiring Challenge</h1>

> Criar uma API seguindo as melhores convenções e padrões do Golang, além de aplicar design patterns e boas práticas de design de API.

## Sobre API

Foi construída utilizando Golang (1.22) juntamente com o framework Echo. Por se tratar de um CRUD simples, não precisou de muitos patterns para deixar o domínio bem dividido, já que possui apenas uma entidade. O grande foco foi em trazer uma estruturação simples para o desafio, mas com boas práticas aplicadas. Em relação à estruturação de pasta, ela segue muito do [Golang Standard](https://github.com/golang-standards/project-layout), que é um repositório público da comunidade de Golang que traz alguns padrões aplicados pela comunidade. Cabe a você decidir se faz sentido ter ou não algumas pastas no seu projeto.

A API foi construída com base no **RESTful**. Principalmente na parte de _update_, foram implementados os dois métodos de atualização de uma entidade, que são **PUT** e **PATCH**, seguindo as regras em que o _PUT_ só realiza a operação se todos os campos forem atualizados. Caso seja uma alteração simples de apenas um campo, você deve utilizar _PATCH_ para realizar a operação com sucesso.

Foi implementado o **graceful shutdown** na aplicação para garantir que, caso tenhamos uma requisição em operação, ela não seja interrompida instantaneamente. Adicionei um context timeout de **15 segundos** e, caso a operação não seja finalizada em até 15s, ela será cancelada automaticamente e o servidor será desligado. Também foi construído um _package_ de logger bem simples para se trabalhar dentro da aplicação, que além de escrever no terminal, também salva no arquivo `server.log`. Assim, podemos analisar caso haja algum problema no momento de execução.

### Como rodar localmente o projeto?

![image makefile](https://github.com/waldrey/eulabs/assets/43473539/3ee0f769-8f90-4ca2-912d-b9f6b66301f3)

Construi um *makefile* para facilitar esse processo de inicialização da aplicação, abaixo tem uma imagem de como é _"cli"_ da aplicação, mas pós clonar repositório. Basta executar os seguintes comandos:

1. `make env` Vai copiar .env.example para .env para carregar na construção dos containers
2. `make install` Será levantado dois containers (mysql & golang[echo]) e automaticamente vai iniciar o servidor HTTP.

#### Para executar os testes

Basta executar o `make test` automaticamente vai executar todos os testes da aplicação🤘🏽

#### Collection & Environments Postman

Na pasta _/api_ no root do repositório você vai encontrar a pasta chamada **postman**  onde contém tanto a collection quanto environments de desenvolvimento para você realizar os testes. Você também pode utilizar **swagger** para realizar os testes caso deseja acessando [Acessando swagger aqui](http://localhost:8080/docs/index.html)