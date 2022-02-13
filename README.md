run postgresql docker container

docker run -d --name postgres-1 -e POSTGRES_DB=lego_db -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 postgres:14.2-alpine

docker run --name postgres-1 --rm -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e PGDATA=/var/lib/postgresql/data/pgdata -v /tmp:/var/lib/postgresql/data -p 5432:5432 -it postgres:14.2-alpine