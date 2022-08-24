package main

import (
	"fmt"
	colly "github.com/gocolly/colly"
	"log"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

/* takes args: " task9.exe wget url1 url2..."*/

func getArgs() (urls []string) {
	args := os.Args
	if len(args) < 3 {
		log.Fatal("usage: wget example.com example2.com ...")
	}

	return args[2:]
}

func CreateFolder(folderName string) {
	_, ok := os.Stat(folderName)
	if os.IsExist(ok) {
		return
	} else if os.IsNotExist(ok) {
		if ok = os.MkdirAll(folderName, os.ModePerm); ok != nil {
			log.Fatal(ok)
		}
	}
}

func main() {
	var (
		folder   = "download"
		mainLink *url.URL
		reg      *regexp.Regexp
		ok       error
		col      *colly.Collector
		ulSet    = make(map[string]struct{})
	)
	CreateFolder(folder)
	urls := getArgs()
	for _, link := range urls {
		link = strings.TrimRight(link, "/")
		if mainLink, ok = url.ParseRequestURI(link); ok != nil {
			log.Fatal(ok)
		}
		if reg, ok = regexp.Compile("https?://([a-z0-9]+[.])*" + mainLink.Host); ok != nil {
			log.Fatal(ok)
		}
		CreateFolder(folder + "/" + mainLink.Host)
		col = colly.NewCollector(colly.URLFilters(reg))
		col.OnHTML("a[href]",
			func(el *colly.HTMLElement) {
				url := el.Request.AbsoluteURL(el.Attr("href"))
				if _, isExist := ulSet[url]; !isExist {
					ulSet[url] = struct{}{}
					_ = col.Visit(url)
				}
			})

		col.OnResponse(func(r *colly.Response) {
			pth := r.Request.URL.Path
			full := folder + "/" + mainLink.Hostname() + pth
			if _, ok := ulSet[full]; !ok {
				ulSet[full] = struct{}{}
			} else {
				return
			}
			if path.Ext(full) == "" {
				CreateFolder(full)
			} else {
				CreateFolder(full[:strings.LastIndexByte(full, '/')])
			}
			if path.Ext(pth) == "" {
				if full[len(full)-1] != '/' {
					full += "/"
				}
				full += "index.html"
				if _, ok := os.Create(full); ok != nil {
					fmt.Println("err file creation", ok)
				}
			}
			fmt.Println("downloaded:", mainLink.Hostname()+pth)
			if ok = r.Save(full); ok != nil {
				log.Fatal(ok)
			}
		})

		if ok = col.Visit(mainLink.String()); ok != nil {
			log.Fatal("err while visiting: " + ok.Error())
		}
		col.Wait()
	}
	fmt.Println("End.")
}
