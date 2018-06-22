package main

type Node struct {
	ID      string
	Types   []string
	Label   string `json:",omitempty"`
	Comment string `json:",omitempty"`

	Attrs map[string][]string `json:",omitempty"`
}

func (n Node) Pluck(key string) string {
	all := n.PluckAll(key)
	if len(all) > 0 {
		return all[0]
	}
	return ""
}

func (n Node) PluckAll(key string) []string {
	return pluck(n.Attrs, key)
}

type Property struct {
	Node

	Domain             string
	Range              []string `json:",omitempty"`
	EquivalentProperty []string `json:",omitempty"`

	SubPropertyOf string      `json:",omitempty"`
	SubProperties []*Property `json:",omitempty"`
}

func (p Property) IsFunctional() bool {
	return has(p.Node.Types, OWL_FunctionalProperty)
}

type Class struct {
	Node

	subClassOf string
	SubClassOf *Class `json:",omitempty"`

	Properties []*Property `json:",omitempty"`
}

type DataType struct {
	Node

	Properties []*Property `json:",omitempty"`
}
