package datetool

import "testing"

func TestDate2Time(t *testing.T) {
	t.Run("test date2time", func(t *testing.T) {
		Date2Time()
	})
	t.Run("test time2Date", func(t *testing.T) {
		Time2Date()
	})
}