FROM golang:1.6.2

WORKDIR /go/src/isiro

ENTRYPOINT ["make"]

CMD ["all"]
