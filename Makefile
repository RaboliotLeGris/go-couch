-include .env

default:
	go run .

start_couchdb:
	docker run --rm -p 5984:5984 -e COUCHDB_USER="$(COUCHDB_USER)" -e COUCHDB_PASSWORD="$(COUCHDB_PASSWORD)" --name go-couch couchdb:3.2.2
init_couchdb:
# TODO: use a more elegant way to pass the password and make it dynamic  
	curl -X PUT -H "Authorization:Basic YWRtaW46cGFzc3dvcmQ" http://127.0.0.1:5984/_users
	curl -X PUT -H "Authorization:Basic YWRtaW46cGFzc3dvcmQ" http://127.0.0.1:5984/_replicator
	curl -X PUT -H "Authorization:Basic YWRtaW46cGFzc3dvcmQ=" http://127.0.0.1:5984/_global_changes

test: lint
	go test ./...

lint:
	golangci-lint run

fmt:
	gofmt -s -w .