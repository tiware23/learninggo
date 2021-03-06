FROM golang:1.14 AS Build
WORKDIR /go/src/
COPY lissajous.go /go/src/
RUN  go build lissajous.go
FROM golang:1.14 AS Runtime
WORKDIR /go/src/
COPY --from=Build /go/src/lissajous /go/src/
EXPOSE 8000
CMD ["/go/src/lissajous"]
