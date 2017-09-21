// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 344.

// Package sexpr provides a means for converting Go objects to and
// from S-expressions.
package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

//!+Decoder

type Decoder struct {
	scan  scanner.Scanner
	token rune
}

func NewDecoder(r io.Reader) *Decoder {
	d := Decoder{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	d.scan.Init(r)
	return &d
}

func (d *Decoder) next()        { d.token = d.scan.Scan() }
func (d *Decoder) text() string { return d.scan.TokenText() }
func (d *Decoder) consume(want rune) {
	if d.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", d.text(), want))
	}
	d.next()
}

//!-Decoder

//!+Unmarshal
// Unmarshal parses S-expression data and populates the variable
// whose address is in the non-nil pointer out.
func Unmarshal(data []byte, out interface{}) (err error) {
	d := NewDecoder(bytes.NewReader(data))
	d.next() // get the first token
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", d.scan.Position, x)
		}
	}()
	read(d, reflect.ValueOf(out).Elem())
	return nil
}

//!-Unmarshal

//!+lexer
type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

//!-lexer

// The read function is a decoder for a small subset of well-formed
// S-expressions.  For brevity of our example, it takes many dubious
// shortcuts.
//
// The parser assumes
// - that the S-expression input is well-formed; it does no error checking.
// - that the S-expression input corresponds to the type of the variable.
// - that all numbers in the input are non-negative decimal integers.
// - that all keys in ((key value) ...) struct syntax are unquoted symbols.
// - that the input does not contain dotted lists such as (1 2 . 3).
// - that the input does not contain Lisp reader macros such 'x and #'x.
//
// The reflection logic assumes
// - that v is always a variable of the appropriate type for the
//   S-expression value.  For example, v must not be a boolean,
//   interface, channel, or function, and if v is an array, the input
//   must have the correct number of elements.
// - that v in the top-level call to read has the zero value of its
//   type and doesn't need clearing.
// - that if v is a numeric variable, it is a signed integer.

//!+read
func read(d *Decoder, v reflect.Value) {
	switch d.token {
	case scanner.Ident:
		// The only valid identifiers are
		// "nil" and struct field names.
		if d.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			d.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(d.text()) // NOTE: ignoring errors
		v.SetString(s)
		d.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(d.text()) // NOTE: ignoring errors
		v.SetInt(int64(i))
		d.next()
		return
	case '(':
		d.next()
		readList(d, v)
		d.next() // consume ')'
		return
	}
	panic(fmt.Sprintf("unexpected token %q", d.text()))
}

//!-read

//!+readlist
func readList(d *Decoder, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(d); i++ {
			read(d, v.Index(i))
		}

	case reflect.Slice: // (item ...)
		for !endList(d) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(d, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name value) ...)
		for !endList(d) {
			d.consume('(')
			if d.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", d.text()))
			}
			name := d.text()
			d.next()
			read(d, v.FieldByName(name))
			d.consume(')')
		}

	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(d) {
			d.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(d, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(d, value)
			v.SetMapIndex(key, value)
			d.consume(')')
		}

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(d *Decoder) bool {
	switch d.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

//!-readlist
