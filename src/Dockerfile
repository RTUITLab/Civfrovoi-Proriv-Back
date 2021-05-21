FROM golang:1.15 AS builder

WORKDIR /go/src/RTUITLab/

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/RTUITLab/main .
CMD ["./main"]