package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Function for converting lowercase letters to uppercase and vice versa
// and remove all symbol:^%$#@#$.... Except: '!'
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

// OpenRead file to string
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
	filename := "data.json"
	directory := "../les6"
	_, err := os.Open(directory + "/" + filename)
	if os.IsNotExist(err) {
		fmt.Printf("The file %s does not exist %s\n", filename, directory)
	} else {
		fmt.Printf("The file %s exists %s\n", filename, directory)
	}
}

func ReadLinebyLineFile() {
	data, err := os.Open("newTextfile.txt")
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
}

func SearchTextInFile() {

	file, err := os.Open("newTextfile.txt")
	check(err)
	defer file.Close()

	someText := "you"
	scanner := bufio.NewScanner(file)
	var count int
	found := false
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), someText) {
			found = true
			count++
		}
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}

	if found {
		fmt.Println("In file:", file.Name(), "Was found text: `",
			someText, "`", count)
	} else {
		fmt.Println("Text `", someText, "` not in file")
	}
}

func BinaryFile() {

	dataBinar := []byte("Data to binary file\nfsdfsfsdf \nfixi dixi\n")
	err := os.WriteFile("Batfile", dataBinar, 0644)
	check(err)

	fopen, err := os.Open("Batfile")
	if err != nil {
		panic(err)
	}
	defer fopen.Close()

	reader := bufio.NewScanner(fopen)
	for reader.Scan() {
		b := reader.Text()
		fmt.Println(b)
	}
}

//Advanced task for file

func main() {

	//BinaryFile()
	//SearchTextInFile()
	//ReadLinebyLineFile()
	//CheckAvailabilityFile()
	//DeletedFile()
	//RenameFile()
	//CopyFile()
	//CreateFile()
	//WriteFile()
	//ReadFile()
	//OpenFile()
	//LineFilter()
}
