# bouncecm - Bounce Change Manager
## Starting the container
Make sure you have docker running on your system first
then run `docker-compose up` if its the first time, use `docker-compose start` if not

Testing Routes For changelog
Get:
`curl -X GET localhost:3000/bounce_rules`

Post:
`curl -X POST -H 'Content-Type: application/json' -d '{"response_code":123, "enhanced_code":"1.2.4", "regex":"testing", "priority":123, "description":"This is for testing", "bounce_action":"AUTOINCREMETTESTING"}' localhost:3000/ bounce_rules/`

Update:
`curl -X PUT -H "Content-Type: application/json" -d '{"id":508, "response_code":123, "enhanced_code":"1.2.4", "regex":"testing", "priority":123, "description":"This is for testing", "bounce_action":"PUTTESTING"}' localhost:3000/bounce_rules/508`

Delete:
`curl -X DELETE localhost:3000/bounce_rules/508`
