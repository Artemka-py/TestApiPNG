#!/bin/bash
app="docker.test"
docker build -t ${app} .
docker run -d -p 56733:80 --memory="2g" --cpuset-cpus="0" \
  --name=${app} \
  -v $PWD:/app ${app}
