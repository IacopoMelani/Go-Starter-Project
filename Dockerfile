# DOCKER FILE EXAMPLE

# MARK: Go build stage

FROM golang:alpine AS golang

WORKDIR /app/go-build

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN go build -o your-awesome-app

# MARK: Prepare prod stage

FROM alpine

WORKDIR /app/your-awesome-app

COPY --from=golang /app/go-build/your-awesome-app .

RUN rm -rf /app/go-build

CMD ["./your-awesome-app", "-fire", "go!"]

EXPOSE 8888
