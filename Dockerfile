FROM golang:1.21.0
WORKDIR /app
ADD . /app
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["prexel-post-api"]
