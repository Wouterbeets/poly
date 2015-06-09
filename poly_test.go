package poly

import (
	"errors"
	"fmt"
	"testing"
)

func TestParsePow(t *testing.T) {
	var tests = []struct {
		input string
		want  int
		err   error
	}{
		{
			input: "X^2",
			want:  2,
			err:   nil,
		},
		{
			input: "X^3",
			want:  0,
			err:   errors.New("degree too high\"X^2\""),
		},
		{
			input: "2",
			want:  0,
			err:   errors.New("argument not of form \"X^2\""),
		},
		{
			input: "",
			want:  0,
			err:   errors.New("argument not of form \"X^2\""),
		},
		{
			input: "-2",
			want:  0,
			err:   errors.New("argument not of form \"X^2\""),
		},
		{
			input: "3",
			want:  0,
			err:   errors.New("argument not of form \"X^2\""),
		},
		{
			input: " 3  ",
			want:  0,
			err:   errors.New("argument not of form \"X^2\""),
		},
		{
			input: "foo",
			want:  0,
			err:   errors.New("argument not of form \"X^2\""),
		},
	}
	for _, test := range tests {
		term := term{}
		err := term.parsePow(test.input)
		if test.want != term.power {
			t.Errorf("input = \n%+v\nexpected output \n%+v\n got \n%+v\n", test.input, test.want, term.power)
		}
		if (test.err == nil && err == nil) || (test.err != nil && err != nil) {
		} else {
			t.Errorf("errors not alike\ninput = \n%+v\nexpected output \n%+v\n got \n%+v\n", test.input, test.err, err)
		}
	}
}

func TestParseMul(t *testing.T) {
	var tests = []struct {
		input string
		want  int
		err   error
	}{
		{
			input: "3",
			want:  3,
			err:   nil,
		},
		{
			input: "",
			want:  0,
			err:   errors.New(""),
		},
		{
			input: "-3",
			want:  -3,
			err:   nil,
		},
		{
			input: " 3  ",
			want:  3,
			err:   nil,
		},
		{
			input: "foo",
			want:  0,
			err:   errors.New(""),
		},
	}
	for _, test := range tests {
		term := term{}
		err := term.parseMul(test.input)
		if test.want != term.multip {
			t.Errorf("input = \n%+v\nexpected output \n%+v\n got \n%+v\n", test.input, test.want, term.multip)
		}
		if (test.err == nil && err != nil) || (test.err != nil && err == nil) {
			t.Errorf("errors not alike\ninput = \n%+v\nexpected output \n%+v\n got \n%+v\n", test.input, test.err, err)
		}
	}
}

func TestParseEq(t *testing.T) {
	var tests = []struct {
		input string
		want  Equa
		err   error
	}{

		{
			input: "2 * X^2 * 3 * X^2 = 4 * X^2",
			want: Equa{
				Lhs: Poly{
					degree:   2,
					eqString: "2 * X^2 * 3 * X^2",
					terms: []term{
						{
							multip: 2,
							indet:  "X",
							power:  2,
						}, {
							multip: 3,
							indet:  "X",
							power:  2,
						},
					},
					operator: []string{
						"*",
					},
				},
				Rhs: Poly{
					degree:   2,
					eqString: "4 * X^2",
					terms: []term{
						{
							multip: 4,
							indet:  "X",
							power:  2,
						},
					},
					operator: []string{},
				},
			},
			err: nil,
		},

		{
			input: "2 * X^2 k 3 * X^2 = 4 - X^2",
			want: Equa{
				Lhs: Poly{
					degree:   0,
					eqString: "2 * X^2 k 3 * X^2 = 4 - X^2",
					terms:    []term{},
					operator: []string{},
				},
				Rhs: Poly{
					degree:   0,
					eqString: "4 - X^2",
					terms:    []term{},
					operator: []string{},
				},
			},
			err: errors.New("exepcting operator, found \"k\""),
		},
	}
	for _, test := range tests {
		eq := Equa{}
		err := eq.ParseEq(test.input)
		if fmt.Sprintf("%+v", eq) != fmt.Sprintf("%+v", test.want) {
			t.Errorf("eq.ParseEq(\"%s\") gives us the struct\n%+v\n Expecting\n%+v\n err = \n%+v\n, test.err = \n%+v\n", test.input, eq, test.want, err, test.err)
		}
	}
}
