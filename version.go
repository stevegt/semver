package version

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// Err holds Version's error messages
type Err struct {
	Msg string
}

func (err *Err) Error() string {
	return err.Msg
}

// Version holds the information to generate a semver string
type Version struct {
	// Major version number (required, must be an integer as string)
	Major string `json:"major"`
	// Minor version number (required, must be an integer as string)
	Minor string `json:"minor"`
	// Patch level (optional, must be an integer as string)
	Patch string `json:"patch,omitempty"`
	// Suffix string, (optional, any string)
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
func Parse(src []byte) (*Version, error) {
	var (
		i   int
		err error
	)
	v := new(Version)
	if bytes.HasPrefix(src, []byte("v")) {
		src = bytes.TrimPrefix(src, []byte("v"))
	}
	parts := strings.Split(string(src), ".")
	if len(parts) > 0 {
		i, err = strconv.Atoi(parts[0])
		if err != nil {
			return nil, &Err{Msg: "Major value must be an integer"}
		}
		v.Major = strconv.Itoa(i)
	} else {
		return nil, &Err{Msg: "Invalid version, expecting semver string"}
	}
	if len(parts) > 1 {
		i, err = strconv.Atoi(parts[1])
		if err != nil {
			return nil, &Err{Msg: "Minor value must be an integer"}
		}
		v.Minor = strconv.Itoa(i)
	} else {
		return nil, &Err{Msg: "Invalid version, expecting semver string"}
	}
	if len(parts) > 2 {
		i, err = strconv.Atoi(parts[2])
		if err != nil {
			return nil, &Err{Msg: "Patch value must be an integer"}
		}
		v.Patch = strconv.Itoa(i)
	}
	if len(parts) > 3 {
		v.Suffix = parts[3]
	}
	return v, nil
}
