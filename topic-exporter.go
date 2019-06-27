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

import (
	"encoding/json"
	"log"
	"github.com/prometheus/client_golang/prometheus"
	"gerrit.opencord.org/kafka-topic-exporter/common/logger"
)

var (
	// voltha kpis
	volthaTxBytesTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "voltha_tx_bytes_total",
			Help: "Number of total bytes transmitted",
		},
		[]string{"logical_device_id", "serial_number", "device_id", "interface_id", "pon_id", "port_number", "title"},
	)
	volthaRxBytesTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "voltha_rx_bytes_total",
			Help: "Number of total bytes received",
		},
		[]string{"logical_device_id", "serial_number", "device_id", "interface_id", "pon_id", "port_number", "title"},
	)
	volthaTxPacketsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "voltha_tx_packets_total",
			Help: "Number of total packets transmitted",
		},
		[]string{"logical_device_id", "serial_number", "device_id", "interface_id", "pon_id", "port_number", "title"},
	)
	volthaRxPacketsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "voltha_rx_packets_total",
			Help: "Number of total packets received",
		},
		[]string{"logical_device_id", "serial_number", "device_id", "interface_id", "pon_id", "port_number", "title"},
	)

	volthaTxErrorPacketsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "voltha_tx_error_packets_total",
			Help: "Number of total transmitted packets error",
		},
		[]string{"logical_device_id", "serial_number", "device_id", "interface_id", "pon_id", "port_number", "title"},
	)

	volthaRxErrorPacketsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "voltha_rx_error_packets_total",
			Help: "Number of total received packets error",
		},
		[]string{"logical_device_id", "serial_number", "device_id", "interface_id", "pon_id", "port_number", "title"},
	)

	// onos kpis
	onosTxBytesTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "onos_tx_bytes_total",
			Help: "Number of total bytes transmitted",
		},
		[]string{"device_id", "port_id"},
	)
	onosRxBytesTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "onos_rx_bytes_total",
			Help: "Number of total bytes received",
		},
		[]string{"device_id", "port_id"},
	)
	onosTxPacketsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "onos_tx_packets_total",
			Help: "Number of total packets transmitted",
		},
		[]string{"device_id", "port_id"},
	)
	onosRxPacketsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "onos_rx_packets_total",
			Help: "Number of total packets received",
		},
		[]string{"device_id", "port_id"},
	)

	onosTxDropPacketsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "onos_tx_drop_packets_total",
			Help: "Number of total transmitted packets dropped",
		},
		[]string{"device_id", "port_id"},
	)

	onosRxDropPacketsTotal = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "onos_rx_drop_packets_total",
			Help: "Number of total received packets dropped",
		},
		[]string{"device_id", "port_id"},
	)

	// onos.aaa kpis
	onosaaaRxAcceptResponses = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_rx_accept_responses",
			Help: "Number of access accept packets received from the server",
		})
	onosaaaRxRejectResponses = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_rx_reject_responses",
			Help: "Number of access reject packets received from the server",
		})
	onosaaaRxChallengeResponses = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_rx_challenge_response",
			Help: "Number of access challenge packets received from the server",
		})
	onosaaaTxAccessRequests = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_tx_access_requests",
			Help: "Number of access request packets sent to the server",
		})
	onosaaaRxInvalidValidators = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_rx_invalid_validators",
			Help: "Number of access response packets received from the server with an invalid validator",
		})
	onosaaaRxUnknownType = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_rx_unknown_type",
			Help: "Number of packets of an unknown RADIUS type received from the accounting server",
		})
	onosaaaPendingRequests = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_pending_responses",
			Help: "Number of access request packets pending a response from the server",
		})
	onosaaaRxDroppedResponses = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_rx_dropped_responses",
			Help: "Number of dropped packets received from the accounting server",
		})
	onosaaaRxMalformedResponses = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_rx_malformed_responses",
			Help: "Number of malformed access response packets received from the server",
		})
	onosaaaRxUnknownserver = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_rx_from_unknown_server",
			Help: "Number of packets received from an unknown server",
		})
	onosaaaRequestRttMillis = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_request_rttmillis",
			Help: "Roundtrip packet time to the accounting server in Miliseconds",
		})
	onosaaaRequestReTx = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_request_re_tx",
			Help: "Number of access request packets retransmitted to the server",
		})
	onosaaaEapolLogoffRx = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_eapol_Logoff_Rx",
			Help: "Number of EAPOL logoff messages received resulting in disconnected state",
		})
	onosaaaEapolResIdentityMsgTrans = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_eapol_Res_IdentityMsg_Trans",
			Help: "Number of authenticating transitions due to EAP response or identity message",
		})
	onosaaaAuthSuccessTrans = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_auth_Success_Trans",
			Help: "Number of authenticated transitions due to successful authentication",
		})
	onosaaaAuthFailureTrans = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_auth_Failure_Trans",
			Help: "Number of transitions to held due to authentication failure",
		})
	onosaaaStartReqTrans = 	prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_start_Req_Trans",
			Help: "Number of transitions to connecting due to start request",
		})
	onosaaaTxAccessReqPkt =	prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_access_Req_Pkt_Tx",
			Help: "Number of access request packets sent",
		})
	onosaaaRxaccessChallPkt = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_access_Chall_Pkt_Rx",
			Help: "Number of access challenge packets received",
		})
	onosaaaEapPktTxauthEap = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_eap_Pkt_Tx_auth_Eap",
			Help: "Number of EAP request packets sent due to the authenticator choosing the EAP method",
		})
	onosaaaTransRespnotNak = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "onosaaa_trans_Resp_not_Nak",
			Help: "Number of transitions to response (received response other that NAK)",
		})
)

