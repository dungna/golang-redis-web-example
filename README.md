# golang-redis-web-example
### INSTALL DRIVER
```go get github.com/mediocregopher/radix.v2```
### ADD DATA FOR REDIS
```127.0.0.1:6379>HMSET album:1 title "Electric Ladyland" artist "Jimi Hendrix" price 4.95 likes 8
127.0.0.1:6379>HMSET album:2 title "Back in Black" artist "AC/DC" price 5.95 likes 3
127.0.0.1:6379>HMSET album:3 title "Rumours" artist "Fleetwood Mac" price 7.95 likes 12
127.0.0.1:6379>HMSET album:4 title "Nevermind" artist "Nirvana" price 5.95 likes 8```

### START
```go run main.go```
### Functions
| Method | Path | Function
| ------ | ------ | ------ |
| GET | /album?id=1 | Show details of a specific album (using the id provided in the query string) |
| POST | /likes | Add a new like for a specific album (using the id provided in the request body) |
| GET | /popular | List the top 3 most liked albums in order |
