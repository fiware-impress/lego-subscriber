package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

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

var mqttClient mqtt.Client

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {

	mqttHostEnv := os.Getenv("MQTT_HOST")
	mqttPortEnv := os.Getenv("MQTT_PORT")
	mqttTopicEnv := os.Getenv("MQTT_TOPIC")
	mqttUserEnv := os.Getenv("MQTT_USER")
	mqttPasswordEnv := os.Getenv("MQTT_PASSWORD")

	ngsiLdHostEnv := os.Getenv("NGSI_LD_HOST")
	ngsiLdPortEnv := os.Getenv("NGSI_LD_PORT")

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

	sub(mqttClient)
	simulateCrane(mqttClient)
}

func sub(client mqtt.Client) {
	topic := topic
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

func simulateCrane(client mqtt.Client) {
	for {
		weight := rand.Float32() * 1000
		text := fmt.Sprintf(`{"weight": %v}`, weight)
		token := client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}
}
