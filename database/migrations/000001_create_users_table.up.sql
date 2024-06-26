CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,

    email VARCHAR(255) NOT NULL,

    password VARCHAR(255) NOT NULL,

    token VARCHAR(255) NULL,
    token_updated_at TIMESTAMP NULL,

    coins BIGINT DEFAULT 0 NULL,

    UNIQUE(email)
);