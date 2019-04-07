package formula

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/kshedden/dstream/dstream"
)

// Tokens that can appear in a formula.
type tokType int

// Allowed token types types.
const (
	vname = iota
	leftp
	rightp
	times
	plus
	icept
	funct
)

// Func is a transformation of a numeric column to a column set.
type Func func(string, []float64) *ColSet

// Operator precedence values; lower number is higher precedence.
var precedence = map[tokType]int{times: 0, plus: 1}

// The token is either a symbol (operator or parentheses), a variable
// name, or a function
type token struct {
	symbol tokType
	name   string // only used if symbol == vname

	// Below are only used for functions
	funcn string
	arg   string
}

// pop removes the last token from the slice, and returns it along
// with the shortened slice.  nil is returned if the slice has length
// zero.
func pop(tokens []*token) ([]*token, *token) {
	if len(tokens) == 0 {
		return nil, nil
	}
	n := len(tokens)
	tok := tokens[n-1]
	tokens = tokens[0 : n-1]
	return tokens, tok
}

// peek returns the last token from the slice.  nil is returned if the
// slice has length zero.
func peek(tokens []*token) *token {
	if len(tokens) == 0 {
		return nil
	}
	n := len(tokens)
	tok := tokens[n-1]
	return tok
}

// push appends the token to the end of the slice and returns the new slice.
func push(tokens []*token, tok *token) []*token {
	return append(tokens, tok)
}

// lex takes a formula and lexes it to obtain an array of tokens.
func lex(input string) ([]*token, error) {

	var tokens []*token
	rdr := strings.NewReader(input)

	isValidContinuation := func(r rune) bool {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			return true
		}
		return false
	}

	for rdr.Len() > 0 {
		r, _, err := rdr.ReadRune()
		if err != nil {
			return nil, err
		}
		switch {
		case r == '(':
			tokens = append(tokens, &token{symbol: leftp})
		case r == ')':
			tokens = append(tokens, &token{symbol: rightp})
		case r == '+':
			tokens = append(tokens, &token{symbol: plus})
		case r == '*':
			tokens = append(tokens, &token{symbol: times})
		case r == '1':
			tokens = append(tokens, &token{symbol: icept})
		case r == ' ':
			// skip whitespace
		case unicode.IsLetter(r) || r == '_':
			name := []rune{r}
			for rdr.Len() > 0 {
				q, _, err := rdr.ReadRune()
				if err != nil {
					panic(err)
				}
				if !isValidContinuation(q) {
					rdr.UnreadRune()
					break
				}
				name = append(name, q)
			}
			tokens = append(tokens, &token{symbol: vname, name: string(name)})
		default:
			return nil, fmt.Errorf("Invalid formula, symbol '%s' is not known.", string(r))
		}
	}

	tokens, err := lexFuncs(tokens)
	return tokens, err
}

func lexFuncs(input []*token) ([]*token, error) {

	output := make([]*token, 0, len(input))
	i := 0
	m := len(input)
	for i < m {
		if i+1 < m && input[i].symbol == vname && input[i+1].symbol == leftp && input[i+3].symbol == rightp {
			// A function
			name := fmt.Sprintf("%s(%s)", input[i].name, input[i+2].name)
			newtok := &token{symbol: funct, name: name, arg: input[i+2].name, funcn: input[i].name}
			output = append(output, newtok)
			i = i + 4
		} else {
			// Not a function
			output = append(output, input[i])
			i++
		}
	}

	return output, nil
}

// isOperator returns true if the token is an opertor (times or plus)
func isOperator(tok *token) bool {
	if tok.symbol == times || tok.symbol == plus {
		return true
	}
	return false
}

// parse converts the formula to RPN
// https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func parse(input []*token) ([]*token, error) {

	var stack []*token
	var output []*token
	var last *token

	for _, tok := range input {

		switch {
		case tok.symbol == vname || tok.symbol == funct || tok.symbol == icept:
			output = append(output, tok)
		case isOperator(tok):
			for {
				last := peek(stack)
				if last == nil || !isOperator(last) {
					break
				}
				if precedence[tok.symbol] > precedence[last.symbol] {
					stack, last = pop(stack)
					output = append(output, last)
				} else {
					break
				}
			}
			stack = push(stack, tok)
		case tok.symbol == leftp:
			stack = push(stack, tok)
		case tok.symbol == rightp:
			for {
				stack, last = pop(stack)
				if last == nil {
					return nil, fmt.Errorf("unbalanced parentheses")
				}
				if last.symbol == leftp {
					break
				} else {
					output = append(output, last)
				}
			}
		}
	}

	for {
		stack, last = pop(stack)
		if last == nil {
			break
		}
		if last.symbol == leftp || last.symbol == rightp {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		output = append(output, last)
	}

	return output, nil
}

