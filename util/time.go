package util

import (
	"time"
)

// bejind 时区
var beijingLocation = time.FixedZone("Asia/Shanghai", 8*60*60)

// Now 获取1970到现在的秒数
func Now() int {
	local, err := time.LoadLocation("Local") //服务器设置的时区
	if err != nil {
		return int(time.Now().Unix())
	}
	now := time.Now().In(local)
	return int(now.Unix())
}

// FormatTime 将参数 t 格式化成北京时间 yyyyMMddHHmmss.
func FormatTime(t time.Time) string {
	return t.In(beijingLocation).Format("20060102150405")
}

// ParseTime 将北京时间 yyyyMMddHHmmss 字符串解析到 time.Time.
func ParseTime(value string) (time.Time, error) {
	return time.ParseInLocation("20060102150405", value, beijingLocation)
}
