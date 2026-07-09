package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`

	Schema KeySpace `json:"keyspace"`
}

type KeySpace struct {
	DB   string `json:"db"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

const (
	PATH = "/home/milrine/_projects/gogo/serve_tcp/config.json"
)

func LoadConfig() *Config {

	log.Println("loading config.json from path:: ", PATH)

	file, err := os.Open(PATH) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	// byte size should equal or greater than config.json file size
	data := make([]byte, 1000) // for the collective good I'm setting it to 1000b
	count, err := file.Read(data)

	if err != nil {
		log.Fatal(err)
	}

	Data := data[:count]
	Conf := &Config{}
	err = json.Unmarshal(Data, Conf)

	if err != nil {
		log.Printf("%s\n", err)
		panic("error while Unmarshalling the config.json file")
	}

	log.Println("Unmarshalled config: ", Conf)

	return Conf
}
