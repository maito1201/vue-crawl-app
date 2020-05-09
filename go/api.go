package main

import (
  "strings"
  "time"
  "net/http"
  "encoding/json"
  "github.com/PuerkitoBio/goquery"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

type Item struct {
  Url string
  Name string
  Price string
  Image string
}

func main() {
  router := gin.Default()

  router.Use(cors.New(cors.Config{
    AllowMethods: []string{
      "POST",
      "GET",
    },
    AllowHeaders: []string{
      "Access-Control-Allow-Headers",
      "Content-Type",
      "Content-Length",
      "Accept-Encoding",
      "X-CSRF-Token",
      "Authorization",
    },
    AllowOrigins: []string{
      "http://app",
    },
    MaxAge: 24 * time.Hour,
  }))

  router.GET("/digimart", func(c *gin.Context) {
    keyword := c.Query("keyword")
    url := "http://www.digimart.net/search?keywordAnd=" + strings.Replace(keyword, " ", "+", -1)
    ret := scrapeDigimart(url)
    c.String(http.StatusOK, string(ret))
  })

  router.GET("/mercari", func(c *gin.Context) {
    keyword := c.Query("keyword")
    url := "https://www.mercari.com/jp/search/?keyword=" + strings.Replace(keyword, " ", "+", -1)
    ret := scrapeMercari(url)
    c.String(http.StatusOK, string(ret))
  })

  router.Run(":3000")
}

func scrapeDigimart(url string) string {
  items := []Item{}
  doc, err := goquery.NewDocument(url)
  if err != nil {
    return("")
  }

  selection := doc.Find("div.itemSearchBlock")
  selection.Each(func(index int, s *goquery.Selection) {
    url := "http://www.digimart.net/" + s.Find("p.ttl").Find("a").AttrOr("href", "")
    name := s.Find("p.ttl").Find("a").Text()
    price := s.Find("p.price").First().Text()
    image := s.Find("div.pic").Find("img").AttrOr("src", "")
    item := Item{ Url: url, Name: name, Price: price, Image: image }
    items = append(items, item)
  })

  json, _ := json.Marshal(items)
  return string(json)
}

func scrapeMercari(url string) string {
  items := []Item{}
  doc, err := goquery.NewDocument(url)
  if err != nil {
    return("")
  }

  selection := doc.Find("div.items-box-content.clearfix").Find("section.items-box")
  selection.Each(func(index int, s *goquery.Selection) {
    url := "https://www.mercari.com" + s.Find("a").AttrOr("href", "")
    name := s.Find("a").Find("h3.items-box-name").Text()
    price := s.Find("a").Find("div.items-box-price").Text()
    image := s.Find("a").Find("figure.items-box-photo").Find("img").AttrOr("data-src", "")
    item := Item{ Url: url, Name: name, Price: price, Image: image }
    items = append(items, item)
  })

  json, _ := json.Marshal(items)
  return string(json)
}