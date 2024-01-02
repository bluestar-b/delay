package main

import (
	"log"
	"os"
	"reflect"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ALLOW_ORIGINS []string `yaml:"ALLOW_ORIGINS"`
	HOST_NAME     string   `yaml:"HOST_NAME"`
	HOST_PORT     string   `yaml:"HOST_PORT"`
	DATA_DIR      string   `yaml:"DATA_DIR"`
}

type ConfigParser struct {
	config Config
	once   sync.Once
}

var instance *ConfigParser

func GetConfigParser() *ConfigParser {
	if instance == nil {
		instance = &ConfigParser{}
	}
	return instance
}

func (c *ConfigParser) LoadConfig(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &c.config)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConfigParser) Get(key string) interface{} {
	c.once.Do(func() {
		err := c.LoadConfig("config.yaml")
		if err != nil {
			log.Fatal("Error loading config:", err)
		}
	})

	val := reflectField(c.config, key)
	if val.IsValid() {
		return val.Interface()
	}

	log.Fatalf("Key '%s' not found in the config", key)
	return nil
}
func reflectField(config Config, key string) reflect.Value {
	val := reflect.ValueOf(config)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	field := val.FieldByName(key)
	return field
}

/*
	configParser := GetConfigParser()

	fooValue, err := configParser.Get("Foo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("FOO:", fooValue)

	barValue, err := configParser.Get("Bar")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("BAR:", barValue)
}
*/
