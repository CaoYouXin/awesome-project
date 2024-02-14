package res

import "fmt"

func wrapText(err error, text string) error {
	return fmt.Errorf("%s(%s)", text, err.Error())
}

func DaoErr1(err error) error {
	return wrapText(err, "DAO error")
}

func ServiceErr1(err error) error {
	return wrapText(err, "SERVICE error")
}

func DaoErr2(err error) (error, string) {
	return wrapText(err, "DAO error"), ""
}

func ServiceErr2(err error) (error, string) {
	return wrapText(err, "SERVICE error"), ""
}
