create table users
(
    id varchar(512),
    email varchar(50),
    password varchar(256),
    username varchar(128),
    created_at date,
    updated_at date
);

create table email_template (
    id varchar(512),
    code varchar(32),
    subject varchar(512),
    body varchar(4048)
)

create table user_emails (
    id varchar(512),
    subject varchar(512),
    body varchar(4048),
    user_id varchar(512),
    date_sent date
)
