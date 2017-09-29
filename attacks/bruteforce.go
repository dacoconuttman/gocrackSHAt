package attacks

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")

//Bruteforce is a bruteforce attack
func Bruteforce(hashFile string, outFile string, len int) {
	foundHashes := make(map[string]string)
	hashMap := make(map[string]string)
	fmt.Println("we in Bruteforce fam")
	makeMap(len)
	hashes := read(hashFile)
	for _, elem := range hashes {
		if val, exists := hashMap[elem]; exists {
			foundHashes[elem] = val
			fmt.Printf("hash: %s, password: %s\n", val, elem)
		}
	}
}

func read(hashFile string) (hashes []string) {
	file, err := ioutil.ReadFile(hashFile)
	if err != nil {
		panic(err)
	}
	hashes = strings.Split(string(file), "\n")
	return
}

func generatePasswords(maxLen int) []string {

	c := make(chan []interface{})
	for i := 1; i <= maxLen; i++ {

	}
}
