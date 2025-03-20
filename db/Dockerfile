FROM postgres:16-alpine

COPY init.sql /docker-entrypoint-initdb.d/init.sql

ENV POSTGRES_USER $POSTGRES_USER
ENV POSTGRES_PASSWORD $POSTGRES_PASSWORD
ENV POSTGRES_DB "edd_db"

VOLUME [ "/var/lib/postgresql/data" ]

EXPOSE 5432
