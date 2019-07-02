package sdk

import (
	"math/big"
	"strconv"
	"strings"
)

type Int struct {
	*big.Int `json:"int"`
}

var (
	ZeroInt = NewInt(0)
	OnInt   = NewInt(1)
	E18Int  = NewInt(1e18)
)

func NewInt(x int64) Int {
	return Int{big.NewInt(x)}
}

func NewIntFromBigInt(i *big.Int) Int {
	return Int{i}
}

func NewIntFromString(s string) (res Int, ok bool) {
	s = strings.ToLower(s)
	if strings.Index(s, "e+") >= 0 {
		arr := strings.Split(s, "e+")
		b, _ := strconv.Atoi(arr[1])
		s = arr[0] + strings.Repeat("0", b)
	}

	i, ok := new(big.Int).SetString(s, 0)
	if !ok {
		return
	}

	return Int{i}, true
}

func add(i, i2 *big.Int) *big.Int { return new(big.Int).Add(i, i2) }

func sub(i, i2 *big.Int) *big.Int { return new(big.Int).Sub(i, i2) }

func mul(i, i2 *big.Int) *big.Int { return new(big.Int).Mul(i, i2) }

func div(i, i2 *big.Int) *big.Int { return new(big.Int).Div(i, i2) }

func equal(i, i2 *big.Int) bool { return i.Cmp(i2) == 0 }

func gt(i, i2 *big.Int) bool { return i.Cmp(i2) == 1 }

func lt(i, i2 *big.Int) bool { return i.Cmp(i2) == -1 }

func neg(i *big.Int) *big.Int { return new(big.Int).Neg(i) }

func abs(i *big.Int) *big.Int { return new(big.Int).Abs(i) }

func (i Int) Add(i2 Int) Int {
	return Int{add(i.Int, i2.Int)}
}

func (i Int) Sub(i2 Int) Int {
	return Int{sub(i.Int, i2.Int)}
}

func (i Int) Mul(i2 Int) Int {
	return Int{mul(i.Int, i2.Int)}
}

func (i Int) Div(i2 Int) Int {
	return Int{div(i.Int, i2.Int)}
}

func (i Int) MulRat(r Rat) Int {
	return Int{div(mul(i.Int, r.Num()), r.Denom())}
}

// Equal compares two Ints
func (i Int) Equal(i2 Int) bool {
	return equal(i.Int, i2.Int)
}

func (i Int) GT(i2 Int) bool {
	return gt(i.Int, i2.Int)
}

func (i Int) GTE(i2 Int) bool {
	return !lt(i.Int, i2.Int)
}

func (i Int) LT(i2 Int) bool {
	return lt(i.Int, i2.Int)
}

func (i Int) LTE(i2 Int) bool {
	return !gt(i.Int, i2.Int)
}

func (i Int) Neg() Int {
	return Int{neg(i.Int)}
}

func (i Int) Abs() Int {
	return Int{abs(i.Int)}
}
