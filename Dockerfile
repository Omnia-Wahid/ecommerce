FROM golang:latest

RUN mkdir /go/src/ecommerce

#ENV GOPATH=/go

WORKDIR /go/src/ecommerce

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ecommerce .

#RUN go build -v -o /usr/local/bin/go/src/ecommerce ./...

EXPOSE 3001

CMD [ "./ecommerce" ]
