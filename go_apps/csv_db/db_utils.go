package main

import (
	"context"
	"embed"
	"fmt"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type ActionLog struct {
	CreatedAt       string `csv:"createdAt"`
	LgateAction     string `csv:"action"`
	UserName        string `csv:"username"`
	FamilyName      string `csv:"familyName"`
	GivenName       string `csv:"givenName"`
	SchoolClassName string `csv:"schoolClassName"`
	SchoolName      string `csv:"schoolName"`
	RemoteAddress   string `csv:"remoteAddress"`
	ContentName     string `csv:"contentName"`
}

type Users struct {
	UserName       string `csv:"username"`
	Password       string `csv:"password"`
	EnabledUser    int    `csv:"enabledUser"`
	IsLocal        int    `csv:"isLocal"`
	SchoolCode     string `csv:"schoolCode"`
	SchoolName     string `csv:"schoolName"`
	FamilyName     string `csv:"familyName"`
	GivenName      string `csv:"givenName"`
	FamilyKanaName string `csv:"familyKanaName"`
	GivenKanaName  string `csv:"givenKanaName"`
	RenewName      int    `csv:"renewName"`
	RenewPassword  int    `csv:"renewPassword"`
	RenewClass     int    `csv:"renewClass"`

	TermName1    string `csv:"termName"`
	ClassName1   string `csv:"className"`
	ClassRole1   string `csv:"classRole"`
	ClassNumber1 string `csv:"classNumber"`

	TermName2    string `csv:"termName"`
	ClassName2   string `csv:"className"`
	ClassRole2   string `csv:"classRole"`
	ClassNumber2 string `csv:"classNumber"`

	TermName3    string `csv:"termName"`
	ClassName3   string `csv:"className"`
	ClassRole3   string `csv:"classRole"`
	ClassNumber3 string `csv:"classNumber"`

	TermName4    string `csv:"termName"`
	ClassName4   string `csv:"className"`
	ClassRole4   string `csv:"classRole"`
	ClassNumber4 string `csv:"classNumber"`

	TermName5    string `csv:"termName"`
	ClassName5   string `csv:"className"`
	ClassRole5   string `csv:"classRole"`
	ClassNumber5 string `csv:"classNumber"`

	TermName6    string `csv:"termName"`
	ClassName6   string `csv:"className"`
	ClassRole6   string `csv:"classRole"`
	ClassNumber6 string `csv:"classNumber"`

	TermName7    string `csv:"termName"`
	ClassName7   string `csv:"className"`
	ClassRole7   string `csv:"classRole"`
	ClassNumber7 string `csv:"classNumber"`

	TermName8    string `csv:"termName"`
	ClassName8   string `csv:"className"`
	ClassRole8   string `csv:"classRole"`
	ClassNumber8 string `csv:"classNumber"`

	TermName9    string `csv:"termName"`
	ClassName9   string `csv:"className"`
	ClassRole9   string `csv:"classRole"`
	ClassNumber9 string `csv:"classNumber"`

	TermName10    string `csv:"termName"`
	ClassName10   string `csv:"className"`
	ClassRole10   string `csv:"classRole"`
	ClassNumber10 string `csv:"classNumber"`

	TermName11    string `csv:"termName"`
	ClassName11   string `csv:"className"`
	ClassRole11   string `csv:"classRole"`
	ClassNumber11 string `csv:"classNumber"`

	TermName12    string `csv:"termName"`
	ClassName12   string `csv:"className"`
	ClassRole12   string `csv:"classRole"`
	ClassNumber12 string `csv:"classNumber"`

	TermName13    string `csv:"termName"`
	ClassName13   string `csv:"className"`
	ClassRole13   string `csv:"classRole"`
	ClassNumber13 string `csv:"classNumber"`

	TermName14    string `csv:"termName"`
	ClassName14   string `csv:"className"`
	ClassRole14   string `csv:"classRole"`
	ClassNumber14 string `csv:"classNumber"`

	TermName15    string `csv:"termName"`
	ClassName15   string `csv:"className"`
	ClassRole15   string `csv:"classRole"`
	ClassNumber15 string `csv:"classNumber"`

	TermName16    string `csv:"termName"`
	ClassName16   string `csv:"className"`
	ClassRole16   string `csv:"classRole"`
	ClassNumber16 string `csv:"classNumber"`

	TermName17    string `csv:"termName"`
	ClassName17   string `csv:"className"`
	ClassRole17   string `csv:"classRole"`
	ClassNumber17 string `csv:"classNumber"`

	TermName18    string `csv:"termName"`
	ClassName18   string `csv:"className"`
	ClassRole18   string `csv:"classRole"`
	ClassNumber18 string `csv:"classNumber"`

	TermName19    string `csv:"termName"`
	ClassName19   string `csv:"className"`
	ClassRole19   string `csv:"classRole"`
	ClassNumber19 string `csv:"classNumber"`

	TermName20    string `csv:"termName"`
	ClassName20   string `csv:"className"`
	ClassRole20   string `csv:"classRole"`
	ClassNumber20 string `csv:"classNumber"`

	TermName21    string `csv:"termName"`
	ClassName21   string `csv:"className"`
	ClassRole21   string `csv:"classRole"`
	ClassNumber21 string `csv:"classNumber"`

	TermName22    string `csv:"termName"`
	ClassName22   string `csv:"className"`
	ClassRole22   string `csv:"classRole"`
	ClassNumber22 string `csv:"classNumber"`
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

type TableName string

var (
	LGATE_ACTIONLOG_TABLE TableName = "lgate_actionlog"
	LGATE_USER_TABLE      TableName = "lgate_users"
)

func flowActionCsvToDB(rows [][]interface{}) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv.DbUser, dbEnv.DbPw, dbEnv.Host, dbEnv.DbPort, dbEnv.DbName))
	if err != nil {
		err := fmt.Errorf("connect db error: %w", err)
		return err
	}
	txn, err := conn.Begin(ctx)
	if err != nil {
		err := fmt.Errorf("begin error : %w", err)
		return err
	}
	rowCounts, err := txn.CopyFrom(ctx, pgx.Identifier{string(LGATE_ACTIONLOG_TABLE)}, []string{"created_at", "action", "user_name", "family_name", "given_name", "school_class_name", "school_name", "remote_address", "content_name"}, pgx.CopyFromRows(rows))

	if err != nil {
		err := fmt.Errorf("copy csv error: %w", err)
		return err
	}
	fmt.Println(rowCounts)
	if err := txn.Commit(ctx); err != nil {
		err := fmt.Errorf("transaction error: %w", err)
		return err
	}
	return nil
}

