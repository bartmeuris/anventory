package ansible

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestInventoryJSONMarshal(t *testing.T) {
	i := NewInventory()
	expect := `{"_meta":{"hostvars":{}},"all":{"children":["ungrouped"]},"ungrouped":{}}`
	if b, err := json.Marshal(&i); err != nil {
		t.Errorf("Could not marshal inventory")
	} else if expect != string(b) {
		fmt.Printf("Expected '%s' != marshalled '%s", expect, string(b))
	}
}
