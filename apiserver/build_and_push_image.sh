#!/usr/bin/env bash
docker build -t dominikmatic/dddns-apiserver:latest .
docker push dominikmatic/dddns-apiserver:latest