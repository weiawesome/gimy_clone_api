package elasticsearch

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/repository/model"
	"api_film_service/utils"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"log"
	"strconv"
	"strings"
)

type SearchResult struct {
	Hits struct {
		Hits []struct {
			Source model.Film `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (r *repository) FilmSave(ctx context.Context, film model.Film) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(film); err != nil {
		return err
	}

	indexName := utils.GetSearchFilmIndex()
	if indexName == "" {
		return errors.New("index name is empty")
	}

	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: film.Id,
		Body:       strings.NewReader(buf.String()),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, r.client)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("error to close buffer " + err.Error())
		}
	}(res.Body)

	if res.IsError() {
		return errors.New("error with code " + strconv.Itoa(res.StatusCode))
	}
	return nil
}

func (r *repository) SearchFilmsGet(ctx context.Context, request *pb.FilmSearchRequest) ([]model.Film, error) {
	var results []model.Film
	var err error

	if request.SearchType != pb.SearchType_TITLE {
		return results, errors.New("error method to search")
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []interface{}{
					map[string]interface{}{
						"multi_match": map[string]interface{}{
							"query":     request.Content,
							"fields":    []string{"Title^3", "Introduction^2", "Actors", "Directors"},
							"type":      "best_fields",
							"fuzziness": "AUTO",
						},
					},
					map[string]interface{}{
						"prefix": map[string]interface{}{
							"Title": strings.ToLower(request.Content),
						},
					},
				},
				"minimum_should_match": 1,
			},
		},
		"from": request.Offset,
		"size": request.Limit,
		"sort": []map[string]interface{}{
			{"_score": map[string]string{"order": "desc"}},
		},
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return results, err
	}

	res, err := r.client.Search(
		r.client.Search.WithContext(ctx),
		r.client.Search.WithIndex(utils.GetSearchFilmIndex()),
		r.client.Search.WithBody(&buf),
		r.client.Search.WithPretty(),
	)
	if err != nil {
		return results, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	if res.IsError() {
		return results, errors.New("error to search films")
	}

	var searchResult SearchResult
	if err := json.NewDecoder(res.Body).Decode(&searchResult); err != nil {
		return results, err
	}
	results = make([]model.Film, len(searchResult.Hits.Hits))
	for i, hit := range searchResult.Hits.Hits {
		results[i] = hit.Source
	}
	return results, err
}

func (r *repository) FilmDelete(id string) error {
	indexName := utils.GetSearchFilmIndex()
	if indexName == "" {
		return errors.New("index name is empty")
	}
	_, err := r.client.Delete(indexName, id)
	return err
}
