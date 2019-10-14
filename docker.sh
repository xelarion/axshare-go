#!/usr/bin/env bash
sudo docker build -t ervincheung/axshare-go .
sudo docker-compose up --detach --build