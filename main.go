package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/martini-contrib/render"
	shuffle "github.com/shogo82148/go-shuffle"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var orm *xorm.Engine

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/lunchapp/")
	viper.AddConfigPath("$HOME/.lunchapp")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	uri := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", viper.GetString("db.user"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.name"))

	log.Info("connecting to db ", uri)
	orm, err = xorm.NewEngine("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}
	orm.ShowSQL(true)
	orm.Logger().SetLevel(core.LOG_DEBUG)

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Put("/employees", AddEmployeeHandler)
	m.Delete("/employees/:id", DeleteEmployeeHandler)
	m.Get("/employees", ListEmployeeHandler)
	m.Get("/groups", GroupEmployeeHandler)
	m.Run()
	orm.Close()
}

func AddEmployeeHandler(req *http.Request, r render.Render) {
	var emp Employee
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&emp)
	if err != nil {
		r.Text(http.StatusBadRequest, err.Error())
		return
	}
	_, err = orm.Insert(emp)
	if err != nil {
		r.Text(http.StatusInternalServerError, err.Error())
		return
	}
	r.Text(http.StatusOK, "OK")
}

func ListEmployeeHandler(params martini.Params, r render.Render) {
	var emps []Employee
	err := orm.Asc("employee_name").Find(&emps)
	if err != nil {
		log.Error(err)
		r.Text(500, err.Error())
		return
	}
	r.JSON(200, emps)
}

func DeleteEmployeeHandler(params martini.Params, r render.Render) {
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		r.Text(http.StatusBadRequest, err.Error())
	}
	emp := Employee{Id: id}
	_, err = orm.Delete(&emp)
	if err != nil {
		log.Error(err)
		r.Text(500, err.Error())
		return
	}

	r.Text(204, "OK")
}

func GroupEmployeeHandler(r render.Render) {
	var emps []Employee
	err := orm.Asc("employee_name").Find(&emps)
	if err != nil {
		log.Error(err)
		r.Text(500, err.Error())
		return
	}
	dist, err := calcGroupDistribution(len(emps))
	if err != nil {
		log.Error(err)
		r.Text(500, err.Error())
		return
	}

	shuffle.Slice(emps)

	var groups []interface{}
	offset := 0
	for groupSize, groupCount := range dist {
		for i := 0; i < groupCount; i++ {
			groups = append(groups, emps[offset:offset+int(groupSize)])
			offset += int(groupSize)
		}
	}

	r.JSON(200, groups)
}
