CREATE TABLE product (
    id serial NOT NULL,
    name varchar(255) NOT NULL,
    description text NOT NULL,
    price numeric(10,2) NOT NULL,
    product_id integer NOT NULL,
    Url varchar(255) NOT NULL,
    CONSTRAINT product_pkey PRIMARY KEY (id)
);