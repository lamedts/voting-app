#!/bin/bash

VOLUME_CONTAINER_NAME=pgdata
CONTAINER_NAME=pgdb
USERNAME=testuser1
PASSWORD=password123!
DB_NAME=pgdata

function start_container {
	docker run --name $CONTAINER_NAME \
	-v $VOLUME_CONTAINER_NAME:/var/lib/postgresql/data \
	-e POSTGRES_USER=$USERNAME -e POSTGRES_PASSWORD=$PASSWORD \
	-e POSTGRES_DB=$DB_NAME \
	-e POSTGRES_ENABLE_SSL \
	-p 5432:5432 \
	-d postgres:9.5.7
}

volumes=$(docker volume ls | tr -s ' ' | cut -d' ' -f2 | tail -n +2)
if [[ -z $(echo $volumes | grep -x "$VOLUME_CONTAINER_NAME") ]]; then
	echo "Creating docker volume $VOLUME_CONTAINER_NAME..."
	docker volume create $VOLUME_CONTAINER_NAME
else
	echo "Use existing docker volume $VOLUME_CONTAINER_NAME"
fi

containers=$(docker ps -a | tr -s ' ' | rev | cut -d' ' -f1 | rev | tail -n +2)
if [[ -z $(echo $containers | grep -x "$CONTAINER_NAME") ]]; then
	echo "Creating docker container $CONTAINER_NAME..."
	start_container
else
	echo "Use existing docker container $CONTAINER_NAME"
fi
