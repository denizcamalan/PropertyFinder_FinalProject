package models

import (
	"log"

	"github.com/denizcamalan/PF_FinalProject/database"
	"github.com/denizcamalan/PF_FinalProject/entities"
)

type CartModel struct{
}

// type Cart interface{
// 	AddItem(id,quantity int64, name,description string, price, vat float64) error
// 	ListItem() ([]entities.Item) 
// 	UpdateItem(id, quantity int64) (error)
// 	SelectItem(id int64) (entities.Item, error)
// 	DeleteItem(id int64) (error)
// }

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

func (*CartModel) ListAll() ([]entities.Item) {

	db, err := database.Get_db()
	if err != nil {
		return nil
	}else {
		values, err := db.Query("SELECT * FROM cart")
		if err != nil {
			return nil
		}else {
			defer values.Close()
			var items []entities.Item
			var item entities.Item
			for values.Next() {
				values.Scan(&item.Id, &item.Name, &item.Price, &item.VAT, &item.Description, &item.Quantity)
				items = append(items, item)
			}
			db.Close()
			return items
		}
	}

}

func (*CartModel) SelectItem(id int64) (entities.Item, error) {

	db, err := database.Get_db()
	if err != nil {
		return entities.Item{}, err
	} else {
		values, err := db.Query("SELECT * FROM cart WHERE id=?", id)
		if err != nil {
			return entities.Item{}, err
		}else {
			defer values.Close()
			var product entities.Item
			for values.Next() {
				values.Scan(&product.Id, &product.Name, &product.Price, &product.VAT,
				&product.Description, &product.Quantity,
				)
			}
			db.Close()
			return product, nil
		}
	}
}

func (*CartModel) UpdateItem(id, quantity int64) (error) {
	//quantity++
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