// FormulaParser takes a formula and dataset, and produces a design
// matrix from them.
type FormulaParser struct {

	// The formula defining the design matrix
	Formulas []string

	// Produces data in chunks
	RawData dstream.Dstream

	// Reference levels for string variables are omitted when
	// forming indicators
	refLevels map[string]string

	// Codes is a map from variable names to maps from variable
	// values to integer codes.  The distinct values of a
	// variable, excluding the reference level, are mapped to the
	// integers 0, 1, ...  Can be set manually, but if it is not
	// will be computed from data.  Not used if all variables are
	// of float64 type.
	codes map[string]map[string]int

	// Map from function name to function.
	funcs map[string]Func

	// Variables to retain that are not in the formula
	keep []string

	// The final data produced by parsing the formula
	Data *ColSet

	ErrorState error

	// Intermediate data
	workData map[string]*ColSet

	facNames map[string][]string
	rpn      [][]*token // separate RPN for each formula
	rawNames []string
	nvar     int
	names    []string
}

// New creates a FormulaParser from a formula and a data stream.
func New(formula string, rawdata dstream.Dstream) *FormulaParser {

	fp := &FormulaParser{
		Formulas: []string{formula},
		RawData:  rawdata,
	}

	return fp
}

// RefLevels specifies the reference levels of a categorical covariate, which are
// omitted when building design matrices.
func (fp *FormulaParser) RefLevels(reflevels map[string]string) *FormulaParser {
	fp.refLevels = reflevels
	return fp
}

// Codes specifies a mapping from variable names (for categorical variables)
// to code maps, which the distinct levels of the variablle distinct integer
// codes.
func (fp *FormulaParser) Codes(codes map[string]map[string]int) *FormulaParser {
	fp.codes = codes
	return fp
}

// Funcs specifies the Go functions that can be used by name in th
func (fp *FormulaParser) Funcs(funcs map[string]Func) *FormulaParser {
	fp.funcs = funcs
	return fp
}

// Keep defines variables that are retained when generating the output
// data stream, even through they are not present in the formula.
func (fp *FormulaParser) Keep(vars ...string) *FormulaParser {

	names := fp.RawData.Names()
	mp := make(map[string]bool)
	for _, na := range names {
		mp[na] = true
	}
	for _, v := range vars {
		if !mp[v] {
			msg := fmt.Sprintf("Formula: variable '%s' not found", v)
			panic(msg)
		}
	}

	fp.keep = vars
	return fp
}

// Done signals that the FormulaParser has been fully configured and is
// ready for use.
func (fp *FormulaParser) Done() dstream.Dstream {
	err := fp.init()
	if err != nil {
		panic(err)
	}
	return fp
}

// NewMulti accepts several formulas and includes all their parsed
// terms in the resulting data set.
func NewMulti(formulas []string, rawdata dstream.Dstream) *FormulaParser {

	fp := &FormulaParser{
		Formulas: formulas,
		RawData:  rawdata,
	}

	return fp
}

func (fp *FormulaParser) Close() {
}

// ColSet represents a design matrix.  It is an ordered set of named
// numeric data columns.
type ColSet struct {
	Names []string
	Data  []interface{}
}

// Extend a ColSet with the data of another ColSet.
func (c *ColSet) Extend(o *ColSet) {

	// Don't add duplicate terms (which may arise when parsing
	// multiple formulas together or when using Keep).
	mp := make(map[string]bool)
	for _, na := range c.Names {
		mp[na] = true
	}

	for j, na := range o.Names {
		if !mp[na] {
			c.Names = append(c.Names, na)
			c.Data = append(c.Data, o.Data[j])
		}
	}
}

// checkConv ensures that the variables with the given names have been
// converted from raw to ColSet form.
func (fp *FormulaParser) checkConv(v ...string) {
	for _, x := range v {
		fp.convertColumn(x)
	}
}

func (fp *FormulaParser) DropNA() {
	panic("FormulaParser does not support missing values")
}

func (fp *FormulaParser) Missing() []bool {
	panic("FormulaParser does not support missing values")
}

func (fp *FormulaParser) NumObs() int {
	panic("FormulaParser does not know the sample size")
}

func (fp *FormulaParser) NumVar() int {
	return len(fp.Data.Data)
}

func (fp *FormulaParser) GetPos(j int) interface{} {
	return fp.Data.Data[j]
}