func exportVolthaKPI(kpi VolthaKPI) {

	for _, data := range kpi.SliceDatas {
		switch title := data.Metadata.Title; title {
		case "Ethernet", "PON":
			volthaTxBytesTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.TxBytes)

			volthaRxBytesTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.RxBytes)

			volthaTxPacketsTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.TxPackets)

			volthaRxPacketsTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.RxPackets)

			volthaTxErrorPacketsTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.TxErrorPackets)

			volthaRxErrorPacketsTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.RxErrorPackets)

			// TODO add metrics for:
			// TxBcastPackets
			// TxUnicastPackets
			// TxMulticastPackets
			// RxBcastPackets
			// RxMulticastPackets

		case "Ethernet_Bridge_Port_History":
			if data.Metadata.Context.Upstream == "True" {
				// ONU. Extended Ethernet statistics.
				volthaTxPacketsTotal.WithLabelValues(
					data.Metadata.LogicalDeviceID,
					data.Metadata.SerialNumber,
					data.Metadata.DeviceID,
					"NA", // InterfaceID
					"NA", // PonID
					"NA", // PortNumber
					data.Metadata.Title,
				).Add(data.Metrics.Packets)

				volthaTxBytesTotal.WithLabelValues(
					data.Metadata.LogicalDeviceID,
					data.Metadata.SerialNumber,
					data.Metadata.DeviceID,
					"NA", // InterfaceID
					"NA", // PonID
					"NA", // PortNumber
					data.Metadata.Title,
				).Add(data.Metrics.Octets)
			} else {
				// ONU. Extended Ethernet statistics.
				volthaRxPacketsTotal.WithLabelValues(
					data.Metadata.LogicalDeviceID,
					data.Metadata.SerialNumber,
					data.Metadata.DeviceID,
					"NA", // InterfaceID
					"NA", // PonID
					"NA", // PortNumber
					data.Metadata.Title,
				).Add(data.Metrics.Packets)

				volthaRxBytesTotal.WithLabelValues(
					data.Metadata.LogicalDeviceID,
					data.Metadata.SerialNumber,
					data.Metadata.DeviceID,
					"NA", // InterfaceID
					"NA", // PonID
					"NA", // PortNumber
					data.Metadata.Title,
				).Add(data.Metrics.Octets)
			}

		case "Ethernet_UNI_History":
			// ONU. Do nothing.

		case "FEC_History":
			// ONU. Do Nothing.

			volthaTxBytesTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.TxBytes)

			volthaRxBytesTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.RxBytes)

			volthaTxPacketsTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.TxPackets)

			volthaRxPacketsTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.RxPackets)

			volthaTxErrorPacketsTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.TxErrorPackets)

			volthaRxErrorPacketsTotal.WithLabelValues(
				data.Metadata.LogicalDeviceID,
				data.Metadata.SerialNumber,
				data.Metadata.DeviceID,
				data.Metadata.Context.InterfaceID,
				data.Metadata.Context.PonID,
				data.Metadata.Context.PortNumber,
				data.Metadata.Title,
			).Set(data.Metrics.RxErrorPackets)

			// TODO add metrics for:
			// TxBcastPackets
			// TxUnicastPackets
			// TxMulticastPackets
			// RxBcastPackets
			// RxMulticastPackets

		case "voltha.internal":
			// Voltha Internal. Do nothing.
		}
	}
}

