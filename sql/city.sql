CREATE SEQUENCE city_city_id_seq
	INCREMENT 1
    MINVALUE 1
    MAXVALUE 999999999999999999
    START 1;

CREATE TABLE public.city
(
    city_id integer NOT NULL DEFAULT nextval('city_city_id_seq'::regclass),
    city_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    province_id integer,
    CONSTRAINT city_pkey PRIMARY KEY (city_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.city
    OWNER to postgres;