package myclient

import (
	"controllers/def/ops"
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"reflect"
	"shared"
	"strings"
)

type MyMQTTClient struct {
	MyClient   mqtt.Client
	Controller ops.IController
}

var subChan = make(chan interface{})
var pubChan = make(chan interface{})

func NewMyMQTTClient(brokerAddress string, port int, controller ops.IController) MyMQTTClient {

	// client configuration
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", brokerAddress, port))
	//opts.SetClientID("")
	opts.SetUsername(shared.UserNameMqtt)
	opts.SetPassword(shared.UserPwdMqtt)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	// create an instance of mqtt client
	myClient := MyMQTTClient{MyClient: mqtt.NewClient(opts), Controller: controller}

	return myClient
}

func (c *MyMQTTClient) Run() {

	// voltage level
	vnew := 0.0
	nodeId := ""
	taskRate := 0.0

	// connect to broker
	if token := c.MyClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// subscribe client to topics
	sub(c.MyClient)

	// loop for receiving messages
	for {

		// receive voltage level from mqtt
		upMsg := <-subChan

		if strings.Index(reflect.TypeOf(upMsg).String(), "Node1") > 0 {
			msgNode := upMsg.(shared.DownDataNode1)
			vnew = msgNode.UplinkMessage.DecodedPayload.Node1BatteryLevel
			nodeId = msgNode.EndDeviceIds.DeviceId
			taskRate = c.Controller.Update(shared.OV, vnew) // goal - optimum voltage, voltage
		} else if strings.Index(reflect.TypeOf(upMsg).String(), "Node2") > 0 {
			msgNode := upMsg.(shared.DownDataNode2)
			vnew = msgNode.UplinkMessage.DecodedPayload.Node2BatteryLevel
			nodeId = msgNode.EndDeviceIds.DeviceId
			taskRate = c.Controller.Update(shared.OV, vnew) // goal - optimum voltage, voltage
		} else {
			fmt.Println("Error: Node unknown...")
			os.Exit(0)
		}

		// publish task rate to mqtt
		msgUp := shared.UpData{TaskRate: taskRate}
		c.MyClient.Publish(shared.BaseTopicName+"/"+nodeId+"/down", 0, false, msgUp)

		fmt.Println("Node: ", nodeId, "Voltage Level: ", vnew, "Task Rate: ", taskRate)
	}

	// disconnect from broker
	c.MyClient.Disconnect(250)
}

func sub(client mqtt.Client) {
	token := client.Subscribe(shared.AnyUpTopicFilter, 0, nil)

	token.Wait()
}

func messagePubHandler(client mqtt.Client, msg mqtt.Message) {

	var r interface{}

	var infoNode01 shared.DownDataNode1
	var infoNode02 shared.DownDataNode2

	// unmarshall message of node 1 - TODO
	err := json.Unmarshal(msg.Payload(), &infoNode01)
	if err != nil {
		fmt.Println("Error:: Something wrong with the message")
		os.Exit(0)
	}

	nodeId := infoNode01.EndDeviceIds.DeviceId
	switch nodeId {
	case shared.Node1Id:
		r = infoNode01
	case shared.Node2Id:
		// unmarshall message of node 2
		err := json.Unmarshal(msg.Payload(), &infoNode02)
		if err != nil {
			fmt.Println("Error:: Something wrong with the message")
			os.Exit(0)
		}
		r = infoNode02
	default:
		fmt.Println("Error:: Node unknown")
		os.Exit(0)
	}

	// send voltage level to client
	subChan <- r
}

func connectHandler(client mqtt.Client) {}

func connectLostHandler(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}
