package slug

import "strings"
import "regexp"

func Slugify(str, replace, separator string) string {
	str = strings.ToLower(str)

	if replace == "" {
		replace = "[^a-z0-9]"
	}

	if separator == "" {
		separator = "-"
	}

	re := regexp.MustCompile(replace)
	str = re.ReplaceAllString(str, " ")

	re = regexp.MustCompile("^ +| +$")
	str = re.ReplaceAllString(str, "")

	re = regexp.MustCompile(" +")
	return re.ReplaceAllString(str, separator)
}
