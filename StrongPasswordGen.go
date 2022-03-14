package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var passwdLength int
var minNumbers int
var minCharacters int
var minSpecialChar int
var help bool

var SmollCharacters string = "abcdefghijklmnopqrstuwvxyz"
var LargeCharacters string = "ABCDEFGHIJKLMNOPQRSTUWVXYZ"
var CijferCharacters string = "123456789"
var SpecialCharacters string = "!@#$%^&*?"

func init() {
	flag.IntVar(&passwdLength, "l", 8, "Set the lenght of your password.")
	flag.IntVar(&minCharacters, "char", 0, "Minimal amount of characters in your password.")
	flag.IntVar(&minNumbers, "num", 0, "Minimal amount of numbers in your password.")
	flag.IntVar(&minSpecialChar, "special", 0, "Minimal amount of special characters in your password.")
	flag.BoolVar(&help, "h", false, "A list of all the flags that can be used.")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(111)
	}

	if passwdLength < 8 {
		fmt.Println("Please only use a password lenght of 8 or more.")
		os.Exit(111)
	}
}

func main() {
	c, e := createPassword()

	if e != nil {
		log.Fatal(e.Error())
	} else {
		fmt.Println("Your costume made password is:", c)
	}
}

func createPassword() (string, error) {
	rand.Seed(time.Now().UnixNano())
	pwdGenLength := passwdLength - 4

	minCharacters = minCharacters / 2

	SC := rand.Intn(pwdGenLength)
	LC := rand.Intn(pwdGenLength - SC)
	N := rand.Intn(pwdGenLength - (SC + LC))
	SPC := pwdGenLength - (SC + LC + N)

	F_SCharacters, e := randomCharacters(0, SC+1)
	F_LCharacters, e := randomCharacters(1, LC+1)
	F_Numbers, e := randomCharacters(2, N+1)
	F_SPCharacters, e := randomCharacters(3, SPC+1)

	password := F_SCharacters + F_LCharacters + F_Numbers + F_SPCharacters
	var passwordCharacters []string

	for _, c := range strings.Split(password, "") {
		passwordCharacters = append(passwordCharacters, c)
	}

	rand.Shuffle(len(passwordCharacters), func(i, j int) {
		passwordCharacters[i], passwordCharacters[j] = passwordCharacters[j], passwordCharacters[i]
	})

	password = ""

	for _, c := range passwordCharacters {
		password += c
	}

	if e != nil {
		return "", fmt.Errorf(e.Error())
	} else {
		return password, nil
	}
}
func randomCharacters(chartype int, amount int) (string, error) {
	var characters string

	switch chartype {
	case 0:
		for i := 0; i < amount; i++ {
			c, e := getCharacters(SmollCharacters)
			if e != nil {
				return "", fmt.Errorf(e.Error())
			} else {
				characters += c
			}
		}

	case 1:
		for i := 0; i < amount; i++ {
			c, e := getCharacters(LargeCharacters)
			if e != nil {
				return "", fmt.Errorf(e.Error())
			} else {
				characters += c
			}
		}
	case 2:
		for i := 0; i < amount; i++ {
			c, e := getCharacters(CijferCharacters)
			if e != nil {
				return "", fmt.Errorf(e.Error())
			} else {
				characters += c
			}
		}
	case 3:
		for i := 0; i < amount; i++ {
			c, e := getCharacters(SpecialCharacters)
			if e != nil {
				return "", fmt.Errorf(e.Error())
			} else {
				characters += c
			}
		}
	}

	return characters, nil
}

func getCharacters(chartype string) (string, error) {
	nChartype := rand.Intn(len(chartype))
	var tChartype []string

	for _, c := range strings.Split(chartype, "") {
		tChartype = append(tChartype, c)
	}

	if len(tChartype) == 0 {
		return "", fmt.Errorf("no characters are given")
	} else {
		return tChartype[nChartype], nil
	}
}
