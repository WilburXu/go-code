package main

import (
	"database/sql"
	"errors"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	ClickHouseClient *sqlx.DB
	ErrRecordNotFound = errors.New("sql: no rows in result set")
)

func OrmErr(err error) error {
	if err != ErrRecordNotFound && err != sql.ErrNoRows {
		return err
	}
	return nil
}

type TestData struct {
	Mid         int64 `db:"mid"`
	Expenditure int64 `db:"expenditure"`
}

func main() {
	var err error
	ClickHouseClient, err = sqlx.Open("clickhouse", "tcp://192.168.64.91:9000?debug=true")
	if err != nil {
		log.Fatal(err)
	}

	sql := "SELECT mid, sum(get_amount) as expenditure FROM live.account_recharge_order group by mid having mid=10012 and status = 2"

	items := new(TestData)
	exist, err := GetOne(items, sql)
	log.Println(exist)
	if err != nil {
		log.Println(err)
	}
}

func GetOne(data interface{}, querySQL string) (bool, error) {
	if ClickHouseClient == nil {
		return false, errors.New("clickhouse cli nil")
	}

	err := ClickHouseClient.Get(data, querySQL)
	if err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

func GetAll(list interface{}, querySQL string) (error) {
	if ClickHouseClient == nil {
		return errors.New("clickhouse cli nil")
	}

	err := ClickHouseClient.Select(list, querySQL)
	if err != nil {
		return err
	}

	return nil
}