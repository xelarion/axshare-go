#!/usr/bin/env bash
docker build -t ervincheung/axshare-go .
docker-compose up --detach --build