# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.11.2 as dev

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace
ADD . /go/src/github.com/aj/my-project/shops

# Setting up working directory
WORKDIR /go/src/github.com/aj/my-project/shops

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shops .


# alpine image
FROM alpine:3.9.2 as prod
# Setting up working directory
WORKDIR /root/
# copy shops binary
COPY --from=dev /go/src/github.com/aj/my-project/shops .
# Service listens on port 8080.
EXPOSE 8080
# Run the shops microservice when the container starts.
ENTRYPOINT ["./shops"]
