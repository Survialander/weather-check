# Weather Check App

## Descrição

Weather Check é um aplicativo que informa a temperatura em dada localização. Basta informar um CEP e ele irá retornar a temperatura em tempo real nos seguintes formatos: 
- Celcius
- Fahrenheit
- Kelvin

## Como utilizar

Para rodar o weather check app precisamos ter instalado o [Docker](https://www.docker.com/) e uma api key válida da [WeatherApi](https://www.weatherapi.com/) 

Primeiro crie um arquivo `.env`, use o `.env.example` como base, depois coloque o valor da sua WeatherApi api key na chave `WEATHER_KEY`, seu arquivo `.env` deve ficar assim:
```
WEATHER_KEY={api key}
```

Agora basta rodar os seguintes comandos:
```
docker build -t weather-check .
```
```
docker run --rm -p 8080:8080 weather-check
```

Ou utilize o docker compose:
```
docker-compose up
```

A aplicação está pronta para ser utilizada, um exemplo de requisição:
```
http://localhost:8080/?cep=17560015
```

A aplicação também pode ser testada sem a necessidade de subir o projeto na sua máquina, ela está hospedada no seguinte endereço:
```
https://weather-cfpxj7qd6a-uc.a.run.app/?cep=
```
