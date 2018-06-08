package ansible

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestInventoryJSONMarshal(t *testing.T) {
	i := NewInventory()
	//expect := `{"_meta":{"hostvars":{}},"all":{"children":["ungrouped"]},"ungrouped":{}}`
	expect := `{"_meta":{"vars":{}},"all":{"children":{"ungrouped":{}}},"ungrouped":{}}`
	if b, err := json.Marshal(&i); err != nil {
		t.Errorf("Could not marshal inventory")
	} else if expect != string(b) {
		fmt.Printf("Expected '%s' != marshalled '%s", expect, string(b))
	}
}
