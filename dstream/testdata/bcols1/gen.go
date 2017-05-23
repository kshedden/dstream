package main

import (
	"compress/gzip"
	"encoding/binary"
	"io"
	"os"

	"github.com/golang/snappy"
)

const (
	n = 1000
)

func writeuvarint(n int, wtr io.Writer) {
	buf := make([]byte, 8)
	var err error
	for k := 0; k < n; k++ {
		x := uint64(k * k)
		m := binary.PutUvarint(buf, x)
		_, err = wtr.Write(buf[0:m])
		if err != nil {
			panic(err)
		}
	}
}

func writeuint64(n int, wtr io.Writer) {
	for k := 0; k < n; k++ {
		x := uint64(k * k)
		err := binary.Write(wtr, binary.LittleEndian, &x)
		if err != nil {
			panic(err)
		}
	}
}

func writeuint32(n int, wtr io.Writer) {
	for k := 0; k < n; k++ {
		x := uint32(k * k)
		err := binary.Write(wtr, binary.LittleEndian, &x)
		if err != nil {
			panic(err)
		}
	}
}

func writeuint16(n int, wtr io.Writer) {
	for k := 0; k < n; k++ {
		x := uint16(k * k)
		err := binary.Write(wtr, binary.LittleEndian, &x)
		if err != nil {
			panic(err)
		}
	}
}

func writeuint8(n int, wtr io.Writer) {
	for k := 0; k < n; k++ {
		x := uint8(k * k)
		err := binary.Write(wtr, binary.LittleEndian, &x)
		if err != nil {
			panic(err)
		}
	}
}

func writefloat64(n int, wtr io.Writer) {
	for k := 0; k < n; k++ {
		x := float64(k * k)
		err := binary.Write(wtr, binary.LittleEndian, &x)
		if err != nil {
			panic(err)
		}
	}
}

func getgz(fname string) (io.WriteCloser, io.WriteCloser) {
	fid, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	gid := gzip.NewWriter(fid)
	return fid, gid
}

func getsz(fname string) (io.WriteCloser, io.WriteCloser) {
	fid, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	sid := snappy.NewBufferedWriter(fid)
	return fid, sid
}

type wg func(string) (io.WriteCloser, io.WriteCloser)

func main() {

	var getfd wg

	for _, tp := range []string{"sz", "gz"} {

		if tp == "sz" {
			getfd = getsz
		} else {
			getfd = getgz
		}

		fid, gid := getfd("uint64.bin." + tp)
		defer fid.Close()
		defer gid.Close()
		writeuint64(n, gid)

		fid, gid = getfd("uint32.bin." + tp)
		defer fid.Close()
		defer gid.Close()
		writeuint32(n, gid)

		fid, gid = getfd("uint16.bin." + tp)
		defer fid.Close()
		defer gid.Close()
		writeuint16(n, gid)

		fid, gid = getfd("uint8.bin." + tp)
		defer fid.Close()
		defer gid.Close()
		writeuint8(n, gid)

		fid, gid = getfd("uvarint.bin." + tp)
		defer fid.Close()
		defer gid.Close()
		writeuvarint(n, gid)

		fid, gid = getfd("float64.bin." + tp)
		defer fid.Close()
		defer gid.Close()
		writefloat64(n, gid)
	}
}
