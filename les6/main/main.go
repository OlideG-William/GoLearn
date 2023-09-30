package main

import (
	"fmt"
	"regexp"
	"time"
)

type interfProduct interface {
	Regexp() string
}

type Product struct {
	NameProduct string
	idProduct   int
	brand_id    string
	Status_id   bool
}

type Client struct {
	Id        int
	FirstName string
	Gender    string
	Contact   string
	Email     string
}

func (p *Client) Regexp() string {
	rNameReg := regexp.MustCompile("[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+")
	TempEmail := p.Email
	p.Email = "Hello pitu gime me your Email " + TempEmail
	FindEmail_str := rNameReg.FindString(p.Email)
	return FindEmail_str
}

func (p *Product) Regexp() string {
	r_nameP, _ := regexp.Compile("name: ([^,]+)")
	p.NameProduct = "your Name: Fixianator but anyway you lost name: LG_Book, Phone: 1234567890"
	r_compli_regex := r_nameP.FindString(p.NameProduct)
	return r_compli_regex
}

func main() {
	start := time.Now()
	fmt.Println()

	namePr := interfProduct(&Product{
		NameProduct: "Iphone",
		idProduct:   398722,
		brand_id:    "Apple",
		Status_id:   true,
	})

	nameCli := interfProduct(&Client{
		Id:        22232,
		FirstName: "Anatoly",
		Gender:    "Men",
		Contact:   "+380956783412",
		Email:     "Faxionds@gmail.com",
	})

	fmt.Println(namePr.Regexp())
	fmt.Println(nameCli.Regexp())

	fmt.Println("\n", time.Since(start))
}
