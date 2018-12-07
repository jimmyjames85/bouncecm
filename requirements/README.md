# Bounce Manager<a id="sec-1" name="sec-1"></a>

<div id="table-of-contents">
<h2>Table of Contents</h2>
<div id="text-table-of-contents">
<ul>
<li><a href="#sec-1">1. Bounce Manager</a>
<ul>
<li><a href="#sec-1-1">1.1. Backend API</a>
<ul>
<li><a href="#sec-1-1-1">1.1.1. Auth</a></li>
<li><a href="#sec-1-1-2">1.1.2. Read</a></li>
<li><a href="#sec-1-1-3">1.1.3. Write</a></li>
</ul>
</li>
<li><a href="#sec-1-2">1.2. Functionality</a>
<ul>
<li><a href="#sec-1-2-1">1.2.1. Logging</a></li>
<li><a href="#sec-1-2-2">1.2.2. Error Handling</a></li>
<li><a href="#sec-1-2-3">1.2.3. Testing</a></li>
<li><a href="#sec-1-2-4">1.2.4. Metrics</a></li>
</ul>
</li>
<li><a href="#sec-1-3">1.3. Stretch Goals</a>
<ul>
<li><a href="#sec-1-3-1">1.3.1. CI/CD</a></li>
</ul>
</li>
</ul>
</li>
</ul>
</div>
</div>

<a id="sec-1-1" name="sec-1-1"></a>
## Backend API

You will need to create an API that provides basic CRUD endpoints to
the bounce and changelog tables. The API is a contract that should be
designed by both the backend and frontend teams, as both teams should
agree on the same protocol. Take a look at
<https://petstore.swagger.io/> for a good example of an API.  For
bouncecm, please default to using `application/json` unless otherwise
specified, e.g. the csv export functionality. **Designing this API
should be priority number one.**

Here is an example of a potential bouncecm API call you might design
to add a new bounce rule. When the frontend client makes the call, the
backend server could respond with a number of different responses.

``` bash
# curl request

curl -X POST bouncehost:8080/bounce_rules/:userid -d '{"response_code":401,"enhanced_code":180,"regex":"^Error .* yahoo","priority":123,"description":"desc","bounce_action":"DROP"}'
```

Possible responses:

-   Bounce Rule Created
``` bash
< HTTP/1.1 201 Created
< Content-Type: application/json
{"bounce_id": 34}
```
-   Bad Request
``` bash
< HTTP/1.1 400 Bad Request
< Content-Type: application/json
{"error": "invalid regex: failed to compile"}
```
-   Unauthorized
``` bash
< HTTP/1.1 401 Unauthorized
< Content-Type: application/json
{"error":"unauthorized"}
```
-   Internal Server Error
``` bash
< HTTP/1.1 500 Internal Server Error
< Content-Type: application/json
{"error":"unknown"}
```

The front end should parse response codes and interpret as follows:

| HTTP Code  | Interpretation                                                                                   |
|------------|--------------------------------------------------------------------------------------------------|
| 2\*\*      | Request was succesful                                                                            |
| 3\*\*      | DONT USE                                                                                         |
| 4\*\*      | Frontend user issued a bad request e.g. `bouce_rule` doesn't exist or `regex` couldn't be parsed |
| 5\*\*      | Backend server is having issues, please try again                                                |

*Please make sure to use HTTP constants defined in go's [net/http package](https://golang.org/src/net/http/status.go#line9).*

---

Your API will need to provide the functionality listed below, but keep
in mind each section/bullet point need not be it's own separate
endpoint. When creating the API, consider changing the functionality
of the endpoint by providing different parameters. For example, the
following request could receive all the bounce_rules in either json or
csv by switching the value passed to `:format`:

``` bash
curl -X GET bouncehost:8080/bounce_rules/:userid/:format
```

The format may not even be passed at all, in which case the default json format should be used.

**Your API will need to provide the following functionality:**

<a id="sec-1-1-1" name="sec-1-1-1"></a>
### Auth

-   grant or deny access by userid
-   only subset of users have read access
-   even smaller subset has write access to create/edit endpoints

<a id="sec-1-1-2" name="sec-1-1-2"></a>
### Read

Provide endpoints that allow you to view current state of both the
bounce table and the changelog table.

-   Get all the current rules in place
-   Get all the current rules in .csv format
-   Get a single rule by rule_id
-   Get the entire changelog
-   Get the entire changelog for just a single rule_id
-   Get the latest N entries from the changelog (provide ability for pagination?)
-   Get the entire changelog in .csv format

