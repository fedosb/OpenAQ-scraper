FROM golang:1.19.0-alpine3.16 AS build
WORKDIR /src

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

RUN go build -a -installsuffix cgo -o bin/app ./cmd/main.go

FROM alpine:3.9 AS bin
COPY --from=build /src/bin/app /

# Export necessary port
EXPOSE 8081
CMD ["./app"]
