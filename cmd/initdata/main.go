package main

import (
	"fmt"
	"github.com/holocycle/holo-back/internal/app/config"
	"github.com/holocycle/holo-back/pkg/db"
	"github.com/holocycle/holo-back/pkg/model"
	"github.com/holocycle/holo-back/pkg/repository"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Failed to load configuration. err=%+v\n", err)
		return
	}

	db, err := db.NewDB(&config.DB)
	defer db.Close()

	userRepository := repository.NewUserRepository()
	_ = userRepository.NewQuery(db).Create(model.NewUser("user1", "test1@test.com", "https://yt3.ggpht.com/a/AATXAJwHPp_TkvcWJyblt9XVYDjNSjrj6KdpQSCQNQ=s288-c-k-c0xffffffff-no-rj-mo"))
	_ = userRepository.NewQuery(db).Create(model.NewUser("user2", "test2@test.com", "https://yt3.ggpht.com/a/AATXAJwHPp_TkvcWJyblt9XVYDjNSjrj6KdpQSCQNQ=s288-c-k-c0xffffffff-no-rj-mo"))
	_ = userRepository.NewQuery(db).Create(model.NewUser("user3", "test3@test.com", "https://yt3.ggpht.com/a/AATXAJwHPp_TkvcWJyblt9XVYDjNSjrj6KdpQSCQNQ=s288-c-k-c0xffffffff-no-rj-mo"))
}
