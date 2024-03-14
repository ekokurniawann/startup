--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Ubuntu 16.2-1.pgdg22.04+1)
-- Dumped by pg_dump version 16.2 (Ubuntu 16.2-1.pgdg22.04+1)

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
-- Name: campaign_images; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.campaign_images (
    id integer NOT NULL,
    campaign_id integer,
    file_name character varying(255),
    is_primary smallint,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.campaign_images OWNER TO ekokurniawan;

--
-- Name: campaign_images_id_seq; Type: SEQUENCE; Schema: public; Owner: ekokurniawan
--

CREATE SEQUENCE public.campaign_images_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.campaign_images_id_seq OWNER TO ekokurniawan;

--
-- Name: campaign_images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ekokurniawan
--

ALTER SEQUENCE public.campaign_images_id_seq OWNED BY public.campaign_images.id;


--
-- Name: campaigns; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.campaigns (
    id integer NOT NULL,
    user_id integer,
    name character varying(255),
    short_description character varying(255),
    description text,
    perks text,
    backer_count integer,
    goal_amount integer,
    current_amount integer,
    slug character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


ALTER TABLE public.campaigns OWNER TO ekokurniawan;

--
-- Name: campaigns_id_seq; Type: SEQUENCE; Schema: public; Owner: ekokurniawan
--

CREATE SEQUENCE public.campaigns_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.campaigns_id_seq OWNER TO ekokurniawan;

--
-- Name: campaigns_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ekokurniawan
--

ALTER SEQUENCE public.campaigns_id_seq OWNED BY public.campaigns.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO ekokurniawan;

--
-- Name: users; Type: TABLE; Schema: public; Owner: ekokurniawan
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(255),
    occupation character varying(255),
    email character varying(255),
    password_hash character varying(255),
    avatar_file_name character varying(255),
    role character varying(255),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.users OWNER TO ekokurniawan;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: ekokurniawan
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_id_seq OWNER TO ekokurniawan;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ekokurniawan
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: campaign_images id; Type: DEFAULT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.campaign_images ALTER COLUMN id SET DEFAULT nextval('public.campaign_images_id_seq'::regclass);


--
-- Name: campaigns id; Type: DEFAULT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.campaigns ALTER COLUMN id SET DEFAULT nextval('public.campaigns_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: campaign_images campaign_images_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.campaign_images
    ADD CONSTRAINT campaign_images_pkey PRIMARY KEY (id);


--
-- Name: campaigns campaigns_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.campaigns
    ADD CONSTRAINT campaigns_pkey PRIMARY KEY (id);


--
-- Name: schema_migration schema_migration_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.schema_migration
    ADD CONSTRAINT schema_migration_pkey PRIMARY KEY (version);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: ekokurniawan
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: campaign_images campaign_images_campaign_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.campaign_images
    ADD CONSTRAINT campaign_images_campaign_id_fkey FOREIGN KEY (campaign_id) REFERENCES public.campaigns(id);


--
-- Name: campaigns campaigns_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ekokurniawan
--

ALTER TABLE ONLY public.campaigns
    ADD CONSTRAINT campaigns_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

