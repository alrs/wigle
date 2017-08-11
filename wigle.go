package wigle

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	MAC_POS int = iota
	SSID_POS
	AUTHMODE_POS
	FIRSTSEEN_POS
	CHANNEL_POS
	RSSI_POS
	CURRENTLATITUDE_POS
	CURRENTLONGITUDE_POS
	ALTITUDEMETERS_POS
	ACCURACYMETERS_POS
	TYPE_POS
	schema = "MAC,SSID,AuthMode,FirstSeen,Channel,RSSI,CurrentLatitude,CurrentLongitude,AltitudeMeters,AccuracyMeters,Type"
)

var (
	badMode  = errors.New("unrecognized mode string")
	badEntry = errors.New("unrecognized entry string")
)

type Network struct {
	MAC          net.HardwareAddr
	SSID         string
	Modes        []string
	FirstSeen    time.Time
	Channel      int
	RSSI         int
	Latitude     float64
	Longitude    float64
	Altitude     int
	Accuracy     int
	PhysicalType string
}

func ParseEntry(entry string) (network Network, err error) {
	schemaSplit := strings.Split(schema, ",")
	elements := strings.Split(entry, ",")
	if len(elements) != len(schemaSplit) {
		return network, badEntry
	}
	fmt.Printf("parsing %s\n", elements[MAC_POS])
	network.MAC, err = net.ParseMAC(elements[MAC_POS])
	if err != nil {
		return
	}
	network.SSID = elements[SSID_POS]
	network.Modes, err = parseModes(elements[AUTHMODE_POS])
	if err != nil {
		return
	}
	network.FirstSeen, err = time.Parse("2006-01-02 15:04:05", elements[FIRSTSEEN_POS])
	if err != nil {
		return
	}
	network.Channel, err = strconv.Atoi(elements[CHANNEL_POS])
	if err != nil {
		return
	}
	network.RSSI, err = strconv.Atoi(elements[RSSI_POS])
	if err != nil {
		return
	}
	network.Latitude, err = strconv.ParseFloat(elements[CURRENTLATITUDE_POS], 64)
	if err != nil {
		return
	}
	network.Longitude, err = strconv.ParseFloat(elements[CURRENTLONGITUDE_POS], 64)
	if err != nil {
		return
	}
	network.Altitude, err = strconv.Atoi(elements[ALTITUDEMETERS_POS])
	if err != nil {
		return
	}
	network.Accuracy, err = strconv.Atoi(elements[ACCURACYMETERS_POS])
	if err != nil {
		return
	}
	network.PhysicalType = elements[TYPE_POS]

	return
}

func parseModes(modes string) ([]string, error) {
	leadSep := "["
	followSep := "]"
	if !(strings.Contains(modes, leadSep) && strings.Contains(modes, followSep)) {
		return []string{}, badMode
	}
	followStripped := strings.Replace(modes, followSep, "", -1)
	leadSplit := strings.Split(followStripped, leadSep)
	return leadSplit[1:], nil
}
