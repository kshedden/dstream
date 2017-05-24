Link to Godoc [documentation](https://godoc.org/github.com/kshedden/dstream/dstream)

To install:

```
go get github.com/kshedden/dstream/dstream
go get github.com/kshedden/dstream/formula
```

__Dstream__ is a package for manipulating streams of typed,
multivariate data in [Go](http://golang.org).  A Dstream is a
[dataframe](http://pandas.pydata.org)-like container that
(conceptually) holds a rectangular array of data in which the columns
are variables and the rows are cases or observations.  The Dstream
framework facilitates processing data of this type, with a special
focus on feeding the data into statistical modeling tools such as
regression analysis.

Dstream is designed to work with large datasets where it is not
possible to load all data for all variables into memory at once.  To
achieve this, Dstream utilizes a _chunked_, _columnar_ representation.
A chunk contains the data for all of the Dstream's variables for a
consecutive subset of rows, stored by variable (column-wise) in typed
arrays.

The chunks are visited in linear order.  When possible, the memory
backing a chunk is re-used for the next chunk.  Therefore, a chunk
must be either completely processed, or copied to independent memory
before subsequent chunks are read.  Random chunk access and sorting
across chunks is not permitted.  Most Dstreams can be Reset and read
multiple times, but this requires all the overhead of the initial read
(e.g. the data will be re-processed from its source).

The typical pattern for working with a Dstream is to visit the chunks
in sequence, extract variables as needed, and perform the desired
processing.  A template for this using a Dstream named _da_ is:

```
for da.Next() {
    x := da.Get("x3").([]float64)
    // do something with x
}
```

### Transformations

A Dstream is processed by applying _transformations_ to it.  Each
transformation yields a new Dstream, so the transformations can be
chained.  Much like Unix pipelines, each transformation performs a
specific (usually simple) modification to the data.  Combining several
such transformations in sequence allows complex manipulations to be
performed.

When possible, the output Dstream of a transformation shares memory
with its input, so references to the input Dstream should not be
retained.  A typical example chaining two transformations would look
like this:

```
ds = DropNA(ds)         // drop all rows with any missing values
ds = Muate(ds, "x1", f) // apply the function f in-place to the variable named "x1"
```

The most common transformations can be grouped as follows:

* _Extension_: add new variables to the dstream, usually defined in
  terms of the existing variables.  Examples include
  [DiffChunk](https://godoc.org/github.com/kshedden/dstream/dstream#DiffChunk),
  [LagChunk](https://godoc.org/github.com/kshedden/dstream/dstream#LagChunk),
  [Apply](https://godoc.org/github.com/kshedden/dstream/dstream#DropNA),
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
  [FilterCol](https://godoc.org/github.com/kshedden/dstream/dstream#FilterCol).

* _Copying_:
  [MemCopy](https://godoc.org/github.com/kshedden/dstream/dstream#DropNA)
  returns an in-memory Dstream that is a copy of a given Dstream.

### Type support

Each column in a Dstream has a fixed type.  When accessing a
variable's values using
[Get](https://godoc.org/github.com/kshedden/dstream/dstream#Get) or
[GetPos](https://godoc.org/github.com/kshedden/dstream/dstream#GetPos),
the data for one variable, in one chunk, is provided as a slice of
values.  To support multiple data types, this slice is returned as an
empty interface{} which can be type-asserted to a concrete type, like
this:

```
x := da.Get("x").([]uint8)
```

Currently, many of the Dstream transformations are only implemented
for a limited range of types, `[]float64` is the most
widely-supported.

### Utility functions

Dstream provides several utility functions for working with Dstreams,
including
[Equal](https://godoc.org/github.com/kshedden/dstream/dstream#Equal)
and
[EqualReport](https://godoc.org/github.com/kshedden/dstream/dstream#EqualReport)
for comparison,
[GetCol](https://godoc.org/github.com/kshedden/dstream/dstream#GetCol)
and
[GetColPos](https://godoc.org/github.com/kshedden/dstream/dstream#GetColPos)
for extracting columns into slices.

### Data sources

A Dstream is created from a data source.  We provide two procedures
[StreamCSV](https://godoc.org/github.com/kshedden/dstream/dstream#StreamCSV)
and
[Bcols](https://godoc.org/github.com/kshedden/dstream/dstream#DropNAhttps://godoc.org/github.com/kshedden/dstream/dstream#Bcols)
for constructing dstreams from certain types of text (csv) and binary
data.  A dstream is based on a minimal Go
[interface](https://golang.org/doc/effective_go.html#interfaces_and_types),
so Dstreams can be obtained from other data sources by implementing a
reader that implements the Dstream interface.
