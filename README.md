# swapi_go
Go application for searching Star Wars movies data

 INTRODUCTION
------------

This is a console application written in Golang
that lets you search for Star Wars characters and
their information such as name, home planet, startships, etc...
It makes use of https://swapi.dev/ which is a Star Wars API


REQUIREMENTS
-------------

Golang 1.20

RUNNING
-------------
To run, simply do
`go run main.go`

You'll be prompted with:
`Enter the Star Wars character name: `

You can enter a letter or a more specfic character name like
* `c`
* `Lando Calrissian`

Entering a full character name will likely return just one character.
Entering a letter or letters will return a more broader result set.

TODO
---------------
* More robust menu and loop to run again
* Unit tests
* Support things beyond character search
* introduce data store to speed up retrieval
* ...



