package main

import (
  "strings"
  "time"
  "net/http"
  "encoding/json"
  "github.com/sclevine/agouti"
  "github.com/PuerkitoBio/goquery"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "fmt"
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

  router.GET("/yahoo", func(c *gin.Context) {
    keyword := c.Query("keyword")
    url := "https://auctions.yahoo.co.jp/search/search?p=" + strings.Replace(keyword, " ", "+", -1)
    ret := scrapeYahoo(url)
    c.String(http.StatusOK, string(ret))
  })

  router.GET("/reverb", func(c *gin.Context) {
    keyword := c.Query("keyword")
    url := "http://reverb.com/marketplace?query=" + strings.Replace(keyword, " ", "+", -1)
    ret := scrapeReverb(url)
    c.String(http.StatusOK, string(ret))
  })

  router.NoRoute(func (c *gin.Context) {
    c.String(http.StatusOK, "")
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

func scrapeYahoo(url string) string {
  items := []Item{}
  resp, err := http.Get(url)
  if err != nil {
    fmt.Printf("error: %v\n", err)
    return("")
  }
  doc, err := goquery.NewDocumentFromResponse(resp)
  if err != nil {
    fmt.Printf("error: %v\n", err)
    return("")
  }
  time.Sleep(time.Second * 10)
  selection := doc.Find("ul.Products__items").Find("li.Product")
  selection.Each(func(index int, s *goquery.Selection) {
    url := s.Find("div.Product__image").Find("a").AttrOr("href", "")
    name := s.Find("h3.Product__title").Text()
    price := s.Find("span.Product__priceValue").First().Text()
    image := s.Find("div.Product__image").Find("img").AttrOr("src", "")
    item := Item{ Url: url, Name: name, Price: price, Image: image }
    items = append(items, item)
  })
  json, _ := json.Marshal(items)
  fmt.Printf("item %v, json %s\n", items, json)
  return string(json)
}

func scrapeReverb(url string) string {
  items := []Item{}
  driver := agouti.ChromeDriver(
    agouti.ChromeOptions("args", []string{
        "--headless",
        "--window-size=30,120",
        "--disable-gpu",                        // ref: https://developers.google.com/web/updates/2017/04/headless-chrome#cli
        "no-sandbox",                           // ref: https://github.com/theintern/intern/issues/878
        "disable-dev-shm-usage",                // ref: https://qiita.com/yoshi10321/items/8b7e6ed2c2c15c3344c6
    }),
  )

  driver.Start()
  defer driver.Stop()
  page, _ := driver.NewPage(agouti.Browser("chrome"))
  page.Navigate(url)
  // 描画の完了を待機
  time.Sleep(7 * time.Second)
  content, _ := page.HTML()
  reader := strings.NewReader(content)
  doc, _ := goquery.NewDocumentFromReader(reader)
  selection := doc.Find("ul.tiles.tiles--four-wide-max").Find("li.tiles__tile")
  selection.Each(func(index int, s *goquery.Selection) {
    url := s.Find("a").AttrOr("href", "")
    name := s.Find("h4.grid-card__title, h3.csp-square-card__title").Text()
    price := s.Find("span.price-display, div.csp-square-card__details__price").Text()
    image := s.Find("img").AttrOr("src", "")
    item := Item{ Url: url, Name: name, Price: price, Image: image }
    if name != "" {
      items = append(items, item)
    }
  })
  json, _ := json.Marshal(items)
  return string(json)
}
