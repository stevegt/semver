package semver

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

func TestVersionCompare(t *testing.T) {
	v1, err := Parse([]byte("v1.1.1"))
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	v2, err := Parse([]byte("v1.2.0"))
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	if Cmp(v1, v2) != -1 {
		t.Errorf("expected -1, got %d", Cmp(v1, v2))
	}
	if Cmp(v2, v1) != 1 {
		t.Errorf("expected 1, got %d", Cmp(v2, v1))
	}
	if Cmp(v1, v1) != 0 {
		t.Errorf("expected 0, got %d", Cmp(v1, v1))
	}
}

func TestVersionUpgrade(t *testing.T) {
	v1, err := Parse([]byte("v1.1.1"))
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	v2, err := Parse([]byte("v1.2.0"))
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	major, minor, patch := Upgrade(v1, v2)
	if major {
		t.Errorf("expected false, got %t", major)
	}
	if !minor {
		t.Errorf("expected true, got %t", minor)
	}
	if !patch {
		t.Errorf("expected true, got %t", patch)
	}
}
