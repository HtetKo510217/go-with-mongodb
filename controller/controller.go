package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/htetko/go-with-mongodb/initializers"
	"github.com/htetko/go-with-mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// insert 1 record helper mothod
func insertOneMovie(movie model.Netflix) {
	insertResult, err := initializers.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

// update 1 record helper method
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}                       // query where id is same as movieId
	update := bson.M{"$set": bson.M{"watched": true}} // set watched to true
	_, err := initializers.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated a single document: ", movieId)
}

// delete 1 record helper method
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	_, err := initializers.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted a single document: ", movieId)
}

// delete all records helper method
func deleteAllMovies() {
	filter := bson.D{{}}
	_, err := initializers.Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted all documents")
}

// get all records helper method
func getAllMovies() []primitive.M {
	var movies []primitive.M
	cursor, err := initializers.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	fmt.Println(movies)

	return movies
}

// acutal controller

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movies := getAllMovies()
	if len(movies) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "No movies found"}`))
		return
	}
	json.NewEncoder(w).Encode(movies)
	w.WriteHeader(http.StatusOK)

}

//create movie

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Movie created successfully"}`))
}

// markaswatched

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Movie marked as watched"}`))
}

// get 1 record

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	var movie model.Netflix
	err := initializers.Collection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Movie not found"}`))
		return
	}
	json.NewEncoder(w).Encode(movie)
	w.WriteHeader(http.StatusOK)
}

// delete 1 record

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Movie deleted successfully"}`))
}

// delete all records
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	deleteAllMovies()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "All movies deleted successfully"}`))
}
