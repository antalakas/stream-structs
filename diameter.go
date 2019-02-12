package stream_structs

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/getsentry/raven-go"
	"strconv"
	"sync"
	"time"
)

type EDREvent struct {
	ID                 		string    `json:"@id"`
	Timestamp               time.Time `json:"@timestamp"`
	AppId                   string    `json:"APP_ID"`
	AppIdReadable           string    `json:"APP_ID_READABLE"`
	CommandCode             string    `json:"COMMAND_CODE"`
	CommandCodeReadable     string    `json:"COMMAND_CODE_READABLE"`
	Connection              string    `json:"CONNECTION"`
	Date                    string    `json:"DATE"`
	Destination             string    `json:"DESTINATION"`
	DestinationCode        	string    `json:"DESTINATION_CODE"`
	DestinationCountry      string    `json:"DESTINATION_COUNTRY"`
	DestinationRegion       string    `json:"DESTINATION_REGION"`
	DestHost                string    `json:"DEST_HOST"`
	DestRealm               string    `json:"DEST_REALM"`
	DiameterCp              string    `json:"DIAMETER_CP"`
	E2EId                   string    `json:"E2E_ID"`
	ErrorFlag               string    `json:"ERROR_FLAG"`
	EventId                 string    `json:"EVENT_ID"`
	ExpResultCode           string    `json:"EXP_RESULT_CODE"`
	FlowName                string    `json:"FLOW_NAME"`
	GeneratedTransaction    string    `json:"GENERATED_TRANSACTION"`
	HBHId                   string    `json:"HBH_ID"`
	HostName                string    `json:"HOST_NAME"`
	HomeIssuedRequest       string    `json:"Home_Issued_Request"`
	IndexName               string    `json:"INDEX_NAME"`
	InternalResult          string    `json:"INTERNAL_RESULT"`
	MessageDirection        string    `json:"MESSAGE_DIRECTION"`
	MsIsdn					string	  `json:"MSISDN"`
	OriginHost              string    `json:"ORIGIN_HOST"`
	OriginRealm             string    `json:"ORIGIN_REALM"`
	ProcessedDate           int    	  `json:"PROCESSED_DATE"`
	RedisKeyRef             string    `json:"REDIS_KEY_REF"`
	RemoteIp                string    `json:"REMOTE_IP"`
	Result                  string    `json:"RESULT"`
	ResultCode              string    `json:"RESULT_CODE"`
	SeqNumber               string    `json:"SEQ_NUMBER"`
	SessionId               string    `json:"SESSION_ID"`
	Source                  string    `json:"SOURCE"`
	SourceCode              string    `json:"SOURCE_CODE"`
	SourceCountry           string    `json:"SOURCE_COUNTRY"`
	SourceFilename          string    `json:"SOURCE_FILENAME"`
	SourceRegion            string    `json:"SOURCE_REGION"`
	Time                    string    `json:"TIME"`
	UserName                string    `json:"USER_NAME"`
	VisitedPLMNId           string    `json:"VISITED_PLMN_ID"`
	VisitedReadablePlmnId   string    `json:"VISITED_READABLE_PLMN_ID"`
	Type                    string    `json:"type"`
	Version                 string    `json:"@version"`
	Customer                string    `json:"@customer"`
}

