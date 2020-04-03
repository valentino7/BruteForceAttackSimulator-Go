#!/bin/sh
if test $1 -lt 0 ; then
	echo "Error parameter, the number of workers must be greater than 0 "
	exit 
fi
if test $# != 1 ; then
	echo "Error parameter 2, enter the number of workers"
	exit 
fi
gnome-terminal -e "go run server.go"
var=16000
for i in $(seq $1)
	do
	    	port=$((var+i))
		gnome-terminal -e "go run worker.go $port"
		var=16000
done

gnome-terminal -e "go run client.go"


