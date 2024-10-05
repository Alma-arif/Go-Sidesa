package config

type Config struct {
	Srv      Server
	DB       Database
	Enk      Enkripsi
	FileConf FileConf
}

type Database struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

type Server struct {
	Host string
	Port string
}

type Enkripsi struct {
	Key string
}

type FileConf struct {
	ImageDerektory string
	FileDerektory  string
	FileType       string
}
