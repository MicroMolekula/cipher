package rsa

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

type Key struct {
	Kn *big.Int
	N  *big.Int
}

func Test() {
	fmt.Println(searchD(big.NewInt(17), big.NewInt(2651394840)).Text(10))
}

func Code(number int, key *Key) *big.Int {
	num := big.NewInt(int64(number))
	cip := num.Exp(num, key.Kn, nil)
	cip.Mod(cip, key.N)
	return cip
}

func Decode(number *big.Int, key *Key) *big.Int {
	return number.Exp(number, key.Kn, key.N)
}

func GenerateKeys(bits int) (*Key, *Key, error) {
	p, err := generatePrime(bits)
	if err != nil {
		return nil, nil, err
	}
	q, err := generatePrime(bits)
	if err != nil {
		return nil, nil, err
	}
	if p.Text(10) == q.Text(10) {
		p, err = generatePrime(bits)
		if err != nil {
			return nil, nil, err
		}
	}
	fmt.Println("P и Q:", p.Text(10), q.Text(10))
	n := big.NewInt(1)
	n = n.Mul(p, q)
	fmt.Println("N:", n.Text(10))
	subP := big.NewInt(1).Sub(p, big.NewInt(1))
	subQ := big.NewInt(1).Sub(q, big.NewInt(1))
	fn := big.NewInt(1).Mul(subP, subQ)
	fmt.Println("Fn:", fn.Text(10))
	e := searchE(fn)
	if e == nil {
		return nil, nil, errors.New("error generate e")
	}
	fmt.Println("E:", e.Text(10))
	d := searchD(e, fn)
	if d == nil {
		return nil, nil, errors.New("error generate d")
	}
	fmt.Println("D:", d.Text(10))
	keyEnc := &Key{
		Kn: e,
		N:  n,
	}
	keyDec := &Key{
		Kn: d,
		N:  n,
	}
	return keyEnc, keyDec, nil
}

func generatePrime(bits int) (*big.Int, error) {
	prime, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	return prime, nil
}

func searchE(fn *big.Int) *big.Int {
	for i := 0; ; i++ {
		exp := big.NewInt(2)
		exp = exp.Exp(exp, big.NewInt(int64(i)), nil)
		e := big.NewInt(2)
		e = e.Exp(e, exp, nil)
		eR := *e.Add(e, big.NewInt(1))
		if e.GCD(nil, nil, e, fn).Text(10) == big.NewInt(1).Text(10) {
			return &eR
		}
		if e.Cmp(fn) == 1 {
			break
		}
	}
	return nil
}

func searchD(e, fn *big.Int) *big.Int {
	return big.NewInt(1).ModInverse(e, fn)
}
