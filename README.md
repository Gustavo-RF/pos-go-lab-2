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
  

O objetivo desse laboratório é retornar as temperaturas de uma localidade buscada pelo cep. Além disso, o deploy será feito pelo google cloud run.
Para buscar o CEP, utilizarei a api do viacep, retornando a localidade desse cep. Com essa localidade, utilizarei a api weather api para buscar
a temperatura em celsius e calcular a temperatura em fahrenheit e kelvin.

A resposta será um objeto com as propriedades
```
{
  "temp_c": float32,
  "temp_f": float32,
  "temp_k": float32
}
```

## Endereço disponível
https://lab1-cloudrun-q4sscskaha-rj.a.run.app/

Para resposta, envie como query a propriedade cep:
```
?cep=cep_somente_numeros
```

### Exemplo
Requisição com sucesso:
https://lab1-cloudrun-q4sscskaha-rj.a.run.app?cep=29092260

Resposta
```
{
  "temp_c": 26.3,
  "temp_f": 79.34,
  "temp_k": 299.3
}
```

Caso envie um cep que não seja válido:
https://lab1-cloudrun-q4sscskaha-rj.a.run.app?cep=asdf

Resposta
```
{
  "message": "Invalid zipcode"
}
```

Caso envie um cep que não exista:
https://lab1-cloudrun-q4sscskaha-rj.a.run.app?cep=12312312

Resposta
```
{
  "message": "zipcode not found"
}
```

Caso não envie um cep:
https://lab1-cloudrun-q4sscskaha-rj.a.run.app

Resposta
```
{
  "message": "Cep is required"
}
```

## Testes locais
No terminal, digite:
```
cp .env.example .env
```
Caso necessário, altere a variável ```WEATHER_API_KEY``` para a sua key do weather api.

Para testar, utilize o docker compose:
```
docker compose up -d --build
```

Acesse localhost:8080 para executar as chamadas.

### Testes unitários
```
go test ./...
```
