CREATE TABLE IF NOT EXISTS state
(
    id        SMALLINT UNSIGNED AUTO_INCREMENT,
    status    TINYTEXT NOT NULL,
    info      TEXT,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

# 初期値
INSERT INTO state (status, info) VALUES ('pause', '準備中...');

CREATE TABLE IF NOT EXISTS comment
(
    id           UUID        NOT NULL DEFAULT UUID(),
    user_id      varchar(32) NOT NULL,
    event_id   UUID        NOT NULL,
    text         TEXT        NOT NULL,
    created_at   DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_anonymous BOOLEAN     NOT NULL,
    color        VARCHAR(7),
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS reaction
(
    id         UUID        NOT NULL DEFAULT UUID(),
    user_id    varchar(32) NOT NULL,
    event_id UUID        NOT NULL,
    stamp_id   UUID        NOT NULL,
    created_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS event
(
    id          UUID         NOT NULL DEFAULT UUID(),
    video_id    varchar(11)  NOT NULL,
    title       varchar(64)  NOT NULL,
    thumbnail   varchar(140) NOT NULL,
    started_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ended_at    DATETIME     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    description TEXT,
    PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS token
(
    token       CHAR(44)    NOT NULL,
    creator_id  VARCHAR(32) NOT NULL,
    user_id     VARCHAR(32) NOT NULL,
    created_at  DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    event_id	UUID        NOT NULL,
    exprie_at   DATETIME    DEFAULT CURRENT_TIMESTAMP,
    description TEXT,
    PRIMARY KEY (`token`)
);
