package main

import (
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateFile() {

	data, err := os.Create("newTextfile.txt")
	check(err)
	str := `For more granular writes, open a file for writing.\n
		It’s idiomatic to defer a Close immediately after opening a file.\n
		defer f.Close()\n
		You can Write byte slices as you’d expect.`

	data.WriteString(str)
}

func ReadFile() {
	data, err := os.ReadFile("newTextfile.txt")
	check(err)
	fmt.Println(string(data))
}

func WriteFile() {
	data, err := os.OpenFile("newTextfile.txt", os.O_WRONLY|os.O_APPEND, 0644)
	check(err)
	defer data.Close()
	data.WriteString("\nhello i remember to add paswd:2222 and email:filson@gmail.com")
}

func CopyFile() {
	//Open initial file
	newdat, err := os.Open("newTextfile.txt")
	check(err)
	// Open another file and we grant file rights
	copt, err := os.OpenFile("copyinffile.txt", os.O_WRONLY, 0666)
	check(err)
	_, err = io.Copy(copt, newdat)
	check(err)

	defer newdat.Close()
	defer copt.Close()

}

func RenameFile() {
	err := os.Chmod("copyinffile.txt", 0777)
	check(err)
	err = os.Rename("copyinffile.txt", "newrenamefile.txt")
	check(err)
}

func DeletedFile() {
	err := os.Chmod("newrenamefile.txt", 0777)
	check(err)
	err = os.Remove("newrenamefile.txt")
	check(err)
}

func CheckAvailabilityFile() {

}

func main() {

	//DeletedFile()
	//RenameFile()
	//CopyFile()
	//CreateFile()
	//WriteFile()
	//ReadFile()
}

// Function for converting lowercase letters to uppercase and vice versa
// and remove all symbol:^%$#@#$.... Except: '!'
/*
func LineFilter() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		runes := []rune(scanner.Text())

		for i, r := range runes {
			if unicode.IsLetter(r) {

				if unicode.IsUpper(r) {
					runes[i] = unicode.ToLower(r)
				} else if unicode.IsLower(r) {
					runes[i] = unicode.ToUpper(r)
				}
			} else if runes[i] == '!' {
				continue
			} else {
				runes[i] = ' '
			}
		}

		str := string(runes)
		CleanString := strings.Join(strings.Fields(str), " ")
		fmt.Println(CleanString)

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
}
*/

//OpenRead file to string
/*
func OpenFile() {

	f, err := os.Open("testino.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	str := make([]byte, 2048)
	_, err = f.Read(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str[6]))
}
*/
