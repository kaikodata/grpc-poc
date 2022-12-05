module kaiko-server

go 1.13

replace kaiko.io => /home/dell/grpc-poc/client/go

require (
	github.com/golang/protobuf v1.3.2
	google.golang.org/grpc v1.25.1
	kaiko.io v0.0.0-00010101000000-000000000000 // indirect
)
