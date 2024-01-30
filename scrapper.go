package main
import(
	"github.com/gocolly/colly"
	"fmt"
	"encoding/json"
)
type Item struct{
	Link string `json:"link"`
	Name string `json:"name"`
	Price string `json:"price"`
	Instock string `json:"instock"`
}
func main()  {
	items:=[]Item{}
	c:=colly.NewCollector(colly.Async(true))
	c.OnHTML("div.side_categories li ul li", func(h *colly.HTMLElement) {
      c.Visit(h.Request.AbsoluteURL(h.ChildAttr("a", "href")))
	})
	c.OnHTML("li.next a", func(h *colly.HTMLElement) {
		c.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	})
	c.OnHTML("article.product_pod", func(h *colly.HTMLElement) {
		item:=Item{
			Link: h.ChildAttr("a", "href"),
			Name: h.ChildAttr("h3 a","title"),
			Price: h.ChildText("p.price_color"),
			Instock: h.ChildText("p.instock"),
		}
		items = append(items, item)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("http://books.toscrape.com/catalogue/category/books/travel_2/index.html")
	c.Wait()
	data,err:=json.MarshalIndent(items,""," ")
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(data))
	fmt.Println(len(items))
}