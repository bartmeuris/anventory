package ansible

import (
	"encoding/json"
	"math"
	"testing"
)

////////////
// Boolean testing

func TestJSONMarshalBoolFalse(t *testing.T) {
	testVarVal(t, NewBool(false), "false")
}

func TestJSONMarshalBoolTrue(t *testing.T) {
	testVarVal(t, NewBool(true), "true")
}

////////////
// Floating point testing

func TestJSONMarshalFloat1dot1(t *testing.T) {
	testVarVal(t, NewFloat(1.1), "1.1")
}

func TestJSONMarshalFloatMax(t *testing.T) {
	testVarVal(t, NewFloat(math.MaxFloat64), "1.7976931348623157e+308")
}

func TestJSONMarshalFloatMin(t *testing.T) {
	testVarVal(t, NewFloat(math.SmallestNonzeroFloat64), "5e-324")
}

func TestJSONMarshalFloatNaN(t *testing.T) {
	f := NewFloat(math.Log(-1.0))
	if _, err := json.Marshal(&f); err == nil {
		t.Errorf("JSON should not be able to parse %s", f.String())
	}
}

func TestJSONMarshalFloatNegInf(t *testing.T) {
	f := NewFloat(math.Inf(-1))
	if _, err := json.Marshal(&f); err == nil {
		t.Errorf("JSON should not be able to parse %s", f.String())
	}
}
func TestJSONMarshalFloatPosInf(t *testing.T) {
	f := NewFloat(math.Inf(1))
	if _, err := json.Marshal(&f); err == nil {
		t.Errorf("JSON should not be able to parse %s", f.String())
	}
}

////////////
// String testing
func TestJSONMarshalStringEmpty(t *testing.T) {
	testVarVal(t, NewString(""), "\"\"")
}
func TestJSONMarshalStringDoubleQuote(t *testing.T) {
	testVarVal(t, NewString("\""), "\"\\\"\"")
}
func TestJSONMarshalStringSingleQuote(t *testing.T) {
	testVarVal(t, NewString("'"), "\"'\"")
}

func TestJSONMarshalString1(t *testing.T) {
	testVarVal(t, NewString("abcde"), "\"abcde\"")
}

////////////
// Integer testing

func TestJSONMarshalInt1(t *testing.T) {
	testVarVal(t, NewInt(1), "1")
}

func TestJSONMarshalIntMin1(t *testing.T) {
	testVarVal(t, NewInt(-1), "-1")
}

func TestJSONMarshalIntMax(t *testing.T) {
	testVarVal(t, NewInt(math.MaxInt64), "9223372036854775807")
}

func TestJSONMarshalIntMin(t *testing.T) {
	testVarVal(t, NewInt(math.MinInt64), "-9223372036854775808")
}

func testVarVal(t *testing.T, v VarValue, expect string) {
	if b, err := json.Marshal(&v); err != nil {
		t.Errorf("Error JSON marshaling %s", v.String())
	} else if string(b) != expect {
		t.Errorf("Marshaled json ('%s') doesn't match expectation ('%s')", string(b), expect)
	}
}
