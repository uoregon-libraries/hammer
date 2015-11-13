// Given a file of hashes, finds any similar items.  File must consist of hex
// digits separated by one or more spaces and then an identifier, such as what
// blockhash produces:
//
//     1cfbcc3b183b1823863780a780ff807f80f7c00860fe85ff3fff0ffc3e600000  one.jpg
//     3cfbdc7b383b002083a780a780ff807f8071c05aa5fd01ff01fe05f00400ffff  two.jpg
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/steakknife/hamming"
)

func usage() {
	fmt.Println("Usage: cat somefile | hammer")
}

func main() {
	raw := readStdin()
	if raw == nil {
		usage()
		os.Exit(1)
	}

	lines := strings.Split(string(raw), "\n")
	hashes := make([]*Hash, 0)
	hLen := -1
	for _, line := range lines {
		if line == "" {
			continue
		}
		hash := HashFromString(line)
		if hash == nil {
			continue
		}

		l := len(hash.Bytes)
		if hLen == -1 {
			hLen = l
		}
		if l != hLen {
			fmt.Println("All hashes must be exactly the same size!")
			usage()
			os.Exit(1)
		}

		hashes = append(hashes, hash)
	}

	if len(hashes) < 2 {
		fmt.Println("Need at least two valid hashes to compare")
		usage()
		os.Exit(1)
	}

	hl := len(hashes)
	hashBits := 8.0 * float64(hLen)
	for i := 0; i < hl; i++ {
		for j := i + 1; j < hl; j++ {
			a := hashes[i]
			b := hashes[j]
			diff := hamming.Bytes(a.Bytes, b.Bytes)
			similarity := 100.0 - (float64(diff) / hashBits * 100.0)
			fmt.Printf("Similarity between %s and %s: %g%% (%d out of %g bits differed)\n", a.Ident, b.Ident, similarity, diff, hashBits)
		}
	}

	fmt.Println("")
}

func readStdin() []byte {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("Error reading from stdin: %s\n\n", err)
			usage()
			os.Exit(1)
		}

		return bytes
	}

	return nil
}
