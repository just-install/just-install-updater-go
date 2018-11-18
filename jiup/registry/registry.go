package registry

import (
	"bytes"
	"encoding/json"
	"errors"
)

// Registry represents a registry.
type Registry struct {
	Schema   string             `json:"$schema"`
	Version  int                `json:"version"`
	Packages map[string]Package `json:"packages"`
}

// Package represents a package
type Package struct {
	Installer struct {
		Interactive *bool         `json:"interactive,omitempty"` // optional: default to false
		Kind        InstallerKind `json:"kind"`
		// Options will either have: (base options set) or (x86 set)
		// or (x86 and x86_64 set). Check the x86 for nil to
		// determine which one.
		Options *struct {
			*Options `json:",omitempty"` // maybe optional
			X86      *Options            `json:"x86,omitempty"`    // maybe optional
			X86_64   *Options            `json:"x86_64,omitempty"` // optional
		} `json:"options,omitempty"` // optional
		X86    *string `json:"x86,omitempty"` // optional, but at least either x86 or x86_64 must be defined
		X86_64 *string `json:"x86_64,omitempty"`
	} `json:"installer"`
	Version string `json:"version"`
}

// InstallerKind represents a type of installer.
type InstallerKind string

// Installer types.
const (
	InstallerKindAdvancedInstaller InstallerKind = "advancedinstaller"
	InstallerKindAsIs                            = "as-is"
	InstallerKindCopy                            = "copy"
	InstallerKindCustom                          = "custom"
	InstallerKindEasyInstall26                   = "easy_install_26"
	InstallerKindEasyInstall27                   = "easy_install_27"
	InstallerKindInnoSetup                       = "innosetup"
	InstallerKindMSI                             = "msi"
	InstallerKindNSIS                            = "nsis"
	InstallerKindZip                             = "zip"
)

// Options represents additional options for a package. All fields are optional.
type Options struct {
	Arguments *[]string `json:"arguments,omitempty"` // optional
	Container *struct {
		Installer     string        `json:"installer"`
		ContainerKind ContainerKind `json:"kind"`
	} `json:"container,omitempty"` // optional
	Destination *string   `json:"destination,omitempty"` // optional
	Extension   *string   `json:"extension,omitempty"`   // optional
	FileName    *string   `json:"filename,omitempty"`    // optional
	Shims       *[]string `json:"shims,omitempty"`       // optional
}

// ContainerKind represents a type of container.
type ContainerKind string

// Container types.
const (
	ContainerKindZip ContainerKind = "zip"
)

// ErrUnsupportedRegistry is returned is the registry version is unsupported.
var ErrUnsupportedRegistry = errors.New("unsupported registry version")

// RegistryVersion is the currently supported registry version.
const RegistryVersion = 4

// New returns a new Registry.
func New() *Registry {
	return &Registry{
		Schema:  "./just-install-schema.json",
		Version: RegistryVersion,
	}
}

// NewFromJSON loads a Registry from a JSON byte array.
func NewFromJSON(jsonBuf []byte) (*Registry, error) {
	r := &Registry{}
	err := json.Unmarshal(jsonBuf, &r)
	if err != nil {
		return nil, err
	}
	if r.Version != RegistryVersion {
		return nil, ErrUnsupportedRegistry
	}
	return r, nil
}

// GetJSON gets the JSON for the Registry.
func (r *Registry) GetJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(r)
	return buffer.Bytes(), err
}
