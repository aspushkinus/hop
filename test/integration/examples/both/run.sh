#!/bin/bash

mix deps.get

cd ping && cargo run $APP_HOST:$APP_PORT
cd ping && cargo run $WEBDRIVER_HOST:$WEBDRIVER_PORT

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
