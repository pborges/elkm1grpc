package elkm1grpc

const elkZoneCount = 8
const elkEnumConst = 48

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