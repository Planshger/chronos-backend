CREATE SCHEMA chronos;

CREATE TABLE chronos.users (
    id              SERIAL        PRIMARY KEY,
    version         BIGINT        NOT NULL DEFAULT 1,
    name       VARCHAR(100)  NOT NULL CHECK (char_length(name) BETWEEN 3 AND 100),

    email           VARCHAR(255)  NOT NULL UNIQUE,  
    password_hash   VARCHAR(255)  NOT NULL,

    role            VARCHAR(20)   NOT NULL DEFAULT 'user' CHECK (role IN ('user', 'admin')),

    plan            VARCHAR(30)   NULL,  
    status          VARCHAR(20)   NOT NULL DEFAULT 'pending-payment' CHECK (status IN ('active', 'blocked', 'pending-payment')),

    created_at      TIMESTAMPTZ   NOT NULL,
    updated_at      TIMESTAMPTZ   NOT NULL
);

CREATE TABLE chronos.tasks (
    id              SERIAL          PRIMARY KEY,
    version         BIGINT          NOT NULL DEFAULT 1,
    title           VARCHAR(100)    NOT NULL CHECK (char_length(title) BETWEEN 1 AND 100),
    description     VARCHAR(1000)   CHECK (char_length(description) BETWEEN 1 AND 1000),
    type            VARCHAR(50)     CHECK (type IN ('задача', 'встреча')),
    completed       BOOLEAN         NOT NULL,

    created_at      TIMESTAMPTZ     NOT NULL,
    completed_at    TIMESTAMPTZ,

    CHECK (
        (completed=FALSE AND completed_at IS NULL)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    ),

    user_id INTEGER NOT NULL REFERENCES chronos.users(id) ON DELETE CASCADE
);

CREATE TABLE chronos.plans (
    id               SERIAL         PRIMARY KEY,
    version          BIGINT         NOT NULL DEFAULT 1,
    name             VARCHAR(100)   NOT NULL UNIQUE CHECK (char_length(name) BETWEEN 1 AND 100),
    price            NUMERIC(10,2)  NOT NULL CHECK (price >= 0),
    description      VARCHAR(500)   NOT NULL CHECK (char_length(description) BETWEEN 1 AND 500),
    duration_days    SMALLINT       NOT NULL CHECK (duration_days > 0),
    popular          BOOLEAN        NOT NULL DEFAULT false,
    active           BOOLEAN        NOT NULL DEFAULT false,
    trial            BOOLEAN        NOT NULL DEFAULT false,

    created_at       TIMESTAMPTZ    NOT NULL DEFAULT now(),
    updated_at       TIMESTAMPTZ    NOT NULL DEFAULT now()
);

CREATE TABLE chronos.plan_features (
    id          SERIAL          PRIMARY KEY,
    position    SMALLINT        NOT NULL,             
    text        VARCHAR(200)    NOT NULL CHECK (char_length(text) BETWEEN 1 AND 200),
    UNIQUE (plan_id, position),

    plan_id     INTEGER         NOT NULL REFERENCES chronos.plans(id) ON DELETE CASCADE
);