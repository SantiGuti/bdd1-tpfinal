# Base de Datos 1 - TP Final

This project consists of a credit card database. Through a CLI application developed in Go the program is managed and through stored procedures and triggers developed in PostgreSQL the logic is generated.

One of the functions of this program is to be able to detect fraud during the verification of the consumption of a credit card. In which, an alert is stored depending on the following situations: - If the card exceeds its purchase limit twice in the same day, which will cause the card status to be changed to suspended. - If two purchases are made in less than one minute at two different merchants with the same zip code. - If two purchases are made in less than five minutes at two different merchants with different zip codes.

In addition, it has the option of generating invoices based on the period, by customer and by each card owned. It has a header and details with the data corresponding to the card, customer and amount to be paid.

Some of the main functions of the program are:

* createDatabase(): connects to postgres, eliminates the database tarjetascredito if it already exists, and creates it.

* autorizar_compras(vnrotarjeta char(16), vcodseguridad char(4), vnrocomercio int, vmonto decimal(7, 2)) It takes a row of the card table, takes its data and makes the corresponding verification with a series of ifs if this card is valid, otherwise it inserts the information of the card in the table rejection. And it will return a boolean on false. is valid, its data will be stored in the table purchase and will return a boolean on true.

* generate_summary(vnroclient int, aniomes char(6)): it is located in the client row, for each client it takes its card and goes through each one of the purchases that it made and the details are incerted in detail. Then it makes the sum of all the purchases and inserts it in the header table.

* function rejection_alert(): trigger in which it will be fired if a new row is entered in the rejection table, it will take its values and insert them in the alert. In addition, if the card was rejected twice by limit, it will also change its status to suspended.

* trigger_authorize_purchases(): trigger in which it takes each of the consumption rows and applies the authorize_purchases function.
