module github.com/gbeletti/service-golang

go 1.19

require (
	github.com/gbeletti/rabbitmq v0.0.10
	github.com/go-chi/chi/v5 v5.0.8
	github.com/joho/godotenv v1.4.0
	github.com/rabbitmq/amqp091-go v1.5.0
	go.mongodb.org/mongo-driver v1.11.1
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.15.13 // indirect
	github.com/montanaflynn/stats v0.6.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/text v0.5.0 // indirect
)

replace github.com/docker/docker => github.com/docker/docker v20.10.3-0.20221013203545-33ab36d6b304+incompatible // 22.06 branch
