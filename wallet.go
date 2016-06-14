// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factom

import (
	"encoding/json"
)

func GenerateFactoidAddress() (*FactoidAddress, error) {
	type addressResponse struct {
		Public string `json:"public"`
		Secret string `json:"secret"`
	}
	
	req := NewJSON2Request("generate-factoid-address", apiCounter(), nil)
	resp, err := walletRequest(req)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}

	a := new(addressResponse)
	if err := json.Unmarshal(resp.JSONResult(), a); err != nil {
		return nil, err
	}
	f, err := GetFactoidAddress(a.Secret)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func GenerateECAddress() (*ECAddress, error) {
	type addressResponse struct {
		Public string `json:"public"`
		Secret string `json:"secret"`
	}
	
	req := NewJSON2Request("generate-ec-address", apiCounter(), nil)
	resp, err := walletRequest(req)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}

	a := new(addressResponse)
	if err := json.Unmarshal(resp.JSONResult(), a); err != nil {
		return nil, err
	}
	e, err := GetECAddress(a.Secret)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func ImportAddresses(addrs ...string) error {
	params := new(importRequest)
	for _, addr := range addrs {
		s := secretRequest{Secret: addr}
		params.Addresses = append(params.Addresses, s)
	}
	req := NewJSON2Request("import-addresses", apiCounter(), params)
	resp, err := walletRequest(req)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return resp.Error
	}
	
	return nil
}

func ListAddresses() ([]string, error) {
	type addressResponse struct {
		Public string `json:"public"`
		Secret string `json:"secret"`
	}
	
	type multiAddressResponse struct {
		Addresses []*addressResponse `json:"addresses"`
	}
	
	req := NewJSON2Request("all-addresses", apiCounter(), nil)
	resp, err := walletRequest(req)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}

	r := make([]string, 0)
	as := new(multiAddressResponse)
	if err := json.Unmarshal(resp.JSONResult(), as); err != nil {
		return nil, err
	}
	for _, adr := range as.Addresses {
		r = append(r, adr.Public)
	}

	return r, nil
}

//func GenerateFactoidAddressFromMnemonic(name string, mnemonic string) (string, error) {
//	name = strings.TrimSpace(name)
//	params := GenerateAddressFromPrivateKeyRequest{Name: name, Mnemonic: mnemonic}
//
//	req := NewJSON2Request("factoid-generate-address-from-token-sale", apiCounter(), params)
//	resp, err := walletRequest(req)
//	if err != nil {
//		return "", err
//	}
//	if resp.Error != nil {
//		return "", resp.Error
//	}
//
//	add := new(GenerateAddressResponse)
//	if err := json.Unmarshal(resp.JSONResult(), add); err != nil {
//		return "", err
//	}
//
//	return add.Address, nil
//}
//
///*
//func DnsBalance(addr string) (int64, int64, error) {
//	fct, ec, err := ResolveDnsName(addr)
//	if err != nil {
//		return 0, 0, err
//	}
//
//	f, err1 := FctBalance(fct)
//	e, err2 := ECBalance(ec)
//	if err1 != nil || err2 != nil {
//		return f, e, fmt.Errorf("%s\n%s\n", err1, err2)
//	}
//
//	return f, e, nil
//}
//
//
//*/
