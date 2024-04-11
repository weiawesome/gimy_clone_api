package mongodb

import (
	pb "api_film_service/proto/film_service"
	"api_film_service/repository/model"
	"api_film_service/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (r *repository) GetBasicFilms(ctx context.Context, request *pb.FilmBasicRequest) ([]model.Film, error) {
	var results []model.Film

	pipeline := mongo.Pipeline{}
	pipeline = sortByUpdate(pipeline)
	pipeline = pagination(pipeline, request.Offset, request.Limit)

	collection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return results, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	err = cursor.All(ctx, &results)
	return results, err
}

func (r *repository) GetPopularFilmsByType(ctx context.Context, request *pb.FilmPopularTypeRequest) ([]model.Film, error) {
	var results []model.Film

	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	pipeline := mongo.Pipeline{}
	pipeline = filterType(pipeline, request.Type)
	pipeline = joinPopularity(pipeline)
	pipeline = filterToday(pipeline, today)
	pipeline = sortByPopularity(pipeline)
	pipeline = pagination(pipeline, request.Offset, request.Limit)

	collection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return results, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	err = cursor.All(ctx, &results)
	return results, err
}

func (r *repository) GetPopularFilmsByCategory(ctx context.Context, request *pb.FilmPopularCategoryRequest) ([]model.Film, error) {
	var results []model.Film

	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	pipeline := mongo.Pipeline{}
	pipeline = filterCategory(pipeline, request.Category)
	pipeline = joinPopularity(pipeline)
	pipeline = filterToday(pipeline, today)
	pipeline = sortByPopularity(pipeline)
	pipeline = pagination(pipeline, request.Offset, request.Limit)

	collection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return results, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	err = cursor.All(ctx, &results)
	return results, err
}

func (r *repository) GetSpecificFilm(ctx context.Context, request *pb.FilmSpecificRequest) (model.Film, error) {
	var result model.Film

	filter := getIdFilter(request.Id)

	collection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	err := collection.FindOne(ctx, filter).Decode(&result)
	return result, err
}

func (r *repository) GetSpecificFilmRoutes(ctx context.Context, request *pb.FilmSpecificRequest) ([]model.FilmRoute, error) {
	var results []model.FilmRoute

	filter := getIdFilter(request.Id)

	collection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmRouteCollection())
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return results, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	err = cursor.All(ctx, &results)
	return results, err
}

func (r *repository) GetRankedFilms(ctx context.Context, request *pb.FilmRankedRequest) ([]model.Film, error) {
	var results []model.Film

	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	oneWeekAgo := today.AddDate(0, 0, -7)

	pipeline := mongo.Pipeline{}
	pipeline = joinPopularity(pipeline)
	pipeline = filterCategory(pipeline, request.Category)
	pipeline = filterOneWeekAgo(pipeline, oneWeekAgo)
	pipeline = calculateTotalPopularity(pipeline)
	pipeline = sortByWeeklyPopularity(pipeline)
	pipeline = pagination(pipeline, request.Offset, request.Limit)

	collection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return results, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	err = cursor.All(ctx, &results)
	return results, err
}

func (r *repository) GetFilterFilms(ctx context.Context, request *pb.FilmFilterRequest) ([]model.Film, error) {
	var results []model.Film

	pipeline := mongo.Pipeline{}
	pipeline = filterType(pipeline, request.Type)
	pipeline = filterCategory(pipeline, request.Category)
	pipeline = filterLocation(pipeline, request.Location)
	pipeline = filterReleaseYear(pipeline, request.ReleaseYear)
	switch request.OrderType {
	case pb.OrderType_UpdateTime:
		pipeline = sortByUpdate(pipeline)

	case pb.OrderType_UploadTime:
		pipeline = sortByUpload(pipeline)
	case pb.OrderType_WEEKLY_POPULARITY:
		now := time.Now().UTC()
		today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		oneWeekAgo := today.AddDate(0, 0, -7)
		pipeline = joinPopularity(pipeline)
		pipeline = filterOneWeekAgo(pipeline, oneWeekAgo)
		pipeline = calculateTotalPopularity(pipeline)
		pipeline = sortByWeeklyPopularity(pipeline)
	case pb.OrderType_TOTAL_POPULARITY:
		pipeline = sortByTotalPopularity(pipeline)
	}
	pipeline = pagination(pipeline, request.Offset, request.Limit)

	collection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return results, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)

	err = cursor.All(ctx, &results)
	return results, err
}

