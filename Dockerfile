FROM golang:1.18 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o myapp .

FROM ubuntu:20.04
COPY --from=build /app/myapp /usr/local/bin
CMD ["myapp"]