CREATE TABLE comments(
    comment_id UUID PRIMARY KEY,
    task_id UUID REFERENCES tasks(task_id) ON DELETE CASCADE,
    creator_id UUID REFERENCES users(user_id) ON DELETE SET NULL,
    message text NOT NULL,
    created_at TIMESTAMP
);