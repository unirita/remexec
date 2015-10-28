package console

import (
	"testing"
)

func init() {
	msgs["TEST001E"] = "test message."
	msgs["TEST002E"] = "test param1[%s] param2[%d]."
}

func TestGetMessage_NoParameter(t *testing.T) {
	m := GetMessage("TEST001E")
	if m != "TEST001E test message." {
		t.Errorf("GetMessage(TEST001E) => %s, wants %s", m, "TEST001E test message.")
	}
}

func TestGetMessage_WithParameters(t *testing.T) {
	m := GetMessage("TEST002E", "abc", 1234)
	if m != "TEST002E test param1[abc] param2[1234]." {
		t.Errorf("GetMessage(TEST002E) => %s, wants %s", m, "TEST002E test param1[abc] param2[1234].")
	}
}
