FROM golang:1.18

ENV DEV_GENERATOR_ENABLED=false

ENV MQTT_HOST="localhost"
ENV MQTT_PORT=1883
ENV MQTT_TOPIC="lego/demonstrator"

ENV NGSI_LD_HOST="localhost"
ENV NGSI_LD_PORT=1026

WORKDIR /go/src/app
COPY ./ ./

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["lego-subscriber"]