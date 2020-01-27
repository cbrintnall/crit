package main

type RawSecret struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

func (p *RawSecret) toSecret() (Secret, error) {
	return Secret{
		Key:   p.Key,
		Value: p.Value,
	}, nil
}
