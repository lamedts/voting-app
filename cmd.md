protoc -I ../voting-app-pb/ ../voting-app-pb/vote.proto --go_out=plugins=grpc:pb
python -m grpc_tools.protoc -I../voting-app-pb --python_out=. --grpc_python_out=. ../voting-app-pb/vote.proto

go run main.go -stderrthreshold=INFO
go build -i -v -o bin/grpc-server .

psql -U $USERNAME -d $DB_NAME < $SQL_ROLE_SETUP_SCRIPT

docker build -t voting-app/worker .
docker run --rm voting-app/worker
