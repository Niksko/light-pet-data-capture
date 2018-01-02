FROM golang:1.8

WORKDIR /go/src/app
COPY --from=builder /go/src/app .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run", "app"]