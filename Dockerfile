FROM golang:1.17.6-alpine AS builder

RUN apk update && apk upgrade && apk add --no-progress --no-cache ca-certificates build-base

WORKDIR /go/src/github.com/jedi4z/vwap-calculator

COPY . .

RUN make test

RUN make build


# Use the alpine image for running the service
FROM alpine:latest AS final_image

RUN apk update && apk upgrade && apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/jedi4z/vwap-calculator/vwap_calculator ./vwap_calculator

CMD ["./vwap_calculator"]