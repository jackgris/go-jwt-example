SERVER_NAME=jwt-server
CLIENT_NAME=jwt-client

build-server:
	$(info Building the app for local testing)
	go mod tidy
	cd $(PWD)/server && go build -o ${SERVER_NAME} .
	mv ./server/${SERVER_NAME} .

run-server:	build-server
	./${SERVER_NAME}

build-client:
	$(info Building the app for local testing)
	go mod tidy
	cd $(PWD)/client && go build -o ${CLIENT_NAME} .
	mv ./client/${CLIENT_NAME} .

run-client:	build-client
	./${CLIENT_NAME}


clean:
	go clean
	rm -f ${SERVER_NAME}
	rm -f ${CLIENT_NAME}
