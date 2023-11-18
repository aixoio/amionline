FROM golang:1.21.4 AS deps

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

FROM deps AS build

COPY . /app
RUN go clean
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /app/server

FROM alpine:latest

COPY --from=build /app/server /

EXPOSE 9090

CMD [ "/server" ]
