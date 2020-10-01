FROM golang:1.14 as builder

ENV APP_USER app
ENV APP_HOME /go/src/gossip

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

WORKDIR $APP_HOME
USER $APP_USER
COPY ./ .

RUN go mod download
RUN go mod verify
RUN go build -o gossip

FROM debian:buster

ENV APP_USER app
ENV APP_HOME /go/src/gossip

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY --chown=0:0 --from=builder $APP_HOME/public $APP_HOME/public
COPY --chown=0:0 --from=builder $APP_HOME/gossip $APP_HOME

EXPOSE 3000
CMD ["./gossip"]
