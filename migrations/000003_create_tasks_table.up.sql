CREATE TABLE tasks(
    task_id UUID NOT NULL PRIMARY KEY,
    project_id UUID NOT NULL REFERENCES projects(project_id) ON DELETE CASCADE,
    creator_id UUID REFERENCES users(user_id) ON DELETE SET NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    performer UUID REFERENCES users(user_id) ON DELETE SET NULL,
    status TEXT NOT NULL CHECK(status in ('Not Started', 'In Progress', 'Completed')),
    deadline TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);