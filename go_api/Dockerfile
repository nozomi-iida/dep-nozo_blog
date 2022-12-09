FROM golang:1.19

ENV APP_ROOT /app

WORKDIR ${APP_ROOT}

RUN go install github.com/cosmtrek/air@latest &&\
    go install github.com/rubenv/sql-migrate/...@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
