package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"reflect"
)

type (
	Config struct {
		HTTP     `yaml:"http"`
		Database `yaml:"mysql"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}

	Database struct {
		URL      string `yaml:"url"`
		DBName   string `yaml:"dbname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
)

func NewConfig(defaultPasswordEnvVar string) (*Config, error) {
	cfg := Config{}

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Database.Password == "" {
		cfg.Database.Password = os.Getenv(defaultPasswordEnvVar)
	}

	return &cfg, nil
}

// Stringers

func (cfg *Config) String() string {
	toRet := ""

	// Рефлексивно вытаскиваем значением из структуры
	val := reflect.ValueOf(*cfg)

	// Получаем количество полей
	numFields := val.NumField()

	// Итерируемся по поялм
	for i := 0; i < numFields; i++ {
		field := val.Field(i)

		//fmt.Print(field.Type())
		//fmt.Print("\t implements fmt.Stringer: ")
		//fmt.Println(field.Type().Implements(reflect.TypeOf((*Stringer)(nil)).Elem()))

		// Проверяем, является ли поле структурой и реализует ли нужный нам интерфейс
		if field.Kind() == reflect.Struct &&
			field.Type().Implements(reflect.TypeOf((*fmt.Stringer)(nil)).Elem()) {
			stringMethod := field.MethodByName("String")
			res := stringMethod.Call([]reflect.Value{})

			toRet += res[0].String() + "\n"

		}

	}
	return toRet
}

func (http HTTP) String() string {
	return "HTTP:\n\tport:" +
		http.Port

}

func (database Database) String() string {
	return "Database:\n" +
		"\turl: " + database.URL + "\n" +
		"\tusername: " + database.Username + "\n" +
		"\tpassword: " + database.Password + "\n"
}

type Stringer interface {
	String() string
}
