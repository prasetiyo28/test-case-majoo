FROM golang:latest



RUN mkdir -p /app

# Run go get swagger from official repository to generate swagger, it will installed in GOPATH, and then we can call binary from /go/bin/swag.
RUN go get -u github.com/swaggo/swag/cmd/swag@v1.6.7

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
# RUN go mod download
# Generate swagger from GOPATH.
RUN /go/bin/swag init -g ./cmd/main.go -o ./docs
RUN go get github.com/alecthomas/template
RUN go mod vendor

RUN go build -o ./app ./cmd/main.go

EXPOSE 8080

ENTRYPOINT ["./app"]