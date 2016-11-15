# urlhortener
[![Build Status](https://travis-ci.org/xlk3099/urlshortener.svg?branch=master)](https://travis-ci.org/xlk3099/urlshortener)
[![Go Report Card](https://goreportcard.com/badge/github.com/xlk3099/urlshortener)](https://goreportcard.com/report/github.com/xlk3099/urlshortener)
[![codecov](https://codecov.io/gh/xlk3099/urlshortener/branch/master/graph/badge.svg)](https://codecov.io/gh/xlk3099/urlshortener)

A simple url shortening service written in Go, it encodes an input long url to a shortened url by base62 encoding its request_id.

* If request_id is in 0 ~ 61, the shortened url  would be 1 char
* If request_id is in 62 ~ 3843, the shortened url would be 2 chars
* If request_id is in 3844 ~ 238327, the shortened url 3 chars
* If request_id is in 238328 ~ 14776335, the shortened url 4 chars
* If request_id is in 14776336 ~ 916132831, shortened url would be 5 chars
* ...

### TechStack:
  - Http server using Gin framework
  - MongoDB to store orignal & shortened url data
  - Dockerfile & docker_compose.yml to help easy deploy

### Installation:
To build the docker image
```sh
$ docker-compose build
```
To start 
```sh
$ docker-compose up
```
The application is running on port 8080, default port of Gin

Example to post a shorten an existing long url request:
```sh
$ curl -sX POST -H 'Content-Type: application/json' 'localhost:8080/shorten' -d '{"url":"http://a.very.long.url"}'
```
Response:
```sh
HTTP 200
'{"short":"http://localhost/a"}'
```
Example to get original url:
```sh
$ curl -sX GET -H 'Content-Type: application/json' 'localhost:8080/original' -d '{"short":"http://localhost/a"}'
```
Response
```sh
HTTP 200
'{"original":"http://a.very.long.url"}'
```

