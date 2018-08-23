package ld

type Namespace struct {
	ID      string
	Short   string
	Props   map[string]string
	Classes map[string]func(Entity) Entity
}
