-- +goose Up
-- +goose StatementBegin
CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE
    EXTENSION IF NOT EXISTS citext;
CREATE
    EXTENSION IF NOT EXISTS moddatetime;

CREATE TABLE users
(
    id                 UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    name               VARCHAR(32)              NOT NULL DEFAULT '',
    email              VARCHAR(64) UNIQUE       NOT NULL CHECK (email <> ''),
    password           VARCHAR(250),
    role               VARCHAR(10)              NOT NULL CHECK (role <> ''),
    avatar             VARCHAR(512),
    status             VARCHAR(250)             NOT NULL CHECK (status <> ''),
    token_join_value   UUID,
    token_join_expires TIMESTAMP WITH TIME ZONE,
    created_at         TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at         TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sessions
(
    id         UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    token      VARCHAR(255) UNIQUE      NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    user_id    UUID                     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    ip         VARCHAR(50)              NOT NULL,
    user_agent VARCHAR(255)             NOT NULL CHECK (user_agent <> ''),
    origin     VARCHAR(100)             NOT NULL CHECK (origin <> ''),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_networks
(
    id          UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    user_id     UUID                     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    client_id   VARCHAR(255) UNIQUE      NOT NULL CHECK (client_id <> ''),
    client_type VARCHAR(100)             NOT NULL CHECK (client_type <> ''),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE currencies
(
    id          UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    title       VARCHAR(10) UNIQUE       NOT NULL CHECK (title <> ''),
    description VARCHAR(100)             NOT NULL CHECK (description <> ''),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE providers
(
    id          UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    title       VARCHAR(100) UNIQUE      NOT NULL CHECK (title <> ''),
    description VARCHAR(250)                      DEFAULT '',
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE instruments
(
    id          UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    title       VARCHAR(100) UNIQUE      NOT NULL CHECK (title <> ''),
    description VARCHAR(250)             NOT NULL CHECK (description <> ''),
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE markets
(
    id            UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    title         VARCHAR(250)             NOT NULL CHECK (title <> ''),
    ticker        VARCHAR(50) UNIQUE       NOT NULL CHECK (ticker <> ''),
    content       TEXT                              DEFAULT '',
    image_url     VARCHAR(1024) CHECK (image_url <> ''),
    currency_id   UUID                     NOT NULL REFERENCES currencies (id) ON DELETE CASCADE,
    instrument_id UUID                     NOT NULL REFERENCES instruments (id) ON DELETE CASCADE,
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE registers
(
    id          UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    identify    VARCHAR(250) UNIQUE      NOT NULL CHECK (identify <> ''),
    provider_id UUID                     NOT NULL REFERENCES providers (id) ON DELETE CASCADE,
    market_id   UUID                     NOT NULL REFERENCES markets (id) ON DELETE CASCADE,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE portfolios
(
    id          UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    title       VARCHAR(250)             NOT NULL,
    active      BOOLEAN                  NOT NULL DEFAULT FALSE,
    user_id     UUID                     NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    currency_id UUID                     NOT NULL REFERENCES currencies (id) ON DELETE CASCADE,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE assets
(
    id           UUID PRIMARY KEY                  DEFAULT uuid_generate_v4(),
    amount       INTEGER                  NOT NULL,
    quantity     INTEGER                  NOT NULL,
    portfolio_id UUID                     NOT NULL REFERENCES portfolios (id) ON DELETE CASCADE,
    market_id    UUID                     NOT NULL REFERENCES markets (id) ON DELETE CASCADE,
    notation_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS users_token_join_value_id_idx ON users (token_join_value);
CREATE INDEX IF NOT EXISTS sessions_token_id_idx ON sessions (token);
CREATE INDEX IF NOT EXISTS user_networks_client_id_id_idx ON user_networks (client_id);
CREATE INDEX IF NOT EXISTS markets_title_id_idx ON markets (title);
CREATE INDEX IF NOT EXISTS markets_ticker_id_idx ON markets (ticker);
CREATE INDEX IF NOT EXISTS registers_identify_id_idx ON registers (identify);

-- index for pagination
CREATE INDEX IF NOT EXISTS users_created_at_id_idx ON users (created_at, id);
CREATE INDEX IF NOT EXISTS markets_created_at_id_idx ON markets (created_at, id);
CREATE INDEX IF NOT EXISTS portfolios_created_at_id_idx ON portfolios (created_at, id);

CREATE TRIGGER users_timestamp BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER sessions_timestamp BEFORE UPDATE ON sessions FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER user_networks_timestamp BEFORE UPDATE ON user_networks FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER currencies_timestamp BEFORE UPDATE ON currencies FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER providers_timestamp BEFORE UPDATE ON providers FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER instruments_timestamp BEFORE UPDATE ON instruments FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER markets_timestamp BEFORE UPDATE ON markets FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER registers_timestamp BEFORE UPDATE ON registers FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER portfolios_timestamp BEFORE UPDATE ON portfolios FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER assets_timestamp BEFORE UPDATE ON assets FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);

INSERT INTO currencies (id, title, description)
VALUES ('e6dffe5f-af39-44c4-a9f2-4938cceb7f7c', 'RUB', 'Рубль'),
       ('fd39d16f-db12-4aa2-80d6-a2917dbc8715', 'USD', 'Доллар'),
       ('9c093338-0079-45af-80b7-c58c991d4535', 'EUR', 'Евро'),
       ('3f909b14-f18b-4b8b-95b3-19a2fcf1f9d7', 'CNY', 'Юань'),
       ('e088d3e2-361f-489a-8646-ec47932c7c4b', 'HKD', 'Гонконгский доллар');

INSERT INTO instruments (id, title, description)
VALUES ('2ca3707d-03b6-4f12-8f1a-6c8ec522ac95', 'STOCK', 'Акции'),
       ('2bbc7edd-8f11-4625-846a-8a98c89e0a29', 'BOND', 'Облигации'),
       ('99a91a87-24eb-4202-af0d-104309a42487', 'ETF', 'Фонды ETF'),
       ('83b6e4ef-0feb-4935-9544-a81d06506d76', 'CURRENCY', 'Валюта'),
       ('cc376387-4f0b-4688-88e6-02c3af93a646', 'CRYPTO', 'Криптовалюта');

INSERT INTO providers (id, title, description)
VALUES ('514edc8f-0921-468e-95f4-2284cba5b7bb', 'tinkoff', 'Тинькофф Инвестиции'),
       ('ba93ed83-8687-41cf-8741-edf79548e7df', 'binance', 'Криптовалютная Биржа Binance');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS sessions CASCADE;
DROP TABLE IF EXISTS user_networks CASCADE;
DROP TABLE IF EXISTS currencies CASCADE;
DROP TABLE IF EXISTS providers CASCADE;
DROP TABLE IF EXISTS instruments CASCADE;
DROP TABLE IF EXISTS markets CASCADE;
DROP TABLE IF EXISTS registers CASCADE;
DROP TABLE IF EXISTS portfolios CASCADE;
DROP TABLE IF EXISTS assets CASCADE;

DROP EXTENSION IF EXISTS "uuid-ossp";
DROP EXTENSION IF EXISTS citext;
DROP EXTENSION IF EXISTS moddatetime;
-- +goose StatementEnd
