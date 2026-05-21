#!/bin/bash

docker stop backend
docker rm backend
docker run -d -p 8080:8080 --name backend $PREVIOUS_IMAGE