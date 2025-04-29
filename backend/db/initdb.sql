CREATE TABLE users(
    user_id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE messages(
    message_id SERIAL PRIMARY KEY,
    source_id INT NOT NULL,
    dest_id INT NOT NULL,
    time VARCHAR(50) NOT NULL,
    text TEXT
);

CREATE TABLE friends(
    user_id INT NOT NULL,
    friend_id INT NOT NULL,
    PRIMARY KEY(user_id, friend_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (friend_id) REFERENCES users(user_id) ON DELETE CASCADE
);