read -p "Quantidade de clientes: " CLIENTS

FILE="$CLIENTS-client.txt"

METADE=$(echo "scale=0; ( $CLIENTS+0.5)/1" | bc)

for index in $(eval echo "{1..$CLIENTS}")
do
	if [ $METADE -eq $index ]; then
		go run clientTCP.go >> "$FILE" &
	else
		go run clientTCP.go &
	fi
done

wait