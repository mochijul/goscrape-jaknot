package main

import (
	"fmt"
	b "goscrape/mysql"
	s "goscrape/scrape"

	"github.com/gocolly/colly"
)

//inisialisasi table detail produk []
//insert into data detail produk jaknot if exist update

//Monitoring process
// get descending index by last_scraped_at => count( last_scraped_at < today)
// scrape priority index produk by last_scraped_at
// error link list

//Data integration
// get data for update

//Set and Get data tokopedia

//Set and Get data shopee

func main() {
	initScrape(false)

}

func initScrape(freshIndexProduct bool) {
	fmt.Println("Starting...")
	c := colly.NewCollector()
	b.RunExec("CREATE DATABASE IF NOT EXISTS goscrape", "Create DB")
	if freshIndexProduct {
		b.RunExec("Drop table if exists goscrape.index_produk", "Drop Table index_produk")
	}
	s.GetIndexProduct(c)
}
