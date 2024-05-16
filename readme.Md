# Endpoint

ALL SUPPORT ROUTES
USER
http://127.0.0.1:8080/login
Body json

```json
{
     "name":"string",
     "password":"string"
}
```
http://127.0.0.1:8080/signup
Body json
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
http://127.0.0.1:8080/review/id
get review: POST 
```json
{
     "campsite_id":"int",
	"user_id":"int",
	"rating":"int",
	"comment":"string",
	"address":"time.Time"
}

```
