FROM golang:latest

RUN set -e \
	&& mkdir -p /go/src/github.com/ieee0824/dummy-app

COPY main.go /go/src/github.com/ieee0824/dummy-app

CMD ["go", "run", "/go/src/github.com/ieee0824/dummy-app/main.go"]
