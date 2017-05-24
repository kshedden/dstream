__Dstream__ is a package for manipulating streams of multivariate data
in [Go](http://golang.org).  A Dstream is a
[dataframe](http://pandas.pydata.org)-like container that holds a
rectangular array of data in which the columns are variables and the
rows are cases or observations.  The Dstream framework facilitates
processing data of this type, with a special focus on feeding the data
into statistical modeling tools such as regression analysis.

__Dstream__ is designed to work with large datasets where it is not
possible to load all data for all variables into memory at once.  To
achieve this, a Dstream utilizes a _chunked_, _columnar_
representation.  A chunk contains the data for all of the Dstream's
variables for a consecutive subset of rows.

The chunks are visited in linear order.  When possible, the memory
backing a chunk is re-used for the next chunk.  Therefore, a chunk
must be either completely processed, or copied to independent memory
before the next chunk is read.  Random chunk access and sorting across
chunks is not permitted.  A Dstream can be _Reset_ and read multiple
times, but this requires all the overhead of the initial read
(i.e. the data will be re-read from disk).

The typical pattern for working with a Dstream is to visit the chunks
in sequence, extract variables as needed, and perform the desired
processing.  A template for this using a Dstream named `da` is:

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

`
ds = DropNA(ds)         // drop all rows with any missing values
ds = Muate(ds, "x1", f) // apply f in-place the variable named "x1"
`

The most common transformations can be divided into the following types:

* _Extension_: add new variables to the dstream, usually defined in
  terms of the existing variables.  Examples include __Diffchunk__,
  __Lagchunk__, __Apply__, and __LinApply__.

* _Re-chunking_: modify the chunk boundaries.  Examples include
  __Segment__ and __Sizechunk__.

* _Mutation_: in-place modifications of the data, examples include
  __Mutate__.

* _Selection_: dropping rows or columns, examples include __Dropna__,
  __Drop__, __FilterCol__.

### Type support

Each column in a Dstream has a fixed type.  When accessing a
variable's values using `Get` or `GetPos`, the data for one variable,
in one chunk, is provided as a slice of values.  To support multiple
data types, this slice is returned as an empty interface{} which can
be type-asserted to a concrete type.  Currently, many of the Dstream
transformations are only implemented for a limited range of types,
`[]float64` is the most widely-used.

### Data sources

A Dstream is created from a data source.  We provide two procedures
__StreamCSV__ and __Bcols__ for constructing dstreams from certain
types of text (csv) and binary data.  A dstream is based on a minimal
Go
[interface](https://golang.org/doc/effective_go.html#interfaces_and_types)
so Dstreams can be obtained from other data sources by implementing a
reader that implements the Dstream interface.