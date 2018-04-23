package main


/*  

<sitemapindex>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
	</sitemap>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
	</sitemap>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
	</sitemap>
</sitemapindex>
*/ 

import (
  "encoding/xml"
  "fmt"
  "io/ioutil"
  "net/http"
)

type Sitemapindex struct {
  Locations []Location `xml:"sitemap"`
}

type Location struct {
  Loc string `xml:"loc"`
}

func (l Location) String() string {
  return fmt.Sprintf(l.Loc)
}

func main() {
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  var s Sitemapindex
  xml.Unmarshal(bytes, &s)
  //fmt.Println(s.Locations)
  fmt.Printf("Here %s some %s", "are", "variables")
  for _, Location := range s.Locations{
  	fmt.Printf("\n%s", Location)
  }
}

/*
package main


/*  

<sitemapindex>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
	</sitemap>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
	</sitemap>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
	</sitemap>
</sitemapindex>
*/ 
/*
import (
  "encoding/xml"
  "fmt"
  "io/ioutil"
  "net/http"
)

type Sitemapindex struct {
  Locations []Location `xml:"sitemap"`
}

type Location struct {
  Loc string `xml:"loc"`
}

func (l Location) String() string {
  return fmt.Sprintf(l.Loc)
}

func main() {
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  var s Sitemapindex
  xml.Unmarshal(bytes, &s)
  //fmt.Println(s.Locations)
  fmt.Printf("Here %s some %s", "are", "variables")
  for _, Location := range s.Locations{
  	fmt.Printf("\n%s", Location)
  }
}
*/

/*
package main 

import "fmt"

func main() {

	a:= 3
	for x:=5; a<25; x+=3{
		fmt.Println("do stuff",x)
		a+=4
	}

}

*/

/*
package main

import "fmt"

func main() {
	x := 5

	for {
		fmt.Println("Do stuff",x)
		x+=3
		if x > 25{
			break
		}
	}

}


*/

