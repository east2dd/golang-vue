package models

func GetCategory(id uint) *Category {
	category := &Category{}

	rows, err := db.Query(`SELECT id, name FROM categories WHERE id = ?`, id)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&category.ID, &category.Name)
		checkErr(err)
	}

	return category
}

func GetCategories() []*Category {
	categories := make([]*Category, 0)

	rows, err := db.Query(`SELECT id, name FROM categories`)
	checkErr(err)

	for rows.Next() {
		var temp = &Category{}
		err = rows.Scan(&temp.ID, &temp.Name)
		checkErr(err)

		categories = append(categories, temp)
	}

	return categories
}

func GetProductsFor(category uint) []*Product {

	products := make([]*Product, 0)

	rows, err := db.Query(`SELECT products.id, products.name FROM products 
													LEFT JOIN categories_products ON (products.id = categories_products.product_id)
													WHERE categories_products.category_id = ? `, category)
	checkErr(err)

	for rows.Next() {
		var temp = &Product{}
		err = rows.Scan(&temp.ID, &temp.Name)
		checkErr(err)

		products = append(products, temp)
	}

	return products
}
