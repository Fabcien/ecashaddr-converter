package common

import "github.com/Fabcien/cashaddr-converter/address"

type Output struct {
	ECashAddr string `json:"ecashaddr,omitempty"`
	BitcoinCashAddr string `json:"bitcoincashaddr,omitempty"`
	Legacy   string `json:"legacy,omitempty"`
	Copay    string `json:"copay,omitempty"`
	Hash     string `json:"hash,omitempty"`
}

func GetAllFormats(addr *address.Address) (*Output, error) {
	ecashaddr, err := addr.ECashAddress()
	if err != nil {
		return nil, err
	}
	ecashaddrstr, err := ecashaddr.Encode()
	if err != nil {
		return nil, err
	}

	bitcoincashaddr, err := addr.BitcoinCashAddress()
	if err != nil {
		return nil, err
	}
	bitcoincashaddrstr, err := bitcoincashaddr.Encode()
	if err != nil {
		return nil, err
	}

	old, err := addr.Legacy()
	if err != nil {
		return nil, err
	}
	oldstr, err := old.Encode()
	if err != nil {
		return nil, err
	}

	copay, err := addr.Copay()
	if err != nil {
		return nil, err
	}
	copaystr, err := copay.Encode()
	if err != nil {
		return nil, err
	}

	return &Output{
		ECashAddr: ecashaddrstr,
		BitcoinCashAddr: bitcoincashaddrstr,
		Legacy:   oldstr,
		Copay:    copaystr,
		Hash:     addr.Hex(),
	}, nil
}
