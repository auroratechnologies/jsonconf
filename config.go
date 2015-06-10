package jsonconf

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Conf is a global variable that will map a given string to an interface generated
// from the JSON file. It will be loaded with values on initial LoadConfig call.
var Conf map[string]interface{}

// LoadConfig will, given a filename, load a json file from the location. It will
// then decode it and store it into the global Conf variable
func LoadConfig(filename string) error {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	var f interface{}
	err := decoder.Decode(&f)
	if err != nil {
		return err
	}
	Conf = f.(map[string]interface{})
	return nil
}

// GetVar will, given the name of a string, try to locate an environment variable
// that matches the string. If that match comes empty, then it will look in the
// configuration file for the string as a key. If that match comes up empty, the
// optional second parameter will be used as a default value.  If it's not present,
// then the call will return an error. Otherwise, it will return the string of
// the interface{} stored in Conf or the default as appropriate
func GetVar(v string, defaultVal ...string) (string, error) {
	env := os.Getenv(v)
	if env != "" {
		return env, nil
	}
	result, found := Conf[v]
	if !found {
		if len(defaultVal) == 0 {
			return "", errors.New("Value not found")
		}
		if len(defaultVal) > 1 {
			return "", errors.New("Only one default value permitted")
		}
		result = defaultVal[0]
	}
	return fmt.Sprintf("%v", result), nil
}
