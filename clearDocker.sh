#!/bin/sh
echo Stopping containers
docker stop $(docker ps -aq)
echo Removing containers
docker rm $(docker ps -aq)
