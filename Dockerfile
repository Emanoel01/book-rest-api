FROM golang:latest as build

WORKDIR /go-app

COPY src/* .

COPY go.mod .

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main


FROM alpine:latest as final

WORKDIR /go-app
EXPOSE 1337
COPY --from=build /go-app/src/main .
CMD ["./main"]