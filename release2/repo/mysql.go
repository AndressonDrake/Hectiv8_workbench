package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"p1-lc03.com/model"
	l "p1-lc03.com/utils/logger"
)

var (
	db *sql.DB
)

func ConnectDB(host, username, password, dbName string, port int) (err error) {
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbName))
	if err != nil {
		l.Log.Error(l.Fields{
			"host":     host,
			"username": username,
			"dbName":   dbName,
			"port":     port,
			"error":    err.Error(),
		}, nil, "failed to connect to database")
		return
	}

	// optional pool config (recommended)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil{
		l.Log.Error(l.Fields{
			"host": host,
			"username": username,
			"dbName": dbName,
			"port": port,
			"error": err.Error(),
		},nil,"failed to PING connect to database")
		return
	}

	l.Log.Info(l.Fields{},nil,"successfully connected to database")

	return
}

func GetTotalGamesSalesReport()(data []model.GetTotalGamesSalesReport,err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query :=  `SELECT gs.name, COUNT(tr.id) AS total
				FROM transactions AS tr
				JOIN games AS gs ON gs.id = tr.game_id
				GROUP BY gs.name
				ORDER BY total DESC`
	
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		l.Log.Error(l.Fields{"error": err.Error()}, nil, "query failed")
		return
	}
	defer rows.Close()

	for rows.Next(){
		var temp model.GetTotalGamesSalesReport

		err = rows.Scan(
			&temp.Name,
			&temp.Total,
		)

		if err != nil {
			l.Log.Error(l.Fields{"error": err.Error()}, nil, "scan failed")
			return
		}

		data = append(data, temp)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

func GetMostPopularGameReport()(data []model.GetMostPopularGameReport,err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query :=  `SELECT 
				g.id,
				g.name,
				COUNT(DISTINCT t.player_id) AS total
			FROM transactions t
			JOIN games g ON g.id = t.game_id
			GROUP BY g.id, g.name
			ORDER BY total DESC`
	
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		l.Log.Error(l.Fields{"error": err.Error()}, nil, "query failed")
		return
	}
	defer rows.Close()

	for rows.Next(){
		var temp model.GetMostPopularGameReport

		err = rows.Scan(
			&temp.ID,
			&temp.Name,
			&temp.Total,
		)

		if err != nil {
			l.Log.Error(l.Fields{"error": err.Error()}, nil, "scan failed")
			return
		}

		data = append(data, temp)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return


}


func GetTotalRevenuePerGameReport()(data []model.GetTotalRevenuePerGameReport,err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT g.name, SUM(g.price) AS total
			FROM transactions AS t
			JOIN games AS g ON g.id = t.game_id
			GROUP BY g.name
			ORDER BY total DESC`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		l.Log.Error(l.Fields{"error": err.Error()}, nil, "query failed")
		return
	}
	defer rows.Close()

	for rows.Next(){
		var temp model.GetTotalRevenuePerGameReport

		err = rows.Scan(
			&temp.Name,
			&temp.Total,
		)

		if err != nil {
			l.Log.Error(l.Fields{"error": err.Error()}, nil, "scan failed")
			return
		}

		data = append(data, temp)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

func GetPlayerCountPerGameReport()(data []model.GetMostPopularGameReport,err error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query :=  `SELECT 
				g.id,
				g.name,
				COUNT(DISTINCT t.player_id) AS total
			FROM transactions t
			JOIN games g ON g.id = t.game_id
			GROUP BY g.id, g.name
			ORDER BY total DESC`
	
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		l.Log.Error(l.Fields{"error": err.Error()}, nil, "query failed")
		return
	}
	defer rows.Close()

	for rows.Next(){
		var temp model.GetMostPopularGameReport

		err = rows.Scan(
			&temp.ID,
			&temp.Name,
			&temp.Total,
		)

		if err != nil {
			l.Log.Error(l.Fields{"error": err.Error()}, nil, "scan failed")
			return
		}

		data = append(data, temp)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

