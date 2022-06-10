FROM golang:1.18.3-alpine3.16

RUN MKDIR /apicontacts

ADD . /apicontacts

WORKDIR /apicontacts

RUN go mod download
RUN go install
RUN go build -o main .

ENTRYPOINT [ "./main" ]