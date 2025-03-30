package portformat

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

const portErrMsg = "invalid port especification"

func dashSplit(args string, ports *[]int) error {

	sp := strings.Split(args, "-")

	if len(sp) != 2 {
		return errors.New(portErrMsg)
	}

	start, err := strconv.Atoi(strings.TrimSpace(sp[0]))
	if err != nil {
		return errors.New(portErrMsg)
	}

	end, err := strconv.Atoi(sp[1])
	if err != nil {
		return errors.New(portErrMsg)
	}

	if start < 1 || end > 65535 {
		return errors.New(portErrMsg)
	}

	if start > end {
		return errors.New(portErrMsg)
	}

	for ; start <= end; start++ {
		*ports = append(*ports, start)
	}

	return nil

}

func convertAndAdd(args string, ports *[]int) error {
	port, err := strconv.Atoi(strings.TrimSpace(args))
	if err != nil {
		return errors.New(portErrMsg)
	}
	if port < 1 || port > 65535 {
		return errors.New(portErrMsg)
	}
	*ports = append(*ports, port)
	return nil
}

func Parse(args string) ([]int, error) {

	var ports []int

	switch {
	case strings.Contains(args, ",") && strings.Contains(args, "-"):
		sp := strings.Split(args, ",")
		for _, arg := range sp {
			if strings.Contains(arg, "-") {
				err := dashSplit(arg, &ports)
				if err != nil {
					return nil, err
				}
			} else {
				err := convertAndAdd(arg, &ports)
				if err != nil {
					return nil, err
				}
			}
		}
	case strings.Contains(args, ","):
		sp := strings.Split(args, ",")
		for _, port := range sp {
			err := convertAndAdd(port, &ports)
			if err != nil {
				return nil, err
			}
		}
	case strings.Contains(args, "-"):
		err := dashSplit(args, &ports)
		if err != nil {
			return nil, err
		}
	default:
		err := convertAndAdd(args, &ports)
		if err != nil {
			return nil, err
		}
	}

	slices.Sort(ports)

	return ports, nil
}
