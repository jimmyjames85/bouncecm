# bouncecm - Bounce Change Manager

## Starting the container

Make sure you have docker running on your system first
then run `docker-compose up` if its the first time, use `./restartDocker.sh` if not,
alternatively you can use `./restartDocker.sh` anytime you want to refresh the docker instances.

## Scripts

There are several scripts within this repo created to be shortcuts for commonly used commands

### connectDatabase.sh

This script allows you to view the MySQL database

There are optional parameters for this script:

`./connectDatabase.sh [dev | test | show] [show] [b | c | u]`

Running the script without any parameters will run: `./connectDatabase.sh dev` which will connect you to the dev database

Running `./connectDatabase.sh show` will show you all the table in the dev database

You can further specify which database to look into using the first parameter, using `dev` for the dev database and `test` for the test database

Using `show` along with `b` or `c` or `u` will show you the bounce_rule, changelog, and user tables accordingly

### clearDocker.sh

This script will stop and remove all docker containers that are currently running

### restartDocker.sh

This script calls `./clearDocker.sh` and then runs `docker-compose up -d` based upon the `docker-compose.yaml` file within the root directory

## Server

First, you'll need to get into the docker container:

`docker exec -it bouncecm_go bash`

From here, check if your server was successfully started you can use one of the following commands:

```bash
Get:
`curl -X GET localhost:3000/bounce_rules`
```

```bash
Post:
`curl -X POST -H 'Content-Type: application/json' -d '{"response_code":123, "enhanced_code":"1.2.4", "regex":"testing", "priority":123, "description":"This is for testing", "bounce_action":"AUTOINCREMETTESTING"}' localhost:3000/bounce_rules/`
```

```bash
Update:
curl -X PUT -H "Content-Type: application/json" -d '{"id":505,"response_code":123,"enhanced_code":"1.2.4","regex":"testing","priority":123,"description":"This is for testing","bounce_action":"AUTOINCREMETTESTING","user_id":0,"comment":"something"}' localhost:3000/bounce_rules/505
```

```bash
Delete:
curl -X DELETE -H "Content-Type: application/json" -d '{"id":354, "user_id":1, "created_at":999999, "comment":"DEletetesting",  "response_code":123, "enhanced_code":"1.2.4", "regex":"testing", "priority":123, "description":"This is for testing", "bounce_action":"PUTTESTING"}' localhost:3000/bounce_rules/354
```

````
Note: /change_logs & /change_logs/{id} require query parameters limit and offset
    example:  /change_logs/?limit=10&offset=0 & /change_logs/{id}?limit=10&offset=0

    /bounce_rules  require query parameters limit and offset
    example:  /bounce_rules/?limit=10&offset=0 ```

## Testing

To run the test suites for the backend, you'll first need to be in the testing container:

`docker exec -it bouncecm_test bash`

and then get into the testing directory

`cd src/github.com/jimmyjames85/bouncecm/internal/integration`

From here you will be able to:

### Run all of the tests

`go test -v`

This will run all of the tests within the directory, which include:

- bounce_rule route tests

- change_log route tests

- database tests

### Run tests for the bounce_rules routes

`go test -v main_test.go bounce_rule_test.go`

This will check the GET, POST, PUT, and DELETE bounce_rule routes

### Run tests for the change_logs routes

`go test -v main_test.go change_log_test.go`

This will check the GET, POST, PUT, and DELETE change_log routes

### Run tests for the database

`go test -v main_test.go database_test.go`

This will check for MySQL commands working with the bounce_rule schema

### Alternatively, you can run the following command to just test all of them once you enter the container

`go test -v src/github.com/jimmyjames85/bouncecm/internal/integration/*.go`

````

### Socket

Concurrent users are implemented through a simple socketed connection
Library used for socket implementation: <https://https://github.com/olahol/melody/>

The server uses a in-memory `map[string]Time` to keep track of rules being edited

The messages emitted by the client should follow the pattern of:
`<COMMAND>:<BOUNCE RULE ID>`

As of right now, there are only three commands that the backend socket takes:

- edit
- release
- check

#### edit

`edit:504`

Returns `EDIT`, `INUSE`

Allocates the timeout for the specified bounce rule in the internal data structure.
Returns `EDIT` if successful and `INUSE` if user is already editing.

#### release

`release:504`

Returns `FREE`

Removes the entry for the specified bounce rule in the map and broadcasts `FREE` to all other connected sessions to update their UIs

#### check

`check:504`

Returns `FREE`

Interface for checking whether or not a certain bounce rule is being edited by another
user. Only returns `FREE` if applicable (No edits / Timeout)
