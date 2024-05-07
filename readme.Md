# Endpoint

ALL SUPPORT ROUTES
http://127.0.0.1:8080/login
Body json

```json
{
     "name":"your_username",
     "password":"your_passwors"
}
```
app.Post("/signup", handler.SignUp)
Body json
``` json
{
     "name":"name",
	"email":"your_email",
	"password":"your_password",
	"phone_number":"phone_number",
	"address":"your_location"
}

```