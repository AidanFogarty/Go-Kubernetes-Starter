FROM golang:latest AS build

LABEL maintainer="Aidan Fogarty"

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

## New stage
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"] 