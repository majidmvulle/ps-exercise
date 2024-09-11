FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN #echo $(ls -la /app/)

RUN CGO_ENABLED=0 GOOS=linux go build -o /service-b ./cmd/server

RUN chmod a+x /service-b

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /service-b /service-b

EXPOSE 8080

CMD ["./service-b"]
