package semver

import (
	"testing"

	. "github.com/stevegt/goadapt"
)

func TestVersion(t *testing.T) {
	expected := "v1.2.3"
	v, err := Parse([]byte(expected))
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	result := v.String()
	if expected != result {
		t.Errorf("expected %q, got %q", expected, result)
	}

	expected = "v1.2"
	v, err = Parse([]byte(expected))
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	result = v.String()
	if expected != result {
		t.Errorf("expected %q, got %q", expected, result)
	}

	// try parsing some version variations
	v, err = Parse([]byte("vA1.2.3"))
	Tassert(t, err == nil, "expected no error, got %s", err)

	expected = "v1.2.3"
	v, err = Parse([]byte("1.2.3"))
	if err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
	result = v.String()
	if expected != result {
		t.Errorf("expected %q, got %q", expected, result)
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
	cmp, err := Cmp(v1, v2)
	Tassert(t, err == nil, "Cmp error: %s", err)
	Tassert(t, cmp == -1, "expected -1, got %d", cmp)
	cmp, err = Cmp(v2, v1)
	Tassert(t, err == nil, "Cmp error: %s", err)
	Tassert(t, cmp == 1, "expected 1, got %d", cmp)
	cmp, err = Cmp(v1, v1)
	Tassert(t, err == nil, "Cmp error: %s", err)
	Tassert(t, cmp == 0, "expected 0, got %d", cmp)

	// try patch version with string
	v3, err := Parse([]byte("v1.2.3-alpha"))
	Tassert(t, err == nil, "Parse error: %s", err)
	v4, err := Parse([]byte("v1.2.3-beta"))
	Tassert(t, err == nil, "Parse error: %s", err)
	cmp, err = Cmp(v3, v4)
	Tassert(t, err == nil, "Cmp error: %s", err)
	Tassert(t, cmp == -1, "expected -1, got %d", cmp)

	// try suffix version as string
	v5, err := Parse([]byte("v1.2.3.alpha"))
	Tassert(t, err == nil, "Parse error: %s", err)
	v6, err := Parse([]byte("v1.2.3.beta"))
	Tassert(t, err == nil, "Parse error: %s", err)
	cmp, err = Cmp(v5, v6)
	Tassert(t, err == nil, "Cmp error: %s", err)
	Tassert(t, cmp == -1, "expected -1, got %d", cmp)

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
	major, minor, patch, suffix, err := Upgrade(v1, v2)
	Tassert(t, err == nil, "Upgrade error: %s", err)
	if major {
		t.Errorf("expected false, got %t", major)
	}
	if !minor {
		t.Errorf("expected true, got %t", minor)
	}
	if !patch {
		t.Errorf("expected true, got %t", patch)
	}
	if !suffix {
		t.Errorf("expected true, got %t", suffix)
	}
}
