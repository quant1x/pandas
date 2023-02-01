package tests

import "testing"

func TestAbstract(t *testing.T) {
	m := &Man{}
	m.Eat()
	m.Run()
	m.Sleep()

}

func TestAbstract2(t *testing.T) {
	akita := NewAkita()
	akita.Eat()

	labrador := NewLabrador()
	labrador.Eat()
}
