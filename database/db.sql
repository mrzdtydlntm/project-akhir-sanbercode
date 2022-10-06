--
-- PostgreSQL database dump
--

-- Dumped from database version 12.12 (Ubuntu 12.12-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 12.12 (Ubuntu 12.12-0ubuntu0.20.04.1)

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

SET default_table_access_method = heap;

--
-- Name: user_table; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_table (
    guid_user character varying(37) NOT NULL,
    email character varying(255) NOT NULL,
    firstname character varying(255) NOT NULL,
    lastname character varying(255),
    password character varying(100) NOT NULL
);


ALTER TABLE public.user_table OWNER TO postgres;

--
-- Data for Name: user_table; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_table (guid_user, email, firstname, lastname, password) FROM stdin;
\.


--
-- Name: user_table user_table_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_table
    ADD CONSTRAINT user_table_email_key UNIQUE (email);


--
-- Name: user_table user_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_table
    ADD CONSTRAINT user_table_pkey PRIMARY KEY (guid_user);


--
-- PostgreSQL database dump complete
--

