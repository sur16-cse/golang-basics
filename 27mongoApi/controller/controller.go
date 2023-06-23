package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/surbhikumari/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://surbhi1611:Sur@1611@cluster0.za3zfxg.mongodb.net/?retryWrites=true&w=majority"
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
func insertOneRecord(movie model.Netflix){
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