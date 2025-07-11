CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE tasks(
    id              UUID            PRIMARY KEY DEFAULT uuid_generate_v4(),
    title           VARCHAR(255)    NOT NULL,
    description     VARCHAR(255)    NOT NULL,
    category_id     UUID            NOT NULL,
    priority        VARCHAR(255)    NOT NULL,
    deadline        TIMESTAMP       NOT NULL,
    created_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP  NOT NULL,
    updated_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP  NOT NULL
);

ALTER TABLE tasks ADD CONSTRAINT fk_tasks_categories FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE;