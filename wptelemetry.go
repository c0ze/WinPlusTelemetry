package wptelemetry

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	ib "github.com/c0ze/iBeagon"
)

type WPTelemetryPacket struct {
	IbPacket *ib.IBeaconPacket
	Battery int
	ButtonPushed bool
}

func NewWPTPacket(line string) *WPTelemetryPacket {
	vals := strings.Split(line, " ")
	return &WPTelemetryPacket{
		IbPacket: ib.NewIBPacket(line),
		Battery: parseBattery(vals),
		ButtonPushed: parseButton(vals)}
}

func IsValid(str string) bool {
	r, err := regexp.Compile(`^04\ 3E\ 2A\ 02\ 01\ .{26}\ 02\ 01\ .{17}\ (30|31)`)
	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return false
	}

	return r.MatchString(str)
}

func (wptp *WPTelemetryPacket) MapKey() string {
	return wptp.IbPacket.MapKey()
}

func parseBattery(vals []string) int {
	battery, _ := strconv.ParseInt(vals[21], 16, 0)
	return int(battery)
}

func parseButton(vals []string) bool {
	return vals[22] == "31"
}

func (wptp *WPTelemetryPacket) ToString() string {
	return fmt.Sprintf("%v BATTERY %v BUTTON %v", wptp.IbPacket.ToString(), wptp.Battery, wptp.ButtonPushed)
}
