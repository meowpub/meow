package rdf

import (
	"github.com/meowpub/meow/ld"
)

type Statement struct {
	*ld.Object
}

func (obj Statement) Obj() *ld.Object {
	return obj.Object
}

func (obj Statement) Predicate() interface{} {
	return obj.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#predicate"]
}

func (obj Statement) Object() interface{} {
	return obj.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#object"]
}

func (obj Statement) Subject() interface{} {
	return obj.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#subject"]
}

type Property struct {
	*ld.Object
}

func (obj Property) Obj() *ld.Object {
	return obj.Object
}

type Bag struct {
	*ld.Object
}

func (obj Bag) Obj() *ld.Object {
	return obj.Object
}

type List struct {
	*ld.Object
}

func (obj List) Obj() *ld.Object {
	return obj.Object
}

func (obj List) Rest() interface{} {
	return obj.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#rest"]
}

func (obj List) First() interface{} {
	return obj.V["http://www.w3.org/1999/02/22-rdf-syntax-ns#first"]
}

type Seq struct {
	*ld.Object
}

func (obj Seq) Obj() *ld.Object {
	return obj.Object
}

type Alt struct {
	*ld.Object
}

func (obj Alt) Obj() *ld.Object {
	return obj.Object
}

var (
	_ ld.Entity = Statement{}
	_ ld.Entity = Property{}
	_ ld.Entity = Bag{}
	_ ld.Entity = List{}
	_ ld.Entity = Seq{}
	_ ld.Entity = Alt{}
)
