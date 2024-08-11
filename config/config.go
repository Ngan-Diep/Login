package config
import(
	os"
    "log"
    "gopkg.in/yaml.v2"
)

type Config struct {
	database Dbconfig `yaml:"database"`
}

type Dbconfig struct {
	User string `yaml:"user"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Name string `yaml:"name"`
}


var Cgf Config
func loadConfig(yaml.Config string ) error{
    file, err := os.Open(yaml.config)
	if err!= nil {
		log.Fatalf("Error opening config file:%v",err)
		return err
	}

	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&Cgf),err!=nil{
		log.Fatal("Error decoding config file:%v",err)
		return err
	}
	return nil