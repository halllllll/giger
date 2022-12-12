package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	fmt.Println("yo")
	// なぜかcsvファイルをDBにぶちこめない謎現象について、
	// ここで検証していく。
	// 直接やってみる。
	filepath := "useractionlog.csv"
	rows, err := readActionLogCsv(filepath)
	if err != nil {
		fmt.Println("fuck!!!")
		fmt.Println(err)
	}
	fmt.Println(rows)

	// ここからがオリジナル
	fmt.Println("cred")
	fmt.Println(dbEnv.DbUser, dbEnv.DbPw, dbEnv.Host, dbEnv.DbPort, dbEnv.DbName)
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv.DbUser, dbEnv.DbPw, dbEnv.Host, dbEnv.DbPort, dbEnv.DbName))
	if err != nil {
		panic(err)
	}

	txn, err := conn.Begin(*ctx)
	if err != nil {
		panic(err)
	}

	xxx, err := txn.CopyFrom(*ctx, pgx.Identifier{string(LGATE_ACTIONLOG_TABLE)}, []string{"created_at", "action", "user_name", "family_name", "given_name", "school_class_name", "school_name", "remote_address", "content_name"}, pgx.CopyFromRows(rows))
	if err != nil {
		fmt.Println(err)
	}
	if err := txn.Commit(*ctx); err != nil {
		err := fmt.Errorf("transaction error: %w", err)
		fmt.Println(err)
	}

	fmt.Println("成功？")
	fmt.Printf("result %d\n", xxx)
	for {
	}
}
