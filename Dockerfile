FROM golang:1.23-bullseye as base

COPY . .

RUN go build -ldflags="-s -w" -o /main main.go

FROM golang:1.23-bullseye

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /main .

CMD ["./main"]