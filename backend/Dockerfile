FROM golang:1.12.1-alpine3.9
RUN apk add --no-cache curl
ENV GOPATH="/go"
RUN ["mkdir", "-p", "/go/src/github.com/cbron/multiappdemo/backend"]
COPY * /go/src/github.com/cbron/multiappdemo/backend/
WORKDIR /go/src/github.com/cbron/multiappdemo/backend
RUN ["go", "build", "-o", "backend"]
CMD ["./backend"]
