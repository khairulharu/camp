# Endpoint

ALL SUPPORT ROUTES

Api for booking campsite use clean architecture pattern, this all we need for creating appliaction booking, if you need business logic for complex code you can edit in `/internal/service`.

# USER
http://127.0.0.1:8080/login
Login User: POST
Content-Type:appliaction/json
Body json
```json
{
     "name":"string",
     "password":"string"
}
```

Reurn Response OK if no error
```json
{
    "token_string_jwt"
}
```
With Return JWT token
http://127.0.0.1:8080/signup
SignUp User: POST
Content-Type:appliaction/json
Body json
``` json
{
     "name":"string",
	"email":"string",
	"password":"string",
	"phone_number":"string",
	"address":"string"
}
```
Return Response OK if no error
```json
{
    "status": "string",
    "message": "string",
    "data": nill,
    "error": "string"
}
```

# REVIEW
http://127.0.0.1:8080/review
Add review: POST 
Content-Type:appliaction/json
Body JSON
```json
{
     "campsite_id":"int",
	"user_id":"int",
	"rating":"int",
	"comment":"string",
	"address":"time.Time"
}

```
Return Response OK if no error
```json
{
    "status": "string",
    "message": "string",
    "data": nill,
    "error": "string"
}
```
http://127.0.0.1:8080/review/id
get review: GET
Return Response OK with Data, if no error
```json
{
    "status": "string",
    "message": "string",
    "data": {
     "campsite_id":"int",
	"user_id":"int",
	"rating":"int",
	"comment":"string",
	"address":"time.Time"
},
    "error": "string"
}
```