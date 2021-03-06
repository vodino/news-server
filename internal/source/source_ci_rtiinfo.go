package source

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"news/internal/store"
	"news/internal/store/schema"
	"news/internal/util"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const RTIInfoName = "RTI Info"

type RTIInfoSource struct {
	*store.NewsArticleSource
	*http.Client
}

func NewRTIInfoSource(source *store.NewsArticleSource) *RTIInfoSource {
	return &RTIInfoSource{
		Client:            http.DefaultClient,
		NewsArticleSource: source,
	}
}

/// LatestPost
///
///u00e0
func (src *RTIInfoSource) LatestPost(ctx context.Context) []*schema.NewsArticlePost {
	response, err := util.RodGetRequest(fmt.Sprintf("%s%s", src.URL, *src.LatestPostURL))
	if err != nil {
		log.Println(err)
		return nil
	}
	data := make(map[string]interface{})
	json.NewDecoder(response).Decode(&data)
	return src.latestPost(data)
}

func (src *RTIInfoSource) latestPost(data map[string]interface{}) []*schema.NewsArticlePost {
	selector := src.LatestPostSelector
	result := make([]*schema.NewsArticlePost, 0)
	for _, value := range data[selector.List[0]].([]interface{}) {
		data := value.(map[string]interface{})
		image := fmt.Sprintf(selector.Image[1], data[selector.Image[0]].(string))
		link := fmt.Sprintf(selector.Link[1], data[selector.Link[0]].(float64))
		title := data[selector.Title[0]].(string)
		date := data[selector.Date[0]].(string)

		link = src.URL + link
		dateTime, _ := util.ParseTime(date)

		result = append(result, &schema.NewsArticlePost{
			Date:   timestamppb.New(dateTime),
			Source: src.Name,
			Logo:   src.Logo,
			Title:  title,
			Image:  image,
			Link:   link,
		})
	}
	return result
}

/// CategoryPost
////////////////
func (src *RTIInfoSource) CategoryPost(ctx context.Context, category string, page int) []*schema.NewsArticlePost {
	category, err := util.ParseCategorySource(src.NewsArticleSource, category)
	if err != nil {
		log.Println(err)
		return nil
	}
	response, err := util.RodGetRequest(fmt.Sprintf("%s%s", src.URL, fmt.Sprintf(*src.CategoryPostURL, category)))
	if err != nil {
		log.Println(err)
		return nil
	}
	data := make(map[string]interface{})
	json.NewDecoder(response).Decode(&data)
	return src.categoryPost(data)
}

func (src *RTIInfoSource) categoryPost(data map[string]interface{}) []*schema.NewsArticlePost {
	selector := src.CategoryPostSelector
	result := make([]*schema.NewsArticlePost, 0)
	for _, value := range data[selector.List[0]].(map[string]interface{})[selector.List[1]].([]interface{}) {
		data := value.(map[string]interface{})
		link := fmt.Sprintf(selector.Link[1], data[selector.Link[0]].(float64))
		title := data[selector.Title[0]].(string)
		image := fmt.Sprintf(selector.Image[1], data[selector.Image[0]].(string))
		date := data[selector.Date[0]].(string)

		link = src.URL + link
		dateTime, _ := util.ParseTime(date)

		result = append(result, &schema.NewsArticlePost{
			Date:   timestamppb.New(dateTime),
			Source: src.Name,
			Logo:   src.Logo,
			Title:  title,
			Image:  image,
			Link:   link,
		})
	}
	return result
}

/// PostArticle
///////////////
func (src *RTIInfoSource) NewsArticle(ctx context.Context, link string) *schema.NewsArticlePost {
	response, err := util.RodGetRequest(link)
	if err != nil {
		log.Println(err)
		return nil
	}
	data := make(map[string]interface{})
	json.NewDecoder(response).Decode(&data)
	return src.newsArticle(data)
}

func (src *RTIInfoSource) newsArticle(data map[string]interface{}) *schema.NewsArticlePost {
	selector := src.ArticleSelector
	description := data[selector.Description[0]].(map[string]interface{})[selector.Description[1]].(string)
	return &schema.NewsArticlePost{
		Description: description,
	}
}
