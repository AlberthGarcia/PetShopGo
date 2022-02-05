package models

type Category struct {
	IdCategory   int64 `json:"idcategory"`
	NameCategory int64 `json:"namecategory"`
}

var SchemaCategory string = `
	create table categories(
		idCategory int(11) unsigned not null primary key auto_increment,
		nameCategory varchar(150) not null,
		current_data timestamp default current_timestamp);`
