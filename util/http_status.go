package util

func GetHttpStatus(code string) int {
	switch {
	case code == "200":
		return 200
	case code == "401":
		return 401
	case code == "400":
		return 400
	default:
		return 500
	}

}
