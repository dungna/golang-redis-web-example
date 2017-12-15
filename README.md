# golang-redis-web-example
### Start
```go run main.go```
### Functions
| Method | Path | Function
| ------ | ------ | ------ |
| GET | /album?id=1 | Show details of a specific album (using the id provided in the query string) |
| POST | /likes | Add a new like for a specific album (using the id provided in the request body) |
| GET | /popular | List the top 3 most liked albums in order |
