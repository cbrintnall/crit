package main

import "fmt"

// SecretFile is the in memory representation of the
// .secret yaml file.
type SecretFile struct {
	RawSecrets    []RawSecret    `yaml:"secrets"`
	GoogleSecrets []GoogleSecret `yaml:"google"`
}

// Secret is a small block representing a key and value
type Secret struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

type ResolveableSecret interface {
	toSecret() (Secret, error)
}

// ResolveAll grabs all secrets defined in a secret file and returns them.
func (s *SecretFile) ResolveAll() ([]Secret, error) {
	secrets := []Secret{}

	for _, r := range s.RawSecrets {
		secret, err := r.toSecret()

		if err != nil {
			return secrets, err
		}

		secrets = append(secrets, secret)
	}

	for _, g := range s.GoogleSecrets {
		secret, err := g.toSecret()

		if err != nil {
			return secrets, err
		}

		secrets = append(secrets, secret)
	}

	return secrets, nil
}

// ToKeyValue turns a secret struct into something a process
// can understand for its environment variables
func (s Secret) ToKeyValue() string {
	return fmt.Sprintf("%s=%s", s.Key, s.Value)
}
