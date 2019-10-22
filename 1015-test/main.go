package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sync"

	"github.com/influxdata/influxdb/client/v2"

	"github.com/go-redis/redis"

	"example.com/m/models"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	MyDB     = "test"
	username = "admin"
	password = ""
)

var (
	cl   *redis.Client
	once sync.Once
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8")
}

func GetClient() *redis.Client {
	once.Do(func() {
		cl = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})
	})
	return cl
}

func main() {
	getRe()
	http.HandleFunc("/login", login)
	http.HandleFunc("/res", res)
	if serve := http.ListenAndServe(":8080", nil); serve != nil {
		log.Fatal(serve.Error())
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var (
		query                 map[string]string
		fields, sortBy, order []string
		offset, limit         int64
	)
	if ml, err := models.GetAllMrzdCashinLog(query, fields, sortBy, order,
		offset, limit); err != nil {
		fmt.Println("eee", err)
	} else {
		logs := []models.MrzdCashinLog{}
		for _, value := range ml {
			cashinLog := reflect.ValueOf(value).Interface().(models.MrzdCashinLog)
			//of := reflect.TypeOf(value)
			//fmt.Println(of)
			logs = append(logs, cashinLog)
		}
		fmt.Fprintf(w, "%+v \r\n", logs)
	}
	fmt.Fprintln(w, "3434")
}

func getRe() {
	cl := GetClient()
	if v, err := cl.Get("mrzd_session_key").Result(); err == nil {
		fmt.Println(v)
	}
}

func getInfluxDb() {
	conn, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://192.168.0.128:8086",
		Username: username,
		Password: password,
	})

	if err != nil {
		log.Fatal("err", err, " \r\n")
	}
	fmt.Println("conn", conn)

	if bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  MyDB,
		Precision: "s",
	}); err != nil {
		log.Fatal("NewBatchPoints err \r\n", err)
	} else {
		fmt.Println("bp", bp, " \r\n")
	}

}

func res(w http.ResponseWriter, r *http.Request) {
	// 构建查询对象
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").
		From("mrzd_cashin_log").
		String()
	fmt.Println(sql)

	// 执行 SQL 语句
	o := orm.NewOrm()
	log1 := []models.MrzdCashinLog{}
	if _, e := o.Raw(sql).QueryRows(&log1); e != nil {
		fmt.Println("err:", e.Error())
	} else {
		fmt.Fprintf(w, "%+v \r\n", log1)
	}
	fmt.Fprintln(w, "55555")
}
