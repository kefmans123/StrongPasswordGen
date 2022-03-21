package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/dlclark/regexp2"
)

var passwdLength int
var minNumbers int
var minCharacters int
var minSpecialChar int
var help bool

var sC int
var lC int
var n int
var spC int

var checkpass string

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
	flag.StringVar(&checkpass, "check", "", "Check your current password if its secure.")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if checkpass == "" {
		if passwdLength < 8 {
			fmt.Println("Please only use a password lenght of 8 or more.")
			os.Exit(0)
		}

		if minCharacters+minNumbers+minSpecialChar > passwdLength {
			fmt.Println("Your minimum requirements are too high for this current lenght. You are trying to use:", minCharacters+minNumbers+minSpecialChar, "characters.\n But your length is only:", passwdLength, "Please use a higher length or make your requirements smaller.")
			os.Exit(0)
		}
	}

}

func main() {
	if checkpass != "" {
		check, _ := CheckPassword(checkpass)

		if check {
			fmt.Printf("Your password is secure.")

		} else {
			fmt.Println("Your password is not secure.")

		}
		os.Exit(0)

	}
	c, e := createPassword()

	if e != nil {
		log.Fatal(e.Error())
	} else {
		fmt.Println("Your costume made password is:", c)
	}
}

func createPassword() (string, error) {
	rand.Seed(time.Now().UnixNano())

	var F_sCharacters string
	var F_lCharacters string
	var F_Numbers string
	var F_spCharacters string

	var e error

	totalMinimal := minCharacters + minNumbers + minSpecialChar
	pwdGenLength := (passwdLength - 4) - totalMinimal

	if totalMinimal > 0 {

		for i := 0; i < pwdGenLength-4; i++ {
			x := rand.Intn(4)

			switch x {

			case 0:
				minCharacters++

			case 1:
				minCharacters++

			case 2:
				minNumbers++
			case 3:
				minSpecialChar++
			}
		}

		if minCharacters > 2 {
			sC++
			lC++

			for i := 0; i < minCharacters-2; i++ {

				x := rand.Intn(2)

				switch x {

				case 0:
					sC++
				case 1:
					lC++
				}
			}
		} else {
			for i := 0; i < minCharacters; i++ {

				x := rand.Intn(2)

				switch x {

				case 0:
					sC++
				case 1:
					lC++
				}
			}
		}

		n = minNumbers
		spC = minSpecialChar

		sC, e = countInt(sC)
		if e != nil {
			return "", fmt.Errorf(e.Error())
		}
		lC, e = countInt(lC)
		if e != nil {
			return "", fmt.Errorf(e.Error())
		}
		n, e = countInt(n)
		if e != nil {
			return "", fmt.Errorf(e.Error())
		}
		spC, e = countInt(spC)
		if e != nil {
			return "", fmt.Errorf(e.Error())
		}

		for i := sC + lC + n + spC; i < passwdLength; i++ {
			x := rand.Intn(4)

			switch x {

			case 0:
				sC++

			case 1:
				lC++

			case 2:
				n++
			case 3:
				spC++
			}
		}

		F_sCharacters, e = randomCharacters(0, sC)
		F_lCharacters, e = randomCharacters(1, lC)
		F_Numbers, e = randomCharacters(2, n)
		F_spCharacters, e = randomCharacters(3, spC)

	} else {
		for i := 0; i < pwdGenLength; i++ {
			x := rand.Intn(4)

			switch x {

			case 0:
				sC++

			case 1:
				lC++

			case 2:
				n++
			case 3:
				spC++
			}
		}

		F_sCharacters, e = randomCharacters(0, sC+1)
		F_lCharacters, e = randomCharacters(1, lC+1)
		F_Numbers, e = randomCharacters(2, n+1)
		F_spCharacters, e = randomCharacters(3, spC+1)

	}

	password := F_sCharacters + F_lCharacters + F_Numbers + F_spCharacters
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

func countInt(interger int) (int, error) {
	if interger == 0 {
		if sC+lC+n+spC >= passwdLength {
			return interger, fmt.Errorf("Your password does not meet the required security standard. Please include atleast 1 small character, 1 capital character, 1 number and 1 special character.")
		} else {
			interger++
			return interger, nil
		}
	} else {
		return interger, nil
	}
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

func CheckPassword(password string) (bool, error) {
	regex, _ := regexp2.Compile("(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*?\\]|]).{8,}$", 0)

	match, _ := regex.MatchString(password)

	return match, nil
}
