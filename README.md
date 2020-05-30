# Skeleton API

* Content root path main : 
```sh
cmd/rest/main.go
```
* To run :
```sh
$ go run cmd/rest/main.go
```
* Script cURL for add message :
```sh
curl --location --request POST 'http://localhost:8080/wp/api/message/add?message=Test'
```
* Script cURL for get message :
```sh
curl --location --request GET 'http://localhost:8080/wp/api/message/get'
```
* open on browser for send message and get retrieve message (WebSocket):
```sh
http://localhost:8080/wp/api/message/client
```
