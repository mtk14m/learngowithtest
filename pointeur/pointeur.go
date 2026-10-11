package pointeur

import (
	"errors"
	"fmt"
)

type Portefeuille struct {
	solde Bitcoin
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (p *Portefeuille) Deposer(montant Bitcoin) {
	p.solde += montant
}

var ErrSoldeInsuffisant = errors.New("le solde du portefeuille est insufissant")

func (p *Portefeuille) Retirer(montant Bitcoin) error {
	if p.solde < montant {
		return ErrSoldeInsuffisant
	}
	p.solde -= montant
	return nil

}

func (p *Portefeuille) Solde() Bitcoin {
	return p.solde
}
