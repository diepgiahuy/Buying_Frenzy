
CREATE TABLE IF NOT EXISTS "user" (
      "id" SERIAL PRIMARY KEY,
      "name" varchar NOT NULL,
      "cash_balance" float NOT NULL ,
      "created_at" timestamptz NOT NULL DEFAULT (now()),
      "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "purchase_history" (
    "id" SERIAL PRIMARY KEY,
    "user_id" integer NOT NULL ,
    "restaurant_name" varchar NOT NULL,
    "dish_name" varchar NOT NULL,
    "transaction_amount" float NOT NULL ,
    "transaction_date" timestamptz NOT NULL DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    CONSTRAINT fk_purchase_history_user
        FOREIGN KEY(user_id)
            REFERENCES "user" (id)
);
