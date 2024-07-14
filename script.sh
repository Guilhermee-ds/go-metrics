#!/bin/bash

URL="http://localhost:8080/payments"
HEADERS="Content-Type: application/x-www-form-urlencoded"

send_request() {
  STATUS=$1
  curl -X POST $URL -H "$HEADERS" -d "status=$STATUS"
}

while true; do
  RANDOM_STATUS=$((RANDOM % 10))

  if [ $RANDOM_STATUS -lt 7 ]; then
    send_request "success"
  else
    send_request "failed"
  fi

  sleep 1

done
