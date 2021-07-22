#!/bin/bash

if [ -z "$1" ]; then
	echo
	echo "Please enter a valid argument"
	echo
	echo "cli : start application in cli mode"
	echo "> ./start.sh cli [help|domain]"
	echo "Options: "
	echo "	help :		print the help" 
	echo "	domains :	print the list of  domains" 
	echo
	echo "web : start application as a web server"
	echo "> ./start.sh web [\$port]"
	echo "Options: "
	echo "	\$port :		Port for the web server"
	echo
else
	if [ $1 = "cli" ]; then
		echo Starting application in cli mode
		echo "> MODE=cli go run main.go $OPTION"
		echo
		MODE=cli go run main.go $2	
	elif [ $1 == "web" ]; then
		echo Starting application as a web server
		echo "> MODE=web go run main.go $OPTION"
		echo
		MODE=web PORT=$2 go run main.go
	fi
fi
