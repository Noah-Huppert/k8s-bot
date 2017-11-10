FROM golang:alpine

# Go Path
WORKDIR /go
ENV GOPATH /go

# Install Git: To install Dep
RUN apk update; apk add git

# Install Dep tool
RUN go get -u github.com/golang/dep/cmd/dep

# Copy source
ENV APP_PATH $GOPATH/src/github.com/Noah-Huppert/kube-bot
RUN mkdir -p "$APP_PATH"
WORKDIR "$APP_PATH"

COPY . .

# Install dependencies
RUN dep ensure

# Start
CMD go run main.go
