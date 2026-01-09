package iso

import "strings"

func Parse(msg string) map[string]string {
	parts := strings.Split(msg, "|")
	fields := make(map[string]string)
	for _, p := range parts[1:] {
		kv := strings.Split(p, "=")
		fields[kv[0]] = kv[1]
	}
	return fields
}
