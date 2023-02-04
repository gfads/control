package shared

import (
	"math/rand"
	"time"
)

const SV = 2.7           // Shutoff voltage (page 17) = 2.7 V
const OV = 3.7           // Optimum voltage (page 17) = 3.7 V
const HYSTERISIS = 0.001 // 10 mV

// Controller types
const BASIC_ONOFF = "OnOff"
const DEAD_ZONE_ONOFF = "OnOffwithDeadZone"
const HYSTERESIS_ONOFF = "OnOffwithHysteresis"

const BASIC_PID = "BasicPID"
const SMOOTHING_PID = "SmoothingDerivativePID"
const INCREMENTAL_FORM_PID = "IncrementalFormPID"
const ERROR_SQUARE_PID = "ErrorSquarePID"
const DEAD_ZONE_PID = "DeadZonePID"

const GAIN_SCHEDULING = "GainScheduling"

// Capacitor behaviour patterns
const INCREASING = 0
const HALF_INCREASING = 1
const QUARTER_INCREASING = 2
const DECREASING = 3
const HALF_DECREASING = 4
const QUARTER_DECREASING = 5
const CONSTANT = 6
const RANDOM = 7

const MAXIMUM_VOLTAGE = 100.0
const STEP_VOLTAGE = 0.001 // in Volts

const ADAPTATION_CYCLES = 131

const MINIMUM_TASK_EXECUTION_RATE = 1    // page 17
const MAXIMUM_TASK_EXECUTION_RATE = 1200 // page 17

const INITIAL_VOLTAGE = 3.3 // 3.3 V - page 19
const INITIAL_RATE = 1

const AVERAGE_CONSUMPTION = 0.01

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

const ANY_UP_TOPIC = BASE_TOPIC_NAME + "/+/up" // any topic

// MQTT CLIENT
const NODE1_ID = "eui-70b3d57ed005867f"
const NODE2_ID = "eui-70b3d57ed0058955"
const BROKER_ADDRESS = "eu1.cloud.thethings.network"
const BROKER_PORT = 1883
const USER_NAME_MQTT = "floodsensor-appid@ttn"
const USER_PWD_MQTT = "NNSXS.GTSVUEBSLQEQEGFB6M7HACHU7C4IAWSPZCWLZUI.UFN7TKM7UJVP6X63Z6HLU2VCC2T2FATWZUW5ZJMTO4GTMD3A4PFA"
const BASE_TOPIC_NAME = "v3/floodsensor-appid@ttn/devices"

type SubDataNode1 struct {
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
type SubDataNode2 struct {
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

type UpData struct {
	TaskRate float64 `json:"task_rate"`
}
