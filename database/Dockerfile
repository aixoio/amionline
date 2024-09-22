FROM mysql:9.0.1-oraclelinux9

WORKDIR /app

ENV MYSQL_ROOT_PASSWORD=amionline
ENV MYSQL_DATABASE=amionline

COPY structure.sql /docker-entrypoint-initdb.d/
