package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imnmania/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://admin:n%40123456@localhost:8083/?ssl=false"
const dbName = "netflix"
const collName = "watchlist"

// ----------------
// * MOST IMPORTANT
// ----------------
var collection *mongo.Collection

// --------------------
// Connect with MongoDB
// --------------------
func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, error := mongo.Connect(context.TODO(), clientOptions)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("MongoDB connection success...")

	// collection instance
	collection = (*mongo.Collection)(client.Database(dbName).Collection(collName))
	fmt.Println("Collection instance/reference is ready...")
}

// ----------------------
// MongoDB Helpers - file
// ----------------------

// insert 1 record
func insertOneMovie(movie model.Netflix) {
	inserted, error := collection.InsertOne(context.Background(), movie)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, error := collection.UpdateOne(context.Background(), filter, update)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, error := collection.DeleteOne(context.Background(), filter)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Movie got deleted with delete count: ", result.DeletedCount)
}

// delete all records
func deleteAllMovies() int {
	filter := bson.D{{}}
	result, error := collection.DeleteMany(context.Background(), filter)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Number of movies deleted: ", result.DeletedCount)
	return int(result.DeletedCount)
}

// get all movies from database
func getAllMovies() []primitive.M {
	cursor, error := collection.Find(context.Background(), bson.D{{}})
	if error != nil {
		log.Fatal(error)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}

// -------------------------
// Actual Controller - file
// -------------------------
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}
