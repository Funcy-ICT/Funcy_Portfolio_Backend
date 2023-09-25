package main

import (
	"backend/app/configs"
	"backend/app/infrastructure"
	"backend/app/interfaces/request"
	"backend/app/packages/utils/auth"
	"backend/app/usecase"
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	// Exit Code
	success = iota
	dbFail
	createUserFail
)

func main() {
	code, err := do()
	if err != nil {
		fmt.Println(err)
		os.Exit(code)
	} else {
		os.Exit(success)
	}
}

func do() (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	mail := getStdinWithDefault(*scanner, "mail address", "funcy@fun.ac.jp")

	db, err := configs.Init()
	if err != nil {
		return dbFail, err
	} else {
		log.Println("db ok")
	}
	authRepository := infrastructure.NewUserRepository(db)
	authUseCase := usecase.NewAuthUseCase(authRepository)
	log.Println("DI ok")

	userExist := false
	userID := ""

	log.Println("check the account exists")
	{
		u, err := authRepository.GetPassword(mail)
		userExist = err == nil
		userID = u.UserID
	}

	if userExist {
		log.Println("the user exist")
		log.Println("skipped: create new super account")
	} else {
		log.Println("create new super account")
		uid, err := authUseCase.CreateAccount(
			request.SignUpRequest{
				Icon:        "",
				FamilyName:  "",
				FirstName:   "",
				Mail:        mail,
				Password:    "",
				Grade:       "",
				Course:      "",
				DisplayName: "",
			},
		)
		if err != nil {
			return createUserFail, err
		}
		userID = uid
	}

	log.Println("create super user token")
	token, err := auth.IssueSuperUserToken(userID)
	if err != nil {
		return 0, err
	}

	fmt.Printf("token: %s\n", token)
	return 0, nil
}

func getStdin(scanner bufio.Scanner, msg string) string {
	fmt.Printf("%s: ", msg)
	scanner.Scan()
	return scanner.Text()
}

func getStdinWithDefault(scanner bufio.Scanner, msg string, defaultVal string) string {
	val := getStdin(scanner, fmt.Sprintf("%s (default: %s)", msg, defaultVal))
	if val == "" {
		return defaultVal
	} else {
		return scanner.Text()
	}
}
