package jsonconf

import(
	"os"
	"log"
	"fmt"
	"errors"
	"encoding/json"
)

/*
Conf is a global variable that will map a given string to an interface generated from the JSON file.

It will be loaded with values on initial LoadConfig call

*/

var Conf map[string]interface{}

/*
	Given a filename, LoadConfig will load a json file from the location. It will then decode it and store it into the global Conf variable
*/

func LoadConfig(filename string) {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	var f interface{}
	err := decoder.Decode(&f)
	if err != nil {
		log.Fatal(err)
	}
	Conf = f.(map[string]interface{})
}

/*
	Given the name of a string, GetVar will try to locate an environment variable that matches the string. If that match comes empty, then it will look in the configuration file for the string as a key. If that match comes up empty, then the call will return an error. Otherwise, it will return the string of the interface{} stored in Conf
*/

func GetVar(v string) (string, error){
	env := os.Getenv(v)
	if (env != ""){
		return env, nil
	}
	result, found := Conf[v]
	if (!found){
		return "", errors.New("Value not found")
	}
	return fmt.Sprintf("%v", result), nil
}
