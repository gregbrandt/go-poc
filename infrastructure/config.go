package infrastructure

import(

    "github.com/spf13/viper"
    "fmt"
    "flag"

)

type Configuration struct {

	DBUser    string   

	DBPassword  string
}

var config *Configuration

func init(){
    fileName := flag.String("envconfig", "config", "a override the environment config file name")
    
    configpath := flag.String("configpath", ".", "a override the environment config file path")
    
    flag.Parse()
    viper.SetConfigName(*fileName) // name of config file (without extension)   
    viper.AddConfigPath(*configpath)     // optionally look for config in the working directory
    err := viper.ReadInConfig() // Find and read the config file
    if err != nil { // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    config = &Configuration{
        DBUser: viper.GetString("DBUser"),
        DBPassword: viper.GetString("DBPassword"),		
	}
}

func GetConfiguration() *Configuration {
	return config
}