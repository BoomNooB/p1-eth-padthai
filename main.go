package main

import (
	"errors"
	"log"
)

var ledger map[string]map[string]int

func main() {
	// init example ledger with 3 addresses
	ledger = map[string]map[string]int{
		"0x123": {
			"tokenA": 100,
			"tokenB": 200,
		},
		"0xabc": {
			"tokenA": 0,
			"tokenC": 150,
		},
		"someAddress": {
			"tokenA": 10,
		},
	}
	log.Printf("Ledger init: %v", ledger)

	// if we try to transfer 50 tokenB from 0x123 to someAddress it should be error someAddress
	err := transfer("0x123", "someAddress", "tokenC", 50)
	if err != nil {
		log.Println("Transfer failed")
	}
	log.Printf("Current ledger: %v", ledger)

	// inefficient balance case
	err = transfer("0xabc", "0x123", "tokenA", 50)
	if err != nil {
		log.Println("Transfer failed")
	}
	log.Printf("Current ledger: %v", ledger)

	// now is success case
	err = transfer("someAddress", "0xabc", "tokenA", 7)
	if err != nil {
		log.Println("Transfer failed")
	}
	log.Printf("Current ledger: %v", ledger)
}

func transfer(sender_address, recipient_address, asset_id string, amount int) error {
	/* Validation state */

	// check if sender has enough balance asset
	if ledger[sender_address][asset_id] < amount {
		log.Println("[Transfer func] Insufficient balance")
		return errors.New("Insufficient balance")
	}

	// check if recipient address are same as sender address
	if sender_address == recipient_address {
		log.Println("[Transfer func] Cannot transfer to same address")
		return errors.New("Cannot transfer to same address")
	}

	// check if asset_id is exist both sender and recipient
	_, isExist := ledger[sender_address][asset_id]
	if !isExist {
		log.Println("[Transfer func] Sender Asset not found")
		return errors.New("Asset not found")
	}

	_, isExist = ledger[recipient_address][asset_id]
	if !isExist {
		log.Println("[Transfer func] Recipent Asset not found")
		return errors.New("Asset not found")
	}

	/* Transfer state */
	ledger[sender_address][asset_id] -= amount
	ledger[recipient_address][asset_id] += amount

	return nil
}
