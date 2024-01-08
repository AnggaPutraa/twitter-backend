CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "bio" varchar,
  "hashed_refresh_token" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "body" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz,
  "user_id" bigint NOT NULL
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "body" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz,
  "user_id" bigint NOT NULL,
  "post_id" bigint NOT NULL,
  "parent_comment_id" bigint
);

CREATE TABLE "notifications" (
  "id" bigserial PRIMARY KEY,
  "body" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "user_id" bigint NOT NULL
);

CREATE TABLE "post_likes" (
  "user_id" bigint NOT NULL,
  "post_id" bigint NOT NULL,
  PRIMARY KEY ("user_id", "post_id")
);

CREATE TABLE "followers" (
  "follower_id" bigint NOT NULL,
  "followes_id" bigint NOT NULL,
  PRIMARY KEY ("follower_id", "followes_id")
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "comments" ("user_id");

CREATE INDEX ON "comments" ("post_id");

CREATE INDEX ON "comments" ("user_id", "post_id");

CREATE INDEX ON "post_likes" ("user_id");

CREATE INDEX ON "post_likes" ("post_id");

CREATE INDEX ON "followers" ("follower_id");

CREATE INDEX ON "followers" ("followes_id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("parent_comment_id") REFERENCES "comments" ("id");

ALTER TABLE "notifications" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "post_likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "post_likes" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "followers" ADD FOREIGN KEY ("follower_id") REFERENCES "users" ("id");

ALTER TABLE "followers" ADD FOREIGN KEY ("followes_id") REFERENCES "users" ("id");
