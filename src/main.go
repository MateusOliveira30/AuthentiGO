package main

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

func main() {

	goth.UseProviders(github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"))

}
