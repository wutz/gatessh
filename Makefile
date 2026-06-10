IMG ?= gatessh-controller:latest

.PHONY: build test crds deploy undeploy

build:
	go build -o bin/gatessh-controller ./cmd/controller/

test:
	go test ./... -v

crds:
	kubectl apply -f config/crds/

deploy: crds
	kubectl apply -f config/crds/
	kubectl apply -f examples/basic.yaml

undeploy:
	kubectl delete -f examples/basic.yaml --ignore-not-found
	kubectl delete -f config/crds/ --ignore-not-found

docker-build:
	docker build -t $(IMG) .

docker-push:
	docker push $(IMG)
