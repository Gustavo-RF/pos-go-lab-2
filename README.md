# Pos Go Lab 2
O objetivo desse laboratório é coletar e visualizar o tempo de execução entre serviços, utilizando o open telemetry.
Além do mais, utilizaremos o Zipkin para visualizar os traces e spans criados na aplicação.

## Arquitetura
Teremos dois serviços:
- Serviço A: Responsável por receber a requisição do cliente, que será o cep, validar e enviar para o serviço B
- Serviço B: Responsável por receber a requisição do serviço A e buscar a localidade utilizando o Via cep e em seguida buscar a temperatura atual
  utilizando o weather api, retornando então um objeto contendo:
  ```
  {
    "city": string,
    "temp_c": float32,
    "temp_f": float32,
    "temp_k": float32
  }
  ```
## Zipkin
O zipkin é um sistema de tracing distribuido. Ele será responsável pela visualização do trace total desde a requisição do
cliente até as chamadas de cep e temperatura.

# Testando
Para testar, execute o comando docker compose
```
docker compose up -d --build
```
Após isso, estará disponível na porta 8080 localmente o endereço para teste.
No terminal, execute um cUrl:
```
 curl --request POST \                                          
  --url http://localhost:8080/ \
  --header 'Content-Type: application/json' \
  --data '{
        "cep": "cep_valido"
}'
```
Essa chamada passará todo o processo descrito acima e retornará a resposta:
  ```
  {
    "city": string,
    "temp_c": float32,
    "temp_f": float32,
    "temp_k": float32
  }
  ```

## Visualizando Trace
Para visualizar o Trace pelo zipkin, acesse:

http://localhost:9411

Assim que a requisição for feita, clique em ```Run Query``` para listar os traces.

![image](https://github.com/user-attachments/assets/adc22259-6ef2-43b9-bf81-6aa295086755)

### Observação
Um tempo de 1 segundo foi acrescido ao inicio e ao fim da execução para melhor visualização dos traces

### Exemplo
Requisição com sucesso:
```
 curl --request POST \                                          
  --url http://localhost:8080/ \
  --header 'Content-Type: application/json' \
  --data '{
        "cep": "29092260"
}'
```

Resposta
```
{
  "city": "Vitória",
  "temp_c": 26.3,
  "temp_f": 79.34,
  "temp_k": 299.3
}
```

Caso envie um cep que não seja válido:
```
 curl --request POST \                                          
  --url http://localhost:8080/ \
  --header 'Content-Type: application/json' \
  --data '{
        "cep": "asdf"
}'
```

Resposta
```
{
  "message": "Invalid zipcode"
}
```

Caso envie um cep que não exista:
```
 curl --request POST \                                          
  --url http://localhost:8080/ \
  --header 'Content-Type: application/json' \
  --data '{
        "cep": "12312312"
}'
```

Resposta
```
{
  "message": "zipcode not found"
}
```

Caso não envie um cep:
```
 curl --request POST \                                          
  --url http://localhost:8080/ \
  --header 'Content-Type: application/json' \
  --data '{
  }'
```

Resposta
```
{
  "message": "Cep is required"
}
```

### Testes unitários
```
go test ./... 
```
