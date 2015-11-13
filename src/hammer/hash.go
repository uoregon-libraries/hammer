package main

import (
	"encoding/hex"
	"log"
	"strings"
)

type Hash struct {
	Bytes []byte
	Ident string
}

func HashFromString(s string) *Hash {
	parts := strings.Fields(s)
	if len(parts) != 2 {
		log.Printf("Error decoding %s: %d parts (we need exactly two)", s, len(parts))
		return nil
	}

	bytes, err := hex.DecodeString(parts[0])
	if err != nil {
		log.Printf("Error decoding %s: %s", s, err)
		return nil
	}

	return &Hash{Bytes: bytes, Ident: parts[1]}
}
