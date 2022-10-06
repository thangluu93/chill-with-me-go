package models

type Config struct {
	HostIP         string   `json:"hostIP"`
	FirebaseConfig string   `json:"firebaseConfig"`
	DB             string   `json:"db"`
	DbName         string   `json:"dbName"`
	DevMode        bool     `json:"devMode"`
	Executors      []string `json:"executors"`
	BucketName     string   `json:"bucketName"`
}
