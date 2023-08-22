
 Start from golang v.1.11 base image
FROM golang:1.11

LABEL maintainer="Valerie Wyns"

# Set current working directory inside the container
WORKDIR ./app

# Copy everything from the current directory to the PWD(present working directory) inside the container
COPY ./go.mod ./go.sum ./

# Download all dependencies
RUN go mod download

COPY ./ ./

# install the package
RUN go install -v ./...

# expose port 8080 to the outside world
EXPOSE 8080

# run executable
CMD ['go', "run", "main.go"]
















# Start from golang v.1.11 base image
#FROM golang:1.11
#
#LABEL maintainer="Valerie Wyns"
#
## Set current working directory inside the container
#WORKDIR ./app
#
## Copy everything from the current directory to the PWD(present working directory) inside the container
#COPY . .
#
## Download all dependencies
#RUN go get -d -v ./...
#
## install the package
#RUN go install -v ./...
#
## expose port 8080 to the outside world
#EXPOSE 8080
#
## run executable
#CMD ['go', "run", "main.go"]