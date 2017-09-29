package attacks

import (
	"fmt"
)

type hashTuple struct {
	plaintext string
	encrypted [32]byte
}

// func fuck(filename string) {
// 	file, err := os.Open(filename)
// }

//Test: test
func Test() {
	fmt.Println("Hieeee") //test
}

//Dictionary: Runs a dictionary attack
func Dictionary(hashFile string, outFile string) {
	fmt.Println("we in Dictionary fam")
}