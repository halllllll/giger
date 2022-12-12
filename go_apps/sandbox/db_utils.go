package main

import (
	"context"
	"embed"
	"fmt"

	"github.com/halllllll/golog"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type ActionLog struct {
	CreatedAt       string `csv:"createdAt"`
	Action          string `csv:"action"`
	UserName        string `csv:"username"`
	FamilyName      string `csv:"familyName"`
	GivenName       string `csv:"givenName"`
	SchoolClassName string `csv:"schoolClassName"`
	SchoolName      string `csv:"schoolName"`
	RemoteAddress   string `csv:"remoteAddress"`
	ContentName     string `csv:"contentName"`
}

type Users struct {
}

//go:embed secret.json
var Env embed.FS

type EnvJson struct {
	Host   string `json:"host"`
	DbPort string `json:"psql_port"`
	DbName string `json:"psql_dbname"`
	DbUser string `json:"psql_user"`
	DbPw   string `json:"psql_pw"`
}

type DbInfo struct {
	host   string
	port   string
	dbName string
	user   string
	pw     string
}

type TableName string

var (
	LGATE_ACTIONLOG_TABLE TableName = "lgate_actionlog"
	LGATE_USER_TABLE     TableName = "lgate_users"
)

func pourActionLog(csvRawData [][]interface{}, txn pgx.Tx, ctx context.Context, tableName string) error {
	golog.InfoLog.Printf("here is ActionLog Method!!!!!! table name: %s\n", tableName)
	_, err := txn.CopyFrom(ctx, pgx.Identifier{tableName}, []string{"created_at", "lgate_action", "user_name", "family_name", "given_name", "school_class_name", "school_name", "remote_address", "content_name"}, pgx.CopyFromRows(csvRawData))

	if err != nil {
		err := fmt.Errorf("actin log csv pouring db error: %w", err)
		return err
	}
	if err := txn.Commit(ctx); err != nil {
		err := fmt.Errorf("transaction error: %w", err)
		return err
	}
	return nil
}