func exportOnosKPI(kpi OnosKPI) {

	for _, data := range kpi.Ports {

		onosTxBytesTotal.WithLabelValues(
			kpi.DeviceID,
			data.PortID,
		).Set(data.TxBytes)

		onosRxBytesTotal.WithLabelValues(
			kpi.DeviceID,
			data.PortID,
		).Set(data.RxBytes)

		onosTxPacketsTotal.WithLabelValues(
			kpi.DeviceID,
			data.PortID,
		).Set(data.TxPackets)

		onosRxPacketsTotal.WithLabelValues(
			kpi.DeviceID,
			data.PortID,
		).Set(data.RxPackets)

		onosTxDropPacketsTotal.WithLabelValues(
			kpi.DeviceID,
			data.PortID,
		).Set(data.TxPacketsDrop)

		onosRxDropPacketsTotal.WithLabelValues(
			kpi.DeviceID,
			data.PortID,
		).Set(data.RxPacketsDrop)
	}
}

func exportImporterKPI(kpi ImporterKPI) {
	// TODO: add metrics for importer data
	logger.Info("To be implemented")
}

func exportOnosAaaKPI(kpi OnosAaaKPI) {

	onosaaaRxAcceptResponses.Set(kpi.RxAcceptResponses)

	onosaaaRxRejectResponses.Set(kpi.RxRejectResponses)

	onosaaaRxChallengeResponses.Set(kpi.RxChallengeResponses)

	onosaaaTxAccessRequests.Set(kpi.TxAccessRequests)

	onosaaaRxInvalidValidators.Set(kpi.RxInvalidValidators)

	onosaaaRxUnknownType.Set(kpi.RxUnknownType)

	onosaaaPendingRequests.Set(kpi.PendingRequests)

	onosaaaRxDroppedResponses.Set(kpi.RxDroppedResponses)

	onosaaaRxMalformedResponses.Set(kpi.RxMalformedResponses)

	onosaaaRxUnknownserver.Set(kpi.RxUnknownserver)

	onosaaaRequestRttMillis.Set(kpi.RequestRttMillis)

	onosaaaRequestReTx.Set(kpi.RequestReTx)

	onosaaaEapolLogoffRx.Set(kpi.RxeapolLogoff)

	onosaaaEapolResIdentityMsgTrans.Set(kpi.EapolResIdentityMsgTrans)

	onosaaaAuthSuccessTrans.Set(kpi.AuthSuccessTrans)

	onosaaaAuthFailureTrans.Set(kpi.AuthFailureTrans)

	onosaaaStartReqTrans.Set(kpi.StartReqTrans)

	onosaaaTxAccessReqPkt.Set(kpi.TxAccessReqPkt)

	onosaaaRxaccessChallPkt.Set(kpi.RxaccessChallPkt)

	onosaaaEapPktTxauthEap.Set(kpi.EapPktTxauthEap)

	onosaaaTransRespnotNak.Set(kpi.TransRespnotNak)
}

func export(topic *string, data []byte) {
	switch *topic {
	case "voltha.kpis":
		kpi := VolthaKPI{}
		err := json.Unmarshal(data, &kpi)
		if err != nil {
			log.Fatal(err)
		}
		exportVolthaKPI(kpi)
	case "onos.kpis":
		kpi := OnosKPI{}
		err := json.Unmarshal(data, &kpi)
		if err != nil {
			log.Fatal(err)
		}
		exportOnosKPI(kpi)
	case "importer.kpis":
		kpi := ImporterKPI{}
		err := json.Unmarshal(data, &kpi)
		if err != nil {
			log.Fatal(err)
		}
		exportImporterKPI(kpi)
	case "onos.aaa.stats.kpis":
		kpi := OnosAaaKPI{}
		err := json.Unmarshal(data, &kpi)
		if err != nil {
			log.Fatal(err)
		}
		exportOnosAaaKPI(kpi)
	default:
		logger.Warn("Unexpected export. Should not come here")
	}
}
