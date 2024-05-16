# Endpoint

ALL SUPPORT ROUTES
USER
http://127.0.0.1:8080/login
Body json

```json
{
     "name":"your_username",
     "password":"your_passwors"
}
```
http://127.0.0.1:8080/signup
Body json
``` json
{
     "name":"name",
	"email":"your_email",
	"password":"your_password",
	"phone_number":"phone_number",
	"address":"your_location"
}

REVIEW
http://127.0.0.1:8080/review
add review: POST 
```json
{
     "campsite_id":"int",
	"user_id":"int",
	"rating":"int",
	"comment":"string",
	"address":"time.Time"
}

```
