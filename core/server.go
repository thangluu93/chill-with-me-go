package core

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
	"io/ioutil"
	"main/models"
	"time"

	firebase "firebase.google.com/go/v4"
)

type Server struct {
	Echo        *echo.Echo
	Auth        *auth.Client
	Firebase    *firebase.App
	Firestore   *firestore.Client
	MongoClient *mongo.Client
	Config      *models.Config
	Executor    *Executor
}

func NewServer() (svr *Server, err error) {
	e := echo.New()

	config := &models.Config{}

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	firebaseConfig := &firebase.Config{}
	err = json.Unmarshal(file, firebaseConfig)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsFile(config.FirebaseConfig)
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	auth, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	firestoreApp, err := firebaseApp.Firestore(context.Background())

	clientOption := options.Client().ApplyURI(config.DB)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, clientOption)

	// init new executor
	executor := NewExecutor(config)

	svr = &Server{
		Echo:        e,
		Auth:        auth,
		Firebase:    firebaseApp,
		MongoClient: mongoClient,
		Executor:    executor,
		Firestore:   firestoreApp,
	}

	return svr, nil

}
