package main

import (
    "fmt"
    "github.com/kylelemons/go-gypsy/yaml"
)

func main() {
    config, err := yaml.ReadFile("conf.yaml")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(config.Get("path")) // string
    fmt.Println(config.GetBool("enabled"))
    fmt.Println(config.GetInt("num"))
}
