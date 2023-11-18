FROM mysql:latest

WORKDIR /app

ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=amionline

COPY structure.sql /docker-entrypoint-initdb.d/
