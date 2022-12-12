SET CLIENT_ENCODING = 'UTF8';
CREATE TABLE IF NOT EXISTS lgate_actionlog (
    created_at TIMESTAMP NOT NULL,
    action VARCHAR(100) NOT NULL,
    user_name VARCHAR(100) NOT NULL,
    family_name VARCHAR(100) NOT NULL,
    given_name VARCHAR(100) NOT NULL,
    school_class_name VARCHAR(500),
    school_name VARCHAR(100) NOT NULL,
    remote_address VARCHAR(100) NOT NULL,
    content_name VARCHAR(100)
);
