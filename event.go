package stream_structs

import "time"

// CollationEvent struct
type CollationEvent struct {
	ID                       string    `json:"@id"`
	Timestamp                time.Time `json:"@timestamp"`
	SourceEvent              string    `json:"@source_event"`
	DeviceSerial             string    `json:"device_serial"`
	Type                     string    `json:"type"`
	Subtype                  string    `json:"subtype"`
	SourceIP                 string    `json:"source_ip"`
	SourcePort               int       `json:"source_port"`
	DestinationIP            string    `json:"destination_ip"`
	DestinationPort          int       `json:"destination_port"`
	CPU                      int       `json:"cpu"`
	Action                   string    `json:"action"`
	Service                  string    `json:"service"`
	Scope                    string    `json:"scope"`
	Direction                string    `json:"direction"`
	OrganizationID           int       `json:"organization_id"`
	BranchID           		 int       `json:"branch_id"`
	ThreatID                 string    `json:"threat_id"`
	TotalBytes               int64     `json:"total_bytes"`
	SourceContinentCode      string    `json:"source_continent_code"`
	SourceCountryName        string    `json:"source_country_name"`
	SourceAsName             string    `json:"source_as_name"`
	SourceAsNumber           string    `json:"source_as_number"`
	SourceLatitude           float64   `json:"source_latitude"`
	SourceLongitude          float64   `json:"source_longitude"`
	DestinationContinentCode string    `json:"destination_continent_code"`
	DestinationCountryName   string    `json:"destination_country_name"`
	DestinationAsName        string    `json:"destination_as_name"`
	DestinationAsNumber      string    `json:"destination_as_number"`
	DestinationLatitude      float64   `json:"destination_latitude"`
	DestinationLongitude     float64   `json:"destination_longitude"`
	AlertID                  string    `json:"alert_id"`
}

// {
// 	// "@version":1,
// 	"@timestamp":"2017-01-20T12:18:17.975Z",
// 	// "@srcevent":"syslog",
// 	// "@vendor":"ws",
// 	// "host":"195.46.9.68",
// 	// "&raw_message_bytesize":725,
// 	// "&logstash_frontend":"eu_ams01_ams-logfe-01",
// 	"@id":"86a406a6-591b-49d7-8b0d-60fa641d6364",
// 	// "device_ip":"195.46.9.68",
// 	"action":"allow",
// 	"severity":1,
// 	"category":"business and economy",
// 	"duration":7,
// 	"destination_port":443,
// 	"destination_ip":"176.9.58.133",
// 	"source_ip":"10.14.4.8",
// 	"source_port":56028,
// 	"sent_bytes":543,
// 	"received_bytes":5864,
// 	"client_type":"mozilla/5.0_(windows_nt_6.1;_wow64)_applewebkit/537.36_(khtml,_like_gecko)_chrome/55.0.2883.87_safari/537.36",
// 	"resource_type":"image/png",
// 	"protocol":"tcp",
// 	"type":"acc_ctrl",
// 	"subtype":"web",
// 	"scope":"action",
// 	"destination_name":"www.belibasakis.gr",
// 	"resource":"/image/cache/catalog/brands/rado-200x100.png",
// 	"source_user":"sofra maria",
// 	"info":"user dn: ldap://cardiff.local ou=stores,ou=dry,ou=cardiff,dc=cardiff,dc=local/",
// 	"reason":"file type permitted",
// 	"device_serial":"AFW01WSAPT",
// 	"organization_id":161,
// 	"branch_id":161,
// 	"direction":"out",
// 	"destination_country_code":"deu",
// 	"destination_country_name":"germany",
// 	"destination_continent_code":"eu",
// 	// "destination_latitude":51.29929999999999,
// 	// "destination_longitude":9.491000000000014,
// 	// "destination_as_name":"hetzner-as",
// 	"destination_as_number":"as24940",
// 	"sent_packets":0,
// 	"received_packets":0,
// 	"total_bytes":6407,
// 	"total_packets":0,
// 	"service":"https",
// 	// "&logstash_febe_latency_sec":0.568000078201294,
// 	// "&logstash_backend":"eu_ams01_ams-logbe-02"
// }

