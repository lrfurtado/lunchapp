--
-- PostgreSQL database dump
--

-- Dumped from database version 10.0
-- Dumped by pg_dump version 10.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: employee; Type: TABLE; Schema: public; Owner: lfurtado
--

CREATE TABLE employee (
    employee_id bigint NOT NULL,
    employee_name character(255) NOT NULL
);


ALTER TABLE employee OWNER TO lfurtado;

--
-- Name: employee_employee_id_seq; Type: SEQUENCE; Schema: public; Owner: lfurtado
--

CREATE SEQUENCE employee_employee_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE employee_employee_id_seq OWNER TO lfurtado;

--
-- Name: employee_employee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: lfurtado
--

ALTER SEQUENCE employee_employee_id_seq OWNED BY employee.employee_id;


--
-- Name: employee employee_id; Type: DEFAULT; Schema: public; Owner: lfurtado
--

ALTER TABLE ONLY employee ALTER COLUMN employee_id SET DEFAULT nextval('employee_employee_id_seq'::regclass);


--
-- Data for Name: employee; Type: TABLE DATA; Schema: public; Owner: lfurtado
--

COPY employee (employee_id, employee_name) FROM stdin;
\.


--
-- Name: employee_employee_id_seq; Type: SEQUENCE SET; Schema: public; Owner: lfurtado
--

SELECT pg_catalog.setval('employee_employee_id_seq', 94, true);


--
-- Name: employee employee_pkey; Type: CONSTRAINT; Schema: public; Owner: lfurtado
--

ALTER TABLE ONLY employee
    ADD CONSTRAINT employee_pkey PRIMARY KEY (employee_id);


--
-- Name: employee; Type: ACL; Schema: public; Owner: lfurtado
--

GRANT SELECT,INSERT,DELETE ON TABLE employee TO test;


--
-- Name: employee_employee_id_seq; Type: ACL; Schema: public; Owner: lfurtado
--

GRANT UPDATE ON SEQUENCE employee_employee_id_seq TO test;


--
-- PostgreSQL database dump complete
--

