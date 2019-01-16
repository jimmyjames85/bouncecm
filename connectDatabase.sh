#!/bin/sh

container=""
commands=""
table=""
if [ $# -eq 0 ]
    then
        echo "\nConnecting to dev MySQL docker container: droprules.database.dev\n"
        container="droprules.database.dev"
fi

if [ "$1" == "show" ]; then
    echo "\nConnecting to dev MySQL docker container: droprules.database.dev\n"
    container="droprules.database.dev"
    echo "\nGrabbing tables from droprules.database.dev\n"
    commands="-e use drop_rules; show tables;"

elif [ "$1" == "dev" ]; then
    echo "\nConnecting to dev MySQL docker container: droprules.database.dev\n"
    container="droprules.database.dev"
    table="drop_rules"
elif [ "$1" == "test" ];then
    echo "\nConnecting to test MySQL docker container: droprules.database.test\n"
    container="droprules.database.test"
    table="test_rules"
fi

if [ "$2" == "show" ]; then
    echo "\nGrabbing tables from "$container"\n"
    commands="-e use "$table"; show tables;"
fi

if [ "$3" == "b" ]; then
    commands+="select * from bounce_rule"
elif [ "$3" == "c" ]; then
    commands+="select * from changelog"
elif [ "$3" == "u" ]; then
    commands+="select * from user"
fi

docker exec -it "$container" mysql -uroot -proot "$commands"