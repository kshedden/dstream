Link to Godoc [documentation](https://godoc.org/github.com/kshedden/dstream/dstream)

To install:

```
go get github.com/kshedden/dstream/dstream
```

__Dstream__ is a package for manipulating streams of typed,
multivariate data in [Go](http://golang.org).  A Dstream is a
[dataframe](http://pandas.pydata.org)-like container that
(conceptually) holds a rectangular array of data in which the columns
are variables and the rows are cases or observations.  The Dstream
framework facilitates processing data of this type, with a primary
focus on feeding the data into statistical modeling tools such as
regression analysis.

Dstream is designed to work with large datasets, where it is not
possible to load all data for all variables into memory at once.  To
achieve this, Dstream utilizes a _chunked_, _columnar_ storage format.
A chunk contains the data for all of the Dstream's variables for a
consecutive subset of rows, stored by variable (column-wise) in typed
arrays.  A single chunk is brought into memory at once.  The Dstream as
a whole will generally not reside in memory.

During data processing, the chunks are visited in linear order.  When
possible, the memory backing a chunk is re-used for the next chunk.
Therefore, a chunk must either be completely processed, or copied to
independent memory before subsequent chunks are read.  Random chunk
access and sorting across chunks is not permitted (if processing of
this type is needed, it should be done with other tools before forming
the Dstream).  Most Dstreams can be Reset and read multiple times, but
this requires all the overhead of the initial read (the data will be
fully re-processed from its source following a call to Reset).

The typical pattern for working with a Dstream is to visit the chunks
in sequence, extract variables as needed, and perform the desired
processing.  A template for this operation using a Dstream named _da_
is:

```
for da.Next() {
    x := da.Get("x3").([]float64) // extract the variable named "x3"
    // do something with x
}
```

Note that the Next method of a dstream attempts to advance to the next
chunk, returning true if successful and false if not.

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
ds = Mutate(ds, "x1", f) // apply the function f in-place to the variable named "x1"
```

If references to the input and output of a transformation are both
retained, e.g. `d2 = transform(d1)`, it is extremely important to
never mix calls on d1 and d2.

The most common transformations can be grouped as follows:

* _Extension_: add new variables to the dstream, usually derived from
  the existing variables.  Examples include
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

### Chunks

A Dstream's chunks have two distinct roles.  First, they serve to
break the data into subsets of manageable size.  Second, for some
analytic procedures, the chunks define meaningful data subsets (e.g. a
chunk may contain all records for a single value of an index
variable).  Here is a pipeline that illustrates both of these roles:

```
da := dstream.FromCSV(r)
// ... more steps to setup the CSV reader
da.SetChunkSize(1000000) // read chunks of 1 million rows at a time
dx = da.Segment(da, []string{"Index"})
dx = dx.DiffChunk(dx, map[string][int]{"Speed", 2})
```

In the above example, we first set up a Dstream to read CSV-formatted
data from an io.Reader, using a chunk size of 1 million to limit the
number of distinct raw reads.  We then use Segment to redefine the
chunk boundaries, so that each chunk contains the values for one level
of the Index variable (note that the data must be pre-sorted by Index
for this to work).  We then difference the Speed variable within each
level of Index (i.e. within each chunk).  Since DiffChunk does not
difference across chunk boundaries, the chunk boundaries are not
merely a computational consideration in this example, they impact the
output of the pipeline.

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

The dataArrays type serves as a reference implementation for a
Dstream.  This implementation uses in-memory sharded arrays to store
the values for each variable.  The dataArrays type is useful for
smaller datasets.  After substantial reduction (e.g. filtering), a
large disk-backed Dstream may be converted to a dataArrays type using
`MemCopy` (much like use of `collect` in Spark).

### Data sources

A Dstream is created from a data source.  We provide two functions
[StreamCSV](https://godoc.org/github.com/kshedden/dstream/dstream#StreamCSV)
and
[Bcols](https://godoc.org/github.com/kshedden/dstream/dstream#DropNAhttps://godoc.org/github.com/kshedden/dstream/dstream#Bcols)
for constructing Dstreams from certain types of text (csv) and binary
data.  A Dstream is based on a minimal Go
[interface](https://golang.org/doc/effective_go.html#interfaces_and_types),
so Dstreams can be obtained from other data sources by implementing a
reader that implements the Dstream interface.

### Status

Dstream is under active development.  Changes that break compatability
are likely.

### Tests

There are many unit tests, but it is likely that at present some core
functionality, and many corner-cases are not well covered by tests.

### Statistical analysis

Dstreams can be fed into statistical modeling tools, including
[goglm](https://github.com/kshedden/goglm),
[dimred](https://github.com/kshedden/dimred) and duration.