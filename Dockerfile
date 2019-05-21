FROM golang
ADD ./src /go/src/chris/go-test-app
RUN go install chris/go-test-app
EXPOSE 80
ENTRYPOINT /go/bin/go-test-app -port 80