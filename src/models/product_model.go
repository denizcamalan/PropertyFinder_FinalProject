package models

import (
	"log"

	"github.com/denizcamalan/PF_FinalProject/database"
	"github.com/denizcamalan/PF_FinalProject/entities"
)


type ProductModel struct {
}

// Add product into database
func (*ProductModel) AddItem(id,quantity int64, name,description string, price, vat float64) error{
	db, err1 := database.Get_db()
	if err1 != nil {
		return err1
	} else {
		_, err := db.Exec("INSERT INTO products (id,name,price,vat,description,quantity) VALUES (?,?,?,?,?,?)",id ,name ,vat ,price,description,quantity) 
		if err != nil {
			log.Println(err, "PRIMARY KEY")
			return err
		}
		db.Close()
		return nil
	}
}

// List all products
func (*ProductModel) ListAll() []interface{}{

	db, err := database.Get_db()
	if err != nil {
		return nil
	} else {
		values, err := db.Query("SELECT * FROM products")
		if err != nil {
			return nil
		}else {
			defer values.Close()
			// product array
			var products_new []entities.Product
			// product struct
			var product_new entities.Product
			// add every product to array for making list
			for values.Next() {
				values.Scan(&product_new.Id, &product_new.Name, &product_new.Price, &product_new.VAT,
				&product_new.Quantity, &product_new.Description,
				)
				products_new = append(products_new, product_new)
			}
			db.Close()
			// convert product_new array to interface
			interface_products := make([]interface{}, len(products_new))
			for index, values := range products_new {
				interface_products[index] = values
			}
			return interface_products
		}
	}

}

// Select items from products database
func (*ProductModel) SelectItem(id int64) interface{} {

	db, err := database.Get_db()
	if err != nil {
		return entities.Product{}
	} else {
		values, err := db.Query("SELECT * FROM products WHERE id=?", id)
		if err != nil {
			return entities.Product{}
		}else {
			defer values.Close()
			var product entities.Product
			for values.Next() {
				values.Scan(&product.Id, &product.Name, &product.Price, &product.VAT,
				&product.Quantity, &product.Description,
				)
			}
			db.Close()
			return product
		}
	}
}

// update products item into the database
func (*ProductModel) UpdateItem(id int64, quantity int64) (error) {
	db, err := database.Get_db()
	if err != nil {
		return err
	} else {
		values, err := db.Prepare("UPDATE products SET quantity=? WHERE id=?")
		if err != nil {
			return err
		}else {
			defer values.Close()
			res, err	:= values.Exec(quantity,id)
			if err !=nil{
				log.Println(err)
			}
			_, err2 := res.RowsAffected()
			if err2 !=nil{
				log.Println(err)
			}
			db.Close()
			return nil
		}
	}
}

// delete product items from database
func (*ProductModel) DeleteItem(id int64) error{

	db, err := database.Get_db()
	if err != nil {
		return err
	} else {
		stmt, err := db.Prepare("DELETE FROM products where id=?") 
		if err != nil{
			log.Println(err,"Delete db")
			return err
		}
		defer stmt.Close()
		_,err2 := stmt.Exec(id)
		if err2 != nil{
			log.Println(err2,"Exec Delete db")
			return err
		}
		db.Close()
		return nil
	}
}