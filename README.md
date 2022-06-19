# restgo

GOLANG practice!

Project created just to practice the new knowledge.

This project is just a CRUD using REST verbs PATCH, DELETE, INSERT, GET to execute the basic operation during a development of backend.
It was used here a PostgreSQL to save the data. Below is the script to create the DB.

-- TO EXECUTE THE Project

# go run .

-- ENDPOINTS AVAILABLE TO PLAY :)

 GET    /persons                  
 GET    /person/:id               
 POST   /person                   
 DELETE /person/:id               
 PATCH  /person/:id               

-- JSON TO PASS ON POST OR PATCH

{
   "id": "0",
    "name": "<NAME>",
    "cpf": "<BRASILIAN DOCUMENT>",
    "age": 18,
    "address": {
        "street": "",
        "number": 0,
        "zip_code": ""
    }
}

-- DDL to create the DATABASE

CREATE DATABASE crud
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1;

CREATE TABLE public.person
(
    id numeric NOT NULL,
    name character varying(256) NOT NULL,
    cpf character varying(16) NOT NULL,
    age numeric NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.person
    OWNER to postgres;

CREATE SEQUENCE public.personseq
    INCREMENT 1;

ALTER SEQUENCE public.personseq
    OWNER TO postgres;    

