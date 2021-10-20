# Hands On Kafka

## Criando uma Rede Docker Comum a Todos os Serviços

Antes de tudo, vamos criar uma rede dentro do docker para que possamos reutilizá-la
entre os nossos diferentes serviços que iremos subir. Para isso, execute:

```
docker network create kafka-hands-on
```

## Subindo o Confluent Platform Stack

1. Google: "kafka all in one"
1. Comentar sobre as possibilidades de configurações
1. Rodar o `docker-compose up`
1. Abir o Control Center: http://localhost:9021
1. Mostrar um pouco da interface do Control Center

## Tópicos

1. Mostrar como cria um Tópico pela interface
1. Criar um tópico manualmente e mostrar na interface:
    ```
    docker-compose exec broker bash
    # kafka- (mostrar que existem vários comandos)
    kafka-topics --create --topic topico-exemplo --bootstrap-server localhost:9092
    ```

## Consumindo e Produzindo

### Sem partições

1. Abrir uma nova aba de terminal
1. Dividir a tela verticalmente (lado a lado)
1. Na esquerda, dividir horizontalmente e iniciar dois consumers
    ```
    docker-compose exec broker bash
    kafka-console-consumer --topic topico-exemplo --bootstrap-server localhost:9092
    ```
1. Na direita, iniciar um producer sem uso de key (usar um pangrama)
    ```
    docker-compose exec broker bash
    kafka-console-producer --topic topico-exemplo --bootstrap-server localhost:9092
    ```
1. Comentar os comportamentos
1. Mostrar as configurações do tópico:
    ```
    kafka-topics --describe --topic topico-exemplo --bootstrap-server localhost:9092
    ```
1. Mostrar essas configurações na interface
1. Mostrar que as mensagens estão persistidas em disco, na interface: filtrar offset 0

### Com partições

1. Deletar o tópico criado:
    ```
    kafka-topics --delete --topic topico-exemplo --bootstrap-server localhost:9092
    ```
1. Criar novamente o tópico, agora com 2 partições:
    ```
    kafka-topics --create --topic topico-exemplo --partitions 2 --bootstrap-server localhost:9092
    ```
1. Vai dar erro, pq o tópico ainda existe. Ao produzir ou consumir em um tópico inexistente, o Kafka cria ele
1. Parar os consumers
1. Deleta e tenta criar outra vez. Se der erro, fecha o CP que ele pode ficar fazendo conexões
1. Descrever o tópico criado e observar que agora existem duas partições:
    ```
    kafka-topics --describe --topic topico-exemplo --bootstrap-server localhost:9092
    ```
1. Reiniciar os dois consumers, do mesmo jeito anterior:
    ```
    kafka-console-consumer --topic topico-exemplo --bootstrap-server localhost:9092
    ```
1. Nada muda. Pare os consumers

### Consumer Groups

1. Agora silumando várias instâncias da mesma aplicacão (consumer groups):
    ```
    kafka-console-consumer --topic topico-exemplo --group awesome-app --bootstrap-server localhost:9092
    # do teclado pro seu coração
    ```
1. Chamar atenção para o algoritmo **Round-Robin**
1. Na direita, dividir horizontalmente
1. Descrever os consumer groups no novo terminal:
    ```
    kafka-consumer-groups --describe --group awesome-app --bootstrap-server localhost:9092
    ```
1. Chamar atenção para o fato de cada partição ter um Consumer ID diferente
1. Matar um dos consumers
1. Descrever os consumer groups novamente:
    ```
    kafka-consumer-groups --describe --group awesome-app --bootstrap-server localhost:9092
    ```
1. Chamar atenção para o fato das partições terem um mesmo Consumer ID
1. Mostrar nos LOGs da outra janela o processo de rebalancing acontecendo
1. Publicar novas mensagens e observar que todas vão para o único consumer de pé:
    ```
    # ela roda a cidade inteira pra ficar comigo
    ```
