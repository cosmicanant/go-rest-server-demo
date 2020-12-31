# go-sample-app
Learning GO

# steps to start local server on port: 3000
- go run main.go

# sample POST request to add a new user
```
curl -X POST \
  http://localhost:3000/users \
  -H 'cache-control: no-cache' \
  -d '
	{"FirstName": "Iron", "LastName": "Man"}
'
```

# sample GET ALL USERS
```
curl -X GET \
  http://localhost:3000/users \
  -H 'accept: application/json'
```