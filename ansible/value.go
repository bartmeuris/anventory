package ansible

import (
    "fmt"
    "encoding/json"
)

type ansibleType int

const (
    aInt ansibleType = iota
    aFloat
    aBool
    aString
)

// VarValue represents the value of an Ansible variable
// This differentiates between string, int, float and bool, in order to parse into the correct JSON representation
type VarValue struct {
	atype ansibleType
	ival int64
	fval float64
	bval bool
	sval string
}

// SetString sets the string value of this variable object
func (a *VarValue) SetString(s string) {
	a.atype = aString
	a.sval = s
}

// NewString creates a new variable value object that represents the provided string value
func NewString(s string) VarValue {
	ret := VarValue{}
	ret.SetString(s)
	return ret
}

// SetInt sets the integer value of this variable object
func (a *VarValue) SetInt(i int64) {
	a.atype = aInt
	a.ival = i
}
// NewInt creates a new variable value object that represents the provided integer value
func NewInt(i int64) VarValue {
	ret := VarValue{}
	ret.SetInt(i)
	return ret
}

// SetFloat sets the float value of this variable object
func (a *VarValue) SetFloat(f float64) {
	a.atype = aFloat
	a.fval = f
}

// NewFloat creates a new variable value object that represents the provided floating point value
func NewFloat(f float64) VarValue {
	ret := VarValue{}
	ret.SetFloat(f)
	return ret
}

// SetBool sets the boolean value of this variable object
func (a *VarValue) SetBool(b bool) {
	a.atype = aBool
	a.bval = b
}

// NewBool creates a new variable value object that represents the provided boolean value
func NewBool(b bool) VarValue {
	ret := VarValue{}
	ret.SetBool(b)
	return ret
}

// String creates the string representation of the object's current value.
func (a *VarValue) String() string {
	switch a.atype {
	case aInt:
		return fmt.Sprintf("%d", a.ival)
	case aFloat:
		return fmt.Sprintf("%f", a.fval)
	case aBool:
		return fmt.Sprintf("%t", a.bval)
	case aString:
		return a.sval
	}
	return ""
}

// MarshalJSON marshals the value of the object according to it's type
func (a *VarValue) MarshalJSON() ([]byte, error) {
	switch a.atype {
	case aInt:
		return json.Marshal(a.ival)
	case aFloat:
		return json.Marshal(a.fval)
	case aBool:
		return json.Marshal(a.bval)
	case aString:
		return json.Marshal(a.sval)
	}
	return nil, fmt.Errorf("Could not marshal unknown AnsibleType %d", a.atype)
}

