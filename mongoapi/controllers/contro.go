package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mongoapi.com/model"
)

const connectionString = "mongodb://localhost:27017"
const dbname = "netflix"
const colname = "watchlist"

var collection *mongo.Collection

// connet with db
// init is not the maian method it runs only one time for the initialization of th program and only one time

func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Mongodb COnnected")
	collection = client.Database(dbname).Collection(colname)

	fmt.Println("Collection instance is ready")

}

// MOnogo helpers
func insert_movie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("INserted one movie in db with id", inserted.InsertedID)

}

// updatre
func updateMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified counnt : ", result.ModifiedCount)

}

func deleteMOvie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	deletCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Movie got deleted: ", deletCount)
}

// deleteall

func deleteAll() int64 {
	filter := bson.D{{}}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("no o fitems deleted are ", deleteResult)
	return deleteResult.DeletedCount

}
func getAllMovies() [] primitive.M  {
	cur,err:=collection.Find(context.Background(),bson.D{{}})
	if err!=nil {
		log.Fatal(err)
		
	}
	var movies [] primitive.M
	for cur.Next(context.Background()){
		var movie bson.M
		err:=cur.Decode(&movie)
		if err!=nil {
			log.Fatal(err)
			
		}
		movies=append(movies, movie)

	}
	defer cur.Close(context.Background())
	return movies
	
}


func GetmyallMOvies(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	allMovies:=getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insert_movie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")
	params:=mux.Vars(r)
	updateMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	

}

func DeleteMovie(w http.ResponseWriter , r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	params:=mux.Vars(r)
	deleteMOvie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
	

	
}
func DeleteAllMovie(w http.ResponseWriter , r *http.Request)  {

	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	
count:=	deleteAll()


	json.NewEncoder(w).Encode(count)
	

	
}

