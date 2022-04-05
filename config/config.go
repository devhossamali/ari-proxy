package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var serviceID string

// SetServiceID ..
func SetServiceID(id string) {
	serviceID = id
}

// GetServiceID ..
func GetServiceID() string {
	return serviceID
}

// Get return OS env variable if found or viper variable
// varName should be passed by default to match standard
// os env variable ex: GIN_HOST == gin.host in viper config.
func Get(varName string) string {

	// System environment variables ..
	if v := os.Getenv(varName); v != "" {
		return v
	}

	systemVarName := strings.Replace(varName, ".", "_", -1)
	if v := os.Getenv(strings.ToUpper(systemVarName)); v != "" {
		return v
	}

	// Local config file .yaml ..
	if v := viper.GetString(varName); v != "" {
		return v
	}

	varName = strings.Replace(varName, "_", ".", 1) // ex. to be "GIN_HOST"
	if v := viper.GetString(varName); v != "" {     // ex. to be "gin.host"
		return v
	}

	log.Printf("Varaible: %s is not defined anywhere!", varName)
	return ""
}

// GetDestinations return service event destinations array
// destination services i.e. CP_MSG, CP_TABLES
// returns array of service names []string{"CP_MSG", "CP_TABLES"}
func GetDestinations(module, operation string) []string {
	varName := "service.destinations." + module + "." + operation
	return strings.Split(viper.GetString(varName), ",")
}
