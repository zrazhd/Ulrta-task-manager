CREATE TABLE project_members(
    project_id UUID NOT NULL REFERENCES projects(project_id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    role TEXT NOT NULL CHECK(role in ('owner', 'member', 'viewer')),
    joined_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (project_id, user_id)
);