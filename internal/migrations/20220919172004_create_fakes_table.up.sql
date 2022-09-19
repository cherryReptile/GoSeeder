create table fakes
(
    id         serial primary key,
    string     text,
    created_at timestamp without time zone not null
);