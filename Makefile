all: build
build: test
	CGO_ENABLED=0 go build 
test:
	go test -v
deps:
	brew install glide || :
	brew install go || :
	glide update
data:
	awk -F: '{print $$1}' /etc/passwd | grep -v '^#' | xargs -L1 -IXXX curl -XPUT -d '{"Name":"XXX"}' http://localhost:3000/employees
groups:
	curl http://localhost:3000/groups | jq .
docker:
	docker build -t lunchapp:latest .
