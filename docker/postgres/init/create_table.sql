SET CLIENT_ENCODING = 'UTF8';
-- create database for metabase
CREATE DATABASE giger_metabase
    ENCODING 'UTF8';


CREATE TABLE IF NOT EXISTS lgate_actionlog (
    created_at TIMESTAMP NOT NULL,
    action VARCHAR(100) NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    family_name VARCHAR(100) NOT NULL,
    given_name VARCHAR(100) NOT NULL,
    school_class_name VARCHAR(100),
    school_name VARCHAR(100) NOT NULL,
    remote_address VARCHAR(100) NOT NULL,
    content_name VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS lgate_users(
    user_name VARCHAR(100) NOT NULL,
    password VARCHAR(100),
    enabled_user SMALLINT,
    is_local SMALLINT,
    school_code VARCHAR(10),
    school_name VARCHAR(100),
    family_name VARCHAR(100),
    given_name VARCHAR(100),
    family_kana_name VARCHAR(100),
    given_kana_name VARCHAR(100),
    renew_name SMALLINT,
    renew_password SMALLINT,
    renew_class SMALLINT,

    term_name1 VARCHAR(100),
    class_name1 VARCHAR(100),
    class_role1 VARCHAR(100),
    class_number1 VARCHAR(100),

    term_name2 VARCHAR(100),
    class_name2 VARCHAR(100),
    class_role2 VARCHAR(100),
    class_number2 VARCHAR(100),

    term_name3 VARCHAR(100),
    class_name3 VARCHAR(100),
    class_role3 VARCHAR(100),
    class_number3 VARCHAR(100),

    term_name4 VARCHAR(100),
    class_name4 VARCHAR(100),
    class_role4 VARCHAR(100),
    class_number4 VARCHAR(100),

    term_name5 VARCHAR(100),
    class_name5 VARCHAR(100),
    class_role5 VARCHAR(100),
    class_number5 VARCHAR(100),

    term_name6 VARCHAR(100),
    class_name6 VARCHAR(100),
    class_role6 VARCHAR(100),
    class_number6 VARCHAR(100),

    term_name7 VARCHAR(100),
    class_name7 VARCHAR(100),
    class_role7 VARCHAR(100),
    class_number7 VARCHAR(100),

    term_name8 VARCHAR(100),
    class_name8 VARCHAR(100),
    class_role8 VARCHAR(100),
    class_number8 VARCHAR(100),

    term_name9 VARCHAR(100),
    class_name9 VARCHAR(100),
    class_role9 VARCHAR(100),
    class_number9 VARCHAR(100),

    term_name10 VARCHAR(100),
    class_name10 VARCHAR(100),
    class_role10 VARCHAR(100),
    class_number10 VARCHAR(100),

    term_name11 VARCHAR(100),
    class_name11 VARCHAR(100),
    class_role11 VARCHAR(100),
    class_number11 VARCHAR(100),

    term_name12 VARCHAR(100),
    class_name12 VARCHAR(100),
    class_role12 VARCHAR(100),
    class_number12 VARCHAR(100),

    term_name13 VARCHAR(100),
    class_name13 VARCHAR(100),
    class_role13 VARCHAR(100),
    class_number13 VARCHAR(100),

    term_name14 VARCHAR(100),
    class_name14 VARCHAR(100),
    class_role14 VARCHAR(100),
    class_number14 VARCHAR(100),

    term_name15 VARCHAR(100),
    class_name15 VARCHAR(100),
    class_role15 VARCHAR(100),
    class_number15 VARCHAR(100),

    term_name16 VARCHAR(100),
    class_name16 VARCHAR(100),
    class_role16 VARCHAR(100),
    class_number16 VARCHAR(100),

    term_name17 VARCHAR(100),
    class_name17 VARCHAR(100),
    class_role17 VARCHAR(100),
    class_number17 VARCHAR(100),

    term_name18 VARCHAR(100),
    class_name18 VARCHAR(100),
    class_role18 VARCHAR(100),
    class_number18 VARCHAR(100),

    term_name19 VARCHAR(100),
    class_name19 VARCHAR(100),
    class_role19 VARCHAR(100),
    class_number19 VARCHAR(100),

    term_name20 VARCHAR(100),
    class_name20 VARCHAR(100),
    class_role20 VARCHAR(100),
    class_number20 VARCHAR(100),

    term_name21 VARCHAR(100),
    class_name21 VARCHAR(100),
    class_role21 VARCHAR(100),
    class_number21 VARCHAR(100),

    term_name22 VARCHAR(100),
    class_name22 VARCHAR(100),
    class_role22 VARCHAR(100),
    class_number22 VARCHAR(100)

);
