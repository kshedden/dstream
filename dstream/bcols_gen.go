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

	b.doneInit = true
}

func (b *bcols) Next() bool {

	if !b.doneInit {
		panic("Call Done before using stream")
	}

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
				x, err := rdr.ReadString('\n')
				if err == io.EOF {
					b.done = true
					break
				} else if err != nil {
					msg := fmt.Sprintf("Error reading variable '%s' at position %d\n", na, b.nobs)
					print(msg)
					panic(err)
				}
				x = strings.TrimRight(x, "\n")
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

func (tb *toBCols) writeDtypes() {

	dtypes := make(map[string]string)

	tb.stream.Reset()
	names := tb.stream.Names()
	tb.stream.Next()
	for j, na := range names {
		u := tb.stream.GetPos(j)
		switch u.(type) {
		case []float64:
			dtypes[na] = "float64"

		case []float32:
			dtypes[na] = "float32"

		case []uint64:
			dtypes[na] = "uint64"

		case []uint32:
			dtypes[na] = "uint32"

		case []uint16:
			dtypes[na] = "uint16"

		case []uint8:
			dtypes[na] = "uint8"

		case []int64:
			dtypes[na] = "int64"

		case []int32:
			dtypes[na] = "int32"

		case []int16:
			dtypes[na] = "int16"

		case []int8:
			dtypes[na] = "int8"

		case []int:
			dtypes[na] = "int"

		case []string:
			dtypes[na] = "string"
		}
	}

	f, err := os.Create(path.Join(tb.path, "dtypes.json"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.Encode(dtypes)
}

func (tb *toBCols) write() {

	tb.stream.Reset()
	names := tb.stream.Names()

	for tb.stream.Next() {
		for j := range names {
			u := tb.stream.GetPos(j)
			switch u.(type) {
			case []float64:
				v := u.([]float64)
				err := binary.Write(tb.wtrs[j], binary.LittleEndian, &v)
				if err != nil {
					panic(err)
				}
			case []float32:
				v := u.([]float32)
				err := binary.Write(tb.wtrs[j], binary.LittleEndian, &v)
				if err != nil {
					panic(err)
				}
			case []uint64:
				v := u.([]uint64)
				err := binary.Write(tb.wtrs[j], binary.LittleEndian, &v)
				if err != nil {
					panic(err)
				}
			case []uint32:
				v := u.([]uint32)
				err := binary.Write(tb.wtrs[j], binary.LittleEndian, &v)
				if err != nil {
					panic(err)
				}
			case []uint16:
				v := u.([]uint16)
				err := binary.Write(tb.wtrs[j], binary.LittleEndian, &v)
				if err != nil {
					panic(err)
				}
			case []uint8:
				v := u.([]uint8)
				err := binary.Write(tb.wtrs[j], binary.LittleEndian, &v)
				if err != nil {
					panic(err)
				}
			case []int64:
				v := u.([]int64)
				err := binary.Write(tb.wtrs[j], binary.LittleEndian, &v)
				if err != nil {
					panic(err)
				}
			case []int32:
				v := u.([]int32)
				err := binary.Write(tb.wtrs[j], binary.LittleEndian, &v)
				if err != nil {
					panic(err)
				}
			case []int16:
				v := u.([]int16)
				err := binary.Write(tb.wtrs[j], binary.LittleEndian, &v)
				if err != nil {
					panic(err)
				}
			case []int8:
				v := u.([]int8)
				err := binary.Write(tb.wtrs[j], binary.LittleEndian, &v)
				if err != nil {
					panic(err)
				}
			case []int:
				v := u.([]int)
				err := binary.Write(tb.wtrs[j], binary.LittleEndian, &v)
				if err != nil {
					panic(err)
				}
			case []string:
				v := u.([]string)
				for _, x := range v {
					_, err := tb.wtrs[j].Write([]byte(x + "\n"))
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
}
