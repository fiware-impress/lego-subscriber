# Lego-Subscriber

Helper application for integrating the lego-crane into the impress-demonstrator. It subscribes to a given mqtt topic and listens to messages of form:
```json
    {"weight": float }
```
The weight is translated into an ngsi-ld entity(type crane, as expected by the impress demonstrator) and persisted to the ngsi-ld broker. 

## Config

| Env-Var | Description | Default |
| ----------------------------------- | ----------------------------------------------- | ------------------------ |
| `DEV_GENERATOR_ENABLED` | For development purposes, the subscriber can populate the topic with generated data itself, when enabled. | `false` |
| `MQTT_HOST` | Host of the mqtt broker for retrieving the crane data.   | `localhost`  |
| `MQTT_PORT` | Port of the mqtt broker for retrieving the crane data.   | 1883  |
| `MQTT_TOPIC` | Topic of the mqtt broker for retrieving the crane data.    | `lego/demonstrator`  |
| `MQTT_USER` | User for authenticating at the broker. If none is provided, we connect anonymous. |  |
| `MQTT_PASSWORD` | Password for authenticating at the broker. If none is provided, we connect anonymous. |  |
| `NGSI_LD_HOST` | Host of the ngsi-ld broker for persisting the crane data.    | `localhost`  |
| `NGSI_LD_PORT` | Port of the ngsi-ld broker for persisting the crane data.    | 1026  |
| `CRANE_ID` | Id to be used for the crane entity. | `urn:ngsi-ld:crane:lego-crane`  |

## Usage

```shell
    docker run quay.io/wi_stefan/lego-subscriber
``` 