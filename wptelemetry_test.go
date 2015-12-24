package wptelemetry

import (
	"strings"
	"testing"
)

// Winplus Telemetry Packet

const WPTelemetryPacket1 = `04 3E 2A 02 01 00 00 3A 08 01 3E 31 5C 1E 02 01 04 1A FF 23 26 5C 30 E2 C5 6D B5 DF FB 48 D2 B0 60 D0 F5 A7 10 96 E0 00 01 00 0A BA D5`

const WPTelemetryPacket2 = `04 3E 2A 02 01 00 00 3A 08 01 3E 31 5C 1E 02 01 04 1A FF 23 26 54 30 E2 C5 6D B5 DF FB 48 D2 B0 60 D0 F5 A7 10 96 E0 00 01 00 0A BA D8`

// invalid packet
const Packet3 = `04 3E 2A 02 01 00 01 AB 0A D3 87 1C DF 1E 02 01 06 1A FF 4C 00 02 16 B9 40 7F 30 F5 F8 46 6E AF F9 25 55 6B 57 FE 6D 0A AB 87 D3 AF AA`

const Packet4 = `04 3E 2A 02 01 00 01 AB 0A D3 87 1C DF 1E 02 01 06 1A FF 4C 00 03 32 B9 40 7F 30 F5 F8 46 6E AF F9 25 55 6B 57 FE 6D 0A AB 87 D3 AF AA`


func TestParseBattery(t *testing.T) {
	battery := parseBattery(strings.Split(WPTelemetryPacket1, " "))
	expected_battery := 92
	if battery != expected_battery {
		t.Errorf("Parsing battery {%v} failed for packet1: %v", expected_battery, battery)
	}

	battery = parseBattery(strings.Split(WPTelemetryPacket2, " "))
	expected_battery = 84
	if battery != expected_battery {
		t.Errorf("Parsing battery {%v} failed for packet2: %v", expected_battery, battery)
	}
}

func TestIsValid(t *testing.T) {
	if !IsValid(WPTelemetryPacket1) {
		t.Errorf("Validation failed for packet1")
	}

	if !IsValid(WPTelemetryPacket2) {
		t.Errorf("Validation failed for packet2")
	}

	if IsValid(Packet3) {
		t.Errorf("Validation passed for packet3")
	}

	if IsValid(Packet4) {
		t.Errorf("Validation passed for packet4")
	}
}

func TestToString(t *testing.T) {
	wptp := NewWPTPacket(WPTelemetryPacket1)
	packet1String := "INT  UUID E2C56DB5-DFFB-48D2-B060-D0F5A71096E0 MAJOR 1 MINOR 10 RSSI -43 BATTERY 92 BUTTON false"
	if packet1String != wptp.ToString() {
		t.Errorf("ToString failed for packet1 \nexpected: %v\ngot: %v", packet1String, wptp.ToString())
	}

	wptp = NewWPTPacket(WPTelemetryPacket2)
	packet2String := "INT  UUID E2C56DB5-DFFB-48D2-B060-D0F5A71096E0 MAJOR 1 MINOR 10 RSSI -40 BATTERY 84 BUTTON false"
	if packet2String != wptp.ToString() {
		t.Errorf("ToString failed for packet2 \nexpected: %v\ngot: %v", packet2String, wptp.ToString())
	}
}

func TestMapKey(t *testing.T) {
	wptp := NewWPTPacket(WPTelemetryPacket1)
	packet1MapKey := "IBE_0000100010"
	if packet1MapKey != wptp.MapKey() {
		t.Errorf("MapKey failed for packet1 \nexpected: %v\ngot: %v", packet1MapKey, wptp.MapKey())
	}

	wptp = NewWPTPacket(WPTelemetryPacket2)
	if packet1MapKey != wptp.MapKey() {
		t.Errorf("MapKey failed for packet2 \nexpected: %v\ngot: %v", packet1MapKey, wptp.MapKey())
	}
}
