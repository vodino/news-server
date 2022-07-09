package trash

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"news/internal/source/sutil"
	"news/internal/store"
	"news/internal/store/schema"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const AbidjanNetName = "Abidjan.Net"

type AbidjanNetSource struct {
	*store.NewsSource
	*http.Client
}

func NewAbidjanNetSource(source *store.NewsSource) *AbidjanNetSource {
	return &AbidjanNetSource{
		Client:     http.DefaultClient,
		NewsSource: source,
	}
}

/// LatestPost
///
///
func (src *AbidjanNetSource) LatestPost(ctx context.Context) []*schema.NewsPost {
	response, err := sutil.RodGetRequest(fmt.Sprintf("%s%s", src.URL, *src.LatestPostURL))
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.latestPost(sutil.NewElement(document.Selection))
}

func (src *AbidjanNetSource) latestPost(document *sutil.Element) []*schema.NewsPost {
	selector := src.LatestPostSelector
	filmList := make([]*schema.NewsPost, 0)

	elementCallBack := func(element *sutil.Element) {
		// category := element.ChildText(selector.Category[0])
		image := element.ChildAttribute(selector.Image[0], selector.Image[1])
		title := element.ChildText(selector.Title[0])
		link := element.Attribute(selector.Link[0])
		date := element.ChildText(selector.Date[0])
		if len(image) == 0 {
			image = element.ChildAttribute(selector.Image[0], selector.Image[2])
			if len(image) == 0 {
				style := element.ChildAttribute(selector.Image[3], "style")
				exp := regexp.MustCompile(`(http(s|)://.*')`)
				image = strings.Trim(exp.FindString(style), "'")
			}
		}
		value := strings.Split(date, "-")
		date = strings.TrimSpace(value[len(value)-1])

		image = sutil.ParseURL(src.URL, image)
		link = sutil.ParseURL(src.URL, link)
		date, _ = sutil.ParseTime(date)

		if strings.Contains(image, "defaut-cover-photo.svg") {
			image = ""
		}

		if !strings.Contains(strings.Join(value, ""), "Fraternité Matin") && len(image) != 0 {
			filmList = append(filmList, &schema.NewsPost{
				Source: src.Name,
				Logo:   src.Logo,
				Image:  image,
				Title:  title,
				Link:   link,
				Date:   date,
			})
		}
	}

	elementCallBack(sutil.NewElement(document.Selection.Find(selector.List[1])))

	document.ForEach(selector.List[0],
		func(i int, element *sutil.Element) {
			elementCallBack(element)
		})
	return filmList
}

/// CategoryPost
///
///
func (src *AbidjanNetSource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsPost {
	category, err := sutil.ParseCategorySource(src.NewsSource, category)
	if err != nil {
		log.Println(err)
		return nil
	}
	response, err := sutil.RodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(*src.CategoryPostURL, category, page)))
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.categoryPost(sutil.NewElement(document.Selection))
}

func (src *AbidjanNetSource) categoryPost(document *sutil.Element) []*schema.NewsPost {
	selector := src.LatestPostSelector
	filmList := make([]*schema.NewsPost, 0)
	document.ForEach(selector.List[0],
		func(i int, element *sutil.Element) {
			// category := element.ChildText(selector.Category[0])
			image := element.ChildAttribute(selector.Image[0], selector.Image[1])
			title := element.ChildText(selector.Title[0])
			link := element.Attribute(selector.Link[0])
			date := element.ChildText(selector.Date[0])

			value := strings.Split(date, "-")
			date = strings.TrimSpace(value[len(value)-1])

			image = sutil.ParseURL(src.URL, image)
			link = sutil.ParseURL(src.URL, link)
			date, _ = sutil.ParseTime(date)

			if strings.Contains(image, "defaut-cover-photo.svg") {
				image = ""
			}

			if !strings.Contains(strings.Join(value, ""), "Fraternité Matin") && len(image) != 0 {
				filmList = append(filmList, &schema.NewsPost{
					Source: src.Name,
					Logo:   src.Logo,
					Image:  image,
					Title:  title,
					Link:   link,
					Date:   date,
				})
			}
		})
	return filmList
}

/// NewsArticle
///
///

func (src *AbidjanNetSource) NewsArticle(ctx context.Context, link string) *schema.NewsArticle {
	response, err := sutil.RodGetRequest(link)
	if err != nil {
		log.Println(err)
		return nil
	}
	document, err := goquery.NewDocumentFromReader(response)
	if err != nil {
		log.Println(err)
		return nil
	}
	return src.newsArticle(sutil.NewElement(document.Selection))
}

func (src *AbidjanNetSource) newsArticle(document *sutil.Element) *schema.NewsArticle {
	selector := src.ArticleSelector
	description := document.ChildOuterHtml(selector.Description[0])
	description = strings.Join(strings.Fields(description), " ")
	return &schema.NewsArticle{
		Description: description,
	}
}