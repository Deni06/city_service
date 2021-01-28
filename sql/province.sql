CREATE SEQUENCE public.province_province_id_seq;

ALTER SEQUENCE public.province_province_id_seq
    OWNER TO postgres;

CREATE TABLE public.province
(
    province_id integer NOT NULL DEFAULT nextval('province_province_id_seq'::regclass),
    province_name character varying(50) COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    CONSTRAINT province_pkey PRIMARY KEY (province_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.province
    OWNER to postgres;