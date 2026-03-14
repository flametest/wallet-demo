CREATE TABLE wallet
(
    id         BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    version    BIGINT UNSIGNED                    NOT NULL,
    name       VARCHAR(64)                        NOT NULL,
    display_id VARCHAR(64)                        NOT NULL,
    balance    DECIMAL(10, 2)                     NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

CREATE UNIQUE INDEX uidx_name ON wallet(name);
CREATE UNIQUE INDEX uidx_display_id ON wallet(display_id);