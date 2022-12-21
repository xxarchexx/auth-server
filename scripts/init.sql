-- creating users table
\c auth;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
CREATE TABLE users(
    --id serial primary key,
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    is_active boolean,  
    username varchar(20),
    first_name varchar(50),
    last_name varchar(50),
    born_at date,
    admin boolean,
    email varchar(50) not null,
    created_at date,
    updated_at date ,
    password varchar(150)
);

-- insert into users (email, username, password) values ('sansolovyov@mail.ru','xxarchexx', '$2a$10$vWnIx47/T9O8sT3/pKdL8eBOnTxOQZt1oXq4FH4T3GZltME9e8kSG');