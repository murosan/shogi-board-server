package usi

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"

	"github.com/murosan/shogi-board-server/app/domain/entity/option"
)

// ParseButton generates new Button and returns it from the given string.
func ParseButton(s string) (*option.Button, error) {
	res := buttonRegex.FindAllStringSubmatch(strings.TrimSpace(s), -1)

	if len(res) == 0 || len(res[0]) < 2 {
		msg := "failed to parse button. given=" + s + "\n" + usiOptionFormat
		return nil, errors.New(msg)
	}

	return &option.Button{Name: res[0][1]}, nil
}

// ParseCheck generates new Check and returns it from the given string.
func ParseCheck(s string) (*option.Check, error) {
	res := checkRegex.FindAllStringSubmatch(strings.TrimSpace(s), -1)

	if len(res) == 0 || len(res[0]) < 3 {
		msg := "failed to parse check. input = " + s + usiOptionFormat
		return nil, errors.New(msg)
	}

	b := res[0][2] == "true"
	return &option.Check{Name: res[0][1], Value: b, Default: b}, nil
}

// ParseRange generates new Range and returns it from the given string.
func ParseRange(s string) (*option.Range, error) {
	res := spinRegex.FindAllStringSubmatch(strings.TrimSpace(s), -1)

	errMsg := "failed to parse range. given=" + s + usiOptionFormat

	if len(res) == 0 || len(res[0]) < 5 {
		return nil, errors.New(errMsg)
	}

	init, err := strconv.Atoi(res[0][2])
	if err != nil {
		msg := errMsg + "\nDefault was not int. Value: " + res[0][2]
		return nil, errors.New(msg)
	}

	min, err := strconv.Atoi(res[0][3])
	if err != nil {
		msg := errMsg + "\nMin was not int. Value: " + res[0][3]
		return nil, errors.New(msg)
	}

	max, err := strconv.Atoi(res[0][4])
	if err != nil {
		msg := errMsg + "\nMax was not int. Value: " + res[0][4]
		return nil, errors.New(msg)
	}

	if min > max {
		msg := fmt.Sprintf(
			"%s\n%s Min value is not lesser than or equal to Max. Min: %d, Max: %d",
			errMsg,
			s,
			min,
			max,
		)
		return nil, errors.New(msg)
	}

	if init < min || init > max {
		msg := fmt.Sprintf(
			"%s\n%s Default value is not in range. Default: %d, Min: %d, Max: %d",
			errMsg,
			s,
			init,
			min,
			max,
		)
		return nil, errors.New(msg)
	}

	return &option.Range{
		Name:    res[0][1],
		Value:   init,
		Default: init,
		Min:     min,
		Max:     max,
	}, nil
}

// ParseSelect generates new Select and returns it from the given string.
func ParseSelect(s string) (*option.Select, error) {
	res := selectRegex.FindAllStringSubmatch(strings.TrimSpace(s), -1)

	errMsg := "failed to parse select. input = " + s + usiOptionFormat

	if len(res) == 0 || len(res[0]) < 4 {
		return nil, errors.New(errMsg)
	}

	init := res[0][2]
	vars := strings.Split(res[0][3], "var")

	if len(vars) < 2 {
		msg := errMsg + "\nNumber of vars must be at least one."
		return nil, errors.New(msg)
	}

	vars = vars[1:] // trim head because it's a space
	valid := false  // flag to check if the vars contains default

	for i, v := range vars {
		vars[i] = strings.TrimSpace(v)
		valid = valid || vars[i] == init
	}

	// vars にデフォルト値がない場合はエラー
	if !valid {
		msg := fmt.Sprintf(
			"%s\n Default value of Select was not in vars. Default: %s, Vars: %v",
			errMsg,
			init,
			vars,
		)
		return nil, errors.New(msg)
	}

	return &option.Select{
		Name:    res[0][1],
		Value:   init,
		Default: init,
		Vars:    vars,
	}, nil
}

// ParseTextFromStringType generates new Text and returns it from the given string.
func ParseTextFromStringType(s string) (*option.Text, error) {
	res := stringRegex.FindAllStringSubmatch(strings.TrimSpace(s), -1)
	errMsg := "failed to parse string. input = " + s + usiOptionFormat
	return parseText(res, errMsg)
}

// ParseTextFromFilenameType generates new Text and returns it from the given string.
func ParseTextFromFilenameType(s string) (*option.Text, error) {
	res := fileNameRegex.FindAllStringSubmatch(strings.TrimSpace(s), -1)
	errMsg := "failed to parse filename. input = " + s + usiOptionFormat
	return parseText(res, errMsg)
}

func parseText(res [][]string, errMsg string) (*option.Text, error) {
	if len(res) == 0 || len(res[0]) < 3 {
		return nil, errors.New(errMsg)
	}

	return &option.Text{
		Name:    res[0][1],
		Value:   res[0][2],
		Default: res[0][2],
	}, nil
}
