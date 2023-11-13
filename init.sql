-- Adminer 4.8.1 PostgreSQL 16.0 dump

DROP SEQUENCE IF EXISTS todo_id_seq;
CREATE SEQUENCE todo_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."todo" (
    "id" integer DEFAULT nextval('todo_id_seq') NOT NULL,
    "title" character varying(50) NOT NULL,
    "description" text NOT NULL,
    "is_done" boolean NOT NULL,
    CONSTRAINT "todo_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


-- 2023-11-12 21:47:40.84607+00
