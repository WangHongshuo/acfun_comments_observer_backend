--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1 (Debian 16.1-1.pgdg120+1)
-- Dumped by pg_dump version 16.2

-- Started on 2024-04-09 13:38:44

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

DROP DATABASE acfun_comm;
--
-- TOC entry 3366 (class 1262 OID 16384)
-- Name: acfun_comm; Type: DATABASE; Schema: -; Owner: -
--

CREATE DATABASE acfun_comm WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


\connect acfun_comm

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
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA public;


--
-- TOC entry 3367 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 215 (class 1259 OID 16398)
-- Name: article; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.article (
    aid bigint NOT NULL,
    last_floor_number integer DEFAULT 0 NOT NULL,
    is_completed boolean DEFAULT false NOT NULL,
    comments_count integer DEFAULT 0 NOT NULL
);


--
-- TOC entry 3368 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN article.aid; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.article.aid IS 'Article ID';


--
-- TOC entry 3369 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN article.last_floor_number; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.article.last_floor_number IS 'Last Comment Floor Number';


--
-- TOC entry 3370 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN article.is_completed; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.article.is_completed IS 'Is get all comments';


--
-- TOC entry 3371 (class 0 OID 0)
-- Dependencies: 215
-- Name: COLUMN article.comments_count; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.article.comments_count IS 'Comments Count';


--
-- TOC entry 216 (class 1259 OID 16404)
-- Name: comment; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.comment (
    cid bigint NOT NULL,
    aid bigint NOT NULL,
    floor_number integer NOT NULL,
    comment text DEFAULT ''::text NOT NULL,
    is_del boolean DEFAULT false NOT NULL,
    harm_info_report_cnt bigint DEFAULT '0'::bigint NOT NULL
);


--
-- TOC entry 3372 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN comment.cid; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.comment.cid IS 'Comment ID';


--
-- TOC entry 3373 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN comment.aid; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.comment.aid IS 'Article ID';


--
-- TOC entry 3374 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN comment.floor_number; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.comment.floor_number IS 'Floor Number';


--
-- TOC entry 3375 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN comment.comment; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.comment.comment IS 'Comment';


--
-- TOC entry 3376 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN comment.is_del; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.comment.is_del IS 'Is Delete';


--
-- TOC entry 3377 (class 0 OID 0)
-- Dependencies: 216
-- Name: COLUMN comment.harm_info_report_cnt; Type: COMMENT; Schema: public; Owner: -
--

COMMENT ON COLUMN public.comment.harm_info_report_cnt IS 'The number of reports of harmful information';


--
-- TOC entry 3213 (class 2606 OID 16413)
-- Name: article article_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.article
    ADD CONSTRAINT article_pkey PRIMARY KEY (aid);


--
-- TOC entry 3217 (class 2606 OID 16415)
-- Name: comment comments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comments_pkey PRIMARY KEY (cid);


--
-- TOC entry 3214 (class 1259 OID 16416)
-- Name: comments_aid_fn_key; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX comments_aid_fn_key ON public.comment USING btree (aid, floor_number);


--
-- TOC entry 3215 (class 1259 OID 16417)
-- Name: comments_aid_key; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX comments_aid_key ON public.comment USING btree (aid);


-- Completed on 2024-04-09 13:38:45

--
-- PostgreSQL database dump complete
--

