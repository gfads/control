/*********************************************************************************
Author: Nelson S Rosa
Description: This program defines constants used in the whole package.
Date: 04/02/2023
*********************************************************************************/

package shared

import (
	"math/rand"
	"time"
)

// AsTAR parameters

const SV = 2.7 // Shutoff voltage (page 17) = 2.7 V
const OV = 3.7 // Optimum voltage (page 17) = 3.7 V

const HYSTERESIS = 0.001 // 10 mV
const MinimumTaskExecutionRate = 1
const MaximumTaskExecutionRate = 1200

const MaximumVoltage = 100.0

const StepVoltage = 0.001 // in Volts

const AdaptationCycles = 131

const InitialVoltage = 3.3 // 3.3 V - page 19
const InitialRate = 1

const SamplingCycleSize = 30

const AverageConsumption = 0.01

// Controller type names

const BasicOnoff = "OnOff"
const DeadZoneOnoff = "OnOffwithDeadZone"
const HysteresisOnoff = "OnOffwithHysteresis"

const BasicPid = "BasicPID"
const SmoothingPid = "SmoothingDerivativePID"
const IncrementalFormPid = "IncrementalFormPID"
const ErrorSquarePid = "ErrorSquarePID"
const DeadZonePid = "DeadZonePID"

const GainScheduling = "GainScheduling"

// Capacitor behaviour patterns/AsTAR

const IncreasingHarvesting = 0
const HalfIncreasingHarvesting = 1
const QuarterIncreasingHarvesting = 2
const DecreasingHarvesting = 3
const HalfDecreasingHarvesting = 4
const QuarterDecreasingHarvesting = 5
const ConstantHarvesting = 6
const RandomHarvesting = 7
const NoHarvesting = 1000

// const AverageConsumption = 0.01

const AnyUpTopicFilter = BaseTopicName + "/+/up" // any topic

// MQTT Client parameters

const Node1Id = "eui-70b3d57ed005867f"
const Node2Id = "eui-70b3d57ed0058955"
const BrokerAddress = "eu1.cloud.thethings.network"
const BrokerPort = 1883
const UserNameMqtt = "floodsensor-appid@ttn"
const UserPwdMqtt = "NNSXS.GTSVUEBSLQEQEGFB6M7HACHU7C4IAWSPZCWLZUI.UFN7TKM7UJVP6X63Z6HLU2VCC2T2FATWZUW5ZJMTO4GTMD3A4PFA"
const BaseTopicName = "v3/floodsensor-appid@ttn/devices"

// Message format stored in "down" topics

type DownDataNode1 struct {
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
type DownDataNode2 struct {
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

// Up topic JSON Message

type UpData struct {
	TaskRate float64 `json:"task_rate"`
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
