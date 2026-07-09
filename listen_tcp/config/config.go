package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func LoadConfig() *Config {

	file, err := os.Open("/home/milrine/_projects/gogo/serve_tcp/config.json") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	Data := data[:count]

	fmt.Printf("read %d bytes: %q\n", count, Data)

	Conf := &Config{}

	err = json.Unmarshal(Data, Conf)

	if err != nil {
		log.Printf("%s\n", err)
		panic("error while Unmarshalling the config.json file")
	}

	return Conf
}
