package myclient

import (
	"controllers"
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"os"
	"shared"
	"time"
)

type MyMQTTClient struct {
	MyClient   mqtt.Client
	Controller controllers.IController
}

type DataNode1 struct {
	EndDeviceIds struct {
		DeviceId       string `json:"device_id"`
		ApplicationIds struct {
			ApplicationId string `json:"application_id"`
		} `json:"application_ids"`
		DevEui  string `json:"dev_eui"`
		JoinEui string `json:"join_eui"`
		DevAddr string `json:"dev_addr"`
	} `json:"end_device_ids"`
	CorrelationIds []string  `json:"correlation_ids"`
	ReceivedAt     time.Time `json:"received_at"`
	UplinkMessage  struct {
		SessionKeyId   string `json:"session_key_id"`
		FPort          int    `json:"f_port"`
		FCnt           int    `json:"f_cnt"`
		FrmPayload     string `json:"frm_payload"`
		DecodedPayload struct {
			Node1BatteryLevel float64 `json:"Node1__Battery_Level"`
			Node1WaterLevel   int     `json:"Node1__WaterLevel"`
		} `json:"decoded_payload"`
		RxMetadata []struct {
			GatewayIds struct {
				GatewayId string `json:"gateway_id"`
				Eui       string `json:"eui"`
			} `json:"gateway_ids"`
			Time         time.Time `json:"time"`
			Timestamp    int       `json:"timestamp"`
			Rssi         int       `json:"rssi"`
			ChannelRssi  int       `json:"channel_rssi"`
			Snr          float64   `json:"snr"`
			UplinkToken  string    `json:"uplink_token"`
			ChannelIndex int       `json:"channel_index"`
			ReceivedAt   time.Time `json:"received_at"`
		} `json:"rx_metadata"`
		Settings struct {
			DataRate struct {
				Lora struct {
					Bandwidth       int    `json:"bandwidth"`
					SpreadingFactor int    `json:"spreading_factor"`
					CodingRate      string `json:"coding_rate"`
				} `json:"lora"`
			} `json:"data_rate"`
			Frequency string    `json:"frequency"`
			Timestamp int       `json:"timestamp"`
			Time      time.Time `json:"time"`
		} `json:"settings"`
		ReceivedAt      time.Time `json:"received_at"`
		ConsumedAirtime string    `json:"consumed_airtime"`
		Locations       struct {
			User struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Source    string  `json:"source"`
			} `json:"user"`
		} `json:"locations"`
		NetworkIds struct {
			NetId          string `json:"net_id"`
			TenantId       string `json:"tenant_id"`
			ClusterId      string `json:"cluster_id"`
			ClusterAddress string `json:"cluster_address"`
		} `json:"network_ids"`
	} `json:"uplink_message"`
}
type DataNode2 struct {
	EndDeviceIds struct {
		DeviceId       string `json:"device_id"`
		ApplicationIds struct {
			ApplicationId string `json:"application_id"`
		} `json:"application_ids"`
		DevEui  string `json:"dev_eui"`
		JoinEui string `json:"join_eui"`
		DevAddr string `json:"dev_addr"`
	} `json:"end_device_ids"`
	CorrelationIds []string  `json:"correlation_ids"`
	ReceivedAt     time.Time `json:"received_at"`
	UplinkMessage  struct {
		SessionKeyId   string `json:"session_key_id"`
		FPort          int    `json:"f_port"`
		FCnt           int    `json:"f_cnt"`
		FrmPayload     string `json:"frm_payload"`
		DecodedPayload struct {
			Node2BatteryLevel float64 `json:"Node2__Battery_Level"`
			Node2WaterLevel   int     `json:"Node2__WaterLevel"`
		} `json:"decoded_payload"`
		RxMetadata []struct {
			GatewayIds struct {
				GatewayId string `json:"gateway_id"`
				Eui       string `json:"eui"`
			} `json:"gateway_ids"`
			Time         time.Time `json:"time"`
			Timestamp    int       `json:"timestamp"`
			Rssi         int       `json:"rssi"`
			ChannelRssi  int       `json:"channel_rssi"`
			Snr          float64   `json:"snr"`
			UplinkToken  string    `json:"uplink_token"`
			ChannelIndex int       `json:"channel_index"`
			ReceivedAt   time.Time `json:"received_at"`
		} `json:"rx_metadata"`
		Settings struct {
			DataRate struct {
				Lora struct {
					Bandwidth       int    `json:"bandwidth"`
					SpreadingFactor int    `json:"spreading_factor"`
					CodingRate      string `json:"coding_rate"`
				} `json:"lora"`
			} `json:"data_rate"`
			Frequency string    `json:"frequency"`
			Timestamp int       `json:"timestamp"`
			Time      time.Time `json:"time"`
		} `json:"settings"`
		ReceivedAt      time.Time `json:"received_at"`
		ConsumedAirtime string    `json:"consumed_airtime"`
		Locations       struct {
			User struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Source    string  `json:"source"`
			} `json:"user"`
		} `json:"locations"`
		NetworkIds struct {
			NetId          string `json:"net_id"`
			TenantId       string `json:"tenant_id"`
			ClusterId      string `json:"cluster_id"`
			ClusterAddress string `json:"cluster_address"`
		} `json:"network_ids"`
	} `json:"uplink_message"`
}

