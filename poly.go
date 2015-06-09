package poly

import (
	"errors"
	"strconv"
	"strings"
)

type Equa struct {
	Lhs Poly
	Rhs Poly
}

func (eq *Equa) ParseEq(str string) []error {
	var retError []error
	strs := strings.Split(str, "=")
	if len(strs) != 2 {
		return append(retError, errors.New("Error: input needs one \"=\" in equation to solve"))
	}
	errChan := make(chan error)
	go eq.Lhs.ParseEq(strings.Trim(strs[0], " "), errChan)
	go eq.Rhs.ParseEq(strings.Trim(strs[1], " "), errChan)
	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			retError = append(retError, err)
		}
	}
	return nil
}

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

func (eq *Poly) ParseEq(str string, errChan chan error) {
	eq.eqString = str
	termStrs := strings.Split(str, " ")
	for i := 0; i <= len(termStrs)-3; i++ {
		t := &term{}
		err := t.parseMul(termStrs[i])
		if err != nil {
			errChan <- err
		}
		err = t.checkMulOp(termStrs[i+1])
		if err != nil {
			errChan <- err
		}
		err = t.parsePow(termStrs[i+2])
		if err != nil {
			errChan <- err
		}
		i = i + 3
		if ((i+1)%4) == 0 && i != 0 && i < len(termStrs) {
			if isOperator(termStrs[i]) {
				eq.operator = append(eq.operator, termStrs[i])
			} else {
				errChan <- errors.New("exepcting operator, found " + termStrs[i])
			}
		}
		eq.terms = append(eq.terms, *t)
		eq.GetDegree(eq.terms)
	}
	errChan <- nil
}

type term struct {
	multip int
	indet  string
	power  int
}

func (t *term) parseMul(mulStr string) error {
	mulStr = strings.Trim(mulStr, " ")
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
	powStr = strings.Trim(powStr, " ")
	indPow := strings.Split(powStr, "^")
	if len(indPow) == 2 {
		t.indet = indPow[0]
		p, err := strconv.Atoi(indPow[1])
		if err != nil {
			return err
		}
		if p < 3 && p > -1 {
			t.power = p
		} else {
			return errors.New("degree too high\"X^2\"")
		}
	} else {
		return errors.New("argument not of form \"X^2\"")
	}
	return nil
}

func isOperator(s string) bool {
	ops := []string{"*", "/", "+", "-"}
	for _, v := range ops {
		if s == v {
			return true
		}
	}
	return false
}