type TDREvent struct {
	IngressSourceFilename	string `json:"INGRESS_SOURCE_FILENAME"`
	IngressSeqNumber        string `json:"INGRESS_SEQ_NUMBER"`
	ProcessedDate           int64  `json:"PROCESSED_DATE"`
	IngressHostName         string `json:"INGRESS_HOST_NAME"`
	//Date                	string `json:"DATE"`
	//Time               		string `json:"TIME"`
	IngressEventID1Ref     	string `json:"INGRESS_EVENTID1_ref"`
	IngressEventID1Rxd      int    `json:"INGRESS_EVENTID1_RXD"`
	IngressEventID2Txd      int    `json:"INGRESS_EVENTID2_TXD"`
	IngressEventID2Ref      string `json:"INGRESS_EVENTID2_ref"`
	IngressTxdRemoteIp      string `json:"INGRESS_Txd_REMOTE_IP"`
	IngressRxdRemoteIp     	string `json:"INGRESS_Rxd_REMOTE_IP"`
	//IngressSourcePeerId		string `json:"INGRESS_SOURCE_PEER_ID"`
	//IngressDestPeerId      	string `json:"INGRESS_DEST_PEER_ID"`
	CommandCode             string `json:"COMMAND_CODE"`
	CommandCodeReadable     string `json:"COMMAND_CODE_READABLE"`
	AppId                 	string `json:"APP_ID"`
	AppIdReadable           string `json:"APP_ID_READABLE"`
	IngressHbhId            string `json:"INGRESS_HBH_ID"`
	E2EId                 	string `json:"E2E_ID"`
	SessionId              	string `json:"SESSION_ID"`
	IngressReqOriginHost    string `json:"INGRESSreq_ORIGIN_HOST"`
	IngressReqOriginRealm   string `json:"INGRESSreq_ORIGIN_REALM"`
	IngressReqDestHost      string `json:"INGRESSreq_DEST_HOST"`
	IngressReqDestRealm     string `json:"INGRESSreq_DEST_REALM"`
	UserName                string `json:"USER_NAME"`
	MsIsdn                 	string `json:"MSISDN"`
	VisitedPlmnId           string `json:"VISITED_PLMN_ID"`
	EgressSourceFilename    string `json:"EGRESS_SOURCE_FILENAME"`
	EgressSeqNumber         string `json:"EGRESS_SEQ_NUMBER"`
	EgressHostname          string `json:"EGRESS_HOSTNAME"`
	EgressEventID3Ref       string `json:"EGRESS_EVENTID3_ref"`
	EgressEventID3Txd       int    `json:"EGRESS_EVENTID3_TXD"`
	EgressEventID4Ref       string `json:"EGRESS_EVENTID4_ref"`
	EgressEventID4Rxd       int    `json:"EGRESS_EVENTID4_RXD"`
	EgressTxdRemoteIp       string `json:"EGRESS_Txd_REMOTE_IP"`
	EgressRxdRemoteIp       string `json:"EGRESS_Rxd_REMOTE_IP"`
	//EgressSourcePeerId      string `json:"EGRESS_SOURCE_PEER_ID`
	//EgressDestPeerId        string `json:"EGRESS_DEST_PEER_ID"`
	EgressHbhId             string  `json:"EGRESS_HBH_ID"`
	EgressAnsOriginHost     string  `json:"EGRESS_Ans_ORIGIN_HOST"`
	EgressAnsOriginRealm    string  `json:"EGRESS_Ans_ORIGIN_REALM"`
	EgressAnsDestHost       string  `json:"EGRESS_Ans_DEST_HOST"`
	EgressAnsDestRealm      string  `json:"EGRESS_Ans_DEST_REALM"`
	VisitedPlmnReadable     string  `json:"VISITED_PLMN_READABLE"`
	IngressAnswerDelay      float64 `json:"INGRESS_ANSWER_DELAY"`
	EgressAnswerDelay       float64 `json:"EGRESS_ANSWER_DELAY"`
	IngressConnection       string  `json:"INGRESS_CONNECTION"`
	EgressConnection        string  `json:"EGRESS_CONNECTION"`
	//IngressCustomer         string `json:"INGRESS_CUSTOMER"`
	//EgressCustomer          string `json:"EGRESS_CUSTOMER"`
	IngressDiameterCp       string `json:"INGRESS_DIAMETER_CP"`
	EgressDiameterCp        string `json:"EGRESS_DIAMETER_CP"`
	ResultCode              string `json:"RESULT_CODE"`
	ExpResult               string `json:"EXP_RESULT"`
	Result                  string `json:"RESULT"`
	InternalResult          int    `json:"INTERNAL_RESULT"`
	//HomeNpId                string `json:"HOME_NP_ID"`
	HomeMno                 string `json:"HOME_MNO"`
	HomeMnoCountry          string `json:"HOME_MNO_COUNTRY"`
	HomeMnoRegion           string `json:"HOME_MNO_REGION"`
	SourceCode              string `json:"SOURCE_CODE"`
	//VisitedNpId             string `json:"VISITED_NP_ID"`
	VisitedMno              string `json:"VISITED_MNO"`
	VisitedMnoCountry       string `json:"VISITED_MNO_COUNTRY"`
	VisitedMnoRegion        string `json:"VISITED_MNO_REGION"`
	DestinationCode        	string `json:"DESTINATION_CODE"`
	//ROAnsDest 			  string `json:"REQUEST_ORIGIN_ANSWERDEST_ID"`
	ROAnsDestMno 		    string `json:"REQUEST_ORIGIN_ANSWERDEST_MNO"`
	ROAnsDestMnoCountry   	string `json:"REQUEST_ORIGIN_ANSWERDEST_MNO_Country"`
	ROAnsDestMnoRegion    	string `json:"REQUEST_ORIGIN_ANSWERDEST_MNO_Region"`
	//RDAnsOriginId         string `json:"REQUEST_DEST_ANSWERORIGIN_ID"`
	RDAnsOriginMno        	string `json:"REQUEST_DEST_ANSWERORIGIN_MNO"`
	RDAnsOriginMnoCountry 	string `json:"REQUEST_DEST_ANSWERORIGIN_MNO_Country"`
	RDAnsOriginMnoRegion  	string `json:"REQUEST_DEST_ANSWERORIGIN_MNO_Region"`
	IndexName          	  	string 	  `json:"INDEX_NAME"`
	Timestamp          	  	time.Time `json:"@timestamp"`
	Type                  	string `json:"type"`
	ID                 	  	string `json:"@id"`
	SourceHost            	string `json:"SOURCE_HOST"`
	DestinationHost       	string `json:"DESTINATION_HOST"`
	Customer              	string `json:"@customer"`
	EdrComplete       	  	bool   `json:"EDR_COMPLETE"`
}

