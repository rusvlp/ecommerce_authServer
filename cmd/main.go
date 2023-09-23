package main

import (
	"ca_kobercams/config"
	"fmt"
)

/*
 TODO:
	1. Посмотреть, как сюда подключается GORM + MySQL
	2. Подробнее разузнать, как работать в контроллерах (посмотреть на структуру HTTP респонзов и реквестов, а также узнать как это все реализовано в GO) [V]
	3. Посмотреть, как работают JWT Access и Refresh токены
	4. Спросить у Димы по поводу слоя репозиториев и их месте в чистой архитектуре
*/

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(cfg)

	/*

		fmt.Print("Server is going to start on port: " + cfg.HTTP.Port)
		if err = app.Run(cfg); err != nil {
			fmt.Print(err)
		}

	*/

	//structureTagTest()
}

/*
func structureTagTest() {
	var andrey user = user{
		Username: "andrey",
		Password: "javagovno",
		Email:    "andrey228@rusvlp.com",
	}

	out, err := json.MarshalIndent(andrey, "", "  ")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(string(out))
}

type user struct {
	Username string `json:"username_32273"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

*/
