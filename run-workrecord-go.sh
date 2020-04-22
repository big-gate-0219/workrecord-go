#!/bin/bash

docker run -it --rm --name go-realizer -p 8080:8080 -v ./src:/go/src/app workrecord-go:1.0
