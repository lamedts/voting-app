 protoc -I pb/ pb/vote.proto --go_out=plugins=grpc:pb

 go run main.go -stderrthreshold=INFO