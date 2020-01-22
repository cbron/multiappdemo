FROM golang:1.12.1-alpine3.9
RUN apk add --no-cache curl
ENV GOPATH="/go"
RUN ["mkdir", "-p", "/go/src/github.com/cbron/multiappdemo/frontend"]
COPY * /go/src/github.com/cbron/multiappdemo/frontend/
WORKDIR /go/src/github.com/cbron/multiappdemo/frontend
RUN ["go", "build", "-o", "frontend"]
CMD ["./frontend"]