// RulesetCollationEvent struct
type RulesetCollationEvent struct {
	ID              string  `json:"@id"`
	AlertID         string  `json:"alert_id"`
	SourceLatitude  float64 `json:"source_latitude"`
	SourceLongitude float64 `json:"source_longitude"`
	OrganizationID  int     `json:"organization_id"`

	// --------------
	// General Fields
	// --------------

	// Device ID
	// Unique identifier of the device
	DeviceSerial string `json:"device_serial"`

	// Type of the event
	Type string `json:"type"`

	// Subtype of the event
	SubType string `json:"subtype"`

	// in|out | prv | unknown
	// Direction
	// Direction of the session
	Direction string `json:"direction"`

	// action|error | state | info
	// Scope
	// Scope of the event
	// Action : e.g. a FW allowed or denied a connection.
	Scope string `json:"scope"`

	// 3-letter country code of the source IP
	SourceCountryCode string `json:"source_country_code"`

	// Country name of the source IP
	SourceCountryName string `json:"source_country_name"`

	// 2-letter continent code of the source IP
	SourceContinentCode string `json:"source_continent_code"`

	// Autonomous System Number of the source IP
	SourceAsNumber string `json:"source_as_number"`

	// 3-letter country code of the destination IP
	DestinationCountryCode string `json:"destination_country_code"`

	// Country name of the destination IP
	DestinationCountryName string `json:"destination_country_name"`

	// 2-letter continent code of the destination IP
	DestinationContinentCode string `json:"destination_continent_code"`

	// Autonomous System Number of the destination IP
	DestinationAsNumber string `json:"destination_as_number"`

	///////////////
	// Traffic Type
	///////////////

	// connection_type
	// session | flow
	// Connection type
	ConnectionType string `json:"connection_type"`

	// start | end
	// Session
	// Whether the session started or ended.
	// Denied sessions are classed as end
	Session               string `json:"session"`
	SourcePort            int    `json:"source_port"`
	SourceIP              string `json:"source_ip"`
	SourceMacAddress      string `json:"source_mac_address"`
	SourceInterface       string `json:"source_interface"`
	SourceZone            string `json:"source_zone"`
	DestinationZone       string `json:"destination_zone"`
	DestinationInterface  string `json:"destination_interface"`
	DestinationMacAddress string `json:"destination_mac_address"`
	DestinationIP         string `json:"destination_ip"`
	DestinationPort       int    `json:"destination_port"`
	Service               string `json:"service"`
	SourceName            string `json:"source_name"`
	DestinationName       string `json:"destination_name"`
	SourceUser            string `json:"source_user"`
	SourceGroup           string `json:"source_group"`
	DestinationUser       string `json:"destination_user"`
	DestinationGroup      string `json:"destination_group"`
	ApplicationName       string `json:"application_name"`
	Category              string `json:"category"`
	Duration              int    `json:"duration"`
	SentBytes             int64  `json:"sent_bytes"`
	ReceivedBytes         int64  `json:"received_bytes"`
	TotalBytes            int64  `json:"total_bytes"`
	SentPackets           int64  `json:"sent_packets"`
	ReceivedPackets       int64  `json:"received_packets"`
	TotalPackets          int64  `json:"total_packets"`

	// snat|dnat
	// snat+dnat|noop
	// Nat Type
	// Type of NAT performed
	NatType string `json:"nat_type"`

	// allow|deny
	// Action
	// Firewall action
	Action string `json:"action"`

	// Sub-Action of the firewall
	// Deny subactions : drop, block
	// Allow subactions : monitor, excempt
	Subaction string `json:"subaction"`

	Protocol string `json:"protocol"`
	Error    string `json:"error"`

	// Flags set for TCP sessions
	// TDB: convert to array of strings
	// String (array)
	TCPFlags string `json:"tcp_flags"`

	// normal|timeout | server-reset | client-reset
	// reason
	// Reason the session closed
	Reason string `json:"reason"`

	SourceVlan      string `json:"source_vlan"`
	DestinationVlan string `json:"destination_vlan"`
	IcmpType        int    `json:"icmp_type"`
	IcmpCode        int    `json:"icmp_code"`

	//////////////
	// Threat Type
	//////////////
	// [anomaly]|malware|ips
	Subtype string `json:"subtype"`
	// SourcePort           int `json:"source_port"`
	// SourceIP             string `json:"source_ip"`
	// SourceInterface      string `json:"source_interface"`
	// DestinationInterface string `json:"destination_interface"`
	// DestinationIP        string `json:"destination_ip"`
	// DestinationPort      int `json:"destination_port"`
	// Service              string `json:"service"`
	// SourceUser           string `json:"source_user"`
	// SourceGroup          string `json:"source_group"`
	// DestinationUser      string `json:"destination_user"`
	// DestinationGroup     string `json:"destination_group"`
	// Protocol   string `json:"protocol"`
	// Direction  string `json:"direction"`
	// Action     string `json:"action"`
	// Subaction  string `json:"subaction"`
	Monitor    string `json:"monitor"`
	ThreatName string `json:"threat_name"`
	FileName   string `json:"file_name"`
	FileType   string `json:"file_type"`
	FileHash   string `json:"file_hash"`
	// Category   string `json:"category"`
	Cve      string `json:"cve"`
	Severity int    `json:"severity"`

	///////////////
	// Email Events
	///////////////
	// spam|regular
	// Subtype              string `json:"subtype"`
	// SourcePort           int    `json:"source_port"`
	// SourceIP             string `json:"source_ip"`
	// SourceInterface      string `json:"source_interface"`
	// DestinationInterface string `json:"destination_interface"`
	// DestinationIP        string `json:"destination_ip"`
	// DestinationPort      int    `json:"destination_port"`
	// Service              string `json:"service"`
	// SourceUser           string `json:"source_user"`
	// SourceGroup          string `json:"source_group"`
	// DestinationUser      string `json:"destination_user"`
	// DestinationGroup string `json:"destination_group"`
	Size       int `json:"size"`
	Attachment string `json:"attachment"`
	// Action           string `json:"action"`
	// Subaction        string `json:"subaction"`

	/////////////
	// VPN Events
	/////////////
	// Subtype              string `json:"subtype"`
	// State                string `json:"state"`
	// Error                string `json:"error"`
	// Action               string `json:"action"`
	// Reason               string `json:"reason"`
	// DestinationIP        string `json:"destination_ip"`
	// SourceIP             string `json:"source_ip"`
	// DestinationPort      int    `json:"destination_port"`
	// SourcePort           int    `json:"source_port"`
	// SourceInterface      string `json:"source_interface"`
	// DestinationInterface string `json:"destination_interface"`
	TunnelIP string `json:"tunnel_ip"`
	// SourceUser       string `json:"source_user"`
	// SourceGroup      string `json:"source_group"`
	// DestinationUser  string `json:"destination_user"`
	// DestinationGroup string `json:"destination_group"`
	// SentBytes     int64  `json:"sent_bytes"`
	// ReceivedBytes int64  `json:"received_bytes"`
	// TotalBytes    int64  `json:"total_bytes"`
	// Duration      int `json:"duration"`
	Mode  string `json:"mode"`
	Phase int    `json:"phase"`

	////////////////
	// System Events
	////////////////
	// Subtype         string `json:"subtype"`
	State string `json:"state"`
	// SourceIP        string `json:"source_ip"`
	// DestinationIP   string `json:"destination_ip"`
	Object string `json:"object"`
	Info   string `json:"info"`
	// SourceInterface string `json:"source_interface"`
	// NatType string `json:"nat_type"`
	// Action   string `json:"action"`
	// Protocol string `json:"protocol"`
	// Error    string `json:"error"`

	/////////////
	// AAA Events
	/////////////
	// Subtype              string `json:"subtype"`
	// Action               string `json:"action"`
	// Error                string `json:"error"`
	// Object               string `json:"object"`
	// Reason               string `json:"reason"`
	// Info                 string `json:"info"`
	// SourceIP             string `json:"source_ip"`
	// DestinationIP        string `json:"destination_ip"`
	// SourceUser           string `json:"source_user"`
	// SourceName           string `json:"source_name"`
	// SourceGroup          string `json:"source_group"`
	// SourceInterface      string `json:"source_interface"`
	// DestinationInterface string `json:"destination_interface"`

	// User Interface
	// Mutually exclusive
	// If accesed through https/ssh/telnet
	// Service string `json:"service"`

	// local | remote
	// User Interface
	// Mutually exclusive
	// If accesed through cli/gui
	UI string `json:"ui"`

	/////////////////////
	// Performance Events
	/////////////////////
	// Subtype       string `json:"subtype"`

	// percentage
	CPU int `json:"cpu"`

	// percentage
	Memory int `json:"memory"`

	TotalSessions int `json:"total_sessions"`

	// ASA
	MaxSessions int `json:"max_sessions"`

	// FGT
	SetupRate int `json:"setup_rate"`

	////////////////////////
	// Access Control Events
	////////////////////////
	// Subtype         string `json:"subtype"`
	// SourcePort      int `json:"source_port"`
	// SourceIP        string `json:"source_ip"`
	// DestinationPort int `json:"destination_port"`
	// DestinationIP   string `json:"destination_ip"`
	// Protocol        string `json:"protocol"`
	// Service         string `json:"service"`
	// Action          string `json:"action"`
	// Subaction       string `json:"subaction"`
	Excempt     string `json:"excempt"`
	RequestType string `json:"request_type"`
	ClientType  string `json:"client_type"`
	ServerType  string `json:"server_type"`
	// DestinationName string `json:"destination_name"`
	// SentBytes       int64 `json:"sent_bytes"`
	// Received_bytes  int64 `json:"received_bytes"`
	// TotalBytes      int64 `json:"total_bytes"`
	// SourceUser      string `json:"source_user"`
	Resource     string `json:"resource"`
	ResourceType string `json:"resource_type"`
	// ApplicationName string `json:"application_name"`
	// Category        string `json:"category"`
}

