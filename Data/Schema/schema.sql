CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    salt TEXT NOT NULL,
    account_email TEXT NOT NULL UNIQUE,
    account_phone TEXT NOT NULL,
    user_group INTEGER NOT NULL,
    totp_key TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NOT NULL,
    created_by TEXT NOT NULL,
    updated_by TEXT NOT NULL,
    deleted_by TEXT NOT NULL,
    created_by_id BIGINT NOT NULL,
    updated_by_id BIGINT NOT NULL,
    deleted_by_id BIGINT NOT NULL,
    is_deleted BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE user_password_resets (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    reset_key TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NOT NULL,
    created_by TEXT NOT NULL,
    updated_by TEXT NOT NULL,
    deleted_by TEXT NOT NULL,
    created_by_id BIGINT NOT NULL,
    updated_by_id BIGINT NOT NULL,
    deleted_by_id BIGINT NOT NULL,
    is_deleted BIGINT NOT NULL DEFAULT 0,
    CONSTRAINT user_password_resets_users_fk
        FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE permissions (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    name TEXT NOT NULL,
    is_create BOOLEAN NOT NULL,
    is_read BOOLEAN NOT NULL,
    is_update BOOLEAN NOT NULL,
    is_delete BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NOT NULL,
    created_by TEXT NOT NULL,
    updated_by TEXT NOT NULL,
    deleted_by TEXT NOT NULL,
    created_by_id BIGINT NOT NULL,
    updated_by_id BIGINT NOT NULL,
    deleted_by_id BIGINT NOT NULL,
    is_deleted BIGINT NOT NULL DEFAULT 0,
    CONSTRAINT permissions_user_fk
        FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE foobar (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
	amount INTEGER NOT NULL,
	currency TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NOT NULL,
    created_by TEXT NOT NULL,
    updated_by TEXT NOT NULL,
    deleted_by TEXT NOT NULL,
    created_by_id BIGINT NOT NULL,
    updated_by_id BIGINT NOT NULL,
    deleted_by_id BIGINT NOT NULL,
    is_deleted BIGINT NOT NULL DEFAULT 0
);