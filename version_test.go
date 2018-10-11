package version

import (
	"testing"
)

func TestVersion(t *testing.T) {
	expected := "v1.1.1"
	v, err := Parse([]byte(expected))
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	result := v.String()
	if expected != result {
		t.Errorf("expected %q, got %q", expected, result)
	}

	expected = "v1.1"
	v, err = Parse([]byte(expected))
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	result = v.String()
	if expected != result {
		t.Errorf("expected %q, got %q", expected, result)
	}

	v, err = Parse([]byte("vA1.2.3"))
	if err == nil {
		t.Errorf("expected an error, returns %s", v.ToJSON())
	}

}
