FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY . .
RUN go get ./...
RUN go build -o ./bin/GoRailwire ./cmd/Railwire/main.go

FROM alpine:3.9
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin/GoRailwire /go/bin/GoRailwire
EXPOSE 9092
ENTRYPOINT /go/bin/GoRailwire
