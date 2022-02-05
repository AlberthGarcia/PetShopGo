package models

type Products struct {
	IdProduct   int64  `json:"idProduct"`
	NameProduct string `json:"nameProduct"`
	Category    int64  `json:"category"`
	Existence   int64  `json:"existence"`
}

var SchemaProducts string = `
	create table products(
		idProduct int(11) unsigned not null primary key auto_increment,
		nameProduct varchar(150) not null,
		category int(11) not null,
		existence int(11) not null,
		current_data timestamp default current_timestamp);`
