#!/bin/bash
#
curl http://localhost:8080/site \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"name": "chuck2"}'#!/bin/bash
