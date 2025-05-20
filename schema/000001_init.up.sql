CREATE TABLE IF NOT EXISTS demo_app.app_user (
    id              SERIAL          PRIMARY KEY,
    name            VARCHAR(50)     NOT NULL,
    username        VARCHAR(50)     NOT NULL UNIQUE,
    password        VARCHAR(250)     NOT NULL
);

CREATE TABLE IF NOT EXISTS demo_app.todo_list (
    id              SERIAL          PRIMARY KEY,
    title           VARCHAR(50)     NOT NULL,
    description     VARCHAR(255)    NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS demo_app.user_list (
    id              SERIAL          PRIMARY KEY,
    user_id         INT             REFERENCES demo_app.app_user(id) ON DELETE CASCADE,
    list_id         INT             REFERENCES demo_app.todo_list(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS demo_app.todo_item (
    id              SERIAL          PRIMARY KEY,
    title           VARCHAR(50)     NOT NULL,
    description     VARCHAR(255),
    done            BOOLEAN         NOT NULL,
    list_id         INT             REFERENCES demo_app.todo_list(id)
);