package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

var topic string = "test/topic"
var mqttHost string = "localhost"
var mqttPort int = 1883
var mqttUser string = ""
var mqttPassword string = ""
var ngsiLdHost string = "localhost"
var ngsiLdPort int = 1026
var ngsiLdUrl string
var craneId string = "urn:ngsi-ld:crane:lego-crane"

var mqttClient mqtt.Client

// Interface to the http-client
type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

/**
* Global http client
 */
var globalHttpClient httpClient = &http.Client{}

type WeightMessage struct {
	Weight JSONFloat `json:"weight"`
}

type JSONFloat struct {
	Value float32
	Valid bool
	Set   bool
}

func (i *JSONFloat) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp float32
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	var weightMessage WeightMessage
	logger.Infof(fmt.Sprint(msg.Payload()))
	err := json.Unmarshal(msg.Payload(), &weightMessage)
	if err != nil {
		logger.Warn("Was unable to unmarshal mqtt message. E: %v", err)
		return
	}
	if weightMessage.Weight.Set && weightMessage.Weight.Valid {
		logger.Infof("Weight %v", weightMessage)
	} else {
		logger.Info("Message is not weight")
		return
	}
	var active bool

	if weightMessage.Weight.Value > 0 {
		active = true
	}

	invoiceTrigger := rand.Float32()
	if invoiceTrigger > 0.95 {
		logger.Warn("++++++++++++++++++++++++++++++++++++++++++++++++++++++++Trigger invoice")
		// in order for the invoice to be generated from time to time, we set the weight to a value that triggers it.
		weightMessage.Weight.Value = 25
	}

	timeNow := time.Now().Format("2006-01-02T15:04:05.002Z")
	locationProperty := `"location": {"type": "GeoProperty", "value": { "type": "Point", "coordinates": [5,5]}}`
	softwareVersionProperty := getStringAsPropertyJson("softwareVersion", "0.0.1", timeNow)
	activeProperty := getBooleanAsPropertyJson("active", active, timeNow)
	maxHookHeightProperty := getNumberAsPropertyJson("maxHookHeight", 130.0, timeNow)
	maxLiftingWeightProperty := getNumberAsPropertyJson("maxLiftingWeight", 25000.0, timeNow)
	maxPayloadProperty := getNumberAsPropertyJson("maxPayload", 25000.0, timeNow)
	payLoadAtTip := getNumberAsPropertyJson("payLoadAtTip", 15000.0, timeNow)
	modelProperty := getStringAsPropertyJson("model", "Euro SSG 130", timeNow)
	// we cut everything after comma for easier handling
	currentWeightProperty := getNumberAsPropertyJson("currentWeight", int(weightMessage.Weight.Value*1000), timeNow)
	inUseProperty := getBooleanAsPropertyJson("inUse", active, timeNow)
	currentConsumption := getNumberAsPropertyJson("currentConsumption", int(rand.Float32()*100), timeNow)

	entityString := fmt.Sprintf(`{
	"id": "%v",
	"type": "crane",
	%v,
	"model": {
		"type": "Property",
		"value": "Lego Crane"
	},
	"healthState": {
		"type": "Property",
		"value": "HEALTHY"
	},
	"currentCost": {
		"type": "Property",
		"value": 40.2
	},
	"generalInformation": {
        "type": "Property",
		"value": {
			%v,
			%v,
			%v,
			%v,
			%v,
			%v,
			%v,
			%v
		}
	},
	%v,
	%v,
	%v
	}`, craneId,
		locationProperty,
		softwareVersionProperty,
		activeProperty,
		maxHookHeightProperty,
		maxLiftingWeightProperty, maxPayloadProperty, modelProperty, currentWeightProperty, payLoadAtTip, inUseProperty, currentWeightProperty, currentConsumption)

	entity := []byte(entityString)
	req, _ := http.NewRequest("POST", ngsiLdUrl+"/entities", bytes.NewBuffer(entity))
	req.Header.Set("NGSILD-Tenant", "impress")
	req.Header.Set("Content-Type", "application/json")

	response, err := globalHttpClient.Do(req)
	if err != nil {
		logger.Warnf("Error: %v", err)
		return
	}

	if response.StatusCode == 409 {
		entityFragment := []byte(fmt.Sprintf(`{
			%v,
			"generalInformation": {
				"type": "Property",
				"value": {
					%v,
					%v,
					%v,
					%v,
					%v,
					%v,
					%v
				}
			},
			%v,
			%v,
			%v
			}`,
			locationProperty,
			softwareVersionProperty,
			activeProperty,
			maxHookHeightProperty,
			maxLiftingWeightProperty, maxPayloadProperty, modelProperty, currentWeightProperty, inUseProperty, currentWeightProperty, currentConsumption))

		req, _ := http.NewRequest("POST", ngsiLdUrl+"/entities/"+craneId+"/attrs/", bytes.NewBuffer(entityFragment))
		req.Header.Set("NGSILD-Tenant", "impress")
		req.Header.Set("Content-Type", "application/json")
		response, err = globalHttpClient.Do(req)
		if err != nil {
			logger.Warnf("Was not able to update. %v", err)
		} else {
			logger.Infof("Update response: %v", response)
		}
	} else {
		logger.Infof("Creation response: %v", response)
	}

}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func getStringAsPropertyJson(name string, a string, t string) string {
	return fmt.Sprintf(`%q: {
		"type": "Property",
		"value": %q,
		"observedAt": %q
	}`, name, a, t)
}

