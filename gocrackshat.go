package main

import (
	"flag"

	"github.com/dacoconuttman/gocrackshat/attacks"
)

func main() {
	hashPtr := flag.String("file", "hashes.txt", "File containing hashes to be cracked")
	attackPtr := flag.String("attack", "bruteforce", "Attack to be used (bruteforce, dictionary, mask)")
	maskPtr := flag.String("mask", "mask.txt", "Mask file. Uses hashcat masks, with support for dictionary words as well.")
	outFilePtr := flag.String("out", "dump.txt", "Output file.")
	lenPtr := flag.Int("len", 6, "Use only with a bruteforce attack. Specifies password length.")

	flag.Parse()

	switch *attackPtr {
	case "bruteforce":
		attacks.Bruteforce(*hashPtr, *outFilePtr, *lenPtr)
		break
	case "dictionary":
		attacks.Dictionary(*hashPtr, *outFilePtr)
		break
	case "mask":
		attacks.Mask(*hashPtr, *maskPtr, *outFilePtr)
		break
	default:
		panic("Please specify an attack type")
		break
	}

}
