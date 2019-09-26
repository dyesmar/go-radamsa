// The goradamsa command is a test driver for the Radamsa package. It is named
// goradamsa so as not to interfere with radamsa, if you have that installed.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yesmar/radamsa"
)

var (
	cmd     = filepath.Base(os.Args[0])
	release string
)

func main() {
	seed := flag.Int64("seed", time.Now().UTC().UnixNano(), "PRNG seed")
	inplace := flag.Bool("inplace", false, "fuzz in place")
	iteration := flag.Int("iter", 0, "zero in on specified iteration")
	n := flag.Int("n", 1, "number of iterations to run")
	verbose := flag.Bool("v", false, "display superfluous detail")
	version := flag.Bool("version", false, "display version information")
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stderr, "%s %s © 2019 Ramsey Dow\n", cmd, release)
		os.Exit(0)
	}

	if *iteration < 0 {
		fmt.Fprintf(os.Stderr, "%s: iteration cannot be negative\n", cmd)
		os.Exit(0)
	}

	if *n < 1 {
		fmt.Fprintf(os.Stderr, "%s: n cannot be negative\n", cmd)
		os.Exit(0)
	}

	if *iteration > *n {
		fmt.Fprintf(os.Stderr, "%s: iteration cannot be greater than n\n", cmd)
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-seed <seed>] [-inplace]\n", cmd)
		os.Exit(1)
	}

	r := radamsa.New(radamsa.WithSeed(*seed), radamsa.InPlace(*inplace))
	xib := []byte(strings.Join(flag.Args()[:], " "))
	xob := make([]byte, 8192)

	for i := 0; i < *n; i++ {
		n, err := r.Fuzz(xib, len(xib), xob, len(xob))
		if err != nil {
			log.Println(err)
			continue
		}
		if *iteration == 0 || *iteration == i+1 {
			if *verbose {
				fmt.Fprintf(os.Stderr, "--> seed %d, iteration %d, %d bytes:\n",
					r.Seed(), r.Iteration(), n)
			}
			fmt.Printf("%s\n", string(xob[:n]))
		}
	}
}
