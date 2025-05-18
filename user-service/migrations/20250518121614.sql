-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "name" text NULL,
  "email" text NULL,
  "address" text NULL,
  "birth" text NULL,
  "gender" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_email" UNIQUE ("email")
);
