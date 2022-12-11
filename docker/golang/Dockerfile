FROM golang:1.19.4-alpine

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

RUN apk update && apk --no-cache add git vim curl jq tree

CMD ["go", "run", "."]