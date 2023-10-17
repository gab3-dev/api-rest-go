# Go API-RESTful with Gin
## Go Lang
Go is a programming language created by Google and released in open source in November 2009. It is a compiled language focused on productivity and concurrency programming.
**[Learn more about Go](https://go.dev/)**.
## Gin Framework
Gin is a web framework written in Golang.
It features a Martini-like API, but with performance up to 40 times faster than Martini.
**[Learn more about Gin](https://gin-gonic.com/)**.
## Creating my GO RESTful API
To contribute to my Back-End studies, I decided to develop two APIs in different technologies. In the future I intend to use both APIs developed in a personal project related to games. 
For now both are being done in the REST architecture, because it is the most common, it is easier to work with, later one of the architectures will change to RPC.
## Run
***For this tutorial, I conclude that you already has installed the Go Lang. If it isn't your case, you can download and install it at this link:*** **[Go Lang Download](https://go.dev/dl/)**
<br>First you must get the packages with this command:
```
go get .
```
Now just run with:
```
go run .
```
## API Methods
GET
```bash
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"
```
GET:id
```bash
curl http://localhost:8080/albums/2 \
    --header "Content-Type: application/json" \
    --request "GET"
```
POST
```bash
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```
DELETE
```bash
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "DELETE" \
    --data '{"id": "1}'
```
