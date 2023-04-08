.PHONT: test

test:
	go test $(shell go list ./... |grep -v ent)
