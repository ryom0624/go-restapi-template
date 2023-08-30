FROM golang:1.19-bullseye as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath ldflags "-s -w" -o app


FROM debian:bullseye-slim as deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]


FROM golang:1.19 as dev

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

COPY build/entrypoint-app.sh /usr/local/bin
RUN chmod +x /usr/local/bin/entrypoint-app.sh

ENTRYPOINT ["entrypoint-app.sh"]
