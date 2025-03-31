package helpers

import(
	"strings"
	"fmt"
	"strconv"

)

func TrimSpace(space string) string {
	removeSpace := strings.TrimSpace(space)
	return removeSpace
}

func ParseMonthYear(input string) (int, int, error) {
	parts := strings.Split(input, "/")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("định dạng không hợp lệ, phải là MM/YYYY")
	}

	month, err1 := strconv.Atoi(parts[0])
	year, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("không thể chuyển đổi month/year sang int")
	}

	return month, year, nil
}
