package main

import (
	"fmt"

	"github.com/svarlamov/goyhfin"
)

func main() {

	var Ativo string
	fmt.Println("Pesquise um ativo por nome:")
	fmt.Scan(&Ativo)

	resp, err := goyhfin.GetTickerData(Ativo, goyhfin.OneMonth, goyhfin.OneDay, false)
	if err != nil {
		fmt.Println("Error fetching Yahoo Finance data:", err)
		panic(err)
	}
	fmt.Println("| Preço mais alto   |         Data             |")
	for ind := range resp.Quotes {
		fmt.Println("|____________________|_________________________|")
		fmt.Println("|", resp.Quotes[ind].High, " | ", resp.Quotes[ind].OpensAt.Day(), "/", resp.Quotes[ind].OpensAt.Month(), "/", resp.Quotes[ind].OpensAt.Year())
	}
}
