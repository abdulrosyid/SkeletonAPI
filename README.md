# Warung Pintar Test

* To run :

```sh
$ go run main.go
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
