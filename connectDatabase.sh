#!/bin/sh

container=""
commands=""
table=""
if [ $# -eq 0 ]
    then
        echo "\nConnecting to dev MySQL docker container: bouncecm_database_1\n"
        container="bouncecm_database_1"
fi

if [ "$1" == "show" ]; then
    echo "\nConnecting to dev MySQL docker container: bouncecm_database_1\n"
    container="bouncecm_database_1"
    echo "\nGrabbing tables from bouncecm_database_1\n"
    commands="-e use drop_rules; show tables;"

elif [ "$1" == "dev" ]; then
    echo "\nConnecting to dev MySQL docker container: bouncecm_database_1v\n"
    container="bouncecm_database_1"
    table="drop_rules"
    commands="-e use drop_rules;"
elif [ "$1" == "test" ];then
    echo "\nConnecting to test MySQL docker container: bouncecm_test_database_1\n"
    container="bouncecm_test_database_1"
    commands="-e use test_rules;"
fi

if [ "$2" == "show" ]; then
    echo "\nGrabbing tables from "$container"\n"
    commands+="show tables;"
elif [ "$2" == "b" ]; then
    commands+="select * from bounce_rule"
elif [ "$2" == "c" ]; then
    commands+="select * from changelog"
elif [ "$2" == "u" ]; then
    commands+="select * from user"
fi

if [ "$3" == "b" ]; then
    commands+="select * from bounce_rule"
elif [ "$3" == "c" ]; then
    commands+="select * from changelog"
elif [ "$3" == "u" ]; then
    commands+="select * from user"
fi

docker exec -it "$container" mysql -uroot -proot "$commands"