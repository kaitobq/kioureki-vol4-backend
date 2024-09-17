CREATE TABLE IF NOT EXISTS user_organization_invitations (
    id char(26) PRIMARY KEY,
    organization_id char(26) NOT NULL,
    user_id char(26) NOT NULL,
    invited_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (organization_id) REFERENCES organizations(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
