// GENERATED CODE, DO NOT EDIT

package dstream

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/golang/snappy"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func (b *bcols) init() {

	fname := path.Join(b.bpath, "dtypes.json")
	fid, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer fid.Close()
	dec := json.NewDecoder(fid)
	dtypes := make(map[string]string)
	err = dec.Decode(&dtypes)
	if err != nil {
		panic(err)
	}
	b.dtypes = dtypes

	// Map from variable names to file names
	fi, err := ioutil.ReadDir(b.bpath)
	if err != nil {
		panic(err)
	}
	vf := make(map[string]string)
	for _, v := range fi {
		na := strings.Split(v.Name(), ".")[0]
		vf[na] = v.Name()
	}

	b.names = b.usenames()

	b.namepos = make(map[string]int)
	for k, na := range b.names {
		b.namepos[na] = k
	}

	b.bdata = b.bdata[0:0]

	for _, na := range b.names {

		// Create concrete types to hold the data
		switch b.dtypes[na] {
		case "float64":
			var x []float64
			b.bdata = append(b.bdata, x)
		case "float32":
			var x []float32
			b.bdata = append(b.bdata, x)
		case "uint64":
			var x []uint64
			b.bdata = append(b.bdata, x)
		case "uint32":
			var x []uint32
			b.bdata = append(b.bdata, x)
		case "uint16":
			var x []uint16
			b.bdata = append(b.bdata, x)
		case "uint8":
			var x []uint8
			b.bdata = append(b.bdata, x)
		case "int64":
			var x []int64
			b.bdata = append(b.bdata, x)
		case "int32":
			var x []int32
			b.bdata = append(b.bdata, x)
		case "int16":
			var x []int16
			b.bdata = append(b.bdata, x)
		case "int8":
			var x []int8
			b.bdata = append(b.bdata, x)
		case "int":
			var x []int
			b.bdata = append(b.bdata, x)
		case "string":
			var x []string
			b.bdata = append(b.bdata, x)
		case "uvarint":
			var x []uint64
			b.bdata = append(b.bdata, x)
		case "varint":
			var x []int64
			b.bdata = append(b.bdata, x)
		case "default":
			panic("bcols.init: unknown type")
		}

		fn, ok := vf[na]
		if !ok {
			msg := fmt.Sprintf("bcols.init: can't find variable %s in %s\n", na, b.bpath)
			os.Stderr.Write([]byte(msg))
			os.Exit(1)
		}

		if strings.HasSuffix(fn, ".bin.gz") {
			// gzip compression
			s := path.Join(b.bpath, fn)
			fid, err := os.Open(s)
			if err != nil {
				panic(err)
			}
			b.toclose = append(b.toclose, fid)
			gid, err := gzip.NewReader(fid)
			if err != nil {
				panic(err)
			}
			b.toclose = append(b.toclose, gid)
			bid := bufio.NewReader(gid) // adds ReadByte
			b.toclose = append(b.toclose, gid)
			b.rdrs = append(b.rdrs, bid)
		} else if strings.HasSuffix(fn, ".bin.sz") {
			// Snappy compression
			s := path.Join(b.bpath, fn)
			fid, err := os.Open(s)
			if err != nil {
				panic(err)
			}
			b.toclose = append(b.toclose, fid)
			gid := snappy.NewReader(fid)
			bid := bufio.NewReader(gid) // adds ReadByte
			b.rdrs = append(b.rdrs, bid)
		} else {
			panic("compression type not recognized")
		}
	}
}

func (b *bcols) Next() bool {

	if b.done {
		return false
	}

	truncate(b.bdata)

	for j, na := range b.names {
		rdr := b.rdrs[j]

		// Handle variable width data
		if b.dtypes[na] == "varint" {
			var v []int64
			for k := 0; k < b.chunksize; k++ {
				var x int64
				x, err := binary.ReadVarint(rdr)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
			continue
		} else if b.dtypes[na] == "uvarint" {
			var v []uint64
			for k := 0; k < b.chunksize; k++ {
				var x uint64
				x, err := binary.ReadUvarint(rdr)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
			continue
		}

		// Handle fixed width data
		switch v := b.bdata[j].(type) {
		case []float64:
			for k := 0; k < b.chunksize; k++ {
				var x float64
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		case []float32:
			for k := 0; k < b.chunksize; k++ {
				var x float32
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		case []uint64:
			for k := 0; k < b.chunksize; k++ {
				var x uint64
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		case []uint32:
			for k := 0; k < b.chunksize; k++ {
				var x uint32
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		case []uint16:
			for k := 0; k < b.chunksize; k++ {
				var x uint16
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		case []uint8:
			for k := 0; k < b.chunksize; k++ {
				var x uint8
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		case []int64:
			for k := 0; k < b.chunksize; k++ {
				var x int64
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		case []int32:
			for k := 0; k < b.chunksize; k++ {
				var x int32
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		case []int16:
			for k := 0; k < b.chunksize; k++ {
				var x int16
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		case []int8:
			for k := 0; k < b.chunksize; k++ {
				var x int8
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		case []int:
			for k := 0; k < b.chunksize; k++ {
				var x int
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		case []string:
			for k := 0; k < b.chunksize; k++ {
				var x string
				err := binary.Read(rdr, binary.LittleEndian, &x)
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				v = append(v, x)
				if j == 0 {
					b.nobs++
				}
			}
			b.bdata[j] = v
		}
	}

	return true
}