type TDRCounterEvent struct {
	Timestamp          	    time.Time `json:"@timestamp"`
	SourceCode              string `json:"SOURCE_CODE"`
	DestinationCode        	string `json:"DESTINATION_CODE"`
	AppId                 	string `json:"APP_ID"`
	VisitedPlmnReadable     string `json:"VISITED_PLMN_READABLE"`
	CommandCode             string `json:"COMMAND_CODE"`
	EgressDiameterCp        string `json:"EGRESS_DIAMETER_CP"`
	IngressDiameterCp       string `json:"INGRESS_DIAMETER_CP"`
	Result                  string `json:"RESULT"`
}

type EDRBuild struct {
	//EDRMap 			map[string]*kafka.Message
	//EDRTimestamp	int64

	EDRMap 			map[string]*kafka.Message
	Message1		*kafka.Message
	Message2		*kafka.Message
	Message3		*kafka.Message
	Message4		*kafka.Message
	EDRTimestamp	int64
}

func NewEDRBuild() *EDRBuild {
	//var edrBuild EDRBuild
	//
	//edrBuild.EDRMap = make(map[string]*kafka.Message)
	//edrBuild.EDRTimestamp = time.Now().Unix()
	//
	//return &edrBuild

	var edrBuild EDRBuild

	edrBuild.EDRMap = make(map[string]*kafka.Message)
	edrBuild.Message1 = nil
	edrBuild.Message2 = nil
	edrBuild.Message3 = nil
	edrBuild.Message4 = nil
	edrBuild.EDRTimestamp = time.Now().Unix()

	return &edrBuild
}

func (edrBuild *EDRBuild) IsComplete() bool {
	//if len(edrBuild.EDRMap) == 4 {
	//	return true
	//}
	//return false

	if (edrBuild.Message1 != nil) && (edrBuild.Message2 != nil) &&
		(edrBuild.Message3 != nil) && (edrBuild.Message4 != nil) {
		return true
	}
	return false
}

