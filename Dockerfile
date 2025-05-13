FROM golang:1.23.6 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o jelly-metrics

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=builder /app/jelly-metrics .

EXPOSE 8097
CMD ["/jelly-metrics"]
