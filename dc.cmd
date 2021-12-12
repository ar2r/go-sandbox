@echo off
REM Командный файл для OS Windows

if "%1"=="up" goto up
if "%1"=="down" goto down
if "%1"=="rebuild" goto rebuild
if "%1"=="flush-db" goto flush_db

:up
    cd devenv
    docker-compose up -d
    cd ..
goto finish

:down
    cd devenv
    docker-compose down
    cd ..
goto finish

:rebuild
    cd devenv
    docker-compose down
    docker-compose up -d --build --remove-orphans
    cd ..
goto finish

:flush_db
    print "Execute dropping..."
    php src/artisan doctrine:oracle:reset -f
    print "Execute migrations..."
    php src/artisan doctrine:migrations:migrate
    print "Execute seeding..."
    php src/artisan db:seed --class=\\Database\\Seeders\\TestSeeder
goto finish

:finish
