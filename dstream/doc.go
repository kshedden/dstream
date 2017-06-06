// dstream is a package for manipulating streams of typed,
// multivariate data.  A Dstream is a data container that
// (conceptually) holds a rectangular array of data in which the
// columns are variables and the rows are cases or observations.  The
// dstream framework facilitates processing data of this type, with a
// focus on feeding the data into statistical modeling tools.
//
// The data held by a dstream is stored column-wise, and each column
// is partitioned into chunks.  A dstream visits its chunks in order.
// In general, a single chunk resides in memory at one point in time.
//
// Most operations on dstreams take the form of a transformation d =
// f(d).  Many transformations are defined in the package, and it is
// easy to add new transformations.  Examples of transformations are
// Mutate (modify a column in-place) and Drop (drop a set of columns
// from the dstream).
package dstream
