FROM golang:1.25-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 go build -o timemcp -trimpath -ldflags="-s -w" .

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder --chown=nonroot:nonroot /app/timemcp /home/nonroot/timemcp

ENV HOST=0.0.0.0

ENV PORT=8000

EXPOSE 8000

USER nonroot

ENTRYPOINT ["/home/nonroot/timemcp"]
