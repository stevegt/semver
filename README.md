# Semver

Semver is a simple, lightweight library for managing and comparing
Semantic Versioning (SemVer) strings in Golang. It provides an
easy-to-use API for parsing, comparing, and upgrading semantic version
strings.

## Features

- Parse SemVer strings into structured Version objects
- Compare two SemVer strings, determining their relative order
- Check for major, minor, and patch-level upgrades between two SemVer strings
- Convert Version objects to JSON and strings

## Usage

To use Semver in your project, first import the library:

```go
import "github.com/stevegt/semver"
```

### Parsing

Parse a SemVer string into a `Version` object:

```go
v, err := semver.Parse([]byte("v1.2.3"))
if err != nil {
    log.Fatal(err)
}
```

### Comparing

Compare two `Version` objects:

```go
cmp := semver.Cmp(version1, version2)
// cmp == 0 if version1 == version2
// cmp == 1 if version1 > version2
// cmp == -1 if version1 < version2
```

### Upgrading

Check if a version is an upgrade (major, minor, or patch) from another version:

```go
major, minor, patch := semver.Upgrade(version1, version2)
// major: true if v2 is a major upgrade from V1
// minor: true if v2 is a minor upgrade from V1
// patch: true if v2 is a patch upgrade from V1
```

### Converting

Convert a `Version` object to a JSON string:

```go
jsonString := version.ToJSON()
```

Convert a `Version` object to a string:

```go
versionString := version.String()
```

