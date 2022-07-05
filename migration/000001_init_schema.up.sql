CREATE TABLE IF NOT EXISTS "restaurant" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "cash_balance" bigint NOT NULL ,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "menu" (
    "id" SERIAL PRIMARY KEY,
    "restaurant_id" integer not null,
    "dish_name" varchar NOT NULL,
    "price" bigint NOT NULL ,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_menu_restaurant
        FOREIGN KEY(restaurant_id)
            REFERENCES restaurant(id)
);

CREATE TABLE IF NOT EXISTS "ops_hour" (
    "id" SERIAL PRIMARY KEY,
    "restaurant_id" integer not null,
    "date" varchar NOT NULL,
    "open_hour" timestamp NOT NULL ,
    "close_hour" timestamp NOT NULL ,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_menu_ops_hour
        FOREIGN KEY(restaurant_id)
            REFERENCES restaurant(id)
);

CREATE EXTENSION pg_trgm;
CREATE EXTENSION btree_gin;

CREATE INDEX idx_restaurant_name ON "restaurant" USING gin ("name");

CREATE INDEX idx_menu_name ON "menu" USING gin ("dish_name");

CREATE INDEX idx_ops_hour_date_open_hour_close_hour On "ops_hour" ("date","open_hour","close_hour");

