CREATE TABLE db_jcc.t_book (
    id INT auto_increment primary key,
    title VARCHAR(255) not null,
    `description` VARCHAR(255) not null,
    image_url VARCHAR(255) not null,
    `release_year` INT not null,
    price VARCHAR(255) not null,
    total_page VARCHAR(255) not null,
    kategori_ketebalan VARCHAR(255) not null,
    created_at DATETIME not null,
    updated_at DATETIME not null
)