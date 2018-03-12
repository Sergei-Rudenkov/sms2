FROM golang:1.9.2

COPY . "/go/src/sms2/"
WORKDIR "/go/src/sms2"

RUN go get
RUN go install
ENTRYPOINT /go/bin/sms2 agile
// ENTRYPOINT /go/bin/sms2 fixed 1000 60

EXPOSE 8080
EXPOSE 5555


// sudo docker build -t sms2 .
// sudo docker run -p 8080:8080 -p 5555:5555 --name sms2 --rm sms2