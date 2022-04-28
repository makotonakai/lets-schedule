# !/bin/bash

# MySQLサーバーが起動するまで待機する
until mysqladmin ping -h ${DB_HOST} -P ${DB_PORT} --silent; do
  echo 'waiting for mysqld to be connectable...'
  sleep 2
done

echo "app is starting...!"
exec go run main.go