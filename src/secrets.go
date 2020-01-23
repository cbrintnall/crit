package main

import "fmt"

const GcpSecret string = "gcp"

// SecretFile is the in memory representation of the
// .secret yaml file.
type SecretFile struct {
	Pullers []PullerParent `yaml:"secrets"`
}

// Secret is a small block representing a key and value
type Secret struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type PullerParent struct {
	From  string `yaml:"from`
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

func (p *PullerParent) toSecret() (Secret, error) {
	return Secret{
		Key:   p.Key,
		Value: p.Value,
	}, nil
}

func (p *PullerParent) isRawSecret() bool {
	return p.From == ""
}

func (p *PullerParent) isGcpSecret() bool {
	return p.From == GcpSecret
}

// ToKeyValue turns a secret struct into something a process
// can understand for its environment variables
func (s Secret) ToKeyValue() string {
	return fmt.Sprintf("%s=%s", s.Key, s.Value)
}
