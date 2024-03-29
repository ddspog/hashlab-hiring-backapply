# Hash Teste Back-end

Aqui segue meu projeto como resposta para o desafio. Utilizei Golang para o Serviço 1 e Python para o Serviço 2. Os dois serviços utilizam um banco de dados compartilhado em PostgreSQL, equipado com uma engine GraphQL.

Este documento primeiro transcreve o desafio, para então explicar a solução aplicada.

## Sobre o Desafio

O teste consiste em escrever 2 microserviços que possibilitam retornar uma lista de produtos com desconto personalizado para cada usuário.

### Restrições

 1. Os serviços desse teste devem ser escritos usando linguagens distintas
 2. Os serviços desse teste devem se comunicar via [gRPC](https://grpc.io/)
 3. Utilize [docker](https://www.docker.com/) para provisionar os serviços
 4. Para facilitar, os serviços podem usar um banco de dados compartilhado

### Serviço 1: Desconto invidual de produto

* Este serviço recebe um id de produto e um id de usuário e retorna um desconto.

Produto exemplo:

```json
{
    id: string
    price_in_cents: int
    title: string
    description: string
    discount: {
        pct: float
        value_in_cents: int
    }
}
```

Usuário exemplo:

```json
{
    id: string
    first_name: string
    last_name: string
    date_of_birth: Date
}
```

* As regras de descontos da aplicação são:
  * Se for aniversário do usuário, o produto terá 5% de desconto
  * Se for black friday (nesse exemplo ela pode ser fixada dia 25/11) o produto terá 10% de desconto
  * O desconto não pode passar de 10%

### Serviço 2: Listagem de produtos

* Expõe uma rota HTTP tal que `GET /product` retorne um json com uma
lista de produtos.

* Essa rota deve receber opcionalmente via header `X-USER-ID` um id de usuário.

* Para obter o desconto personalizado este serviço deve utilizar o serviço 1.

* Caso o serviço 1 retorne um erro, a lista de produtos ainda precisa ser retornada, porém com esse produto que deu erro sem desconto.

* Se o serviço de desconto (1) cair, o serviço de lista (2) tem que continuar funcionando e retornando a lista normalmente, só não vai aplicar os descontos.

## Sobre minha solução

![](https://raw.githubusercontent.com/ddspog/hashlab-hiring-backapply/master/img/services-comm.png)

Conforme a figura acima, meu projeto faz a comunicação esperada entre os serviços, definindo o código para 3 docker containers nas seguintes pastas:

* **database**: Configura o database definindo o `serve-config` container, que irá disponibilizar arquivos de inicialização do banco de dados utilizado.
* **service1**: Contém todo o código que compõe o `service1` container, escrito em Golang.
* **service2**: Armazena o servidor `service2`, escrito em Python.

### Como Executar

Utilizando o `docker-compose`, é possível configurar o database e os dois serviços, de modo que comuniquem entre si, e possam ser acessados externamente.

É só executar:

```sh
docker-compose up -d
```

Que o serviço 1 fica disponível em `<docker-machine-ip>:5001` e o serviço 2 em `<docker-machine-ip>:5002`. Para saber qual `<docker-machine-ip>` usar, utilize o comando:

```sh
docker-machine ip
```

Com os serviços em execução, pode-se testar o serviço 1 por gRPC, utilizando [BloomRPC](https://github.com/uw-labs/bloomrpc) por exemplo. E o serviço 2 via [Postman](https://www.getpostman.com).

Para fins de teste, populei o banco de dados com 4 Produtos e 3 Usuários. Abaixo segue o id dos usuários para testes:

 Usuário | Aniversário
 ------- | -----------
 543.004.756-11 | 15/06
 432.888.752-30 | 25/11
 879.451.123-99 | 14/03

E abaixo segue o id dos produtos para testes:

 Produto | Preço
 ------- | -----
 1a184013-d405-4da1-b956-4781e5e4d256 | R$ 1,00
 84b21225-9be4-4d9c-9bf0-6e5b321e2ca4 | R$ 26,60
 71bbc322-ffef-47af-8d87-c4bc596900af | R$ 30,00
 1cd9ab74-56ba-444e-8666-c00eb9597e69 | R$ 26,90

 **OBS**: Para poder testar os dois serviços, aconselho configurar a variável `MOCKED_TODAY_DATE` na execução do compose, com uma data desejada no formato `2019-08-07` para o dia 7 de agosto de 2019. A execução padrão dos serviços leva em conta o dia atual.

 Como esperado, o serviço 1 configura um server gRPC, que segue o protocolo em `discount.proto`. O serviço 2 é um REST server que espera chamadas GET /product, com o id de usuário fornecido no header: `X-USER-ID`.

### Como o BD funciona

![](https://raw.githubusercontent.com/ddspog/hashlab-hiring-backapply/master/img/database-setup.png)

Como mostrado na imagem, o database é PostgreSQL gerenciado por um motor GraphQL desenvolvido pela [Hasura](https://hasura.io), definidos nos containers `postgres`e `graphql-engine` respectivamente. 

Na pasta `database` há o arquivo `schema.up.sql` que detalha o Schema das tabelas utilizadas no BD, este arquivo é lido pelo container `graphql-engine` em sua inicialização, para criar o database.

Utilizo um programa escrito em Golang, no container `populate` para rodar a mutation definida em `database/populate.gql`. Antes disso, o programa verifica se as tabelas já existem com a query `database/check-tables.gql`, e espera serem definidas.

Os arquivos necessários para configuração do database são todos disponibilizados pelo container `serve-config` que é um servidor de arquivos estáticos, permitindo o donwload dos arquivos nos outros containers.

Essa separação foi feita, para que eu configurasse o BD a partir de duas imagens apenas: `ddsdok/hasura-graphql-engine` e `ddsdok/graphql-populate`. Pretendo utilizar as mesmas imagens em outros projetos, bastando-me apenas criar outro `serve-config` container para disponibilizar os arquivos definindo o BD.

### Os Serviços

![](https://raw.githubusercontent.com/ddspog/hashlab-hiring-backapply/master/img/service1.png)

No serviço 1, organizei o projeto de acordo com o estilo Golang: utilizar bastante código aberto, e testar bastante o código antes de mandar para produção.

Para isso, separei o projeto em 4 pacotes, e utilizei a framework [Cobra](https://github.com/spf13/cobra) para usar o programa em CLI. Utilizei o máximo de meu conhecimento possível em Golang, utilizando inclusive, pacotes meus como base. Os testes foram feitos por BDD + TDD + Table-Driven-Tests.

O pacote **server** é responsável por definir o server gRPC, e seus métodos; além de disponibilizar alguns valores globais para o servidor, como nome, host, port.

No pacote **model**, as entidades usuário e produtos são definidas, com métodos para executarem queries para um servidor GraphQL, e então recuperarem informação do Banco de Dados.

Alguns métodos de checagem são definidos no pacote **today**, o qual verifica se a data atual corresponde ao Black Friday, ou a o aniversário do usuário.

A entrada do servidor é definida na pacote **cmd**, que utiliza a framework Cobra para processar argumentos CLI.

![](https://raw.githubusercontent.com/ddspog/hashlab-hiring-backapply/master/img/service2.png)

Já no serviço 2, por ser Python, optei pela simplicidade. Ao invés de pacotes, dividi o projeto em poucos arquivos (módulos) no mesmo nível, apenas encapsulando o código gerado pelo protocolo gRPC em um pacote único. Os testes foram simples, feitos em BDD.

O servidor é iniciado no módulo **app** que usa a framework Flask para configurar o REST server, e a rota GET /product.

O módulo **app** usa os módulos **db** e **grpc_client** para pegar informação de todos os produtos, e possíveis descontos para o usuário atual, respectivamente.

O módulo **grpc_client** utiliza o módulo **discount** que basicamente contém o código gerado pelo protocolo `discount.proto`.

O serviço 2, por ser extremamente simples e curto em linhas, pode evoluir separando para cada módulo um pacote único, dividindo em quantos módulos forem necessários.

<!-- Envie o resultado do seu desafio para dev@hash.com.br (ele pode ser open source!). Em até uma semana marcaremos uma conversa com você após analisarmos seu desafio. -->

<!-- ## Avaliação

1. Conversaremos sobre a estrutura do código, escolha do banco, e outras decisões que foram tomadas
2. Discutiremos como esse sistema evoluiria ao longo do tempo
3. Considere que as regras de descontos irão mudar com o tempo
4. Considere que mais pessoas irão trabalhar junto com você nesse projeto -->
