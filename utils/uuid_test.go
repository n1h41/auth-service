package utils

import "testing"

func TestGenerateUUID(t *testing.T) {
	uuid := GenerateUUID()
	if len(uuid) != 36 {
		t.Errorf("GenerateUUID() = %v, want %v", uuid, 36)
	}
  println(uuid)
}
