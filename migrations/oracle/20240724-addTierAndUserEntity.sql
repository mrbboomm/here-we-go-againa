-- +migrate Up
CREATE TABLE tier (
    id NUMBER GENERATED BY DEFAULT ON NULL AS IDENTITY PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    name CLOB
);
CREATE TABLE user (
    id NUMBER GENERATED BY DEFAULT ON NULL AS IDENTITY PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    username VARCHAR2(255) NOT NULL,
    password VARCHAR2(255) NOT NULL,
    tier_id NUMBER,
    CONSTRAINT fk_user_tier FOREIGN KEY (tier_id) REFERENCES tier(id)
);

-- +migrate Down
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS tier;