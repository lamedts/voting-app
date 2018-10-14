 protoc -I pb/ pb/vote.proto --go_out=plugins=grpc:pb

 go run main.go -stderrthreshold=INFO

 psql -U $USERNAME -d $DB_NAME < $SQL_ROLE_SETUP_SCRIPT
