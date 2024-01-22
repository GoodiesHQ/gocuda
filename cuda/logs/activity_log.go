package logs

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Implements LogEntry
type ActivityLogEntry struct {
	Timestamp   time.Time // log received
	Level       string    // severity
	Action      string    // action taken
	Type        string    // activity type
	IPProtocol  string    // tcp, udp, icmp, etc...
	SrcIF       string    // source interface
	SrcIP       string    // source IP
	SrcNAT      string    // translated source IP
	SrcPort     uint16    // source port
	SrcMAC      string    // source MAC
	DstIF       string    // destination interface
	DstIP       string    // destination IP
	DstNAT      string    // translated destination IP
	DstPort     uint16    // destination port
	Service     string    // service (friendly name)
	RuleName    string    // firewall rule name
	Info        string    // unknown...
	Duration    uint64    // connection duration
	Count       uint64    // packet count in this session
	RXBytes     uint64    // received bytes
	TXBytes     uint64    // sent bytes
	RXPackets   uint64    // received packets
	TXPackets   uint64    // sent packets
	User        string    // user of an authenticated session if exists
	Protocol    string    // not sure, but I know it differs from the IPProtocol
	Application string    // traffic application name if detected
	Target      string    // ... possibly IDS related?
	Content     string    // content type if detected
	URLCategory string    // url category if detected
}

var fields = [...]string{
	"Timestamp", "Level", "Action", "Type", "IP Protocol",
	"Source Iface", "Source IP", "Source NAT", "Source Port", "Source MAC",
	"Destination Iface", "Destination IP", "Destination NAT", "Destination Port",
	"Service", "Rule Name", "Info", "Duration", "Count",
	"RX Bytes", "TX Bytes", "RX Packets", "TX Packets", "User", "Protocol",
	"Application", "Target", "Content", "URL Category"}

// timestamp format: "YYYY MM DD hh:mm:ss OFFSET"
const timestampFmt = "2006 01 02 15:04:05 -07:00"

const pattern = `` +
	`(\d{4}\s\d{2}\s\d{2}\s\d{2}:\d{2}:\d{2}\s\S+)` + // timestamp
	`\s+` +
	`(\w+)` + // level
	`\s+` +
	`(\w+):` + // action
	`\s+` +
	`(.*)\s*$` // piped values
var re = regexp.MustCompile(pattern)

func (log *ActivityLogEntry) CSV() string {
	return strings.Join(
		[]string{
			log.Timestamp.String(), log.Level, log.Action, log.Type, log.IPProtocol,
			log.SrcIF, log.SrcIP, log.SrcNAT, fmt.Sprintf("%d", log.SrcPort), log.SrcMAC,
			log.DstIF, log.DstIP, log.DstNAT, fmt.Sprintf("%d", log.DstPort), log.Service,
			log.RuleName, log.Info, fmt.Sprintf("%d", log.Duration), fmt.Sprintf("%d", log.Count),
			fmt.Sprintf("%d", log.RXBytes), fmt.Sprintf("%d", log.TXBytes),
			fmt.Sprintf("%d", log.RXPackets), fmt.Sprintf("%d", log.TXPackets),
			log.User, log.Protocol, log.Application, log.Target, log.Content, log.URLCategory,
		},
		",",
	)
}

func (log *ActivityLogEntry) String() string {
	return ""
}

type ActivityLogParser struct{}

func (log *ActivityLogParser) Fields() []string {
	return fields[:]
}

func (log *ActivityLogParser) FieldsCSV() string {
	return strings.Join(log.Fields(), ",")
}

func (parser *ActivityLogParser) ParseLine(line string) (LogEntry, *ParseFailure) {
	// regex for the full log entry
	line = strings.TrimSpace(line)

	matches := re.FindStringSubmatch(line)
	if matches == nil || len(matches) != 5 {
		return nil, NewParseFailureMessage(line, "unexpected number of groups in the line")
	}

	timestamp, err := time.Parse(timestampFmt, matches[1])
	if err != nil {
		return nil, NewParseFailureMessage(line, "failed to process the timestamp")
	}

	values := strings.Split(matches[4], "|")
	if len(values) != 26 {
		return nil, NewParseFailureMessage(line, "unexpected number of piped values")
	}

	srcPort, err := strconv.ParseUint(values[4], 10, 16)
	if err != nil {
		return nil, NewParseFailureMessage(line, "failed to parse the source port")
	}

	dstPort, err := strconv.ParseUint(values[7], 10, 16)
	if err != nil {
		return nil, NewParseFailureMessage(line, "failed to parse destination port")
	}

	duration, err := strconv.ParseUint(values[14], 10, 64)
	if err != nil {
		return nil, NewParseFailureMessage(line, "failed to parse duration")
	}

	count, err := strconv.ParseUint(values[15], 10, 64)
	if err != nil {
		return nil, NewParseFailureMessage(line, "failed to parse count")
	}

	rxBytes, err := strconv.ParseUint(values[16], 10, 64)
	if err != nil {
		return nil, NewParseFailureMessage(line, "failed to parse received/rx byte count")
	}

	txBytes, err := strconv.ParseUint(values[17], 10, 64)
	if err != nil {
		return nil, NewParseFailureMessage(line, "failed to parse transmitted/tx byte count")
	}

	rxPackets, err := strconv.ParseUint(values[18], 10, 64)
	if err != nil {
		return nil, NewParseFailureMessage(line, "failed to parse received/rx packet count")
	}

	txPackets, err := strconv.ParseUint(values[19], 10, 64)
	if err != nil {
		return nil, NewParseFailureMessage(line, "failed to parse transmitted/tx packet count")
	}

	return &ActivityLogEntry{
		Timestamp:   timestamp,
		Level:       matches[2],
		Action:      matches[3],
		Type:        values[0],
		IPProtocol:  values[1],
		SrcIF:       values[2],
		SrcIP:       values[3],
		SrcNAT:      values[12],
		SrcPort:     uint16(srcPort),
		SrcMAC:      values[5],
		DstIF:       values[9],
		DstIP:       values[6],
		DstNAT:      values[13],
		DstPort:     uint16(dstPort),
		Service:     values[8],
		RuleName:    values[10],
		Info:        values[11],
		Duration:    duration,
		Count:       count,
		RXBytes:     rxBytes,
		TXBytes:     txBytes,
		RXPackets:   rxPackets,
		TXPackets:   txPackets,
		User:        values[20],
		Protocol:    values[21],
		Application: values[22],
		Target:      values[23],
		Content:     values[24],
		URLCategory: values[25],
	}, nil
}
