package models

func GetProduct(id uint) *Product {
	product := &Product{}

	rows, err := db.Query(`SELECT id, name FROM products WHERE id = ?`, id)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name)
		checkErr(err)
	}

	return product
}

func GetProducts() []*Product {
	products := make([]*Product, 0)

	rows, err := db.Query(`SELECT id, name FROM products`)
	checkErr(err)

	for rows.Next() {
		var temp = &Product{}
		err = rows.Scan(&temp.ID, &temp.Name)
		checkErr(err)

		products = append(products, temp)
	}

	return products
}

func GetCategoriesFor(product uint) []*Category {
	categories := make([]*Category, 0)
	rows, err := db.Query(`SELECT categories.id, categories.name FROM categories 
			LEFT JOIN categories_products ON (categories.id = categories_products.category_id)
			WHERE categories_products.category_id = ? `, product)

	checkErr(err)

	for rows.Next() {
		var temp = &Category{}
		err = rows.Scan(&temp.ID, &temp.Name)
		checkErr(err)

		categories = append(categories, temp)
	}

	return categories
}