func (edrBuild *EDRBuild) IsExpired() bool {
	return false
}

func (edrBuild *EDRBuild) AppendEDR(eventID string,
	msg *kafka.Message) {
	//edrBuild.EDRMap[eventID] = msg

	switch eventID {
	case "1":
		edrBuild.Message1 = msg
		break
	case "2":
		edrBuild.Message2 = msg
		break
	case "3":
		edrBuild.Message3 = msg
		break
	case "4":
		edrBuild.Message4 = msg
		break
	}
}

func (edrBuild *EDRBuild) Pack() {
	if edrBuild.Message1 != nil {
		edrBuild.EDRMap["1"] = edrBuild.Message1
	}
	if edrBuild.Message2 != nil {
		edrBuild.EDRMap["2"] = edrBuild.Message2
	}
	if edrBuild.Message3 != nil {
		edrBuild.EDRMap["3"] = edrBuild.Message3
	}
	if edrBuild.Message1 != nil {
		edrBuild.EDRMap["4"] = edrBuild.Message4
	}
}

type E2EMessage struct {
	E2eID		string
	IsComplete	bool
}

func NewE2EMessage(e2eID string, isComplete bool) *E2EMessage {
	var e2eMessage E2EMessage

	e2eMessage.E2eID = e2eID
	e2eMessage.IsComplete = isComplete

	return &e2eMessage
}

type EDRSmall struct {
	E2EId                   string    `json:"E2E_ID"`
	EventId                 string    `json:"EVENT_ID"`
	IndexName               string    `json:"INDEX_NAME"`
}

type MediationDiameterMessage struct {
	EdrSmall 	EDRSmall
	Msg 		*kafka.Message
}

func NewMediationDiameterMessage(edrSmall EDRSmall,
	msg *kafka.Message) *MediationDiameterMessage {
	var mediationMessage MediationDiameterMessage

	mediationMessage.EdrSmall = edrSmall
	mediationMessage.Msg = msg

	return &mediationMessage
}

