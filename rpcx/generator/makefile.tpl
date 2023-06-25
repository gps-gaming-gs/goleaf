rpc:
	goleaf rpc protoc {{.serviceName}}.proto --zrpc_out=. --home template

cocos:
	goleaf rpc protoc {{.serviceName}}.proto --zrpc_out=. --home template --cocos={{.serviceName}}

run:
	go run main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/{{.serviceName}}_service main.go