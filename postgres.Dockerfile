FROM postgres:15.4

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        postgresql-contrib \
        postgresql-15-postgis \
    && rm -rf /var/lib/apt/lists/*

COPY ./scripts/init-postgis.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
