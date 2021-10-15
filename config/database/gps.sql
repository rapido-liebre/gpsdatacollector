--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.5
-- Dumped by pg_dump version 13.4

-- Started on 2021-10-12 23:29:04 CEST

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2127 (class 1262 OID 16384)
-- Name: gps; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE gps WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE gps OWNER TO postgres;

\connect gps

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

--
-- TOC entry 186 (class 1259 OID 24578)
-- Name: coordinates; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.coordinates (
    id integer NOT NULL,
    device_id character(50) NOT NULL,
    point point NOT NULL
);


ALTER TABLE public.coordinates OWNER TO admin;

--
-- TOC entry 185 (class 1259 OID 24576)
-- Name: coordinates_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.coordinates_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.coordinates_id_seq OWNER TO admin;

--
-- TOC entry 2128 (class 0 OID 0)
-- Dependencies: 185
-- Name: coordinates_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.coordinates_id_seq OWNED BY public.coordinates.id;


--
-- TOC entry 2002 (class 2604 OID 24581)
-- Name: coordinates id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.coordinates ALTER COLUMN id SET DEFAULT nextval('public.coordinates_id_seq'::regclass);


--
-- TOC entry 2004 (class 2606 OID 24583)
-- Name: coordinates coordinates_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.coordinates
    ADD CONSTRAINT coordinates_pkey PRIMARY KEY (id);


-- Completed on 2021-10-12 23:29:10 CEST

--
-- PostgreSQL database dump complete
--

