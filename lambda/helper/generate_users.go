package helper

import (
	"github.com/achaquisse/skulla-api/model"
	"github.com/jaswdr/faker"
)

func GenerateUsers(n int) []model.User {
	fake := faker.New()

	var users []model.User
	for i := 0; i < n; i++ {
		user := model.User{
			ID:       fake.Int(),
			Name:     fake.Person().FirstName(),
			Username: fake.Internet().User(),
			Email:    fake.Person().Contact().Email,
		}

		users = append(users, user)
	}

	return users
}