func (fp *FormulaParser) Get(na string) interface{} {

	for j, nm := range fp.Data.Names {
		if nm == na {
			return fp.GetPos(j)
		}
	}

	msg := fmt.Sprintf("Formula: variable '%s' not found", na)
	panic(msg)
}

// setCodes inspects the data to determine integer codes for the
// distinct, non-reference levels of each categorical (string type)
// variable.
func (fp *FormulaParser) setCodes() {

	fp.codes = make(map[string]map[string]int)
	fp.facNames = make(map[string][]string)

	// Codes requires resettable data, since we must read through
	// all the data to get the code information.
	fp.RawData.Reset()
	names := fp.RawData.Names()

	for fp.RawData.Next() {

		nvar := fp.RawData.NumVar()
		for j := 0; j < nvar; j++ {
			v := fp.RawData.GetPos(j)
			if v == nil {
				break
			}
			na := names[j]
			switch v := v.(type) {
			case []string:
				// Get the category codes for this
				// variable.  If this is the first
				// chunk, start from scratch.
				codes, ok := fp.codes[na]
				if !ok {
					codes = make(map[string]int)
					fp.codes[na] = codes
				}

				ref := fp.refLevels[na]
				for _, x := range v {
					if x == ref {
						continue
					}
					_, ok := codes[x]
					if !ok {
						// New code
						fm := fmt.Sprintf("%s[%s]", na, x)
						fp.facNames[na] = append(fp.facNames[na], fm)
						codes[x] = len(codes)
					}
				}
			}
		}
	}

	fp.RawData.Reset()
}

// codeStrings creates a ColSet from a string array, creating
// indicator variables for each distinct value in the string array,
// except for ref (the reference level).
func (fp *FormulaParser) codeStrings(na, ref string, s []string) {

	// Get the category codes for this variable
	codes := fp.codes[na]

	var dat []interface{}
	for _, _ = range codes {
		dat = append(dat, make([]float64, len(s)))
	}

	for i, x := range s {
		if x == ref {
			continue
		}
		c := codes[x]
		//_ = dat[c] TODO delete
		//_ = dat[c][i] TODO delete
		dat[c].([]float64)[i] = 1
	}

	fp.workData[na] = &ColSet{Names: fp.facNames[na], Data: dat}
}

func (fp *FormulaParser) getRawCol(na string) interface{} {

	j := -1
	for i, x := range fp.rawNames {
		if x == na {
			j = i
			break
		}
	}

	if j == -1 {
		msg := fmt.Sprintf("Formula: variable '%s' not found.", na)
		panic(msg)
	}

	return fp.RawData.GetPos(j)
}

// convertColumn converts the raw data column with the given name to a
// ColSet object.
func (fp *FormulaParser) convertColumn(na string) {

	// Only need to convert once
	_, ok := fp.workData[na]
	if ok {
		return
	}

	s := fp.getRawCol(na)
	switch s := s.(type) {
	case []string:
		ref := fp.refLevels[na]
		fp.codeStrings(na, ref, s)
	case []float64:
		fp.workData[na] = &ColSet{
			Names: []string{na},
			Data:  []interface{}{s},
		}
	default:
		panic(fmt.Sprintf("unknown type %T in convertColumn", s))
	}
}

// doPlus creates a new ColSet by adding the columnsets named 'a' and
// 'b'.  Addition of two ColSet objects produces a new ColSet with
// columns comprising the union of the two arguments.
func (fp *FormulaParser) doPlus(a, b string) *ColSet {

	ds1 := fp.workData[a]
	ds2 := fp.workData[b]

	var names []string
	var dat []interface{}

	names = append(names, ds1.Names...)
	names = append(names, ds2.Names...)
	dat = append(dat, ds1.Data...)
	dat = append(dat, ds2.Data...)

	return &ColSet{Names: names, Data: dat}
}

// doTimes creates a new ColSet by multiplying the columnsets named
// 'a' and 'b'.  Multiplication produces a new ColSet with columns
// comprising all pairwise product of the two arguments.
func (fp *FormulaParser) doTimes(a, b string) *ColSet {

	ds1 := fp.workData[a]
	ds2 := fp.workData[b]

	var names []string
	var dat []interface{}

	for j1, na1 := range ds1.Names {
		for j2, na2 := range ds2.Names {
			d1 := ds1.Data[j1].([]float64)
			d2 := ds2.Data[j2].([]float64)
			x := make([]float64, len(d1))
			for i, _ := range x {
				x[i] = d1[i] * d2[i]
			}
			names = append(names, na1+":"+na2)
			dat = append(dat, x)
		}
	}

	return &ColSet{names, dat}
}

