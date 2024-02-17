package semver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Version holds the information to generate a semver string
type Version struct {
	// Major version number (required)
	Major string `json:"major"`
	// Minor version number (required)
	Minor string `json:"minor"`
	// Patch level (optional)
	Patch string `json:"patch,omitempty"`
	// Suffix string, (optional)
	Suffix string `json:"suffix,omitempty"`
	// Timestamp (optional, a timestamp in form of YYYY-MM-DD HH:MM:SS)
}

func (v *Version) String() string {
	if v.Patch == "" {
		return "v" + v.Major + "." + v.Minor
	}
	if v.Suffix == "" {
		return "v" + v.Major + "." + v.Minor + "." + v.Patch
	}
	return "v" + v.Major + "." + v.Minor + "." + v.Patch + v.Suffix
}

// ToJSON takes a version struct and returns JSON as byte slice
func (v *Version) ToJSON() []byte {
	src, _ := json.Marshal(v)
	return src
}

// Parse takes a byte slice and returns a version struct,
// and an error value.
func Parse(src []byte) (v *Version, err error) {
	v = &Version{}
	if bytes.HasPrefix(src, []byte("v")) {
		src = bytes.TrimPrefix(src, []byte("v"))
	}

	parts := strings.Split(string(src), ".")
	switch len(parts) {
	case 4:
		v.Suffix = parts[3]
		fallthrough
	case 3:
		v.Patch = parts[2]
		fallthrough
	case 2:
		v.Minor = parts[1]
		fallthrough
	case 1:
		v.Major = parts[0]
	default:
		return nil, fmt.Errorf("Invalid version, expecting semver string: %s, parts: %#v", src, parts)
	}
	return v, nil
}

// Cmp takes two version structs and returns an integer value
// indicating the relationship between the two versions.
// 0: v1 == v2
// 1: v1 > v2
// -1: v1 < v2
func Cmp(v1, v2 *Version) (int, error) {
	majorCmp, err := CmpPart(v1.Major, v2.Major)
	if err != nil {
		return 0, err
	}
	minorCmp, err := CmpPart(v1.Minor, v2.Minor)
	if err != nil {
		return 0, err
	}
	patchCmp, err := CmpPart(v1.Patch, v2.Patch)
	if err != nil {
		return 0, err
	}
	suffixCmp, err := CmpPart(v1.Suffix, v2.Suffix)
	if err != nil {
		return 0, err
	}

	intCmp := majorCmp*1000 + minorCmp*100 + patchCmp*10 + suffixCmp
	if intCmp > 0 {
		return 1, nil
	}
	if intCmp < 0 {
		return -1, nil
	}
	return 0, nil
}

// Upgrade takes two version structs and returns three bools
// indicating the relationship between the two versions.
// 1. major: true if v2 is a major upgrade from V1
// 2. minor: true if v2 is a minor upgrade from V1
// 3. patch: true if v2 is a patch upgrade from V1
// 4. suffix: true if v2 is a suffix upgrade from V1
func Upgrade(v1, v2 *Version) (major, minor, patch, suffix bool, err error) {
	majorUp, err := CmpPart(v2.Major, v1.Major)
	if err != nil {
		return false, false, false, false, err
	}
	minorUp, err := CmpPart(v2.Minor, v1.Minor)
	if err != nil {
		return false, false, false, false, err
	}
	patchUp, err := CmpPart(v2.Patch, v1.Patch)
	if err != nil {
		return false, false, false, false, err
	}
	suffixUp, err := CmpPart(v2.Suffix, v1.Suffix)
	if err != nil {
		return false, false, false, false, err
	}
	if majorUp > 0 {
		major = true
		minor = true
		patch = true
		suffix = true
	} else if minorUp > 0 {
		minor = true
		patch = true
		suffix = true
	} else if patchUp > 0 {
		patch = true
		suffix = true
	} else if suffixUp > 0 {
		suffix = true
	}
	return
}

// CmpPart takes two version components (e.g. Major, Minor, Patch,
// Suffix) and returns an integer value indicating the relationship
// between the two components.
// 0: v1 == v2
// 1: v1 > v2
// -1: v1 < v2
func CmpPart(v1, v2 string) (cmp int, err error) {
	v1Int, v1Str, err := splitParts(v1)
	if err != nil {
		return 0, err
	}
	v2Int, v2Str, err := splitParts(v2)
	if err != nil {
		return 0, err
	}
	if v1Int > v2Int {
		return 1, nil
	}
	if v1Int < v2Int {
		return -1, nil
	}
	if v1Str > v2Str {
		return 1, nil
	}
	if v1Str < v2Str {
		return -1, nil
	}
	return 0, nil
}

// splitParts takes a version string and returns an integer and a
// string representing the leading integer and the remaining
// characters in the string.
func splitParts(v string) (intPart int, strPart string, err error) {
	// split the string into two parts using a regex
	pat := `^(\d*)(.*?)$`
	re := regexp.MustCompile(pat)
	parts := re.FindStringSubmatch(v)
	if len(parts) != 3 {
		return 0, "", fmt.Errorf("Invalid version part: %s", v)
	}
	if len(parts[1]) == 0 {
		intPart = 0
	} else {
		intPart, err = strconv.Atoi(parts[1])
		if err != nil {
			return 0, "", fmt.Errorf("Invalid integer part: %s", v)
		}
	}
	strPart = parts[2]
	return
}
