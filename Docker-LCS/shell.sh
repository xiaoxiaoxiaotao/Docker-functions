#!/bin/bash
docker build -t my-go-app:v1.0 .
docker run -p 8080:8080 --name my-gin-container -v "$(pwd):/Docker_gin_functions" my-go-app:v1.0