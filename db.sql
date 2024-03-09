-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-03-06T17:41:08.733Z

CREATE TYPE "post_status" AS ENUM (
  'draft',
  'published',
  'private'
);

CREATE TABLE "users" (
  "id" integer,
  "username" varchar,
  "role" varchar,
  "created_at" timestamp
);

CREATE TABLE "posts" (
  "id" integer PRIMARY KEY,
  "title" varchar,
  "body" text,
  "user_id" integer,
  "status" post_status,
  "created_at" timestamp
);

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
