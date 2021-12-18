#!/bin/bash
set -e

# Создать пользователя
# Установить пароль
# Назначить роль Суперпользователя
# Для созданной БД активировать расширение для генерации GUID сресдтвами БД

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER docker;

    ALTER USER docker WITH PASSWORD '4757';
    ALTER USER docker WITH SUPERUSER;

    CREATE DATABASE $DATABASE_NAME;

    GRANT ALL PRIVILEGES ON DATABASE $DATABASE_NAME TO docker;

    \connect $DATABASE_NAME;
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL
