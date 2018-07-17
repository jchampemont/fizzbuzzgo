# FizzBuzzGo

[![Build Status](https://travis-ci.org/jchampemont/fizzbuzzgo.svg?branch=master)](https://travis-ci.org/jchampemont/fizzbuzzgo)

FizzBuzzGo is a simple fizz-buzz REST server written in Go.

(Please be kind, this is my very first Go program.)

It exposes a REST API endpoint (/fizzBuzz) accepting five query parameters :
  - string1 (default value "fizz")
  - string2 (default value "buzz")
  - int1 (default value 3)
  - int2 (default value 5)
  - limit (default value 100)

and it returns a list of strings with numbers from 1 to limit where :

  - all multiples of int1 are replaced by string1
  - all multiples of int2 are replaced by string2
  - all multiples of int1 and int2 are replaces by string1string2


It uses a single dependency (github.com/gorilla/mux) but don't forget to run `go get ./...` beforehand.

## Examples
````
$ curl http://localhost:8080/fizzBuzz
["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz","13","14","fizzbuzz","16","17","fizz","19","buzz","fizz","22","23","fizz","buzz","26","fizz","28","29","fizzbuzz","31","32","fizz","34","buzz","fizz","37","38","fizz","buzz","41","fizz","43","44","fizzbuzz","46","47","fizz","49","buzz","fizz","52","53","fizz","buzz","56","fizz","58","59","fizzbuzz","61","62","fizz","64","buzz","fizz","67","68","fizz","buzz","71","fizz","73","74","fizzbuzz","76","77","fizz","79","buzz","fizz","82","83","fizz","buzz","86","fizz","88","89","fizzbuzz","91","92","fizz","94","buzz","fizz","97","98","fizz","buzz"]

$ curl http://localhost:8080/fizzBuzz\?limit\=12
["1","2","fizz","4","buzz","fizz","7","8","fizz","buzz","11","fizz"]

$ curl http://localhost:8080/fizzBuzz\?limit\=12\&string1\=fuzz
["1","2","fuzz","4","buzz","fuzz","7","8","fuzz","buzz","11","fuzz"]

$ curl http://localhost:8080/fizzBuzz\?limit\=a
among expected integer parameters (int1, int2 and limit) some were not integers
````