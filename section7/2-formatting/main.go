package main

import "fmt"

type ConfigItem struct {
	Key   string
	Value interface{} // any
	IsSet bool
}

/*
%v - default format
%+v - default format with struct name
%#v - Go-syntax format
%T - type
%t - boolean
%s - string
%q - double-quoted string with Go syntax
%f (%.2f) - float
%% - percent
*/

func (c ConfigItem) String() string {
	return fmt.Sprintf("Key: %s, Value: %s, IsSet: %t", c.Key, c.Value, c.IsSet)
}

func main() {
	appName := "EnvParser"
	version := 1.2
	port := 3000
	isEnabled := true

	status := fmt.Sprintf("Application: %s (Version: %.1f) running on port %d. Enabled: %t", appName, version, port, isEnabled)
	fmt.Println(status)

	item1 := ConfigItem{Key: "API_URL", Value: "http://localhost:3000/api", IsSet: true}
	item2 := ConfigItem{Key: "TIMEOUT_MS", Value: 1000, IsSet: true}
	item3 := ConfigItem{Key: "DEBUG_MODE", Value: false, IsSet: true}

	fmt.Printf("Item 1 (%%v): %v\n", item1)
	fmt.Printf("Item 2 (%%+v): %+v\n", item2)
	fmt.Printf("Item 3 (%%#v): %#v\n", item3)
}
