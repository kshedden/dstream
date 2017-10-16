// Package dstream provides facilities for manipulating streams of
// typed, multivariate data.  A Dstream is a data container that
// (conceptually) holds a rectangular array of data in which the
// columns are variables and the rows are cases or observations.  The
// dstream framework facilitates processing data of this type in a
// streaming manner, with a focus on feeding the data into statistical
// modeling tools.
//
// The data held by a Dstream is stored as chunks of contiguous rows.
// Within each chunk, the data are stored column-wise.  A Dstream
// visits its chunks in order.  When processing a Dstream, call Next
// to advance to the next chunk, then call Get to retrieve the data
// for one column.
//
// Most operations on Dstreams take the form of a transformation d =
// f(d), where d is a Dstream.  Many transformations are defined in
// the package, and it is easy to add new transformations.  Examples
// of transformations are Mutate (modify a column in-place) and Drop
// (drop one or more columns from the Dstream).
package dstream
