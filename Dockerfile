FROM golang:latest as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

ENV CGO_ENABLED=0

COPY . .
RUN go build -ldflags="-w -s" -o btctouah .

FROM golang:latest AS runner
WORKDIR /build
COPY --from=builder /build/btctouah .
EXPOSE 8000
CMD ["./btctouah"]

