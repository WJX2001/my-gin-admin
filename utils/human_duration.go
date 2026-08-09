package utils

import (
	"strconv"
	"strings"
	"time"
)

func ParseDuration(d string) (time.Duration, error) {
	// 去掉首尾空格
	d = strings.TrimSpace(d)
	// 先用标准库解析，能认 2h 30m 1h30m 等
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}

	// 配置文件里带 d
	if strings.Contains(d, "d") {
		// 找到 d 出现的位置 7d -> index = 1
		index := strings.Index(d, "d")

		// 7d -> 7
		dayCount, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(dayCount)
		// 解析 d 后面的部分 7d2h -> 2h
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
