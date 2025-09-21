# Build
FROM golang:1.22 AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/order-api ./main.go

# Runtime
FROM gcr.io/distroless/base-debian12
ENV GIN_MODE=release
COPY --from=build /bin/order-api /order-api
EXPOSE 8081
ENTRYPOINT ["/order-api"]
