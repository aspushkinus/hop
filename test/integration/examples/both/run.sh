#!/bin/bash

# exit on fail
set -e

mix deps.get

echo "====== Ping dependencies ====="
cd support/ping && cargo build
cd support/ping && cargo run $APP_HOST:$APP_PORT
cd support/ping && cargo run $WEBDRIVER_HOST:$WEBDRIVER_PORT

# # Wait for Elm
# until nc -z $APP_HOST $APP_PORT; do
#   echo "Waiting for $APP_HOST ..."
#   sleep 1
# done

# # Wait for Webdriver
# until nc -z $WEBDRIVER_HOST $WEBDRIVER_PORT; do
#   echo "Waiting for $WEBDRIVER_HOST ..."
#   sleep 1
# done

mix test test/$TEST_DIR
