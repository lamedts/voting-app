
protoc -I /usr/local/include -I. \
  -I ./vendor \
  -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  ./pb/vote.proto

protoc -I/usr/local/include -I. \
  -I ./vendor \
  -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  ./pb/vote.proto




python -m grpc_tools.protoc -I../voting-app-pb --python_out=. --grpc_python_out=. ../voting-app-pb/vote.proto

go run main.go -stderrthreshold=INFO
go build -i -v -o bin/grpc-server .

psql -U $USERNAME -d $DB_NAME < $SQL_ROLE_SETUP_SCRIPT

docker build -t voting-app/voter .
docker run --rm -p 8080:8080/tcp voting-app/voter

docker build -t voting-app/dashboard .
docker run --rm -p 8081:8081/tcp voting-app/dashboard

docker-compose up --build
docker-compose up
docker-compose down -v
