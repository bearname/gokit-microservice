create table products
(
    id              uuid                                       not null
        constraint products_pkey
            primary key,
    title           varchar(256)                               not null,
    sku             varchar(256)                               not null
        constraint products_sku_key
            unique,
    price           numeric                                    not null,
    available_count integer      default 0                     not null,
    image_url       varchar(1024)                              not null,
    image_width     integer                                    not null,
    image_height    integer                                    not null,
    color           varchar(100) default ''::character varying not null,
    material        varchar(100) default ''::character varying not null
);

alter table products
    owner to postgres;

INSERT INTO public.products (id, title, sku, price, available_count, image_url, image_width, image_height, color, material) VALUES ('c51747a8-d264-11eb-b8bc-0242ac130003', 't-shirt', '1234', 1200, 200, 'https://www.pexels.com/photo/graceful-young-woman-swimming-in-rippling-sea-at-sundown-6770410/', 1920, 1080, 'white', 'lion');