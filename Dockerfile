FROM golang:1.13.6

RUN mkdir -p $GOPATH/src/casual-talk
WORKDIR $GOPATH/src/casual-talk
COPY . $GOPATH/src/casual-talk
RUN go get github.com/go-sql-driver/mysql
RUN go build .

EXPOSE 8080

ENTRYPOINT ["./casual-talk"]

