// Package dqrack is a helping wrapper for Dgraph.
// It provides a load of useful "marshalling" functions
// to allow lazy work.
//
// Pronounced "d-crack"
package dqrack

import (
	"encoding/json"
	"reflect"
	"strings"

	dgraph "github.com/dgraph-io/dgraph/client"
	"github.com/jmoiron/sqlx/reflectx"
)

// Dqrack is a set of cheap tricks, like any other cheap trick library.
type Dqrack struct {
	Dgraph *dgraph.Dgraph
	Mapper *reflectx.Mapper
}

// Qrackable types let us get information out of structs in a non-arbitrary fashion.
type Qrackable interface {
	GetName() string
}

func New(dg *dgraph.Dgraph) *Dqrack {
	return &Dqrack{
		Dgraph: dg,
		Mapper: reflectx.NewMapper("dq"),
	}
}

// GetNode creates a node with edges from the struct.
func (dq *Dqrack) GetNode(v Qrackable) (n dgraph.Node, err error) {
	vName := v.GetName()
	n, err = dq.Dgraph.NodeBlank(vName)
	if err != nil {
		return
	}

	rt := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)
	fm := dq.Mapper.FieldMap(rv)

	var e dgraph.Edge

	// basic edges
	// type makes it searchable by the struct name
	e = n.Edge("_type")
	e.SetValueString(strings.ToLower(rt.Name()))
	e.ConnectTo(n)

	// identity is it's own name, or specific identity.
	e = n.Edge("_identity")
	e.SetValueString(vName)
	e.ConnectTo(n)

	for key, rfv := range fm {
		sf, _ := rt.FieldByName(key)
		skey := sf.Tag.Get("dq")

		e = n.Edge(skey)
		setEdge(e, rfv)
		e.ConnectTo(n)
	}

	return
}

// setEdge figures out the best edge type for some struct value
// TODO: forcibly define default with struct tags.
func setEdge(e dgraph.Edge, v reflect.Value) error {
	switch v.Type().Kind() {
	case reflect.String:
		return e.SetValueString(v.String())
	case reflect.Bool:
		return e.SetValueBool(v.Bool())
	case reflect.Int:
		return e.SetValueInt(v.Int())
	default:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return err
		}
		return e.SetValueDefault(string(b))
	}
}