func getNumberAsPropertyJson(name string, a int, t string) string {
	return fmt.Sprintf(`%q: {
		"type": "Property",
		"value": %d,
		"observedAt": %q
	}`, name, a, t)
}

func getBooleanAsPropertyJson(name string, a bool, t string) string {
	return fmt.Sprintf(`%q: {
		"type": "Property",
		"value": %t,
		"observedAt": %q
	}`, name, a, t)
}

func main() {

	devGeneratorEnabled := os.Getenv("DEV_GENERATOR_ENABLED")

	mqttHostEnv := os.Getenv("MQTT_HOST")
	mqttPortEnv := os.Getenv("MQTT_PORT")
	mqttTopicEnv := os.Getenv("MQTT_TOPIC")
	mqttUserEnv := os.Getenv("MQTT_USER")
	mqttPasswordEnv := os.Getenv("MQTT_PASSWORD")

	ngsiLdHostEnv := os.Getenv("NGSI_LD_HOST")
	ngsiLdPortEnv := os.Getenv("NGSI_LD_PORT")

	craneIdEnv := os.Getenv("CRANE_ID")

	if craneIdEnv != "" {
		logger.Infof("Setting the crane id to %v", craneIdEnv)
		craneId = craneIdEnv
	}

	if mqttUserEnv != "" {
		logger.Infof("Authorizing as user %v to the mqtt-broker.", mqttUserEnv)
		mqttUser = mqttUserEnv
	}

	if mqttPasswordEnv != "" {
		logger.Info("Setting password.")
		mqttPassword = mqttPasswordEnv
	}

	if mqttHostEnv != "" {
		logger.Infof("Setting mqtt host: %v", mqttHostEnv)
		mqttHost = mqttHostEnv
	}

	if mqttPortEnv != "" {
		mqttPortInt, err := strconv.Atoi(mqttPortEnv)
		if err == nil {
			logger.Infof("Setting mqtt port: %v", mqttPortInt)
			mqttPort = mqttPortInt
		}
	}

	if mqttTopicEnv != "" {
		logger.Infof("Setting mqtt topic: %v", mqttTopicEnv)
		topic = mqttTopicEnv
	}

	if ngsiLdHostEnv != "" {
		logger.Infof("Setting ngis-ld host: %v", ngsiLdHostEnv)
		ngsiLdHost = ngsiLdHostEnv
	}

	if ngsiLdPortEnv != "" {
		ngsiLdPortInt, err := strconv.Atoi(ngsiLdPortEnv)
		if err != nil {
			logger.Infof("Setting ngsi-ld port: %v", ngsiLdPortEnv)
			ngsiLdPort = ngsiLdPortInt
		}
	}

	ngsiLdUrl = fmt.Sprintf("http://%v:%v/ngsi-ld/v1", ngsiLdHost, ngsiLdPort)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", mqttHost, mqttPort))
	opts.SetClientID("go_mqtt_client")
	if mqttUser != "" && mqttPassword != "" {
		logger.Infof("Creating an authenticated connection for user %v.", mqttUser)
		opts.SetUsername(mqttUser)
		opts.SetPassword(mqttPassword)
	}
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// run the crane simulator internally
	if devGeneratorEnabled != "" {
		enabled, err := strconv.ParseBool(devGeneratorEnabled)
		if err == nil && enabled {
			simulateCrane(mqttClient)
		}
	}

	topic := topic
	token := mqttClient.Subscribe(topic, 1, nil)
	token.Wait()
	if token.Error() != nil {
		fmt.Print(token.Error().Error())
	}
	fmt.Printf("Subscribed to topic: %s", topic)
	for {
		fmt.Printf("Wait")
		time.Sleep(time.Second)
	}
}

func simulateCrane(client mqtt.Client) {
	for {
		weight := rand.Float32()
		text := fmt.Sprintf(`{"weight": %v}`, weight)
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}
