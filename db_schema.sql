--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

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
-- Name: article; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.article (
    aid bigint NOT NULL,
    last_floor_number integer DEFAULT 0 NOT NULL,
    is_completed boolean DEFAULT false NOT NULL,
    comments_count integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.article OWNER TO postgres;

--
-- Name: COLUMN article.aid; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.article.aid IS 'Article ID';


--
-- Name: COLUMN article.last_floor_number; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.article.last_floor_number IS 'Last Comment Floor Number';


--
-- Name: COLUMN article.is_completed; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.article.is_completed IS 'Is get all comments';


--
-- Name: COLUMN article.comments_count; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.article.comments_count IS 'Comments Count';


--
-- Name: comment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.comment (
    cid bigint NOT NULL,
    aid bigint NOT NULL,
    floor_number integer NOT NULL,
    comment text DEFAULT ''::text NOT NULL,
    is_del boolean DEFAULT false NOT NULL,
    harm_info_report_cnt bigint DEFAULT '0'::bigint NOT NULL
);


ALTER TABLE public.comment OWNER TO postgres;

--
-- Name: COLUMN comment.cid; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.comment.cid IS 'Comment ID';


--
-- Name: COLUMN comment.aid; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.comment.aid IS 'Article ID';


--
-- Name: COLUMN comment.floor_number; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.comment.floor_number IS 'Floor Number';


--
-- Name: COLUMN comment.comment; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.comment.comment IS 'Comment';


--
-- Name: COLUMN comment.is_del; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.comment.is_del IS 'Is Delete';


--
-- Name: COLUMN comment.harm_info_report_cnt; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.comment.harm_info_report_cnt IS 'The number of reports of harmful information';


--
-- Name: article article_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.article
    ADD CONSTRAINT article_pkey PRIMARY KEY (aid);


--
-- Name: comment comments_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.comment
    ADD CONSTRAINT comments_pkey PRIMARY KEY (cid);


--
-- Name: comments_aid_fn_key; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX comments_aid_fn_key ON public.comment USING btree (aid, floor_number);


--
-- Name: comments_aid_key; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX comments_aid_key ON public.comment USING btree (aid);


--
-- PostgreSQL database dump complete
--

