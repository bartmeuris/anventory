package ansible

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestAnsible(t *testing.T) {
	expect := `{"_meta":{"vars":{"localhost":{"ansible_host":"localhost","ansible_user":"bart","has_something":false,"some_number":50}}},"all":{"hosts":{"localhost":{}},"children":{"ungrouped":{}}},"test":{"hosts":{"localhost":{}}},"ungrouped":{}}`
	i := NewInventory()
	g, _ := i.AddGroup("test")
	g.AddHost("localhost")
	hv := i.GetHostVars("localhost")

	hv["ansible_host"] = NewString("localhost")
	hv["ansible_user"] = NewString(os.Getenv("USER"))
	hv["has_something"] = NewBool(false)
	hv["some_number"] = NewInt(50)

	if b, err := json.Marshal(&i); err != nil {
		t.Errorf("Could not marshal inventory: %v", err)
	} else if string(b) != expect {
		fmt.Printf("Expected '%s' != marshalled '%s", expect, string(b))
	}
	//ioutil.WriteFile("../inventory.json", b, 0644)
}
