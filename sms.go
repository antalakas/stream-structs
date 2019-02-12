package stream_structs

import (
	"encoding/json"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/getsentry/raven-go"
	"sync"
	"time"
)

//### API Record
//{
//	"source" => "api"
//	"@id" => "5f81781e-83e9-491b-a621-3e76d0871608",
//	"converted_timestamp" => "2019-01-10T09:57:38.511Z",
//}
//
//### SMS Record
//{
//	"source": "sms",
//	"@id": "5f81781e-83e9-491b-a621-3e76d0871608",
//	"index_name": "sms-2019.01.09"
//}
type SMSEvent struct {
	ID                 	string 		`json:"@id"`
	Source             	string 		`json:"source"`
	ConvertedTimestamp 	time.Time 	`json:"converted_timestamp"`
	IndexName			string 		`json:"index_name"`
}

//{
//	"converted_timestamp": "2019-01-10T09:57:38.511Z",
//	"@id": "5f81781e-83e9-491b-a621-3e76d0871608",
//	"index_name": "sms-2019.01.09",
//	"converted": "true"
//}
type ConvertedEvent struct {
	ID                 	string 		`json:"@id"`
	ConvertedTimestamp 	time.Time 	`json:"converted_timestamp"`
	IndexName          	string 		`json:"index_name"`
	Converted			bool 		`json:"converted"`
}

type SMSBuild struct {
	SMSMap 			map[string]*kafka.Message
	ApiMessage		*kafka.Message
	SmsMessage		*kafka.Message
	SMSTimestamp	int64
}

func NewSMSBuild() *SMSBuild {
	var smsBuild SMSBuild

	smsBuild.SMSMap = make(map[string]*kafka.Message)
	smsBuild.ApiMessage = nil
	smsBuild.SmsMessage = nil
	smsBuild.SMSTimestamp = time.Now().Unix()

	return &smsBuild
}

func (smsBuild *SMSBuild) IsComplete() bool {

	if (smsBuild.ApiMessage != nil) && (smsBuild.SmsMessage != nil) {
		return true
	}
	return false
}

func (smsBuild *SMSBuild) IsExpired() bool {
	return false
}

func (smsBuild *SMSBuild) AppendSMS(eventID string,
	msg *kafka.Message) {
	switch eventID {
	case "1":
		smsBuild.ApiMessage = msg
		break
	case "2":
		smsBuild.SmsMessage = msg
		break
	}
}

func (smsBuild *SMSBuild) Pack() {
	if smsBuild.ApiMessage != nil {
		smsBuild.SMSMap["1"] = smsBuild.ApiMessage
	}
	if smsBuild.SMSMap != nil {
		smsBuild.SMSMap["2"] = smsBuild.SmsMessage
	}
}

type S2SMessage struct {
	S2sID		string
	IsComplete	bool
}

func NewS2SMessage(s2sID string, isComplete bool) *S2SMessage {
	var s2sMessage S2SMessage

	s2sMessage.S2sID = s2sID
	s2sMessage.IsComplete = isComplete

	return &s2sMessage
}

type MediationSMSMessage struct {
	SmsEvent 	SMSEvent
	Msg 		*kafka.Message
}

func NewMediationSMSMessage(smsEvent SMSEvent,
	msg *kafka.Message) *MediationSMSMessage {
	var mediationMessage MediationSMSMessage

	mediationMessage.SmsEvent = smsEvent
	mediationMessage.Msg = msg

	return &mediationMessage
}

func GetConvertedFromS2SID(s2sID string, isComplete bool,
	syncSmsMap *sync.Map) (ConvertedEvent, error) {
	// this is the struct to store the constructed ConvertedEvent
	var convertedEvent ConvertedEvent

	// retrieve fields of key: S2SID
	result, _ := syncSmsMap.Load(s2sID)

	if result == nil {
		// A Converted has been formed by this s2sID
		return convertedEvent,
			errors.New("A Converted has been formed by this s2sID")
	}

	smsBuild := result.(*SMSBuild)

	//fmt.Println("--------------")
	//fmt.Printf("%v\n", smsBuild.ApiMessage)
	//fmt.Printf("%v\n", smsBuild.SmsMessage)
	//fmt.Println("++++++++++++++")

	var ts time.Time
	var eventId string
	var indexName string

	// iterate in fields of key: S2SID
	for _, v := range smsBuild.SMSMap {
		// create an SMSEvent object from string
		var smsEvent SMSEvent

		if v == nil {
			// A Converted has been formed by this e2eID
			return convertedEvent,
				errors.New("A Converted has been formed by this s2sID")
		}

		err := json.Unmarshal([]byte(v.Value), &smsEvent)

		//messageContext := fmt.Sprintf("%s/%d/%d\t%s\t%s\n",
		//	v.Topic, v.Partition, v.Offset, v.Key, v.Value)

		//fmt.Println("--------------")
		////fmt.Println(messageContext)
		//fmt.Printf("%v\n", smsEvent)
		//fmt.Println("++++++++++++++")

		// check for error
		if err != nil {
			raven.CaptureErrorAndWait(err,
				map[string]string{"Converted generation": "Could not " +
					"unmarshal SMS field"})
		}

		// populate smsMap
		if smsEvent.Source == "api" {
			eventId = smsEvent.ID
			ts = smsEvent.ConvertedTimestamp
		}

		if smsEvent.Source == "sms" {
			indexName = smsEvent.IndexName
		}
	}

	convertedEvent.ID = eventId
	convertedEvent.ConvertedTimestamp = ts
	convertedEvent.IndexName = indexName
	convertedEvent.Converted = isComplete

	// Currently, do not permit sending the converted sms if it is not
	// completed - for sure IndexName would be empty
	if isComplete == false {
		return convertedEvent,
			errors.New("Actually there was no match")
	}

	//fmt.Println("--------------")
	//fmt.Printf("%v\n", convertedEvent)
	//fmt.Println("++++++++++++++")

	return convertedEvent, nil
}
