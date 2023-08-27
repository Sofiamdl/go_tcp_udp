#!/bin/bash
#please run "chmod +x test.sh" beforehand

echo "🐰 Pressupondo que o RabbitMQ esteja rodando em localhost:5672... 🐰"
read -p "Quantos clientes? " CLIENTS_QUANT
read -p "Quantas execuções? " EXEC_QUANT

FILE="$CLIENTS_QUANT.txt"

go run server.go &

for index in $(eval echo "{1..$EXEC_QUANT}")
do
	go run client.go -clients=$CLIENTS_QUANT >> "$FILE" &
done

if [ ! -d "tests" ]; then
    mkdir tests
fi

mv "$FILE" tests/

wait