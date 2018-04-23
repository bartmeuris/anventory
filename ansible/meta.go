package ansible

// Meta special "group"
type Meta struct {
	Hostvars map[string](map[string]VarValue) `json:"hostvars"`
}

// IsGroup indicates if this instance is a group or not. For ansible.Meta types, this always returns false
func (g Meta) IsGroup() bool {
	return false
}

// GetHost returns the variablename/value map for a specific host in the special "_meta" group
func (g *Meta) GetHost(name string) (map[string]VarValue) {
	if _, ok := g.Hostvars[name]; !ok {
		// Create map if doesn't exist yet
		g.Hostvars[name] = make(map[string]VarValue)
	}
	return g.Hostvars[name]
}

