// GENERATED CODE, DO NOT EDIT
package dstream

// MemCopy returns a Dstream that copies the provided Dstream into
// in-memory storage.  The provided Dstream is copied from its current
// position.  To copy an entire Dstream either pass a newly-created
// Dstream or call Reset before calling MemCopy.
func MemCopy(data Dstream) Dstream {

	nvar := data.NumVar()
	bdata := make([][]interface{}, nvar)

	for data.Next() {
		for j := 0; j < nvar; j++ {
			var y interface{}
			v := data.GetPos(j)
			switch v := v.(type) {
			case []float64:
				z := make([]float64, len(v))
				copy(z, v)
				y = z
			case []float32:
				z := make([]float32, len(v))
				copy(z, v)
				y = z
			case []uint64:
				z := make([]uint64, len(v))
				copy(z, v)
				y = z
			case []uint32:
				z := make([]uint32, len(v))
				copy(z, v)
				y = z
			case []uint16:
				z := make([]uint16, len(v))
				copy(z, v)
				y = z
			case []uint8:
				z := make([]uint8, len(v))
				copy(z, v)
				y = z
			case []string:
				z := make([]string, len(v))
				copy(z, v)
				y = z
			}
			bdata[j] = append(bdata[j], y)
		}
	}

	da := &dataArrays{
		rawData: bdata,
		xform: xform{
			names: data.Names(),
		},
	}

	da.init()

	return da
}
