package poly

import (
	"fmt"
	"testing"
)

func TestParseEq(t *testing.T) {
	var tests = []struct {
		input string
		want  Poly
		err   error
	}{

		{
			input: "2 * X^2 * 3 * X^2 = 4 * X^2",
			want: Poly{
				degree:   2,
				eqString: "2 * X^2 * 3 * X^2 = 4 * X^2",
				terms: []term{
					{
						multip: 2,
						indet:  "X",
						power:  2,
					}, {
						multip: 3,
						indet:  "X",
						power:  2,
					}, {
						multip: 4,
						indet:  "X",
						power:  2,
					},
				},
				operator: []string{
					"*",
					"=",
				},
			},
			err: nil,
		},

		{
			input: "2 * X^2 * 3 * X^2",
			want: Poly{
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
			err: nil,
		},

		{
			input: "2 + X^2",
			want: Poly{
				degree:   0,
				eqString: "2 + X^2",
				terms:    []term{},
				operator: []string{},
			},
		},

		{
			input: "2 + X^",
			want: Poly{
				degree:   0,
				eqString: "2 + X^",
				terms:    []term{},
				operator: []string{},
			},
		},

		{
			input: "foo bar",
			want: Poly{
				degree:   0,
				eqString: "foo bar",
				terms:    []term{},
				operator: []string{},
			},
		},

		{
			input: "2 + X^2 * 3 * X^2",
			want: Poly{
				degree:   0,
				eqString: "2 + X^2 * 3 * X^2",
				terms:    []term{},
				operator: []string{},
			},
		},
	}

	for _, test := range tests {
		eq := Poly{}
		eq.ParseEq(test.input)
		if fmt.Sprintf("%+v", eq) != fmt.Sprintf("%+v", test.want) {
			t.Errorf("eq.ParseEq(\"%s\") gives us the struct\n%+v \n Expecting\n%+v\n", test.input, eq, test.want)
		}
	}
}