func GetTDRFromE2EID(e2eID string, isComplete bool, tdrMap *sync.Map) (TDREvent,
	map[int]EDREvent) {
	// this is the struct to store the constructed tdr
	var tdrEvent TDREvent

	// retrieve fields of key: E2EID
	result, _ := tdrMap.Load(e2eID)

	if result == nil {
		// A TDR has been formed by this e2eID
		return tdrEvent, nil
	}

	edrBuild := result.(*EDRBuild)

	edrMap := make(map[int]EDREvent)

	// iterate in fields of key: E2EID
	for _, v := range edrBuild.EDRMap {
		// create an EDR object from string
		var edrEvent EDREvent

		if v == nil {
			// A TDR has been formed by this e2eID
			return tdrEvent, nil
		}

		err := json.Unmarshal([]byte(v.Value), &edrEvent)

		// check for error
		if err != nil {
			raven.CaptureErrorAndWait(err,
				map[string]string{"TDR generation": "Could not " +
					"unmarshal EDR field"})
		}

		// populate edrMap
		eventId, _ := strconv.Atoi(edrEvent.EventId)
		edrMap[eventId] = edrEvent
	}

	tdrEvent.IngressEventID1Rxd = 0
	tdrEvent.IngressEventID2Txd = 0
	tdrEvent.EgressEventID3Txd = 0
	tdrEvent.EgressEventID4Rxd = 0
	tdrEvent.InternalResult = 0

	tdrEvent.DestinationHost = ""

	if edr, ok := edrMap[1]; ok {
		tdrEvent.Customer = edr.Customer
		tdrEvent.IngressSourceFilename = edr.SourceFilename
		tdrEvent.IngressSeqNumber = edr.SeqNumber
		tdrEvent.ProcessedDate = time.Now().Unix()
		tdrEvent.IngressHostName = edr.HostName
		tdrEvent.IngressEventID1Ref = edr.ID
		tdrEvent.IngressEventID1Rxd = 1
		tdrEvent.IngressRxdRemoteIp = edr.RemoteIp
		tdrEvent.CommandCode = edr.CommandCode
		tdrEvent.CommandCodeReadable = edr.CommandCodeReadable
		tdrEvent.AppId = edr.AppId
		tdrEvent.AppIdReadable = edr.AppIdReadable
		tdrEvent.IngressHbhId = edr.HBHId
		tdrEvent.SessionId = edr.SessionId
		tdrEvent.IngressReqOriginHost = edr.OriginHost
		tdrEvent.IngressReqOriginRealm = edr.OriginRealm
		tdrEvent.IngressReqDestHost = edr.DestHost
		tdrEvent.IngressReqDestRealm = edr.DestRealm
		tdrEvent.UserName = edr.UserName
		tdrEvent.MsIsdn = edr.MsIsdn
		tdrEvent.VisitedPlmnId = edr.VisitedPLMNId
		tdrEvent.VisitedPlmnReadable = edr.VisitedReadablePlmnId
		tdrEvent.IngressDiameterCp = edr.DiameterCp
		tdrEvent.IndexName = edr.IndexName
		tdrEvent.Timestamp = edr.Timestamp
		tdrEvent.IngressConnection = edr.Connection
		atId := fmt.Sprintf("%s_%s_%s",
			"tdr", edr.SourceFilename, edr.SeqNumber)
		tdrEvent.ID = atId

		tdrEvent.ROAnsDestMno = edr.Source
		tdrEvent.ROAnsDestMnoCountry = edr.SourceCountry
		tdrEvent.ROAnsDestMnoRegion = edr.SourceRegion
		tdrEvent.RDAnsOriginMno = edr.Destination
		tdrEvent.RDAnsOriginMnoCountry = edr.DestinationCountry
		tdrEvent.RDAnsOriginMnoRegion = edr.DestinationRegion

		switch edr.HomeIssuedRequest {
		case "Yes":
			tdrEvent.HomeMno = edr.Source
			tdrEvent.HomeMnoCountry = edr.SourceCountry
			tdrEvent.HomeMnoRegion = edr.SourceRegion
			tdrEvent.VisitedMno = edr.Destination
			tdrEvent.VisitedMnoCountry = edr.DestinationCountry
			tdrEvent.VisitedMnoRegion = edr.DestinationRegion
			tdrEvent.SourceCode = edr.SourceCode
			tdrEvent.DestinationCode = edr.DestinationCode
		case "No":
			tdrEvent.HomeMno = edr.Destination
			tdrEvent.HomeMnoCountry = edr.DestinationCountry
			tdrEvent.HomeMnoRegion = edr.DestinationRegion
			tdrEvent.VisitedMno = edr.Source
			tdrEvent.VisitedMnoCountry = edr.SourceCountry
			tdrEvent.VisitedMnoRegion = edr.SourceRegion
			tdrEvent.SourceCode = edr.DestinationCode
			tdrEvent.DestinationCode = edr.SourceCode
		}

		tdrEvent.SourceHost = edr.OriginHost
		tdrEvent.DestinationHost = edr.DestHost
	}

	if edr, ok := edrMap[2]; ok {
		tdrEvent.IngressEventID2Txd = 1
		tdrEvent.IngressEventID2Ref = edr.ID
		tdrEvent.IngressTxdRemoteIp = edr.RemoteIp

		// Filling the following values means that later, redis
		// persistence for counters will find proper values
		if edr1, ok := edrMap[1]; ok {
			//edr.Source = edr1.Source
			edr.SourceCode = edr1.SourceCode
			//edr.SourceCountry = edr1.SourceCountry
			//edr.SourceRegion = edr1.SourceRegion

			//edr.Destination = edr1.Destination
			edr.DestinationCode = edr1.DestinationCode
			//edr.DestinationCountry = edr1.DestinationCountry
			//edr.DestinationRegion = edr1.DestinationRegion
		}

		if tdrEvent.DestinationHost == "" {
			tdrEvent.DestinationHost = edr.OriginHost
		}

		if tdrEvent.MsIsdn == "" {
			tdrEvent.MsIsdn = edr.MsIsdn
		}
	}

	if edr, ok := edrMap[3]; ok {
		tdrEvent.EgressSourceFilename = edr.SourceFilename
		tdrEvent.EgressSeqNumber = edr.SeqNumber
		tdrEvent.EgressHostname = edr.HostName
		tdrEvent.EgressEventID3Ref = edr.ID
		tdrEvent.EgressEventID3Txd = 1
		tdrEvent.EgressTxdRemoteIp = edr.RemoteIp
		tdrEvent.EgressHbhId = edr.HBHId
		tdrEvent.EgressDiameterCp = edr.DiameterCp

		// Filling the following values means that later, redis
		// persistence for counters will find proper values
		if edr1, ok := edrMap[1]; ok {
			edr.SourceCode = edr1.SourceCode
			edr.DestinationCode = edr1.DestinationCode
		}
	}

	if tdrEvent.IngressDiameterCp == "" && tdrEvent.EgressDiameterCp == "" {
		// We do not need to persist this kind of message
		return tdrEvent, nil
	}

	if edr, ok := edrMap[4]; ok {
		tdrEvent.EgressEventID4Ref = edr.ID
		tdrEvent.EgressEventID4Rxd = 1
		tdrEvent.EgressRxdRemoteIp = edr.RemoteIp
		tdrEvent.EgressAnsOriginHost= edr.OriginHost
		tdrEvent.EgressAnsOriginRealm = edr.OriginRealm
		tdrEvent.EgressAnsDestHost = edr.DestHost
		tdrEvent.EgressAnsDestRealm = edr.DestRealm

		// Filling the following values means that later, redis
		// persistence for counters will find proper values
		if edr1, ok := edrMap[1]; ok {
			edr.SourceCode = edr1.SourceCode
			edr.DestinationCode = edr1.DestinationCode
		}
	}

	if tdrEvent.IngressEventID1Rxd == 1 &&
		tdrEvent.IngressEventID2Txd == 1 {
		tdrEvent.IngressAnswerDelay =
			float64(edrMap[2].Timestamp.UnixNano() -
				edrMap[1].Timestamp.UnixNano()) / 1000000000.0
	}

	if tdrEvent.EgressEventID3Txd == 1 &&
		tdrEvent.EgressEventID4Rxd == 1 {
		tdrEvent.EgressAnswerDelay =
			float64(edrMap[4].Timestamp.UnixNano() -
				edrMap[3].Timestamp.UnixNano()) / 1000000000.0
		tdrEvent.EgressConnection = edrMap[3].HostName + "_" +
			edrMap[4].RemoteIp
	}

	if tdrEvent.IngressEventID1Rxd == 1 &&
		tdrEvent.IngressEventID2Txd == 1 &&
		tdrEvent.EgressEventID3Txd == 1 &&
		tdrEvent.EgressEventID4Rxd == 1 {

		tdrEvent.InternalResult = 1
		tdrEvent.Result = edrMap[2].Result
		tdrEvent.ResultCode = edrMap[2].ResultCode
		tdrEvent.ExpResult = edrMap[2].ExpResultCode
	}

	tdrEvent.E2EId = e2eID
	tdrEvent.Type = "diameter-tdr"
	tdrEvent.EdrComplete = isComplete

	return tdrEvent, edrMap
}

// DiameterCounterMessage struct
type DiameterCounterMessage struct {
	Key    string
	Metric string
}

// NewDiameterCounterMessage function
func NewDiameterCounterMessage(key string, metric string) *DiameterCounterMessage {
	var diameterCounterMessage DiameterCounterMessage

	diameterCounterMessage.Key = key
	diameterCounterMessage.Metric = metric

	return &diameterCounterMessage
}
