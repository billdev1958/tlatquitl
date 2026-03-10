CREATE TABLE IF NOT EXISTS roles (
	id UUID PRIMARY KEY,
	code VARCHAR(100) NOT NULL UNIQUE,
	name VARCHAR(100) NOT NULL UNIQUE,
	description TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS permissions (
	id UUID PRIMARY KEY,
	code VARCHAR(100) NOT NULL UNIQUE,
	name VARCHAR(100) NOT NULL,
	description TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS user_roles (
	id UUID PRIMARY KEY,
	user_id UUID NOT NULL,
	role_id UUID NOT NULL,
	-- branch_scope_id UUID,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	UNIQUE(user_id, role_id, branch_scope_id)
);

CREATE TABLE IF NOT EXISTS role_permissions (
	id UUID PRIMARY KEY,
	role_id UUID NOT NULL,
	permission_id UUID NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	UNIQUE(role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY,
	name VARCHAR(75),
	lastname1 VARCHAR(75),
	lastname2 VARCHAR(75),
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS account (
	id UUID PRIMARY KEY,
	user_id UUID NOT NULL,
	email VARCHAR(75) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	is_verified BOOLEAN NOT NULL DEFAULT FALSE,
	version INTEGER DEFAULT 1,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	password_change_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	deleted_at TIMESTAMPTZ
);



ALTER TABLE account
ADD CONSTRAINT fk_user
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE user_roles
ADD CONSTRAINT fk_user_roles_user
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE user_roles
ADD CONSTRAINT fk_user_roles_role
FOREIGN KEY (role_id) REFERENCES roles(id);

ALTER TABLE role_permissions
ADD CONSTRAINT fk_role_permissions_role
FOREIGN KEY (role_id) REFERENCES roles(id);

ALTER TABLE role_permissions
ADD CONSTRAINT fk_role_permissions_permission
FOREIGN KEY (permission_id) REFERENCES permissions(id);
