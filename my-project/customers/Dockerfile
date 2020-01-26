# golang image where workspace (GOPATH) configured at /go.

FROM golang:1.11.2 as dev
MAINTAINER shubham jagdhane

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace
ADD . /go/src/github.com/aj/my-project/customers

# Setting up working directory
WORKDIR /go/src/github.com/aj/my-project/customers

# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o customers .


# alpine image
FROM alpine:3.9.2 as prod
MAINTAINER shubham jagdhane

# install vim, curl and git
RUN apk update
RUN apk add vim
RUN apk add curl
RUN apk add git

# Setting up working directory
WORKDIR /root/
# copy movies binary
COPY --from=dev /go/src/github.com/aj/my-project/customers .
# Service listens on port 8080.
EXPOSE 8080
# Run the movies microservice when the container starts.
ENTRYPOINT ["./customers"]
