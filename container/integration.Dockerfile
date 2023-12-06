FROM golang:1.21

WORKDIR /go/src/github.com/TrevorEdris/about-me

COPY . .

ENTRYPOINT [ "/go/src/github.com/TrevorEdris/about-me/tests/entrypoint_integration.sh" ]
