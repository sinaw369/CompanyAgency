package utility

import "errors"

func Alpha(value string) (bool, error) {
	for _, val := range value {
		if val < 65 && val > 90 {
			return false, errors.New("the string does not consist of letters")
		}

	}
	return true, nil
}
