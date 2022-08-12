# PropertyFinder_FinalProject

This project is a basket servis projcet. Used REST API developed in Go.

Customers will be able to purchase existing products. Use cases such as creating, updating, deleting products or creating users etc.

## The functions of this service are as follows;
1. List Products
- Users should be able to list all products.
2. Add To Cart
- Users can add their products to the basket and the total of the basket changes accordingly.
3. Show Cart
- Users can list the products they have added to their cart and total price and VAT of the cart.
4. Delete Cart Item
- Users can remove the products added from the cart.
5. Complete Order
- Users can create an order with the products they add to their cart.

## Some business rules
1. Products always have price and VAT (Value Added Tax, or KDV). VAT might be different for different products. Typical VAT percentage is %1, %8 and %18.
2. There might be discount in following situations:
a. Every fourth order whose total is more than given amount may have discount depending on products. Products whose VAT is %1 don't have any discount but products whose VAT is %8 and %18 have discount of %10 and %15 respectively.
b. If there are more than 3 items of the same product, then fourth and subsequent ones would have %8 off.
c. If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off.
d. Only one discount can be applied at a time. Only the highest discount should be applied.


## Usage 

    git clone https://github.com/denizcamalan/PropertyFinder_FinalProject.git

open project and run:

     docker-compose up -d

     $GOBIN/PF_FinalProject

and database is uploaded.

for ading products to cart:

    POST http://localhost:8088/users/cart/add?id=2

 then list cart:

    GET http://localhost:8088/users/cart

delete item:

    POST http://localhost:8088/users/cart/remove?id=1

buy item into card :

    POST http://localhost:8088/users/cart/buy

list orders :

    GET http://localhost:8088/users/orders

get products :

    GET http://localhost:8088/users/productlist

add product :

    POST http://localhost:8088/users/product/add?id=9&quantity=60&name=Harman Kardon Go Plus&description=speaker&price=4550.6&vat=0.18

delete product :

    POST http://localhost:8088/users/product/delete?id=9
