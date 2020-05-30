FROM golang as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build

FROM scratch
COPY --from=builder /app/edgetpu-exporter /app/

ENTRYPOINT [ "/app/edgetpu-exporter" ]
