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
	cloudStorage "firebase.google.com/go/v4/storage"
)

type Server struct {
	Echo         *echo.Echo
	Auth         *auth.Client
	DBName       string
	Firebase     *firebase.App
	Firestore    *firestore.Client
	CloudStorage *cloudStorage.Client
	MongoClient  *mongo.Client
	Config       *models.Config
	Executor     *Executor
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

	cloudStorageApp, err := firebaseApp.Storage(context.Background())

	clientOption := options.Client().ApplyURI(config.DB)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		return nil, err
	}

	svr = &Server{
		Echo:         e,
		Auth:         auth,
		Config:       config,
		DBName:       config.DbName,
		Firebase:     firebaseApp,
		CloudStorage: cloudStorageApp,
		MongoClient:  mongoClient,
		Firestore:    firestoreApp,
	}

	return svr, nil
}

func (s *Server) Start() (err error) {
	return s.Echo.Start(s.Config.HostIP)
}
