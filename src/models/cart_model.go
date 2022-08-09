package models

import (
	"log"

	"github.com/denizcamalan/PF_FinalProject/database"
	"github.com/denizcamalan/PF_FinalProject/entities"
)

type CartModel struct{
}

// Add items to cart from products database
func (*CartModel) AddItem(id,quantity int64, name,description string, price, vat float64) error{
	db, err1 := database.Get_db()
	if err1 != nil {
		return err1
	} else {
		_, err := db.Exec("INSERT INTO cart (id,name,price,vat,description,quantity) VALUES (?,?,?,?,?,?)",id ,name ,vat ,price,description,quantity) 
		if err != nil {
			log.Println(err, "PRIMARY KEY")
			return err
		}
		return nil
	}
}

// List all of the selected items into the cart
func (*CartModel) ListAll() []interface{} {

	db, err := database.Get_db()
	if err != nil {
		return nil
	}else {
		values, err := db.Query("SELECT * FROM cart")
		if err != nil {
			return nil
		}else {
			defer values.Close()
			// cart array
			var items_new []entities.Item
			// item struct
			var item_new entities.Item
			// add every item to struct to make list
			for values.Next() {
				values.Scan(&item_new.Id, &item_new.Name, &item_new.Price, &item_new.VAT, &item_new.Description, &item_new.Quantity)
				items_new = append(items_new, item_new)
			}
			db.Close()
			// convert cart array to interface 
			interface_items := make([]interface{}, len(items_new))
			for index, values := range items_new {
				interface_items[index] = values
			}
			return interface_items
		}
	}

}

// select item from database by item's id
func (*CartModel) SelectItem(id int64) interface{} {

	db, err := database.Get_db()
	if err != nil {
		return entities.Item{}
	} else {
		values, err := db.Query("SELECT * FROM cart WHERE id=?", id)
		if err != nil {
			return entities.Item{}
		}else {
			defer values.Close()
			var product entities.Item
			for values.Next() {
				values.Scan(&product.Id, &product.Name, &product.Price, &product.VAT,
				&product.Description, &product.Quantity,
				)
			}
			db.Close()
			return product
		}
	}
}

// update cart table 
func (*CartModel) UpdateItem(id, quantity int64) (error) {

	db, err := database.Get_db()
	if err != nil {
		return err
	} else {
		values, err := db.Prepare("UPDATE cart SET quantity=? WHERE id=?")
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

// delete items from cart table
func (*CartModel) DeleteItem(id int64) error{

	db, err := database.Get_db()
	if err != nil {
		return err
	} else {
		stmt, err := db.Prepare("DELETE FROM cart where id=?") 
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
