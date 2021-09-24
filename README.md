# Hands On Kafka

## Subindo o Confluent Platform Stack

1. Google: "kafka all in one"
1. Pegar o [link do arquivo raw](https://raw.githubusercontent.com/confluentinc/cp-all-in-one/6.2.0-post/cp-all-in-one/docker-compose.yml) do `docker-compose.yml` do cp-all-in-one
1. Baixar localmente dentro da pasta `demo`: `wget https://raw.githubusercontent.com/confluentinc/cp-all-in-one/6.2.0-post/cp-all-in-one/docker-compose.yml`
1. Abrir no VS Code: `code .`
1. Remover as seções:
    - `schema-registry`
    - `connect`
    - `ksqldb-server`
    - `ksqldb-cli`
    - `ksql-datagen`
1. Tirar os `depends_on` do `control-center`: `schema-registry`, `connect` e `sqldb-server`
1. Tirar os `depends_on` do `rest-proxy`: `schema-registry`
1. Comentar as variáveis de ambiente:
    - `KAFKA_CONFLUENT_SCHEMA_REGISTRY_URL`
    - `CONTROL_CENTER_CONNECT_CONNECT-DEFAULT_CLUSTER`
    - `KAFKA_REST_SCHEMA_REGISTRY_URL`
1. Abir o Control Center: http://localhost:9021