/*
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	vnew := getVoltageLevel(msg)
	fmt.Println("Voltage Level: ", vnew)
}
*/

var ch = make(chan float64)

func messagePubHandler(client mqtt.Client, msg mqtt.Message) {
	vnew := getVoltageLevel(msg)

	// send voltage level to client
	ch <- vnew
}

/*
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Client connected...")
}
*/

func connectHandler(client mqtt.Client) {}

/*
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}
*/

func connectLostHandler(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func NewMyMQTTClient(brokerAddress string, port int, controller controllers.IController) MyMQTTClient {

	// client configuration
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", brokerAddress, port))
	//opts.SetClientID("")
	opts.SetUsername(shared.USER_NAME_MQTT)
	opts.SetPassword(shared.USER_PWD_MQTT)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	// create an instance of mqtt client
	myClient := MyMQTTClient{MyClient: mqtt.NewClient(opts), Controller: controller}

	return myClient
}

func (c *MyMQTTClient) Run() {

	// connect to broker
	if token := c.MyClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// subscribe client to topics
	sub(c.MyClient)

	for {

		// receive voltage level from mqtt
		vnew := <-ch
		taskRate := c.Controller.Update(shared.OV, vnew) // goal - optimum voltage, voltage

		fmt.Println("Voltage Level: ", vnew, "Task Rate: ", taskRate)
	}

	// disconnect from broker
	c.MyClient.Disconnect(250)
}

func sub(client mqtt.Client) {
	topic01 := "v3/floodsensor-appid@ttn/devices/eui-70b3d57ed005867f/up" // device 1
	topic02 := "v3/floodsensor-appid@ttn/devices/eui-70b3d57ed0058955/up" // device 2
	topic03 := "v3/floodsensor-appid@ttn/devices"                         // all devices
	token01 := client.Subscribe(topic01, 0, nil)
	token02 := client.Subscribe(topic02, 0, nil)
	token03 := client.Subscribe(topic03, 0, nil)

	token01.Wait()
	token02.Wait()
	token03.Wait()
}

func getVoltageLevel(msg mqtt.Message) float64 {
	var infoNode01 DataNode1
	var infoNode02 DataNode2
	var vnew float64

	// unmarshall message of node 1 - TODO
	err := json.Unmarshal(msg.Payload(), &infoNode01)
	if err != nil {
		fmt.Println("Error:: Something wrong with the message")
		os.Exit(0)
	}

	nodeId := infoNode01.EndDeviceIds.DeviceId
	switch nodeId {
	case shared.NODE1_ID:
		vnew = infoNode01.UplinkMessage.DecodedPayload.Node1BatteryLevel
	case shared.NODE2_ID:
		// unmarshall message of node 2
		err := json.Unmarshal(msg.Payload(), &infoNode02)
		if err != nil {
			fmt.Println("Error:: Something wrong with the message")
			os.Exit(0)
		}
		vnew = infoNode02.UplinkMessage.DecodedPayload.Node2BatteryLevel
	default:
		fmt.Println("Error:: Node unknown")
		os.Exit(0)
	}
	return vnew
}
