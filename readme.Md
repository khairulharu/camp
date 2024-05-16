# Endpoint

ALL SUPPORT ROUTES

USER
http://127.0.0.1:8080/login
Login User: POST
Content-Type:appliaction/json
Body json
With Return JWT token
```json
{
     "name":"string",
     "password":"string"
}
```
http://127.0.0.1:8080/signup
SignUp User: POST
Content-Type:appliaction/json
Body json
Return Response OK if no error
``` json
{
     "name":"string",
	"email":"string",
	"password":"string",
	"phone_number":"string",
	"address":"string"
}

REVIEW
http://127.0.0.1:8080/review
Add review: POST 
Content-Type:appliaction/json
Body JSON
Reurn Response OK if no error
```json
{
     "campsite_id":"int",
	"user_id":"int",
	"rating":"int",
	"comment":"string",
	"address":"time.Time"
}

```
http://127.0.0.1:8080/review/id
get review: POST
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
