# run postgresql docker container (--rm causes Docker to automatically remove the container when it exits.)

docker run -d --name postgres-1 -e POSTGRES_DB=lego_db -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -p 5432:5432 postgres:14.2-alpine

docker run --name postgres-1 --rm -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e PGDATA=/var/lib/postgresql/data/pgdata -v /tmp:/var/lib/postgresql/data -p 5432:5432 -it postgres:14.2-alpine

# run pdadmin docker container (connect to postgres container use it's ip (eg: 172.17.0.2) use "docker inspect postgres-1 | grep IPAddress ")

docker run --name pgadmin --rm -p 8080:80 -e 'PGADMIN_DEFAULT_EMAIL=qi.xu@carrier.com' -e 'PGADMIN_DEFAULT_PASSWORD=user' -d dpage/pgadmin4

# exec to postgresql

docker exec -it postgres-1 bash

# login to database

psql --username=postgres --dbname=postgres

# show connect info

\c

# show table list

\dt

# show content in "lego" table

SELECT \* FROM lego;

# exit

\q
