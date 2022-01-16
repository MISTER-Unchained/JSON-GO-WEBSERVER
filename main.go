package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func q(e bool) {
	if e == false {
		panic("value does not exist")
	}
}

type setupfile struct {
	Id            int               `json:"id"`
	Name          string            `json:"name"`
	Bindip        string            `json:"bindip"`
	Hostnames     []string          `json:"hostnames"`
	Routes        map[string]string `json:"routes"`
	Staticcontent map[string]string `json:"staticcontent"`
	Port          int               `json:"port"`
}

func readserverfile(s string) setupfile {
	var result setupfile
	dat, err := os.ReadFile(s)
	check(err)
	json.Unmarshal(dat, &result)
	return result
}

func log(s string) {
	fmt.Println(fmt.Sprintf("[DEV] " + s))
}

func main() {
	setup := readserverfile("./setup.json")
	gin.SetMode(gin.ReleaseMode)
	f, _ := os.Create(fmt.Sprintf("./log/" + strconv.Itoa(setup.Id) + "-" + setup.Name + "-gin.log"))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()

	for k, v := range setup.Staticcontent {
		r.Static(k, v)
		log(fmt.Sprintf("setup static content route: <\"" + k + "\">|<\"" + v + "\">"))
	}
	fmt.Println()
	for k, v := range setup.Routes {
		r.StaticFile(k, v)
		log(fmt.Sprintf("setup static route: <\"" + k + "\">|<\"" + v + "\">"))
	}

	r.GET("/test/json/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	addr := setup.Bindip + ":" + strconv.Itoa(setup.Port)
	log(fmt.Sprintf("serving on:  \"" + addr + "\"  "))

	for i := 0; i < 10; i++ {
		fmt.Println("not working dammit")
	}

	r.Run(addr)
}
