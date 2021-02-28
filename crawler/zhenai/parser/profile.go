package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
)

var ageReg = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightReg = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeReg = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {

	result := engine.ParseResult{}
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageReg))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightReg))
	if err == nil {
		profile.Height = height
	}

	income := extractString(contents, incomeReg)
	if err == nil {
		profile.Income = income
	}

	result.Items = append(result.Items, profile)
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
