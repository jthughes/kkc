-- +goose up
-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2026-02-19T11:19:10.176Z

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "username" varchar UNIQUE NOT NULL
);

CREATE TABLE "players" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "user_id" integer NOT NULL,
  "game_id" integer NOT NULL,
  "name" varchar,
  "skindancer" bool DEFAULT false,
  "class" varchar
);

CREATE TABLE "turns" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "game_id" integer,
  "name" varchar,
  "writeup" varchar,
  "term" integer,
  "month" integer
);

CREATE TABLE "player_status" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "turn_id" integer,
  "alive" bool DEFAULT true,
  "sane" bool DEFAULT true,
  "crockery" bool DEFAULT false,
  "lodging" bool DEFAULT true,
  "imre" bool DEFAULT false,
  "university" bool DEFAULT true,
  "medica" bool DEFAULT false,
  "coin" float DEFAULT 0,
  "ep_linguistics" integer DEFAULT 0,
  "ep_arithmetics" integer DEFAULT 0,
  "ep_rhetoric_and_logic" integer DEFAULT 0,
  "ep_archives" integer DEFAULT 0,
  "ep_sympathy" integer DEFAULT 0,
  "ep_physicking" integer DEFAULT 0,
  "ep_alchemy" integer DEFAULT 0,
  "ep_artificery" integer DEFAULT 0,
  "ep_naming" integer DEFAULT 0
);

CREATE TABLE "games" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "game_master" integer NOT NULL,
  "name" varchar,
  "type" varchar,
  "type_number" varchar
);

CREATE TABLE "actions" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "player_turn_id" integer,
  "lodging" varchar,
  "visit_imre" bool DEFAULT false,
  "attend_university" bool DEFAULT false
);

CREATE TABLE "player_turn" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "player_id" integer,
  "turn_id" integer,
  "posts" integer DEFAULT 0,
  "private_messages" integer DEFAULT 0,
  "quality_posts" integer DEFAULT 0,
  "quality_rp" integer DEFAULT 0
);

CREATE TABLE "complaints" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "action_id" integer,
  "target_id" integer
);

CREATE TABLE "imre_actions" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "action_id" integer,
  "eolian_pipes" bool,
  "eolian_practice" bool,
  "giles_loan" float,
  "devi_loan" float,
  "dice_bet" float,
  "nox_nahlrout" integer,
  "nox_couriers" integer,
  "nox_bloodless" integer,
  "nox_gram" integer,
  "black_market_mommet" integer,
  "black_market_bodyguard" bool,
  "black_market_bodygaurd_target" integer,
  "black_market_assassin" bool,
  "black_market_assassin_target" integer,
  "black_market_contracts" bool
);

CREATE TABLE "contracts" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "game_id" integer,
  "name" varchar,
  "description" varchar,
  "requesting_player_id" integer,
  "accepted_player_id" integer,
  "bounty_coin" float,
  "bounty_item" varchar
);

CREATE TABLE "imre_contract_action" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "imre_action_id" integer,
  "contract_id" integer
);

CREATE TABLE "elevation_points" (
  "id" integer PRIMARY KEY,
  "created_at" timestamp NOT NULL,
  "action_id" integer,
  "ep_linguistics" integer,
  "ep_arithmetics" integer,
  "ep_rhetoric_and_logic" integer,
  "ep_archives" integer,
  "ep_sympathy" integer,
  "ep_physicking" integer,
  "ep_alchemy" integer,
  "ep_artificery" integer,
  "ep_naming" integer
);

ALTER TABLE "turns" ADD FOREIGN KEY ("game_id") REFERENCES "games" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "player_status" ADD FOREIGN KEY ("turn_id") REFERENCES "turns" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "actions" ADD FOREIGN KEY ("player_turn_id") REFERENCES "player_turn" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "player_turn" ADD FOREIGN KEY ("player_id") REFERENCES "players" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "player_turn" ADD FOREIGN KEY ("turn_id") REFERENCES "turns" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "complaints" ADD FOREIGN KEY ("action_id") REFERENCES "actions" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "complaints" ADD FOREIGN KEY ("target_id") REFERENCES "players" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "imre_actions" ADD FOREIGN KEY ("action_id") REFERENCES "actions" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "imre_actions" ADD FOREIGN KEY ("black_market_mommet") REFERENCES "players" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "imre_actions" ADD FOREIGN KEY ("black_market_bodygaurd_target") REFERENCES "players" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "imre_actions" ADD FOREIGN KEY ("black_market_assassin_target") REFERENCES "players" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "contracts" ADD FOREIGN KEY ("game_id") REFERENCES "games" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "contracts" ADD FOREIGN KEY ("requesting_player_id") REFERENCES "players" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "contracts" ADD FOREIGN KEY ("accepted_player_id") REFERENCES "players" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "imre_contract_action" ADD FOREIGN KEY ("imre_action_id") REFERENCES "imre_actions" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "imre_contract_action" ADD FOREIGN KEY ("contract_id") REFERENCES "contracts" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "elevation_points" ADD FOREIGN KEY ("action_id") REFERENCES "actions" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "players" ADD FOREIGN KEY ("game_id") REFERENCES "games" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "players" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "games" ADD FOREIGN KEY ("game_master") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;
