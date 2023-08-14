read -p "Quantidade de clientes: " CLIENTS

FILE="$CLIENTS-client.txt"

METADE=$(echo "scale=0; ( $CLIENTS+0.5)/1" | bc)

if [ $CLIENTS -eq 1 ] ; then
	METADE=1
fi

for index in $(eval echo "{1..$CLIENTS}")
do
	if [ $METADE -eq $index ]; then
		go run RPCclient.go >> "$FILE" &
	else
		go run RPCclient.go &
	fi
done

wait