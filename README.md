[![Build Status](https://travis-ci.com/kshedden/dstream.svg?branch=master)](https://travis-ci.com/kshedden/dstream)
[![Go Report Card](https://goreportcard.com/badge/github.com/kshedden/dstream)](https://goreportcard.com/report/github.com/kshedden/dstream)
[![codecov](https://codecov.io/gh/kshedden/dstream/branch/master/graph/badge.svg)](https://codecov.io/gh/kshedden/dstream)
[![GoDoc](https://godoc.org/github.com/kshedden/dstream?status.png)](https://godoc.org/github.com/kshedden/dstream)

# Preliminaries

To install:

```
go get github.com/kshedden/dstream...
```

# Introduction

__Dstream__ is a package for manipulating streams of typed,
multivariate data in [Go](http://golang.org).  A Dstream is a
[dataframe](http://pandas.pydata.org)-like container that holds a
rectangular array of data in which the columns are variables and the
rows are cases or observations.

Dstream is designed to handle large datasets, where it is not
possible to load all data for all variables into memory at once.  To
achieve this, Dstream utilizes a _chunked_, _column-based_ storage format.
A chunk contains the data for a contiguous block of rows.
The data are stored by variable (column-wise) in typed
Go slices.  Only one chunk of the Dstream is held in memory at one
time.

During data processing, the chunks are visited in order.  The `Next`
method advances the Dstream to the next chunk.  When possible, the
memory backing a chunk is re-used for the next chunk.  Therefore, a
chunk must either be completely processed, or copied to independent
memory before making subsequent calls to `Next`.
Most Dstreams can be reset with the `Reset` method and
read multiple times, but this requires all the overhead of the initial
read (the data will be fully re-processed from its source following a
call to `Reset`).

The typical pattern for working with a Dstream is to visit the chunks
in sequence, extract variables as needed, and perform the desired
processing.  A template for this operation using a Dstream named _ds_
is:

```
for ds.Next() {
    x := ds.Get("x3").([]float64) // extract the variable named "x3"
    // do something with x
}
```

### Transformations

A Dstream is processed by applying _transformations_ to it.  Each
transformation yields a new Dstream, so the transformations can be
chained.  Much like Unix pipelines, each transformation performs a
specific (usually simple) modification to the data.  Chaining several
such transformations in sequence allows complex manipulations to be
performed.

Since the output Dstream of a transformation may share memory with its
input, references to the input Dstream should not be retained.  A
typical example chaining two transformations would look like this:

```
ds = DropNA(ds)          // drop all rows with any missing values
ds = Mutate(ds, "x1", f) // apply the function f in-place to the variable named "x1"
```

The most common transformations can be grouped as follows:

* _Extension_: add new variables to the dstream, usually derived from
  the existing variables.  Examples include
  [Generate](https://godoc.org/github.com/kshedden/dstream/dstream#Generate),
  [DiffChunk](https://godoc.org/github.com/kshedden/dstream/dstream#DiffChunk),
  [LagChunk](https://godoc.org/github.com/kshedden/dstream/dstream#LagChunk),
  and
  [LinApply](https://godoc.org/github.com/kshedden/dstream/dstream#LinApply).

* _Re-chunking_: modify the chunk boundaries.  Examples include
  [Segment](https://godoc.org/github.com/kshedden/dstream/dstream#Segment)
  and [SizeChunk](https://godoc.org/github.com/kshedden/dstream/dstream#SizeChunk).

* _Mutation_: in-place modifications of the data, examples include
  [Mutate](https://godoc.org/github.com/kshedden/dstream/dstream#Mutate).

* _Selection_: dropping rows or columns, examples include
  [DropNa](https://godoc.org/github.com/kshedden/dstream/dstream#DropNA),
  [Drop](https://godoc.org/github.com/kshedden/dstream/dstream#Drop),
  [Filter](https://godoc.org/github.com/kshedden/dstream/dstream#Filter).

* _Copying_:
  [MemCopy](https://godoc.org/github.com/kshedden/dstream/dstream#DropNA)
  returns an in-memory Dstream that is a copy of a given Dstream.

* _Type conversion_:
  [Convert](https://godoc.org/github.com/kshedden/dstream/dstream#Convert)
  converts among numeric types.

### Chunks

A Dstream's chunks have two distinct roles.  First, they serve to
break the data into subsets of manageable size.  Second, for some
analytic procedures, the chunks define meaningful data subsets (e.g. a
chunk may contain all records for a single value of an index
variable).  Here is a pipeline that illustrates both of these roles:

```
tc := dstream.CSVTypeConf {
	Float64: []string{"Index", "Speed"}
	Flaot64Pos: []int{0, 1}
}
da := dstream.FromCSV(r).TypeConf(tc).ChunkSize(1000).Done()
dx = da.Segment(da, "Index")
dx = dx.DiffChunk(dx, map[string][int]{"Speed", 2})
```

In the above example, we first set up a Dstream to read CSV-formatted
data from an io.Reader, using a chunk size of 1 million to limit the
number of rows held in memory at one time.  We then use Segment to redefine the
chunk boundaries, so that each chunk contains the values for one level
of the Index variable (note that the data must be pre-sorted with
respect to the Index variable
for this to work).  We then difference the Speed variable within each
level of Index (i.e. within each chunk).  Since DiffChunk does not
difference across chunk boundaries, the chunk boundaries are not
merely a computational consideration in this example, they impact the
output of the pipeline.

### Types

Each column in a Dstream has a fixed type.  The core of the package
supports 10 numeric types (1, 2, 3, and 8 byte signed and unsigned
integers, and 4 and 8 byte floating point values), along with strings
and time values.

When accessing a variable's values using
[Get](https://godoc.org/github.com/kshedden/dstream/dstream#Get) or
[GetPos](https://godoc.org/github.com/kshedden/dstream/dstream#GetPos),
the data for one variable, in one chunk, is provided as a slice of
values.  This slice is returned as an empty interface{} which can be
type-asserted to a concrete type, like this:

```
x := da.Get("x").([]uint8)
```

Conversion from any numeric type to any other numeric type can be
carried out using the `Convert` transformation, for example:

```
da = Convert(da, "x1", "int32")
```

Currently, many of the Dstream transformations are only implemented
for a limited range of types, `[]float64` is the most widely-supported
type.  We plan to support for all types in all transformations in the
near future.  We are using
[templates](https://golang.org/pkg/text/template) and code generation
to do this without too much source bloat.

### Utility functions

Dstream provides several utility functions for working with Dstreams,
including
[Equal](https://godoc.org/github.com/kshedden/dstream/dstream#Equal)
and
[EqualReport](https://godoc.org/github.com/kshedden/dstream/dstream#EqualReport)
for making comparisons,
[GetCol](https://godoc.org/github.com/kshedden/dstream/dstream#GetCol)
and
[GetColPos](https://godoc.org/github.com/kshedden/dstream/dstream#GetColPos)
for extracting columns into slices.

The [Join](https://godoc.org/github.com/kshedden/dstream/dstream#Join)
framework allows several Dstreams to be joined at the chunk level
based on a shared index variable.

### Exported types and the xform implementation

A Dstream is any Go struct that implements the Dstream interface,
which has seven methods: Next, Names, Get, GetPos, NumVar, NumObs, and
Reset.  Most concrete Dstream types are returned by exported
functions, but the types themselves are not exported from the package.
Thus, the caller sees a Dstream value as a Dstream interface type, not
as its underlying concrete type.  Using interface types allows
interoperability between different concrete Dstream types when working
with transformations.  An exception to this rule is that a few Dstream
types have additional methods that are not part of the Dstream
interface (such as the CSV reader).  To access these methods, the
value must be stored in a variable having the appropriate concrete
type.

To simplify implementation of Dstreams and transformations, a
prototypical transformation called `xform` is provided.  The `xform`
type fully implements the Dstream interface in a trivial way.  Most of
the transforms embed an xform, and re-implement some but not all of
its methods as needed.

### The memory-backed reference implementation

The `DataFrame` type serves as a reference implementation for a
Dstream.  This implementation uses in-memory sharded arrays to store
the values for each variable.  The `DataFrame` type is useful for
smaller datasets.  After substantial reduction (e.g. filtering), a
large disk-backed Dstream may be converted to a `DataFrame` type using
`MemCopy` (much like use of `collect` in Spark).

### Input/output and data sources

A Dstream is created from a data source.  We provide three frameworks for
serializing data to and from files.  The easiest approach is to use the
[NewLoad](https://godoc.org/github.com/kshedden/dstream/dstream#NewLoad) and
[Save](https://godoc.org/github.com/kshedden/dstream/dstream#Save)
functions.  The whole dstream is serialized and stored in a single compressed
file.  The file format uses Go [gobs](https://blog.golang.org/gobs-of-data).
The files are read and written by chunk,
so this format can be used for large data sets that do not fit into memory.

[StreamCSV](https://godoc.org/github.com/kshedden/dstream/dstream#StreamCSV)
can be used to read and write text/csv files.

[Bcols](https://godoc.org/github.com/kshedden/dstream/dstream#Bcols)
is a binary format that stores the data in a hierarchy of directories and files in raw native form.

Since a Dstream is based on a minimal Go
[interface](https://golang.org/doc/effective_go.html#interfaces_and_types),
Dstreams readers for other data sources can be easily implemented.

# Dataframes

A `DataFrame` is a memory-backed Dstream.  If `df` is any Dstream that can fit into memory,
it can be converted to a `DataFrame` using `MemCopy`:

```
da := dstream.MemCopy(df)
```

Since a DataFrame allocates independent memory for every chunk, it is safe to process
the chunnks of a DataFrame in parallel.

### Status

Dstream is under active development.  Changes that break compatability
are likely.

### Tests

There are many unit tests, but it is likely that at present some core
functionality, and many corner-cases are not well covered by tests.

### Statistical analysis

Dstreams can be fed into statistical modeling tools, including
[glm](https://github.com/kshedden/statmodel/tree/master/glm),
[dimred](https://github.com/kshedden/dimred) and
[duration](https://github.com/kshedden/statmodel/duration).
