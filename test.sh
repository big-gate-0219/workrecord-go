#!/bin/bash

# Signup
echo "##### Signup"
curl -s -X POST http://127.0.0.1:8080/api/accounts/signup -H 'Content-Type:application/json' -d '{"userId":"dummy","email":"dummy@dummy.com", "userName":"dummy dummy", "password":"password"}' > result.json
cat result.json

# Signin
echo ""
echo "##### Signin"
curl -s -X POST http://127.0.0.1:8080/api/accounts/signin -H 'Content-Type:application/json' -d '{"name":"test1@mail.com","Password":"password"}' > result.json
cat result.json
jwt_token=`cat result.json | grep "token" | sed -e 's/{"token":"//' -e 's/"}//'`

# start of work
echo ""
echo "#### start of work"
curl -s -X POST  http://127.0.0.1:8080/api/work-records/today/start -H "x-auth-token: $jwt_token"

# end of work
echo ""
echo "#### end of work"
curl -s -X POST  http://127.0.0.1:8080/api/work-records/today/end -H "x-auth-token: $jwt_token"

# getWorkRecord
echo ""
echo "##### Get work record"
curl -s -X GET  http://127.0.0.1:8080/api/work-records -H "x-auth-token: $jwt_token"

# get my groups
echo ""
echo "##### Get my groups"
curl -s -X GET  http://127.0.0.1:8080/api/groups -H "x-auth-token: $jwt_token" > result.json
cat result.json

# get workrecords group today
echo ""
echo "#### Get workrecord group today"
curl -s -X GET  http://127.0.0.1:8080/api/work-records/groups/1/today -H "x-auth-token: $jwt_token"

# get group users
echo ""
echo "#### Get group users"
curl -s -X GET  http://127.0.0.1:8080/api/groups/1 -H "x-auth-token: $jwt_token"

# add group
echo ""
echo "#### Get group users"
curl -s -X POST http://127.0.0.1:8080/api/groups -H "x-auth-token: $jwt_token" -H 'Content-Type:application/json' -d  '{"group_name":"dummygroup","a":"a"}' > result.json
cat result.json
