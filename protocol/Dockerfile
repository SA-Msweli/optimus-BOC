FROM golang:1.20-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app ./...

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /app /app
EXPOSE 8080
CMD ["/app"]