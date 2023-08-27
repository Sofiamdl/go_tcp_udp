#!/bin/bash

echo "🐰 Pressupondo que o RabbitMQ esteja rodando em localhost:5672... 🐰"
read -p "Quantos clientes? " CLIENTS_QUANT
read -p "Quantas execuções? " EXEC_QUANT

FILE="$CLIENTS_QUANT.txt"

METADE=$(echo "scale=0; ( $EXEC_QUANT+0.5)/1" | bc) #TODO: adicionar a logica de guardar metade das execucoes

go run receive.go -consumers=$CLIENTS_QUANT -executions=$EXEC_QUANT >> "$FILE" &
go run send.go -executions=$EXEC_QUANT &

if [ ! -d "tests" ]; then
    mkdir tests
fi

mv "$FILE" tests/

wait