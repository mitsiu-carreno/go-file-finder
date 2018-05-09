package main

import(
	"os"
	"log"
	"fmt"
	"time"
	"flag"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/cheggaaa/pb.v1"
	models "github.com/mitsiu-carreno/go-file-finder/declarations"
)

// Log sets the logger to store missing filenames
var (
	Log *log.Logger
)

// NewLog creates a file on the specified path and writes all missing files there
func NewLog(logpath string){
	file, err := os.Create(logpath)
	Check(err)
	Log = log.New(file, "", log.LstdFlags)
}

// Check function 
func Check(e error){
	if e != nil{
		fmt.Println(e)
		panic(e)
	}
}

func main(){
	// Config log
	var logpath = flag.String("logpath", "./missings.log", "Log Path")
	NewLog(*logpath)


	// DATABASE
	var (
		hosts		= os.Getenv("MAIN_DB_HOST")
		database 	= os.Getenv("MAIN_DB_DB")
		username 	= os.Getenv("MAIN_DB_USER")
		password 	= os.Getenv("MAIN_DB_PASSWORD")
		collection	= os.Getenv("MAIN_DB_COLLECTION")
		filePath 	= os.Getenv("FILE_INPUT")
	)

	// Mongo connection
	info := &mgo.DialInfo{
		Addrs		: []string{hosts},
		Timeout 	: 60 * time.Second,
		Database 	: database,
		Username	: username,
		Password 	: password,
	}

	session, err := mgo.DialWithInfo(info)
	Check(err)
	defer session.Close()

	col := session.DB(database).C(collection)

	var documents []models.Declarations

	// Retreive all documents from mongo
	err = col.Find(bson.M{}).All(&documents) 

	progressBar := pb.StartNew(len(documents)) 

	for _, entry := range documents{
		//var entryNum = i+1
		progressBar.Increment()

		// Check if file exists
		_, err := os.Stat(filePath + entry.ARCHIVO)
		if os.IsNotExist(err){
			Log.Println(entry.ARCHIVO)
			continue
		}
		Check(err)
	}

	progressBar.FinishPrint("Check missiongs.log")
}