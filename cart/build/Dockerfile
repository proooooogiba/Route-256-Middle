FROM golang:1.22-alpine as builder

WORKDIR /build

COPY bin/cart /server

FROM scratch
COPY --from=builder server /bin/server

ENTRYPOINT ["/bin/server"]
