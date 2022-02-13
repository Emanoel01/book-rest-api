FROM golang:latest as build

WORKDIR /go-app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main


FROM alpine:latest as final

WORKDIR /go-app
EXPOSE 1337
COPY --from=build /go-app/main .
CMD ["./main"]