package dota2ru

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/stewkk/iu9-networks/lab2/internal/errors"
	"golang.org/x/net/html"
)

// Service encapsulates usecase logic of dota2ru.
type Service interface {
	// ParseHeadings returns headings from dota2.ru forum.
	ParseHeadings(page int) ([]Heading, error)
}

// NewService returns new Service object.
func NewService() Service {
	return &service{}
}

type service struct{}

func (s *service) ParseHeadings(page int) ([]Heading, error) {
	url := "https://dota2.ru/forum/forums/zhelezo-novosti-i-obsuzhdenija.166/"
	if page != 1 {
		url += "page-"
		url += strconv.Itoa(page)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("can't download webpage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %v", errors.ErrServerStatusNotOK, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't parse webpage: %w", err)
	}

	var headings = make([]Heading, 0, 30)
	var f func(*html.Node) bool
	f = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "ul" && getAttr(n, "class") == "forum-section__list" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && c.Data == "li" &&
					strings.Index(getAttr(c, "class"), "forum-section__item") != -1 &&
					strings.Index(getAttr(c, "class"), "forum-section__item--first") == -1 {

					titleBlock := getChild(getChild(getChild(c, "div", "forum-section__col-2"), "div", "forum-section__title"), "a", "forum-section__title-unlogged")
					link :=  "https://dota2.ru" + getAttr(titleBlock, "href")
					title := strings.TrimSpace(titleBlock.FirstChild.Data)

					headings = append(headings, Heading{
						Title: title,
						Link:  link,
					})
				}
			}
			return true
		}
		isFound := false
		for c := n.FirstChild; c != nil && !isFound; c = c.NextSibling {
			isFound = f(c)
		}
		return isFound
	}
	if !f(doc) {
		return nil, fmt.Errorf("webpage parse failed: can't find headings")
	}
	return headings, nil
}

func getAttr(n *html.Node, name string) string {
	for _, attr := range n.Attr {
		if attr.Key == name {
			return attr.Val
		}
	}
	return ""
}

func getChild(n *html.Node, block string, class string) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == block && getAttr(c, "class") == class {
			return c
		}
	}
	return nil
}

// Heading represents forum thread heading.
type Heading struct {
	Title string
	Link  string
}
