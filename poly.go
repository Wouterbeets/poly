package poly

import (
	"errors"
	"strconv"
	"strings"
)

type Poly struct {
	degree   int
	eqString string
	terms    []term
	operator []string
}

func (eq *Poly) GetDegree(terms []term) {
	for _, v := range terms {
		if eq.degree < v.power {
			eq.degree = v.power
		}
	}
}

func (eq *Poly) ParseEq(str string) error {
	eq.eqString = str
	termStrs := strings.Split(str, " ")
	for i := 0; i < len(termStrs); i++ {
		t := &term{}
		err := t.parseMul(termStrs[i])
		if err != nil {
			return err
		}
		err = t.checkMulOp(termStrs[i+1])
		if err != nil {
			return err
		}
		err = t.parsePow(termStrs[i+2])
		if err != nil {
			return err
		}
		i = i + 3
		if ((i+1)%4) == 0 && i != 0 && i < len(termStrs) {
			eq.operator = append(eq.operator, termStrs[i])
		}
		eq.terms = append(eq.terms, *t)
		eq.GetDegree(eq.terms)
	}
	return nil
}

type term struct {
	multip int
	indet  string
	power  int
}

func (t *term) parseMul(mulStr string) error {
	mul, err := strconv.Atoi(mulStr)
	if err != nil {
		return err
	}
	t.multip = mul
	return nil
}

func (t *term) checkMulOp(opStr string) error {
	if opStr == "*" {
		return nil
	} else {
		return errors.New("operator before indeterminate must be \"*\"")
	}
}

func (t *term) parsePow(powStr string) error {
	indPow := strings.Split(powStr, "^")
	t.indet = indPow[0]
	p, err := strconv.Atoi(indPow[1])
	if err != nil {
		return err
	}
	t.power = p
	return nil
}
