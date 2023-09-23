do-unknown-request:
	curl --request POST --url http://localhost:3001/auth/error
do-specific-request:
	curl --request POST --url http://localhost:3001/auth/error --header 'Accept-Language: ru'
run:
	go run cmd/main.go