// createIcept inserts an intercept (array of 1's) into the dataset
// being constructed and returns true if an intercept is not already
// included, otherwise returns false.
func (fp *FormulaParser) createIcept() bool {

	_, ok := fp.workData["icept"]
	if ok {
		return false
	}

	// Get the length of the data set.
	var n int
	nvar := fp.RawData.NumVar()
	for j := 0; j < nvar; j++ {
		x := fp.RawData.GetPos(j)
		switch x := x.(type) {
		case []float64:
			n = len(x)
		case []string:
			n = len(x)
		}
		break
	}

	x := make([]float64, n)
	for i, _ := range x {
		x[i] = 1
	}
	fp.workData["icept"] = &ColSet{Names: []string{"icept"}, Data: []interface{}{x}}

	return true
}

func (fp *FormulaParser) Names() []string {
	return fp.names
}

// init performs lexing and parsing of the formula, only done once.
func (fp *FormulaParser) init() error {

	for _, fml := range fp.Formulas {
		fmx, err := lex(fml)
		if err != nil {
			return err
		}
		rpn, err := parse(fmx)
		if err != nil {
			return err
		}
		fp.rpn = append(fp.rpn, rpn)
	}

	if fp.codes == nil {
		fp.setCodes()
	}

	// Read one chunk to get the number of variables
	ok := fp.Next()
	if !ok {
		return fmt.Errorf("Unable to read data")
	}
	if fp.ErrorState != nil {
		return fp.ErrorState
	}
	fp.nvar = len(fp.Data.Names)
	fp.names = fp.Data.Names
	fp.RawData.Reset()

	return nil
}

func (fp *FormulaParser) doFormula(rpn []*token) bool {

	fp.runFuncs(rpn)

	// Special case a single variable with no operators
	if len(rpn) == 1 {
		na := rpn[0].name
		fp.checkConv(na)
		fp.Data.Extend(fp.workData[na])
		fp.workData = nil
		return true
	}

	var stack []string

	for ix, tok := range rpn {
		last := ix == len(rpn)-1
		switch {
		case isOperator(tok):
			if len(stack) < 2 {
				fp.ErrorState = fmt.Errorf("not enough arguments")
				return false
			}

			// Pop the last two arguments off the stack
			arg2 := stack[len(stack)-1]
			arg1 := stack[len(stack)-2]
			stack = stack[0 : len(stack)-2]

			fp.checkConv(arg1, arg2)
			var rslt *ColSet
			switch tok.symbol {
			case plus:
				rslt = fp.doPlus(arg1, arg2)
			case times:
				rslt = fp.doTimes(arg1, arg2)
			default:
				panic("invalid symbol")
			}
			if last {
				// The last thing computed is the result
				fp.Data.Extend(rslt)
			}
			nm := fmt.Sprintf("tmp%d", ix)
			fp.workData[nm] = rslt
			stack = append(stack, nm)
		case tok.symbol == icept:
			q := fp.createIcept()
			if q {
				stack = append(stack, "icept")
			}
		case tok.symbol == vname:
			fp.checkConv(tok.name)
			stack = append(stack, tok.name)
		case tok.symbol == funct:
			stack = append(stack, tok.name)
		}
	}

	if len(stack) != 1 {
		fp.ErrorState = fmt.Errorf("invalid formula [2]")
		return false
	}

	return true
}

// Parse builds a design matrix out of the formula and raw data.
func (fp *FormulaParser) Next() bool {

	fp.Data = new(ColSet)
	if fp.RawData.Next() == false {
		return false
	}

	fp.rawNames = fp.RawData.Names()

	for _, rpn := range fp.rpn {
		fp.workData = make(map[string]*ColSet)
		fp.doFormula(rpn)
	}

	for _, na := range fp.keep {
		fp.Data.Names = append(fp.Data.Names, na)
		fp.Data.Data = append(fp.Data.Data, fp.RawData.Get(na))
	}

	fp.workData = nil

	return true
}

func (fp *FormulaParser) runFuncs(rpn []*token) {

	for _, tok := range rpn {
		if tok.symbol != funct {
			continue
		}

		f := fp.funcs[tok.funcn]
		x := fp.getRawCol(tok.arg)
		switch x := x.(type) {
		case []float64:
			fp.workData[tok.name] = f(tok.name, x)
		default:
			panic("funtions can only be applied to numeric data")
		}
	}
}

// Reset changes the state of the formula parser so that subsequent
// reads start from the beginning of the dataset.
func (fp *FormulaParser) Reset() {
	fp.ErrorState = nil
	fp.RawData.Reset()
}

func find(s []string, x string) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}
