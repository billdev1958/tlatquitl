CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY,
	name VARCHAR(75),
	lastname1 VARCHAR(75),
	lastname2 VARCHAR(75),
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS roles (
	id UUID PRIMARY KEY,
	code VARCHAR(100) NOT NULL UNIQUE,
	name VARCHAR(100) NOT NULL UNIQUE,
	description TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS permission (
	id UUID PRIMARY KEY,
	code VARCHAR(100) NOT NULL UNIQUE,
	name VARCHAR(100) NOT NULL,
	description TEXT,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
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

CREATE TABLE IF NOT EXISTS user_roles (
	id UUID PRIMARY KEY,
	user_id UUID NOT NULL,
	role_id UUID NOT NULL,
	branch_scope_id UUID,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	UNIQUE(user_id, role_id)
);

CREATE TABLE IF NOT EXISTS role_permissions (
	id UUID PRIMARY KEY,
	role_id UUID NOT NULL,
	permission_id UUID NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	UNIQUE(role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS user_role_assignments (
	id UUID PRIMARY KEY,
	user_id UUID NOT NULL,
	role_id UUID NOT NULL,
	scope_type VARCHAR(50) NOT NULL DEFAULT 'GLOBAL',
	scope_id UUID NULL,
	effect VARCHAR(20) NOT NULL DEFAULT 'ALLOW',
	valid_from TIMESTAMPTZ NULL,
	valid_until TIMESTAMPTZ NULL,
	assigned_by UUID NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
	CHECK (
		(scope_type = 'GLOBAL' AND scope_id IS NULL)
		OR (scope_type <> 'GLOBAL' AND scope_id IS NOT NULL)
	)
);

ALTER TABLE account
ADD CONSTRAINT fk_account_user
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
FOREIGN KEY (permission_id) REFERENCES permission(id);

ALTER TABLE user_role_assignments
ADD CONSTRAINT fk_user_role_assignments_user
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE user_role_assignments
ADD CONSTRAINT fk_user_role_assignments_role
FOREIGN KEY (role_id) REFERENCES roles(id);

ALTER TABLE user_role_assignments
ADD CONSTRAINT fk_user_role_assignments_assigned_by
FOREIGN KEY (assigned_by) REFERENCES users(id);
