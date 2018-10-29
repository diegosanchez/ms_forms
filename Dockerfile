FROM golang:1.8

WORKDIR /go/src/app

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN echo "gopath -> $GOPATH"
RUN echo "pwd -> $PWD"
RUN ls -la 

RUN dep ensure
RUN go build app/cmd/forms-server
RUN go install app/cmd/forms-server

EXPOSE 8080/tcp

CMD /go/bin/forms-server
