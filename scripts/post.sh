#!/bin/bash
#
# Populates the site tables
#

curl http://localhost:8080/site \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "site 1", "created_at": "today"}'

curl http://localhost:8080/site \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "site 2", "created_at": "today"}'

curl http://localhost:8080/site \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "site 3", "created_at": "today"}'

curl http://localhost:8080/site \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "site 4", "created_at": "today"}'

