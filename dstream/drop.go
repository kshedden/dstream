package dstream

import "fmt"

type drop struct {
	xform

	dropVars []string
	keepPos  []int
}

func (d *drop) init() {

	hna := make(map[string]bool)
	for _, v := range d.source.Names() {
		hna[v] = true
	}

	dmp := make(map[string]bool)
	for _, na := range d.dropVars {
		if !hna[na] {
			msg := fmt.Sprintf("cannot drop '%s', no such column", na)
			panic(msg)
		}
		dmp[na] = true
	}

	for k, na := range d.source.Names() {
		if !dmp[na] {
			d.keepPos = append(d.keepPos, k)
			d.names = append(d.names, na)
		}
	}
}

// Drop removes the given variables from a Dstream.
func Drop(data Dstream, vars []string) Dstream {
	d := &drop{
		xform: xform{
			source: data,
		},
		dropVars: vars,
	}
	d.init()
	return d
}

func (d *drop) GetPos(j int) interface{} {
	return d.source.GetPos(d.keepPos[j])
}
