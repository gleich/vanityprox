# syntax=docker/dockerfile:1
FROM golang:1.23 AS build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/vanityprox .

FROM alpine:3.20.2

RUN apk update && apk add --no-cache ca-certificates=20241121-r1 tzdata=2024b-r0

WORKDIR /src
COPY --from=build /bin/vanityprox /bin/vanityprox
RUN touch .env

CMD ["/bin/vanityprox"]