1. Subir novamente o consumidor que tinha caído
1. Publicar novas mensagens e observar que o rebalanceamento foi feito:
    ```
    # eu sou seu esquema preferido
    # ela dispensa a balada, as amigas, pra ficar comigo
    ```
1. Descrever o consumer group e observar que os Consumer IDs são distintos novamente:
    ```
    kafka-consumer-groups --describe --group awesome-app --bootstrap-server localhost:9092
    ```
1. Subindo um novo consumer, na esperança de ir mais rápido. Produzir novas mensagens
    ```
    kafka-console-consumer --topic topico-exemplo --group awesome-app --bootstrap-server localhost:9092
    # é só tu ligar pra mim que eu não resisto
    ```

### Usando Chaves

1. Produzindo com chaves:
    ```
    kafka-console-producer --topic topico-exemplo --bootstrap-server localhost:9092 --property parse.key=true --property key.separator=:
    ```

## Integrando Kafka em Aplicações

### Consumer em Go

[Documentação da Biblioteca da Confluent](https://docs.confluent.io/platform/current/clients/confluent-kafka-go/index.html)

[Documentação de Configurações dos Consumers](https://kafka.apache.org/documentation.html#consumerconfigs)

1. Subir inicialmente tentando se comunicar com a porta errada (`9092`) antes de usar a certa (`29092`)
1. Mostrar que não é possível subir sem informar o `group.id`
1. Mostrar que é possível consumir um ou vários tópicos: `.Subscribe()` ou `.SubscribeTopics()`
1. Produzir várias mensagens em um tópico particionado e observar que o consumo será meio bagunçado
1. Mostrar as diferenças entre as opções de `earliest` e `latest` do reset do offset
1. Encerrar o consumidor, produzir mensagens, voltar com ele e observar que continua onde parou
1. Mostrar o uso de expressão regular no subscribe

### Producer em PHP

[Documentação da Biblioteca Laravel Kafka](https://github.com/mateusjunges/laravel-kafka)

[Documentação da Biblioteca librdkafka em C](https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md)

#### Comandos Úteis

```sh
# Executar obrigatoriamente
cp .env.example .env
docker-compose build
docker run -it --rm --name php-producer-app -v $(pwd):/var/www/html php-producer-app composer install
docker run -it --rm --name php-producer-app -v $(pwd):/var/www/html php-producer-app php artisan key:generate
docker-compose up
docker-compose exec app sh
php artisan migrate

# Requisições para a API
http POST http://localhost:8000/api title="Forrest Gump" release_year=1994 description="Um cara que conta historias"
http POST http://localhost:8000/api title="The Shawshank Redemption" release_year=1994

# Caso seja preciso acessar o app
docker-compose exec app sh

# Caso seja preciso acessar a database e fazer mudanças
docker-compose exec db mysql -u root

UPDATE films SET title = 'Big Hero 6', release_year = 2014, description = 'Uma animacao daora' where id = 1;
```

## Connectors

[Debezium Connector for MySQL](https://debezium.io/documentation/reference/1.2/connectors/mysql.html)

Caminho da pasta no WSL: `\\wsl$`
Completo: `\\wsl$\Ubuntu-20.04\home\claudson\Code\kafka-hands-on\src\connectors`

## Possíveis Dúvidas

### Preciso de um acompanhamento/monitoramento constante para garantir que o processo do broker está rodando no docker?

Você geralmente não vai rodar o Kafka com Docker. Talvez um serviço gerenciado ou no bare metal. Mínimo de 3 máquinas (brokers) no cluster.

Existe um controlador do Kafka para o Kubernetes, é um caminho possível, mas geralmente um serviço autogerenciado ou você subir suas máquinas e configurar o Kafka é mais comum.

### Qual o Use Case dos Consumer Groups?

Distribuir o processamento das mensagens. Aumentar o throughput da leitura das mensagens.

