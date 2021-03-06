PROJ=boidstorm
EXEDEST=/usr/bin

GOC=go build

.PHONY: build fmt get install precommit run test uninstall vet

build:
	${GOC}

fmt:
	go fmt $$(go list ./... | grep -v /${PROJ}/vendor/)

get:
	dep ensure

install:
	install ${PROJ} ${EXEDEST}/

precommit: fmt vet test

run: build
	./${PROJ}

test:
	go test $$(go list ./... | grep -v /${PROJ}/vendor/)

uninstall:
	rm ${EXEDEST}/${PROJ}

vet:
	go vet $$(go list ./... | grep -v /${PROJ}/vendor/)