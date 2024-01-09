CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "name" varchar NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "bio" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()) 
);

CREATE TABLE "posts" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "body" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "user_id" uuid NOT NULL
);

CREATE TABLE "comments" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "body" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "user_id" uuid NOT NULL,
  "post_id" uuid NOT NULL,
  "parent_comment_id" uuid
);

CREATE TABLE "notifications" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "body" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "user_id" uuid NOT NULL
);

CREATE TABLE "post_likes" (
  "user_id" uuid NOT NULL,
  "post_id" uuid NOT NULL,
  PRIMARY KEY ("user_id", "post_id")
);

CREATE TABLE "followers" (
  "follower_id" uuid NOT NULL,
  "followes_id" uuid NOT NULL,
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

CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at := NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_users_updated_at
BEFORE UPDATE ON "users"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trigger_update_posts_updated_at
BEFORE UPDATE ON "posts"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

CREATE TRIGGER trigger_update_comments_updated_at
BEFORE UPDATE ON "comments"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at();

