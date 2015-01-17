package reports

const schema string = `

DROP SCHEMA IF EXISTS inbox CASCADE;
CREATE SCHEMA inbox;


CREATE TABLE inbox.items (
       id    uuid        PRIMARY KEY,
       title text
);



`