func (r *repository) GetSearchFilms(ctx context.Context, request *pb.FilmSearchRequest) ([]model.Film, error) {
	var results []model.Film
	var err error

	if request.SearchType != pb.SearchType_CELEBRITY {
		return results, errors.New("error method to search")
	}

	pipeline := mongo.Pipeline{}
	pipeline = filterCelebrity(pipeline, request.Content)
	pipeline = pagination(pipeline, request.Offset, request.Limit)

	collection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return results, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx)
	err = cursor.All(ctx, &results)
	return results, err
}

func (r *repository) SaveFilm(ctx context.Context, request *pb.FilmSaveRequest) error {
	if _, err := r.GetSpecificFilm(ctx, &pb.FilmSpecificRequest{Id: request.Id}); err == nil {
		return errors.New("the id have been use")
	}
	film := model.Film{
		Id:           request.Id,
		Title:        request.Title,
		Type:         request.Type.String(),
		Category:     request.Category.String(),
		Language:     request.Language,
		Resource:     request.Resource,
		ReleaseYear:  request.ReleaseYear,
		State:        request.State,
		Location:     request.Location.String(),
		Introduction: request.Introduction,
		Directors:    request.Directors,
		Actors:       request.Actors,
		Popularity:   0,
		CreatedAt:    time.Now().UTC(),
		UpdateAt:     time.Now().UTC(),
	}
	collection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	_, err := collection.InsertOne(ctx, film)
	return err
}

func (r *repository) SaveFilmEpisode(ctx context.Context, request *pb.FilmSaveEpisodeRequest) error {
	filmCollection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	filmRouteCollection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmRouteCollection())

	opts := options.Update().SetUpsert(true)
	update := bson.M{
		"$addToSet": bson.M{"episodes": request.Episode},
	}
	filter := bson.M{"id": request.Id, "route": request.Route}

	_, err := filmRouteCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	_, err = filmCollection.UpdateOne(
		ctx,
		getIdFilter(request.Id),
		bson.M{
			"$set": bson.M{
				"state":     request.State,
				"update_at": time.Now().UTC(),
			},
		},
	)
	return err
}

func (r *repository) DeleteFilmEpisode(ctx context.Context, request *pb.FilmSaveEpisodeRequest) error {
	filmCollection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	filmRouteCollection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmRouteCollection())
	opts := options.Update().SetUpsert(false)
	filter := bson.M{"id": request.Id, "route": request.Route}
	update := bson.M{
		"$pull": bson.M{"episodes": request.Episode},
	}

	_, err := filmRouteCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	_, err = filmCollection.UpdateOne(
		ctx,
		getIdFilter(request.Id),
		bson.M{
			"$set": bson.M{
				"state":     request.State,
				"update_at": time.Now().UTC(),
			},
		},
	)
	return err

}

func (r *repository) DeleteFilm(ctx context.Context, request *pb.FilmSpecificRequest) error {
	filmCollection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	filmRoutesCollection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmRouteCollection())
	filmPopulationCollection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmPopularityCollection())
	filter := bson.M{"id": request.Id}

	if _, err := filmPopulationCollection.DeleteMany(ctx, filter); err != nil {
		return err
	}

	if _, err := filmRoutesCollection.DeleteMany(ctx, filter); err != nil {
		return err
	}
	if _, err := filmCollection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}

func (r *repository) AddPopularity(ctx context.Context, request *pb.FilmSpecificRequest) error {
	now := time.Now().UTC()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	filmCollection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmCollection())
	filmPopularityCollection := r.client.Database(utils.EnvMongoDb()).Collection(utils.GetFilmPopularityCollection())
	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"id":         request.Id,
		"created_at": today,
	}
	update := bson.M{
		"$inc": bson.M{"popularity": 1},
	}
	_, err := filmPopularityCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	idFilter := getIdFilter(request.Id)
	_, err = filmCollection.UpdateOne(ctx, idFilter, update, opts)
	return err
}
