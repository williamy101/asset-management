create database asset_management

use asset_management

create table asset_categories (
    category_id int auto_increment primary key,
    category_name varchar(100) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);

create table statuses (
	status_id int auto_increment primary key,
	status_name varchar(100) not null,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp on update current_timestamp
);

create table roles (
    role_id int auto_increment primary key,
    role_name varchar(100) not null,
    created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp on update current_timestamp
);

create table assets (
	asset_id int auto_increment primary key,
	asset_name varchar(100) not null,
	category_id int,
    status_id int not null,
    last_maintenance date,
    next_maintenance date,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    foreign key (category_id) references asset_categories(category_id),
   	foreign key (status_id) references statuses(status_id)
);

create table users (
    user_id int auto_increment primary key,
    name varchar(100) not null,
    email varchar(100) unique not null,
    password varchar(255) not null, -- password nantinya dalam hash
    role_id int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    foreign key (role_id) references roles(role_id)
);

create table maintenances (
    maintenance_id int auto_increment primary key,
    asset_id int not null,
    user_id int not null,
    description text,
    cost decimal(10, 2),
    status_id int not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
    FOREIGN KEY (asset_id) REFERENCES assets(asset_id) on delete cascade,
    FOREIGN KEY (user_id) REFERENCES users(user_id) on delete cascade,
    FOREIGN KEY (status_id) REFERENCES statuses(status_id)
);

insert into roles (role_name) values 
('Admin'), ('User');

insert into asset_categories (category_name) values 
('Perlengkapan IT'), ('Furnitur'), ('Kendaraan');

insert into statuses (status_name) values 
('available'), ('in use'), ('in maintenance'), ('scheduled'), ('completed');

insert into users (name, email, password, role_id) values
('Admin User', 'admin@email.com', 'Adm5321', 1) -- dummy data untuk admin pertama