package osuapi

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGetUser(t *testing.T) {
	c := testingGenClient()
	user, err := c.GetUser("peppy", OsuStandard)
	if err != nil {
		panic(err)
	}
	printUser(user)
}
func TestGetUserByID(t *testing.T) {
	c := testingGenClient()
	user, err := c.GetUserByID(140148, OsuMania)
	if err != nil {
		panic(err)
	}
	printUser(user)
}
func TestGetUserByUsername(t *testing.T) {
	c := testingGenClient()
	user, err := c.GetUserByUsername("Ikkun", Taiko)
	if err != nil {
		panic(err)
	}
	printUser(user)
}
func TestGetUserFull(t *testing.T) {
	c := testingGenClient()
	user, err := c.GetUserFull("803484", CatchTheBeat, "id", 20)
	if err != nil {
		panic(err)
	}
	printUser(user)
}

func printUser(u User) {
	fmt.Printf(`Username: %v
	ID: %v
	Count300: %v
	Count100: %v
	Count50: %v
	PlayCount: %v
	RankedScore: %v
	TotalScore: %v
	Rank: %v
	Level: %v
	PP: %v
	Accuracy: %v
	CountRankSS: %v
	CountRankS: %v
	CountRankA: %v
	Country: %v
	CountryRank: %v
	Events: %v
`, u.Username, u.ID, u.Count300, u.Count100, u.Count50, u.PlayCount, u.RankedScore, u.TotalScore, u.Rank, u.Level, u.PP, u.Accuracy, u.CountRankSS, u.CountRankS, u.CountRankA, u.Country, u.CountryRank, u.Events)
}
func testingGenClient() *APIClient {
	data, err := ioutil.ReadFile("osukey.txt")
	if err != nil {
		panic(err)
	}
	k := strings.Trim(string(data), "\r\n\t ")
	return NewClient(k)
}
