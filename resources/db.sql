CREATE TABLE db_golang_yasir.t_user (
    id INT auto_increment primary key,
    username VARCHAR(255) not null UNIQUE,
    `password` VARCHAR(255) not null,
    full_name VARCHAR(255) not null,
    birth_date VARCHAR(255) not null,
    image_url VARCHAR(255) not null,
    phone VARCHAR(255) not null,
    email VARCHAR(255) not null,
    `address` VARCHAR(255) not null,
    role_id INT not null,
    is_active INT not null,
    created_at DATETIME not null,
    updated_at DATETIME not null
);


CREATE TABLE db_golang_yasir.t_role (
    id INT auto_increment primary key,
    title VARCHAR(255) not null,
    detail VARCHAR(255) not null,
    created_at DATETIME not null,
    updated_at DATETIME not null
);

CREATE TABLE db_golang_yasir.t_shop (
    id INT auto_increment primary key,
    tag VARCHAR(255) not null UNIQUE,
    `name` VARCHAR(255) not null,
    detail VARCHAR(255) not null,
    image_url VARCHAR(255) not null,
    phone VARCHAR(255) not null,
    email VARCHAR(255) not null,
    `address` VARCHAR(255) not null,
    is_active INT not null,
    `user_id` INT not null,
    created_at DATETIME not null,
    updated_at DATETIME not null
);

CREATE TABLE db_golang_yasir.t_product (
    id INT auto_increment primary key,
    `name` VARCHAR(255) not null,
    detail VARCHAR(255) not null,
    category VARCHAR(255) not null,
    image_url VARCHAR(255) not null,
    price DECIMAL not null,
    stock INT not null,
    shop_id INT not null,
    created_at DATETIME not null,
    updated_at DATETIME not null
);

CREATE TABLE db_golang_yasir.t_review (
    id INT auto_increment primary key,
    note VARCHAR(255) not null,
    response VARCHAR(255) not null,
    media_url VARCHAR(255) not null,
    rating INT not null,
    `user_id` INT not null,
    product_id INT not null,
    created_at DATETIME not null,
    updated_at DATETIME not null
);

CREATE TABLE db_golang_yasir.t_session (
    id INT auto_increment primary key,
    courier VARCHAR(255) not null,
    note VARCHAR(255) not null,
    delivery_cost DECIMAL not null,
    `user_id` INT not null,
    status_id INT not null,
    created_at DATETIME not null,
    updated_at DATETIME not null
);

CREATE TABLE db_golang_yasir.t_status (
    id INT auto_increment primary key,
    title VARCHAR(255) not null,
    detail VARCHAR(255) not null,
    created_at DATETIME not null,
    updated_at DATETIME not null
);

CREATE TABLE db_golang_yasir.t_cart (
    id INT auto_increment primary key,
    product_name VARCHAR(255) not null,
    note VARCHAR(255) not null,
    product_price DECIMAL not null,
    product_price_mod DECIMAL not null,
    amount INT not null,
    `session_id` INT not null,
    created_at DATETIME not null,
    updated_at DATETIME not null
);


INSERT INTO db_golang_yasir.t_user (username, `password`, full_name, birth_date, image_url, phone, email, `address`, `role_id`, is_active, created_at, updated_at) values('master', '$2a$10$brdQAGKIU1wq5Wq2kUzZ..XWIYZRWyb9GladTwHIIqDPhMOm/WWpS', 'Irfan Yas', '1995-08-02', 'https://ui-avatars.com/api/?name=Irfan%20Yas&background=random', '085155112341', 'irfanyasiras@gmail.com', 'Kota Bandung', 1, 1, NOW(), NOW());
INSERT INTO db_golang_yasir.t_role (title, detail, created_at, updated_at) values('ADMIN', 'Role for developer', NOW(), NOW());
INSERT INTO db_golang_yasir.t_role (title, detail, created_at, updated_at) values('USER', 'Role for customers and sellers', NOW(), NOW());
INSERT INTO db_golang_yasir.t_status (title, detail, created_at, updated_at) values('CART', 'Masih dalam keranjang pembeli', NOW(), NOW());
INSERT INTO db_golang_yasir.t_status (title, detail, created_at, updated_at) values('PREPARING', 'Dalam persiapan untuk dikirim oleh penjual', NOW(), NOW());
INSERT INTO db_golang_yasir.t_status (title, detail, created_at, updated_at) values('ON DELIVERY', 'Dalam pengiriman', NOW(), NOW());
INSERT INTO db_golang_yasir.t_status (title, detail, created_at, updated_at) values('FINISHED', 'Berhasil dikirim', NOW(), NOW());
INSERT INTO db_golang_yasir.t_status (title, detail, created_at, updated_at) values('CANCELED', 'Dibatalkan', NOW(), NOW());
