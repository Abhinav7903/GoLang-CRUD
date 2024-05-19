package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Abhinav7903/mongo/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://db:hello123@cluster0.awms5vz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const collName = "watchlist"

var collection *mongo.Collection

// Create a connection to the MongoDB database
func init() {
	// Create a new MongoClient
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collName)

}

func insertOne(movie model.Netflix){
	inserted,err:=collection.InsertOne(context.Background(),movie)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Inserted one movie with id: ",inserted.InsertedID)
}

func updateone(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}
	result,err:=collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Updated one movie with id: ",result.ModifiedCount)

}

func deleteone(movieId string){
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	result,err:=collection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Deleted one movie with id: ",result.DeletedCount)
}

func deleteAll() int64{
	result,err:=collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Deleted all movies with count: ",result.DeletedCount)
	return result.DeletedCount
}

func getAllMovie() []primitive.M{
	cursor,err:=collection.Find(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}
	var movies []primitive.M
	for cursor.Next(context.Background()){
		var movie bson.M
		err:=cursor.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}

		movies=append(movies,movie)
	}

	defer cursor.Close(context.Background())
	return movies

}

func GetMyAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	movies:=getAllMovie()
	fmt.Println(movies)
	json.NewEncoder(w).Encode(movies)
	
}

func CreateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods","POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOne(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarksAsWatched(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods","PUT")

	params:=mux.Vars(r)
	updateone(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods","DELETE")

	params:=mux.Vars(r)
	deleteone(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods","DELETE")

	count:=deleteAll()
	json.NewEncoder(w).Encode(count)
	
}