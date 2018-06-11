package ansible

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
	"flag"
	"path/filepath"
	"bytes"
)

// see https://medium.com/@povilasve/go-advanced-tips-tricks-a872503ac859 for testing tips/tricks
var update = flag.Bool("update", false, "update .golden files")

func TestAnsible(t *testing.T) {
	//expect := `{"_meta":{"vars":{"localhost":{"ansible_host":"localhost","ansible_user":"bart","has_something":false,"some_number":50}}},"all":{"hosts":{"localhost":{}},"children":{"ungrouped":{}}},"test":{"hosts":{"localhost":{}}},"ungrouped":{}}`
	i := NewInventory()
	g, _ := i.AddGroup("test")
	g.AddHost("localhost")
	hv := i.GetHostVars("localhost")

	hv["ansible_host"] = NewString("localhost")
	//hv["ansible_user"] = NewString(os.Getenv("USER"))
	hv["ansible_user"] = NewString("ansible")
	hv["has_something"] = NewBool(false)
	hv["some_number"] = NewInt(50)
	if b, err := json.Marshal(&i); err != nil {
		t.Errorf("Could not marshal inventory: %v", err)
	} else {
		golden := filepath.Join("testdata", t.Name() + ".golden.json")
		if *update {
			ioutil.WriteFile(golden, b, 0644)
		}
		expected, _ := ioutil.ReadFile(golden)
		
		if !bytes.Equal(b, expected) {
			fmt.Printf("Expected '%s' != marshalled '%s", string(expected), string(b))
		}
	}
}
