#!/bin/bash

# Wait for Elm
echo "Looking for app at $APP_HOST:$APP_PORT"
until nc -z $APP_HOST $APP_PORT; do
  echo "Waiting for $APP_HOST ..."
  sleep 1
done
echo "Found $APP_HOST"

# Wait for Webdriver
echo "Looking for webdriver at $WEBDRIVER_HOST:$WEBDRIVER_PORT"
until nc -z $WEBDRIVER_HOST $WEBDRIVER_PORT; do
  echo "Waiting for $WEBDRIVER_HOST ..."
  sleep 1
done
echo "Found $WEBDRIVER_HOST"

ginkgo
