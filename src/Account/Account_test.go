package Account

import (
	"testing"
)

func TestGetB64Hash(t *testing.T) {
	VaildHash := "r/FgKFuU0VX7JLueIQItPqgW4M1gcjcrIf4ninYUMNT2ct2LuNDwbfwmfdhqSkP/vQ5sEzIVHWXsYLif59KeYA=="
	Got := GetB64Hash("mailmix")

	if Got != VaildHash {
		t.Errorf("TestHash > invaild hash")
	}
}
