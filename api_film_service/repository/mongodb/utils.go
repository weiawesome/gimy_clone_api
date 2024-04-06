package mongodb

import (
	pb "api_film_service/proto/film_service"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func getIdFilter(id string) bson.D {
	return bson.D{{"id", id}}
}
func filterType(pipeline []bson.D, mediaType pb.MediaType) []bson.D {
	if mediaType != pb.MediaType_HOME {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"type", mediaType.String()}}}})
	}
	return pipeline
}
func filterCategory(pipeline []bson.D, mediaCategory pb.MediaCategory) []bson.D {
	if mediaCategory != pb.MediaCategory_ALL_CATEGORY {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"category", mediaCategory.String()}}}})
	}
	return pipeline
}
func filterLocation(pipeline []bson.D, mediaLocation pb.MediaLocation) []bson.D {
	if mediaLocation != pb.MediaLocation_ALL_LOCATION {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"location", mediaLocation.String()}}}})
	}
	return pipeline
}
func filterReleaseYear(pipeline []bson.D, releaseYear uint32) []bson.D {
	if releaseYear != 0 {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"release_year", releaseYear}}}})
	}
	return pipeline
}
func filterCelebrity(pipeline []bson.D, celebrity string) []bson.D {
	pipeline = append(pipeline, bson.D{
		{"$match", bson.D{
			{"$or", []bson.D{
				{{"actors", celebrity}},
				{{"directors", celebrity}},
			}},
		}},
	})
	return pipeline
}
func joinPopularity(pipeline []bson.D) []bson.D {
	pipeline = append(pipeline, bson.D{{"$lookup", bson.D{
		{"from", "FilmPopularity"},
		{"localField", "id"},
		{"foreignField", "id"},
		{"as", "popularityDetails"},
	}}})
	return pipeline
}
func filterToday(pipeline []bson.D, today time.Time) []bson.D {
	pipeline = append(pipeline, bson.D{
		{"$addFields", bson.D{
			{"popularityDetails", bson.D{
				{"$filter", bson.D{
					{"input", "$popularityDetails"},
					{"as", "item"},
					{"cond", bson.D{{"$eq", bson.A{"$$item.created_at", today}}}},
				}},
			}},
		}},
	})
	pipeline = append(pipeline, bson.D{{"$addFields", bson.D{{"isTodayPopularity", bson.D{{"$cond", bson.A{bson.D{{"$gt", bson.A{bson.D{{"$size", "$popularityDetails"}}, 0}}}, true, false}}}}}}})
	return pipeline
}

func filterOneWeekAgo(pipeline []bson.D, oneWeekAgo time.Time) []bson.D {
	pipeline = append(pipeline, bson.D{{"$addFields", bson.D{
		{"popularityDetails", bson.D{
			{"$filter", bson.D{
				{"input", "$popularityDetails"},
				{"as", "popularityItem"},
				{"cond", bson.D{{"$gte", bson.A{"$$popularityItem.createdAt", oneWeekAgo}}}},
			}},
		}},
	}}})
	return pipeline
}
func calculateTotalPopularity(pipeline []bson.D) []bson.D {
	pipeline = append(pipeline, bson.D{{"$addFields", bson.D{
		{"totalPopularity", bson.D{{"$sum", "$popularityDetails.popularity"}}},
	}}})
	return pipeline
}
func sortByPopularity(pipeline []bson.D) []bson.D {
	pipeline = append(pipeline, bson.D{{"$sort", bson.D{{"isTodayPopularity", -1}, {"popularityDetails.popularity", -1}, {"popularity", -1}}}})
	return pipeline
}
func sortByWeeklyPopularity(pipeline []bson.D) []bson.D {
	pipeline = append(pipeline, bson.D{{"$sort", bson.D{{"totalPopularity", -1}, {"popularity", -1}}}})
	return pipeline
}
func sortByUpload(pipeline []bson.D) []bson.D {
	pipeline = append(pipeline, bson.D{{"$sort", bson.D{{"created_at", -1}}}})
	return pipeline
}
func sortByUpdate(pipeline []bson.D) []bson.D {
	pipeline = append(pipeline, bson.D{{"$sort", bson.D{{"update_at", -1}}}})
	return pipeline
}
func sortByTotalPopularity(pipeline []bson.D) []bson.D {
	pipeline = append(pipeline, bson.D{{"$sort", bson.D{{"popularity", -1}}}})
	return pipeline
}
func pagination(pipeline []bson.D, offset int32, limit int32) []bson.D {
	pipeline = append(pipeline, bson.D{{"$skip", offset}})
	pipeline = append(pipeline, bson.D{{"$limit", limit}})
	return pipeline
}
