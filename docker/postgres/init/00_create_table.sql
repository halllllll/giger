SET CLIENT_ENCODING = 'UTF8';
-- create database for metabase
CREATE DATABASE giger_metabase
    ENCODING 'UTF8';

-- いちおう複合キーをつけとく
-- ただし同一csvファイル内でレコードが重複することがあるので、それはアプリ側ではじく
CREATE TABLE IF NOT EXISTS lgate_actionlog (
    created_at TIMESTAMP NOT NULL,
    action TEXT NOT NULL,
    user_name TEXT NOT NULL,
    family_name TEXT,
    given_name TEXT,
    school_class_name TEXT,
    school_name TEXT NOT NULL,
    remote_address TEXT NOT NULL,
    content_name TEXT,
    PRIMARY KEY(created_at, user_name, remote_address)
);

-- ユーザーデータはログを取る必要がないので、取得したらまるっとDELETEして更新する
CREATE TABLE IF NOT EXISTS lgate_users(
    user_name TEXT NOT NULL,
    password TEXT,
    enabled_user SMALLINT,
    is_local SMALLINT,
    school_code TEXT,
    school_name TEXT,
    family_name TEXT,
    given_name TEXT,
    family_kana_name TEXT,
    given_kana_name TEXT,
    renew_name SMALLINT,
    renew_password SMALLINT,
    renew_class SMALLINT,

    term_name1 TEXT,
    class_name1 TEXT,
    class_role1 TEXT,
    class_number1 TEXT,

    term_name2 TEXT,
    class_name2 TEXT,
    class_role2 TEXT,
    class_number2 TEXT,

    term_name3 TEXT,
    class_name3 TEXT,
    class_role3 TEXT,
    class_number3 TEXT,

    term_name4 TEXT,
    class_name4 TEXT,
    class_role4 TEXT,
    class_number4 TEXT,

    term_name5 TEXT,
    class_name5 TEXT,
    class_role5 TEXT,
    class_number5 TEXT,

    term_name6 TEXT,
    class_name6 TEXT,
    class_role6 TEXT,
    class_number6 TEXT,

    term_name7 TEXT,
    class_name7 TEXT,
    class_role7 TEXT,
    class_number7 TEXT,

    term_name8 TEXT,
    class_name8 TEXT,
    class_role8 TEXT,
    class_number8 TEXT,

    term_name9 TEXT,
    class_name9 TEXT,
    class_role9 TEXT,
    class_number9 TEXT,

    term_name10 TEXT,
    class_name10 TEXT,
    class_role10 TEXT,
    class_number10 TEXT,

    term_name11 TEXT,
    class_name11 TEXT,
    class_role11 TEXT,
    class_number11 TEXT,

    term_name12 TEXT,
    class_name12 TEXT,
    class_role12 TEXT,
    class_number12 TEXT,

    term_name13 TEXT,
    class_name13 TEXT,
    class_role13 TEXT,
    class_number13 TEXT,

    term_name14 TEXT,
    class_name14 TEXT,
    class_role14 TEXT,
    class_number14 TEXT,

    term_name15 TEXT,
    class_name15 TEXT,
    class_role15 TEXT,
    class_number15 TEXT,

    term_name16 TEXT,
    class_name16 TEXT,
    class_role16 TEXT,
    class_number16 TEXT,

    term_name17 TEXT,
    class_name17 TEXT,
    class_role17 TEXT,
    class_number17 TEXT,

    term_name18 TEXT,
    class_name18 TEXT,
    class_role18 TEXT,
    class_number18 TEXT,

    term_name19 TEXT,
    class_name19 TEXT,
    class_role19 TEXT,
    class_number19 TEXT,

    term_name20 TEXT,
    class_name20 TEXT,
    class_role20 TEXT,
    class_number20 TEXT,

    term_name21 TEXT,
    class_name21 TEXT,
    class_role21 TEXT,
    class_number21 TEXT,

    term_name22 TEXT,
    class_name22 TEXT,
    class_role22 TEXT,
    class_number22 TEXT

);
