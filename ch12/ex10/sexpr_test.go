// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package sexpr

import (
	"reflect"
	"testing"
)

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

func TestEx03(t *testing.T) {
	want := "t"
	data, err := Marshal(true)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != want {
		t.Error("want: %s, actual: %s\n", want, data)
	}

	want = "nil"
	data, err = Marshal(false)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != want {
		t.Errorf("want: %s, actual: %s\n", want, data)
	}

	want = "2.46"
	data, err = Marshal(2.46)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != want {
		t.Errorf("want: %s, actual: %s\n", want, data)
	}

	want = "#C(1 2)"
	data, err = Marshal(1 + 2i)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != want {
		t.Errorf("want: %s, actual: %s\n", want, data)
	}

	var i interface{} = []int{1, 2, 3}
	want = "(\"[]int\" (1 2 3))"
	data, err = Marshal(&i)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != want {
		t.Errorf("want: %s, actual: %s\n", want, data)
	}
}

func TestEx08(t *testing.T) {
	data, err := Marshal(true)
	if err != nil {
		t.Fatal(err)
	}
	var outBool bool
	err = Unmarshal(data, &outBool)
	if err != nil {
		t.Fatal(err)
	}

	want := "nil"
	data, err = Marshal(false)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != want {
		t.Errorf("want: %s, actual: %s\n", want, data)
	}

	wantFloat := 2.46
	data, err = Marshal(wantFloat)
	if err != nil {
		t.Fatal(err)
	}
	var outFloat float64
	err = Unmarshal(data, &outFloat)
	if err != nil {
		t.Fatal(err)
	}
	if outFloat != wantFloat {
		t.Errorf("want: %g, actual: %g\n", wantFloat, outFloat)
	}

	wantComplex := 1 + 2i
	data, err = Marshal(wantComplex)
	if err != nil {
		t.Fatal(err)
	}
	var outComplex complex128
	err = Unmarshal(data, &outComplex)
	if err != nil {
		t.Fatal(err)
	}
	if outComplex != wantComplex {
		t.Errorf("want: %g, actual: %g\n", wantFloat, outFloat)
	}

	var i interface{} = []int{1, 2, 3}
	want = "(\"[]int\" (1 2 3))"
	data, err = Marshal(&i)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != want {
		t.Errorf("want: %s, actual: %s\n", want, data)
	}
}
