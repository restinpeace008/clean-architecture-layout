FROM golang:1.20.1-alpine as builder

WORKDIR /source

COPY . .

RUN go mod download

RUN mkdir -p bin && go build -o /bin/app ./cmd/app/

FROM alpine

WORKDIR /app

COPY --from=builder /bin/app ./app
COPY --from=builder /source/configs/config.json ./configs/config.json

CMD [ "./app" ]