func flowUsersCsvToDB(rows [][]interface{}) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv.DbUser, dbEnv.DbPw, dbEnv.Host, dbEnv.DbPort, dbEnv.DbName))
	if err != nil {
		err := fmt.Errorf("connect db error: %w", err)
		return err
	}

	txn, err := conn.Begin(ctx)
	if err != nil {
		err := fmt.Errorf("begin error : %w", err)
		return err
	}
	_, err = txn.Exec(ctx, fmt.Sprintf("DELETE FROM %s;", string(LGATE_USER_TABLE)))
	if err != nil {
		err := fmt.Errorf("delete user table error: %w", err)
		return err
	}
	rowCounts, err := txn.CopyFrom(ctx, pgx.Identifier{string(LGATE_USER_TABLE)}, []string{"user_name", "password", "enabled_user", "is_local", "school_code", "school_name", "family_name", "given_name", "family_kana_name", "given_kana_name", "renew_name", "renew_password", "renew_class", "term_name1", "class_name1", "class_role1", "class_number1", "term_name2", "class_name2", "class_role2", "class_number2", "term_name3", "class_name3", "class_role3", "class_number3", "term_name4", "class_name4", "class_role4", "class_number4", "term_name5", "class_name5", "class_role5", "class_number5", "term_name6", "class_name6", "class_role6", "class_number6", "term_name7", "class_name7", "class_role7", "class_number7", "term_name8", "class_name8", "class_role8", "class_number8", "term_name9", "class_name9", "class_role9", "class_number9", "term_name10", "class_name10", "class_role10", "class_number10", "term_name11", "class_name11", "class_role11", "class_number11", "term_name12", "class_name12", "class_role12", "class_number12", "term_name13", "class_name13", "class_role13", "class_number13", "term_name14", "class_name14", "class_role14", "class_number14", "term_name15", "class_name15", "class_role15", "class_number15", "term_name16", "class_name16", "class_role16", "class_number16", "term_name17", "class_name17", "class_role17", "class_number17", "term_name18", "class_name18", "class_role18", "class_number18", "term_name19", "class_name19", "class_role19", "class_number19", "term_name20", "class_name20", "class_role20", "class_number20", "term_name21", "class_name21", "class_role21", "class_number21", "term_name22", "class_name22", "class_role22", "class_number22"}, pgx.CopyFromRows(rows))

	if err != nil {
		err := fmt.Errorf("copy csv error: %w", err)
		return err
	}
	fmt.Println(rowCounts)
	if err := txn.Commit(ctx); err != nil {
		err := fmt.Errorf("transaction error: %w", err)
		return err
	}

	return nil
}
