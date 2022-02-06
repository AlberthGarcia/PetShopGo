package models

import (
	"fmt"
	"package/db"
)

//Struct for a Category
type Category struct {
	IdCategory   int64  `json:"idcategory"`
	NameCategory string `json:"namecategory"`
}

//Type to list all categories
type listCategories []Category

//Schema to create a table for categories
var SchemaCategory string = `
	create table categories(
		idCategory int(11) unsigned not null primary key auto_increment,
		nameCategory varchar(150) not null,
		current_data timestamp default current_timestamp
);`

//func to create our object of category
func BuildCategory(nameCategory string) Category {
	return Category{
		NameCategory: nameCategory,
	}
}

//method to insert into table category
func (cat *Category) InsertCategory() {
	sql := "INSERT categories SET nameCategory=?"
	if result, err := db.Exec(sql, cat.NameCategory); err != nil {
		panic(err)
	} else {
		cat.IdCategory, _ = result.LastInsertId()
	}
}

//Func to save the category
func SaveCategory(nameCategory string) Category {
	category := BuildCategory(nameCategory)
	category.InsertCategory()
	return category
}

//func to get a category
func GetCategoryById(idCategory int64) (Category, error) {
	sql := "select idCategory, nameCategory from categories where idCategory=?"
	category := Category{}
	if rows, err := db.Query(sql, idCategory); err != nil {
		return category, err
	} else {
		for rows.Next() {
			rows.Scan(&category.IdCategory, &category.NameCategory)
		}
		return category, nil
	}
}

//Func to list all categories
func GetCategories() (listCategories, error) {
	sql := "select idCategory, nameCategory from categories"
	categories := listCategories{}
	if rows, err := db.Query(sql); err != nil {
		return categories, err
	} else {
		for rows.Next() {
			category := Category{}
			rows.Scan(&category.IdCategory, &category.NameCategory)
			categories = append(categories, category)
		}
		return categories, nil
	}

}

//func to delete a category
func (cat *Category) DeleteCategoryById() {
	sql := "delete from categories where idCategory=?"

	if _, err := db.Exec(sql, cat.IdCategory); err != nil {
		panic(err)
	} else {
		fmt.Println("Categoria eliminada")
	}
}

//Method to update a category
func (cat *Category) UpdateCategory() {
	sql := "UPDATE categories SET nameCategory=? where idCategory=?"
	if _, err := db.Exec(sql, cat.NameCategory, cat.IdCategory); err != nil {
		panic(err)
	} else {
		fmt.Println(cat)
	}
}

//method to save or update a category, depence the ID
func (cat *Category) Save() {
	if cat.IdCategory == 0 {
		cat.InsertCategory()
	} else {
		cat.UpdateCategory()
	}
}
