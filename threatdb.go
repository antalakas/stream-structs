package stream_structs

// AlertParam struct
type AlertParam struct {
	ID           string        `json:"id"`
	Title        string        `json:"title"`
	Body         string        `json:"body"`
	Criticality  int           `json:"criticality"`
	EventID      string        `json:"event_id"`
	UserIDs      []string      `json:"user_ids"`
	Devid        string        `json:"device_serial"`
	BranchID     int       	   `json:"branch_id"`
	Country      string        `json:"country"`
	AsnRegistry  string        `json:"asn_registry"`
	Asn          string        `json:"asn"`
	Continent    string        `json:"continent"`
	Geolocation  string        `json:"geolocation"`
	Reasoning    string        `json:"reasoning"`
	OrgID        int           `json:"organization_id"`
	AlertType    string        `json:"alert_type"`
	Sources      []SourceParam `json:"sources"`
	Entity       string        `json:"entity"`
	EntityType   string        `json:"entity_type"`
	Entity2      string        `json:"entity2"`
	EntityType2  string        `json:"entity_type2"`
	Subtype      string        `json:"subtype"`
	TimeDetected int64         `json:"time_detected"`
	TimeOccurred int64         `json:"time_occured"`
	TimeContext  string        `json:"time_context"`
	CanBeGrouped bool          `json:"can_be_grouped"`
	Blocked 	 bool          `json:"blocked"`
	Tri          float64       `json:"tri"`
	Service      string        `json:"service"`
}

// NewAlertParam function
func NewAlertParam() *AlertParam {
	var alertParam AlertParam
	return &alertParam
}

// AlertMessage struct
type AlertMessage struct {
	Alert *AlertParam    	`json:"alert"`
	Event *CollationEvent 	`json:"event"`
}

// NewAlertMessage function
func NewAlertMessage(alert *AlertParam, event *CollationEvent) *AlertMessage {
	var alertMessage AlertMessage

	alertMessage.Alert = alert
	alertMessage.Event = event

	return &alertMessage
}

// SourceParam struct
type SourceParam struct {
	Source           *Source `json:"source"`
	SourceOccurences int     `json:"occurence"`
}

// NewSourceParam function
func NewSourceParam(source *Source, sourceOccurences int) *SourceParam {
	var sourceParam SourceParam

	sourceParam.Source = source
	sourceParam.SourceOccurences = sourceOccurences

	return &sourceParam
}

// Source struct
type Source struct {
	File        string `json:"file"`
	SourceType  string `json:"type"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	URL         string `json:"url"`
	Downloaded  string `json:"downloaded"`
	Category    string `json:"category"`
}

// NewSource function
func NewSource() *Source {
	var source Source
	return &source
}

// Device struct
type Device struct {
	DevInfo string `json:"device_name"`
}
