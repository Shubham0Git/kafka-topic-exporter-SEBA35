// Copyright 2018 Open Networking Foundation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

// configuration
type BrokerInfo struct {
	Name			string `yaml: name`
	Host			string `yaml: host`
	Description		string `yaml: description`
	Topics		  []string `yaml: topics`
}

type LoggerInfo struct {
	LogLevel		string `yaml: loglevel`
	Host			string `yaml: host`
}

type TargetInfo struct {
	Type			string `yaml: type`
	Name			string `yaml: name`
	Port			int    `yaml: port`
	Description		string `yaml: description`
}

type Config struct {
	Broker		BrokerInfo `yaml: broker`
	Logger		LoggerInfo `yaml: logger`
	Target		TargetInfo `yaml: "target"`
}

// KPI Events format
type Metrics struct {
	TxBytes            float64 `json:"tx_bytes"`
	TxPackets          float64 `json:"tx_packets"`
	TxErrorPackets     float64 `json:"tx_error_packets"`
	TxBcastPackets     float64 `json:"tx_bcast_packets"`
	TxUnicastPackets   float64 `json:"tx_ucast_packets"`
	TxMulticastPackets float64 `json:"tx_mcast_packets"`
	RxBytes            float64 `json:"rx_bytes"`
	RxPackets          float64 `json:"rx_packets"`
	RxErrorPackets     float64 `json:"rx_error_packets"`
	RxBcastPackets     float64 `json:"rx_bcast_packets"`
	RxMulticastPackets float64 `json:"rx_mcast_packets"`

	// ONU Ethernet_Bridge_Port_history
	Packets            float64 `json:"packets"`
	Octets             float64 `json:"octets"`
}

type Context struct {
	InterfaceID string `json:"intf_id"`
	PonID       string `json:"pon_id"`
	PortNumber  string `json:"port_no"`

	// ONU Performance Metrics
	ParentClassId string `json:"parent_class_id"`
	ParentEntityId string `json:"parent_entity_id"`
	Upstream    string `json:"upstream"`
}

type Metadata struct {
	LogicalDeviceID string   `json:"logical_device_id"`
	Title           string   `json:"title"`
	SerialNumber    string   `json:"serial_no"`
	Timestamp       float64  `json:"ts"`
	DeviceID        string   `json:"device_id"`
	Context         *Context `json:"context"`
}

type SliceData struct {
	Metrics  *Metrics  `json:"metrics"`
	Metadata *Metadata `json:"metadata"`
}

type VolthaKPI struct {
	Type       string       `json:"type"`
	Timestamp  float64      `json:"ts"`
	SliceDatas []*SliceData `json:"slice_data"`
}

type OnosPort struct {
	PortID        string  `json:"portId"`
	RxPackets     float64 `json:"pktRx"`
	TxPackets     float64 `json:"pktTx"`
	RxBytes       float64 `json:"bytesRx"`
	TxBytes       float64 `json:"bytesTx"`
	RxPacketsDrop float64 `json:"pktRxDrp"`
	TxPacketsDrop float64 `json:"pktTxDrp"`
}

type OnosKPI struct {
	DeviceID string      `json:"deviceId"`
	Ports    []*OnosPort `json:"ports"`
}

type ImporterKPI struct {
	DeviceID string 	`json: "deviceId"`
	// TODO: add metrics data
}

type OnosAaaKPI struct {
	RxAcceptResponses    float64 `json:"acceptResponsesRx"`
	RxRejectResponses    float64 `json:"rejectResponsesRx"`
	RxChallengeResponses float64 `json:"challengeResponsesRx"`
	TxAccessRequests     float64 `json:"accessRequestsTx"`
	RxInvalidValidators  float64 `json:"invalidValidatorsRx"`
	RxUnknownType        float64 `json:"unknownTypeRx"`
	PendingRequests      float64 `json:"pendingRequests"`
	RxDroppedResponses   float64 `json:"droppedResponsesRx"`
	RxMalformedResponses float64 `json:"malformedResponsesRx"`
	RxUnknownserver      float64 `json:"unknownServerRx"`
	RequestRttMillis     float64 `json:"requestRttMillis"`
	RequestReTx          float64 `json:"requestReTx"`
}
