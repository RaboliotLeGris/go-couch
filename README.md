# go-couch

I cut a few corners to do the test in an acceptable time.

* On the performance side, it shouldn't be that great due to a lot of mapping between models.
* As said in the code, the bulk processing is not that great since the user will get an answer before it's effectively in the DB
* We should use context when possible
* Add Graceful shutdown of the worker with the router API could also be nice
* We should add a CI

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