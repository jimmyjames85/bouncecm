#!/bin/sh
./clearDocker.sh
echo Composing
docker-compose up -d 
if [ -z `docker ps -q --no-trunc | grep $(docker-compose ps -q bouncecm)` ]; then
  echo "No, it's not running."
else
  echo "Go server started, access through: \n  docker exec -it bouncecm_go bash"
fi
