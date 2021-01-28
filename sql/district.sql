CREATE SEQUENCE district_district_id_seq
	INCREMENT 1
    MINVALUE 1
    MAXVALUE 999999999999999999
    START 1;

CREATE TABLE public.district
(
    district_id integer NOT NULL DEFAULT nextval('district_district_id_seq'::regclass),
    district_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
	city_id integer,
    CONSTRAINT district_pkey PRIMARY KEY (district_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.district
    OWNER to postgres;