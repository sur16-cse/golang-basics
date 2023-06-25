package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/surbhikumari/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://surbhi1611:Surbhi1543@cluster0.za3zfxg.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const collName = "watchlist"

//Most Important,
var collection *mongo.Collection

func init() {
	//client option
	clientOption:=options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client,err:=mongo.Connect(context.TODO(),clientOption)

	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection=client.Database(dbName).Collection(collName)

	//collection instance 
	fmt.Println("Collection instance is ready")
}

//MongoDb helpers - file

//insert 1 record
func insertOneMovie(movie model.Netflix){
	inserted,err:=collection.InsertOne(context.Background(),movie)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ",inserted.InsertedID)
}

//update 1 record
func updateOneMovie(movieId string)  {
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	update:=bson.M{"$set":bson.M{"watched":true}}

	result,err:=collection.UpdateOne(context.Background(),filter,update)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("modified count:",result.ModifiedCount)
}

//delete 1 record
func deleteOneMovie(movieId string)  {
	id,_:=primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{"_id":id}
	deleteCount,err:=collection.DeleteOne(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Movie got delete with deleteCount: ",deleteCount)
}

// delete all record from mongoDb
func deleteAllMovie() int64{
	deleteResult,err:=collection.DeleteMany(context.Background(),bson.D{{}},nil)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Number of movies delete: ",deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

//get all movies from database
func getAllMovies() []primitive.M {
	curr,err:=collection.Find(context.Background(),bson.D{})
	if err!=nil{
		log.Fatal(err)
	}

	var movies []primitive.M
	for curr.Next(context.Background()){
		var movie bson.M
		err:=curr.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}
		movies=append(movies, movie)
	}
	defer curr.Close(context.Background())
	return movies
}

//exporting function
func GetMyAllMovies(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	allMovies:= getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	_=json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	params:= mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params:=mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteALLMovies(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	count:=deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}