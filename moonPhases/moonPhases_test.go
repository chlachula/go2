package moonPhases

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= 1.0e-2
}
func TestMoonAgePhaseNumbers(t *testing.T) {
	mAngleN, mAngle, mAngleP, mAgeN, mAge, mAgeP := moonAgePhaseNumbers(0)
	wAngleN, wAngle, wAngleP, wAgeN, wAge, wAgeP := 353.57, 0.0, 6.43, 29.0, 0.0, 0.53

	//if mAngleN != wAngleN || mAngle != wAngle || mAngleP != wAngleP || mAgeN != wAgeN || mAge != wAge || mAgeP != wAgeP {
	if !almostEqual(mAngleN, wAngleN) {
		t.Errorf("Unexpected mAngleN want=%.2f got=%.2f", mAngleN, wAngleN)
	}
	if !almostEqual(mAngle, wAngle) {
		t.Errorf("Unexpected mAngle want=%.2f got=%.2f", mAngle, wAngle)
	}
	if !almostEqual(mAngleP, wAngleP) {
		t.Errorf("Unexpected mAngleP want=%f got=%f", mAngleP, wAngleP)
	}
	if !almostEqual(mAgeN, wAgeN) {
		t.Errorf("Unexpected mAgeN want=%.2f got=%.2f", mAgeN, wAgeN)
	}
	if !almostEqual(mAge, wAge) {
		t.Errorf("Unexpected mAge want=%.2f got=%.2f", mAge, wAge)
	}
	if !almostEqual(mAgeP, wAgeP) {
		t.Errorf("Unexpected mAgeP want=%.2f got=%.2f", mAgeP, wAgeP)
	}
}