<a id="sec-1-1-3" name="sec-1-1-3"></a>
### Write

Provide endpoints that allow creation and modification of bounce
rules. For any modification made to the bounce table, the **regex must
validated**, and the following must be logged to the change table:
-   Author
-   Timestamp
-   Rule ID
-   Change Comment (cannot be empty string)
-   Snapshot of the change that was made (i.e. all the same columns from the bounce table)

Create/Edit a rule

-   Create a new rule (**Validate input** e.g. no empty values. Every column from bounce table must be provided, except for auto-generated id)
-   Update an existing rule (At least on column must be provided. Validate input.)
-   Delete an existing rule
-   Revert a rule to a previous state. This should not be a separate endpoint. Considered this a roll forward. Use the same endpoint as "Update an existing rule".

<a id="sec-1-2" name="sec-1-2"></a>
## Functionality

Designing the API will inherently desing the functionality of your
application. However, you'll also want to add a `/healthcheck`
endpoint to the API. While the product manager has designed the
product features for the bounce manager, it's our responsibility to
make it work. If and *when* something goes wrong we're on the hook! So
we want to add functionality that gives us visibility into the health
and preformance metrics of our application. How many requests per
second are we handling?  What kind of errors are we seeing? What is
the status of our dependencies? The following requirements are
something we put into all of our services, and allow us us to monitor
and guage system health. **Adding functionality to help maintain and
monitor system health should be priority number two.**

<a id="sec-1-2-1" name="sec-1-2-1"></a>
### Logging

Logs should go to stderr, by using `log.Printf()` or
`log.Println()`. We want to know what events have happened throughout
the lifetime of the application. *What events do you think we should
log*? We definately want to log errors, and any modifications made to
the bounce table. Logs should be emitted as json (consider a logEvent
function), and should have the following parameters:

-   unix timestamp
-   what endpoint was hit
-   what userid hit the endpoint
-   if bounce rule was modified what was modified
-   if user without proper permission, reject request, and log userid

<a id="sec-1-2-2" name="sec-1-2-2"></a>
### Error Handling

Any time there is an error it should also be logged. Use judgment to
decide where in the callstack to log it. For example, lets say the DAO
encounters an error talking to the database, this should be bubbled up
to the calling function in the bouncecm server. However, we want to
include as much info as can. You can do this by using the
`errors.Wrap()` function, e.g.

``` go
return nil, errors.Wrap(err, "attempting to fetch rule_id from bounce table")
```

Once the error is bubbled up back to the server, the application
should then log the error (consider a logEvent function where
event=error).

For errors that *are* logged, they should have the following parameters:

-   unix timestamp
-   the error that occurred
-   if possible userid
-   if possible what endpoint was hit
-   don't log bad requests i.e 400 errors e.g. a request for non-existing ruleid should not be logged as an error

<a id="sec-1-2-3" name="sec-1-2-3"></a>
### Testing

We want an acceptance testing framework that should spin up a
dockerized mysql instance with a bounce table and changelog
table. Each of your tests should populate the table, hit an endpoint,
then verify the expected outcome occurred...or fail the test. You may
have to reset or tear down any test data so the next test can run. For
example, let's say we are testing a `/delete` endpoint for the bounce
table. The test might:

1.  Create a bounce rule in the bounce table with `rule_id=180`
2.  Hit the delete endpoint with `rule_id=80`. Verify rule 180 still exists. Verify client received `404 Status Not Found` response code.
3.  Hit the delete endpoint with `rule_id=180`. Verify rule 180 was deleted and client received `202 Status Accepted` response code.
4.  Hit the delete endpoint with `rule_id=180` again. Verify client received `404 Status Not Found` response code, because the rule no longer exists.
5.  Tear down for next test. In the event the test fails, rule 180 may still be in table, make sure to remove it for next test.

<a id="sec-1-2-4" name="sec-1-2-4"></a>
### Metrics

Keep in mind this is part of our MVP, however, it has lower
priority. We will further build out these requirements, as you near
completion.

Emit metrics where applicable, specifically
-   how long it took to process an entire request (consider middleware)
-   db response time
-   requests per second
-   error rate

<a id="sec-1-3" name="sec-1-3"></a>
## Stretch Goals

<a id="sec-1-3-1" name="sec-1-3-1"></a>
### CI/CD
