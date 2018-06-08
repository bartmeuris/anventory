package ansible

import (
	"encoding/json"
	"fmt"
	//"encoding/json"
)

// Inventory represents the complete inventory that can be marshaled to JSON and pumped into Ansible
type Inventory map[string]InventoryEntry

// InventoryEntry represents groups or the special "_meta" group in an Ansible inventory
type InventoryEntry interface {
	IsGroup() bool
}

// NewInventory creates a new empty Ansible Inventory instance with special groups already initialized
func NewInventory() Inventory {
	ret := Inventory(make(map[string]InventoryEntry))

	// Add special Meta group:
	meta := Meta{Vars: make(map[string](map[string]VarValue))}
	ret["_meta"] = &meta
	ret.AddGroup("all")
	ret.AddGroup("ungrouped")
	ret.GetGroup("all").AddChild("ungrouped")
	return ret
}

// AddGroup adds a group to the inventory. If the group already exists, it returns an error.
func (i Inventory) AddGroup(name string) (*Group, error) {
	if _, ok := i[name]; ok {
		return nil, fmt.Errorf("Group '%s' already exists", name)
	}
	g := NewGroup()
	i[name] = g
	return g, nil
}

// GetHostVars returns the variable/value map for a certain host from the "_meta" entry.
func (i Inventory) GetHostVars(hostname string) map[string]VarValue {
	// Get the hostvars map from the special _meta group
	if meta, ok := i["_meta"].(*Meta); ok {
		return meta.GetHost(hostname)
	}
	return nil
}

// GetGroup returns the group with the specified name if it exists, otherwise 'nil'
func (i Inventory) GetGroup(name string) *Group {
	if val, ok := i[name]; ok {
		if g, ok := val.(*Group); ok {
			return g
		}
	}
	return nil
}

// MarshalJSON implements a custom marshalling
func (i Inventory) MarshalJSON() ([]byte, error) {
	// make sure all hosts are added to the "all" group.
	allGroup, allOk := i["all"].(*Group)
	if !allOk {
		return nil, fmt.Errorf("Internal inventory error: 'all' entry is not a group")
	}
	for g := range i {
		if g == "all" {
			continue
		} else if i[g].IsGroup() {
			for _, gHost := range i[g].(*Group).Hosts {
				allGroup.AddHost(gHost)
			}
		} else if meta, ok := i[g].(*Meta); ok {
			// Treat Meta group
			for hv := range meta.Vars {
				allGroup.AddHost(hv)
			}
		}
	}
	return json.Marshal((map[string]InventoryEntry)(i))
}
