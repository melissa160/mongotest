FROM golang:1.11.0-stretch
MAINTAINER Juliana Davila "juliana.davila@liftit.co"

COPY ./ci_config /root/.ssh/
RUN chmod 600 /root/.ssh/id_rsa
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

RUN mkdir -p $GOPATH/liftitapp/mongotest
WORKDIR $GOPATH/liftitapp/mongotest
COPY . .

RUN export GO111MODULE=on
RUN go build -o main

EXPOSE 3000
CMD ./main