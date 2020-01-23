package main

import "fmt"

// SecretFile is the in memory representation of the
// .secret yaml file.
type SecretFile struct {
	Secrets []Secret `yaml:"secrets"`
}

// Secret is a small block representing a key and value
type Secret struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type PullerParent struct {
	From string `yaml:"from`
	Key  string `yaml:"key"`
}

type Puller interface {
	getCreds() Secret
}

type GcpSecretsManagerPuller struct{}

func (a GcpSecretsManagerPuller) getCreds() (Secret, error) {
	return Secret{}, nil
}

// ToKeyValue turns a secret struct into something a process
// can understand for its environment variables
func (s Secret) ToKeyValue() string {
	return fmt.Sprintf("%s=%s", s.Key, s.Value)
}
