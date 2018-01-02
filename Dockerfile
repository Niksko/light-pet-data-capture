FROM golang:1.8

WORKDIR /go/src/light-pet-data-capture
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run", "light-pet-data-capture"]