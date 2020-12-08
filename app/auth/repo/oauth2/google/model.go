package google

import (
	"errors"
	"red-auth/domain/auth"
	"strconv"
	"strings"
	"time"
)

type AuthorizedUserInfo struct {
	ResourceName   string         `json:"resourceName"`
	Etag           string         `json:"etag"`
	Names          []Name         `json:"names"`
	Photos         []Photo        `json:"photos"`
	Genders        []Gender       `json:"genders"`
	Birthdays      []Birthday     `json:"birthdays"`
	EmailAddresses []EmailAddress `json:"emailAddresses"`
}

type Birthday struct {
	Metadata BirthdayMetadata `json:"metadata"`
	Date     Date             `json:"date"`
}

type Date struct {
	Year  int64 `json:"year"`
	Month int64 `json:"month"`
	Day   int64 `json:"day"`
}

type Gender struct {
	FormattedValue string `json:"formattedValue"`
	Value          string `json:"value"`
}

type BirthdayMetadata struct {
	Primary bool   `json:"primary"`
	Source  Source `json:"source"`
}

type Source struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type EmailAddress struct {
	Metadata EmailAddressMetadata `json:"metadata"`
	Value    string               `json:"value"`
}

type EmailAddressMetadata struct {
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
	Source   Source `json:"source"`
}

type Name struct {
	Metadata             BirthdayMetadata `json:"metadata"`
	DisplayName          string           `json:"displayName"`
	FamilyName           string           `json:"familyName"`
	GivenName            string           `json:"givenName"`
	DisplayNameLastFirst string           `json:"displayNameLastFirst"`
	UnstructuredName     string           `json:"unstructuredName"`
}

type Photo struct {
	Metadata BirthdayMetadata `json:"metadata"`
	URL      string           `json:"url"`
	Default  bool             `json:"default"`
}

func (g *AuthorizedUserInfo) toUserInfo() (*auth.UserInfo, error) {
	if len(g.Names) < 1 {
		return nil, errors.New("names empty array")
	}
	if len(g.EmailAddresses) < 1 {
		return nil, errors.New("emailAddresses empty array")
	}
	if len(g.Photos) < 1 {
		return nil, errors.New("photos empty array")
	}

	if len(g.Genders) < 1 {
		return nil, errors.New("genders empty array")
	}
	gender := ""
	switch g.Genders[0].FormattedValue {
	case "Male":
		gender = "M"
	case "Female":
		gender = "F"
	default:
		gender = ""
	}

	if len(g.Birthdays) < 1 {
		return nil, errors.New("birthdays empty array")
	}
	gDay := leftPad2Len(strconv.FormatInt(g.Birthdays[0].Date.Day, 10), "0", 2)
	gMonth := leftPad2Len(strconv.FormatInt(g.Birthdays[0].Date.Month, 10), "0", 2)
	gYear := strconv.FormatInt(g.Birthdays[0].Date.Year, 10)
	t, _ := time.Parse("01022006", gMonth+gDay+gYear)

	return &auth.UserInfo{
		ID:       g.Names[0].Metadata.Source.ID,
		AuthType: "Google",
		Name:     g.Names[0].DisplayName,
		Email:    g.EmailAddresses[0].Value,
		Birthday: t,
		Gender:   gender,
		PhotoURL: g.Photos[0].URL,
	}, nil
}

func leftPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}
