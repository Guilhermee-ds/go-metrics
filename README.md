
# PROJETO DE METRICAS COM GO, PROMETHEUS E GRAFANA

Este projeto envolve um sistema de pagamentos no qual eu gero diversos pagamentos por minuto, com estados aleatórios de sucesso e falha. Esses dados são então exibidos em um dashboard usando Prometheus e Grafana para análise detalhada.




## Stacks Ultilizadas

- Docker
- Docker-Composer
- Golang >= go1.22.1
- Grafana
- Prometheus 

####

## Dependências

- Docker
- Docker-Composer
- Golang >= go1.22.1


####

## Como Rodar

Depois de ter todas a dependências informadas anteriormente

Inicie o container

```shell
docker compose up
```

Inicie o shell para gerar os pagamentos

```shell
./script.sh  
```

## Grafana

Agora acesse o Grafana no seu navegador atraves dessa url

```shell
http://localhost:3000/
```
Credenciais do Grafana

usuario: Admin 

Senha : Admin


Inicie uma nova dashboard e adicione uma nova visualização 
Com Prometheus

para criar uma nova dashboard adicione esse PromQL query

Primeiro Dashboard
```shell
sum(irate(ecomerce_payments_total[$__interval]))by (status)
```

Segundo Dashboard
```shell
rate(ecommerce_http_duration_sum[$__rate_interval]) / rate(ecommerce_http_duration_count[$__rate_interval])
```
# Finalizado!!!
Pronto! Agora a sua aplicação está sendo monitorada por dois painéis de controle, proporcionando uma análise mais detalhada e um monitoramento eficaz.
####






