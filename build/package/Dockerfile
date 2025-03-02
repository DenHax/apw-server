FROM golang:1.23-alpine AS builder

WORKDIR /usr/local/src

RUN apk add --no-cache bash git make gcc gettext musl-dev

COPY ./go.mod ./go.sum .

RUN go mod download

COPY ./ ./

RUN go build -o ./bin/apw ./cmd/apw/main.go


FROM alpine:3.20 AS runner

RUN apk add --no-cache ca-certificates postgresql-client

COPY ./scripts/wait-storage.sh /tmp/wait.sh

RUN chmod 744 /wait.sh

COPY ./configs/config.yaml /config.yaml

ENV CONFIG_PATH /config.yaml

COPY --from=builder /usr/local/src/bin/apw /apw

CMD ["/apw"]
