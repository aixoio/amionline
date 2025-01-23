FROM mysql:9.2.0-oraclelinux9

WORKDIR /app

ENV MYSQL_ROOT_PASSWORD=amionline
ENV MYSQL_DATABASE=amionline

COPY structure.sql /docker-entrypoint-initdb.d/
