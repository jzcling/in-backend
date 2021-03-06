FROM golang:1.15-alpine as builder
LABEL stage=builder
WORKDIR /go/src/in-backend/

# restore dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY ./helpers/ ./helpers/
COPY ./scheduler/worker ./scheduler/worker
COPY ./services/assessment/pb ./services/assessment/pb
RUN cd scheduler/worker/cmd && \
    CGO_ENABLED=1 && \
    GOOS=linux && \
    GOARCH=amd64 && \
    go build -o /build/dist/main

FROM alpine as release
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /build/dist/ .
RUN chmod +x ./main
ENTRYPOINT ["./main"]