package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/PuerkitoBio/goquery"
)
func main(){
  res, err := http.Get("https://www.kitapsec.com/mobil/49-gunde-tyt-matematik-kampi-eko-tonguc-akademi_urn613687.html")
  if err != nil {
    log.Fatal(err)
    return
  }
  doc, err := goquery.NewDocumentFromReader(res.Body)
  price := doc.Find(".dty_urunFiyati").Text()
		fmt.Printf("Review: %s\n",price)
}
