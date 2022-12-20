# go-couch

## Required
* Go - 1.19
* Docker
* Make
* Curl

## How to launch

* Copy .env.template to .env (and set the wanted password (it will be easier to use `admin:password`))
* In a terminal, start CouchDB with `make start_couchdb`
* Then in another terminal execute `make init_couchdb`, it will creates the required databases
* Now you can launch the API with `make`

## To improve
* Better startup of the couchDB database
* Pass the password more cleaning in the db init phase
* Use JWT instead of basic auth