# bouncecm - Bounce Change Manager
## Starting the container
Make sure you have docker running on your system first
then run `docker-compose up` if its the first time, use `docker-compose start` if not


10/22/18 - Added basic CRUD for bounce_rules.

GET /bounce_rules  List's all bounce rules
GET /bounce_rules/{id} Get single bounce rule
POST /bounce_rules  Create a bounce rule 
PUT /bounce_rules/{id} Update bounce rule

Send POST/PUT as json
Example:
    {response_code":"916","enhanced_code":"916","regex":"916","priority":1,"description":"Email was received","bounce_action":"return"}


