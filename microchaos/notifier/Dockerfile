# Build
FROM golang:1.22 AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/notifier ./main.go

# Runtime
FROM gcr.io/distroless/base-debian12
COPY --from=build /bin/notifier /notifier
EXPOSE 8082
ENTRYPOINT ["/notifier"]
