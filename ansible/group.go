package ansible

import (
	"sync"
)

// Group represents a group in Ansible
type Group struct {
	Hosts []string `json:"hosts,omitempty"`
	Children []string `json:"children,omitempty"`
	Vars map[string]VarValue `json:"vars,omitempty"`
	m *sync.Mutex
}

// IsGroup indicates if this instance is a group or not. For ansible.Group types, this always returns true
func (g *Group) IsGroup() bool {
	return true
}

// AddHost adds a host to this group
func (g *Group) AddHost(host string) {
	for _, s := range g.Hosts {
		if s == host {
			return
		}
	}
	// Wrap in mutex to prevent race conditions
	g.m.Lock()
	g.Hosts = append(g.Hosts, host)
	g.m.Unlock()
}

// AddChild Adds a child group to this group
func (g *Group) AddChild(child string) {
	for _, s := range g.Children {
		if s == child {
			return
		}
	}
	// Wrap in mutex to prevent race conditions
	g.m.Lock()
	g.Children = append(g.Children, child)
	g.m.Unlock()
}

// NewGroup creates a new Group instance
func NewGroup() (*Group) {
	ret := &Group{}
	ret.m = &sync.Mutex{}
	ret.Hosts = make([]string, 0)
	ret.Children = make([]string, 0)
	ret.Vars = make(map[string]VarValue, 0)
	return ret
}

