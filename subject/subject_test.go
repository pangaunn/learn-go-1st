package subject

import "testing"

func TestGetSubject(t *testing.T) {
	expected := "Subject"
	if GetSubject() != expected {
		t.Error("not match ja")
	}
}
