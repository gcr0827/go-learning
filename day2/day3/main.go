package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	/*p := Person{
		Name: "John Doe",
		Age:  20,
	}
	data, err := json.Marshal(p)
	fmt.Println(string(data), err)

	jsonStr := `{"Name":"Bob","Age":25}`
	var p2 Person
	err = json.Unmarshal([]byte(jsonStr), &p2)
	fmt.Println(p2, err)
	*/

	config, err := loadConfig("config.json")
	fmt.Println(config, err)
}

type Person struct {
	Name string
	Age  int
}

type Config struct {
	AppName    string `json:"app_name"`
	Port       int    `json:"port"`
	TimeoutSec int    `json:"timeout_sec"`
	Debug      bool   `json:"debug"`
}

func loadConfig(path string) (Config, error) {
	config := Config{}

	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("读取文件失败: %w", err)
	}

	// 针对文件解析失败
	if err = json.Unmarshal(contentBytes, &config); err != nil {
		return config, fmt.Errorf("解析配置文件失败: %w", err)
	}
	return config, err
}
