CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    is_private BOOLEAN DEFAULT FALSE,
    created_by INT REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE group_users (
    group_id INT REFERENCES chat_groups(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (group_id, user_id)
);

CREATE TABLE group_messages (
    id SERIAL PRIMARY KEY,
    group_id INT REFERENCES chat_groups(id) ON DELETE CASCADE,
    sender_id INT REFERENCES users(id),
    message TEXT,
    media_url TEXT, -- for images, files, etc.
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE direct_messages (
    id SERIAL PRIMARY KEY,
    sender_id INT REFERENCES users(id),
    receiver_id INT REFERENCES users(id),
    message TEXT,
    media_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE message_status (
    message_id INT,
    user_id INT,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    PRIMARY KEY (message_id, user_id)
    -- message_id can refer to direct OR group message depending on design
);

CREATE TABLE media_uploads (
    id SERIAL PRIMARY KEY,
    uploaded_by INT REFERENCES users(id),
    file_url TEXT NOT NULL,
    file_type VARCHAR(50),
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
