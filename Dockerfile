FROM golang:1.9 as base
WORKDIR /usr/src
COPY . .
RUN go get -d -v github.com/gorilla/mux
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o main

FROM scratch
COPY --from=base /usr/src/main /go-http-microservice
CMD ["/go-http-microservice"]