// CounterCollationEvent struct
type CounterCollationEvent struct {
	// Timestamp when the log event received in listeners
	Timestamp time.Time `json:"@timestamp"`
	// Unique identifier for the log event
	ID string `json:"@id"`
	// Organization identification
	OrgID int `json:"organization_id"`
	// Organization identification
	BranchID int `json:"branch_id"`
	// Unique identifier of the device
	DeviceSerial string `json:"device_serial"`
	// Type of the event
	Type string `json:"type"`
	// Subtype of the event
	SubType string `json:"subtype"`
	// Firewall action
	// allow|deny
	Action string `json:"action"`
	// Service of the session. If service Service unavailable, it’s
	// proto/destination_port.
	Service string `json:"service"`
	// Scope of the event
	// action|error | state | info
	// Action : e.g. a FW allowed or denied a connection.
	Scope string `json:"scope"`
	// session | flow
	ConnectionType string `json:"connection_type"`
	// Whether the session started or ended. Denied sessions are classed as end
	// start | end
	Session string `json:"session"`
	// Sender’s email address in case of threat through email
	From string `json:"from"`
	// Recipient’s email address in case of threat through email
	To string `json:"to"`
	// Direction of the session
	// in|out | prv | unknown
	Direction string `json:"direction"`
	// IP address related to the event
	SourceIP string `json:"source_ip"`
	// IP address related to the event
	DestinationIP string `json:"destination_ip"`
	//Source port of the session
	SourcePort int `json:"source_port"`
	// Destination port of the session
	DestinationPort int `json:"destination_port"`
	// Country name of the source IP
	SourceCountryCode string `json:"source_country_code"`
	// Country name of the destination IP
	DestinationCountryCode string `json:"destination_country_code"`
	SourceAsNumber         string `json:"source_as_number"`
	DestinationAsNumber    string `json:"destination_as_number"`
	ThreatID               string `json:"threat_id"`
	// Total bytes
	TotalBytes int `json:"total_bytes"`
	Sample int `json:"&sample"`
}

// NewCollationEvent function
func NewCollationEvent() *CollationEvent {
	var collationEvent CollationEvent
	return &collationEvent
}

// NewRulesetCollationEvent function
func NewRulesetCollationEvent() *RulesetCollationEvent {
	var rulesetCollationEvent RulesetCollationEvent
	return &rulesetCollationEvent
}

// CounterMessage struct
type CounterMessage struct {
	Key       string
	CountType string
	Totalbyte int
}

// NewCounterMessage function
func NewCounterMessage(key string, totalbyte int) *CounterMessage {
	var counterMessage CounterMessage

	counterMessage.Key = key
	counterMessage.Totalbyte = totalbyte

	return &counterMessage
}

// ThreatIQCounterMessage struct
type ThreatIQCounterMessage struct {
	Key    string
	Metric string
}

// NewThreatIQCounterMessage function
func NewThreatIQCounterMessage(key string, metric string) *ThreatIQCounterMessage {
	var threatIQCounterMessage ThreatIQCounterMessage

	threatIQCounterMessage.Key = key
	threatIQCounterMessage.Metric = metric

	return &threatIQCounterMessage
}
