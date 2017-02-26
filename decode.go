package elkm1grpc

import (
	"strconv"
)

const elkZoneCount = 8
const elkEnumConst = 48 - 1 // because the real enums start at 1

func decodeArmingStatusReport(text string) (a *ArmingStatusReport) {
	a = new(ArmingStatusReport)
	a.Areas = make([]*ArmingStatusArea, elkZoneCount)
	for i := 0; i < elkZoneCount; i++ {
		a.Areas[i] = new(ArmingStatusArea)
		a.Areas[i].Area = int32(i + 1)
		a.Areas[i].ArmingStatus = ArmingStatusArea_ArmingStatus(int(text[4 + i]) - elkEnumConst)
		a.Areas[i].ArmUpState = ArmingStatusArea_ArmUpState(int(text[12 + i]) - elkEnumConst)
		a.Areas[i].AlarmState = ArmingStatusArea_AlarmState(int(text[20 + i]) - elkEnumConst)
	}
	return
}

func decodeZoneUpdate(text string) (a *Zone) {
	a = new(Zone)
	zone, _ := strconv.Atoi(text[4:7])
	a.Zone = int32(zone)
	status, sub := decodeZoneStatus(int(text[7]))
	a.Status = Zone_ZoneStatus(status)
	a.SubStatus = Zone_ZoneSubStatus(sub)
	return
}

func decodeZoneStatus(s int) (status, sub int) {
	status = s & 0x0F >> 2
	sub = s & 0x03
	return
}
