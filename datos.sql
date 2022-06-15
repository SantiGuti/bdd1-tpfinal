/*Insertamos datos de 20 clientes*/
insert into cliente values(1, 'Juan', 'Rosas', 'Serano 701', '011-68943567'); 
insert into cliente values(2, 'Martin', 'Valdez', 'Mendoza 293', '011-78908583'); 
insert into cliente values(3, 'Nicolas', 'Sanchez', 'Pringles 1073', '011-33531935'); 
insert into cliente values(4, 'Esteban', 'Alpa', 'Dorrego 489', '011-78479803'); 
insert into cliente values(5, 'Tomas', 'Torrellas', 'Urquiza 644', '011-53508576'); 
insert into cliente values(6, 'Ruben', 'Beserra', 'Parana 872', '011-39265304'); 
insert into cliente values(7, 'Ricardo', 'Gonzalez', 'Gral. Rivas 302', '011-31351376'); 
insert into cliente values(8, 'Hernan', 'Ramirez', 'Sgto. Cabral 704', '011-54903717'); 
insert into cliente values(9, 'Fernando', 'Hernandez', 'Reconquista 462', '011-60197456'); 
insert into cliente values(10, 'Mariela', 'Ferreira', 'Marcos Sastre 1923', '011-98791451'); 
insert into cliente values(11, 'Florencia', 'Molina', 'Cnel. Charlone 1587', '011-74107190'); 
insert into cliente values(12, 'Clara', 'Olivieri', 'Gral. Las Heras 2093', '011-53663515'); 
insert into cliente values(13, 'Lourdes', 'Caballero', 'Av. Congreso 3541', '011-58155389'); 
insert into cliente values(14, 'Sol', 'Loza', 'Paraguay 526', '011-53612190'); 
insert into cliente values(15, 'Mercedes', 'Ferreiro', 'Libertad 1221', '011-49197929'); 
insert into cliente values(16, 'Belen', 'Fazzi', 'Salerno 2434', '011-90515713'); 
insert into cliente values(17, 'Federica', 'Funes', 'Las Delicias 1983', '011-20950617'); 
insert into cliente values(18, 'Jimena', 'Venditto', 'Gelly y Obes 941', '011-71630935'); 
insert into cliente values(19, 'Trinidad', 'Precedo', 'Urquiza 1284', '011-62026547'); 
insert into cliente values(20, 'Delfina', 'Pacioni', 'Junin 2764', '011-35627028'); 

/*Insertamos datos de 20 comercios*/
insert into comercio values(01, 'Libreria El patito feo', 'Av. San Luis 1687', 'B1663HGK', '011-93155601');
insert into comercio values(02, 'Heladeria Gustavo', 'Serrano 1523', 'B1722NHC', '011-97684470');
insert into comercio values(03, 'Carniceria El cordero feliz', 'Ituzaingo 4896', 'B1669FUE', '011-40346435');
insert into comercio values(04, 'Libreria Dalessandro', 'Mendoza 1962', 'B1621CCP', '011-30148678');
insert into comercio values(05, 'Almacen Soledad 24hs', 'Saavedra 1131', 'B1623GTS', '011-66350988');
insert into comercio values(06, 'Verduleria Tutti frutti', 'Av. Pueyrredon 1528', 'C1118AAS', '011-21803607');
insert into comercio values(07, 'Bar Expresso portal', 'Vicente F. Lopez 768', 'B1640EUJ', '011-75439970');
insert into comercio values(08, 'Restaurant Finoli finoli', 'Av. Colon 1223', 'B7600FXF', '011-48065565');
insert into comercio values(09, 'Servicio Tecnico Esteban', 'Av. Primera Junta 459', 'B1663KHE', '011-46079515');
insert into comercio values(010, 'Almacen Los chinos', 'La Pampa 3042', 'B1844GBD', '011-49358570');
insert into comercio values(011, 'Peluqueria Enjoy', 'Cordoba 1233', 'B1825DMG', '011-92305948');
insert into comercio values(012, 'Farmacia San luis', 'Felipe Amoedo 1402', 'B1878AJD', '011-83204395');
insert into comercio values(013, 'Taller Mecanico Inyex', '11 de Septiembre 3873', 'B1666DMM', '011-93053019');
insert into comercio values(014, 'Casa de electronica Electrofer', 'Gral. Belgrano 1955', 'B1722CWM', '011-99256007');
insert into comercio values(015, 'Ferreteria Gerardo', 'Entre Rios 1016', 'B1611FZO', '011-29124305');
insert into comercio values(016, 'Pintureria Ruiseñor', 'Catamarca 3023', 'B1663NXJ', '011-11486352');
insert into comercio values(017, 'Kiosco Los amigos', 'Italia 473', 'B1663NXJ', '011-70546601');
insert into comercio values(018, 'Rotiseria La dorada', 'Avellaneda 3712', 'B1708GHY', '011-51365309');
insert into comercio values(019, 'Electrodomesticos Fraverino', 'El Churrinche 3084', 'B1834COJ', '011-74885170');
insert into comercio values(020, 'Casa de ropa Mimo', 'Derqui 946', 'B1804EMT', '011-74777097');

/*Insertamos datos de las tarjetas de los clientes*/
insert into tarjeta values(4756326984155476, 1, '201807', '202302', '6713', 500000.00, 'vigente');
insert into tarjeta values(4532969538877007, 2, '202003', '202504', '6646', 200000.00, 'suspendida');
insert into tarjeta values(4929941716451245, 3, '202204', '202702', '2312', 100000.00, 'vigente');
insert into tarjeta values(4823836840552412, 4, '202105', '202612', '8748', 250000.00, 'vigente');
insert into tarjeta values(4556919139852637, 5, '202012', '202506', '3354', 100000.00, 'vigente');
insert into tarjeta values(4024007184081593, 6, '201910', '202304', '5737', 135000.00, 'vigente');
insert into tarjeta values(5130558181821199, 6, '202002', '202405', '6997', 400000.00, 'vigente');
insert into tarjeta values(5251225878053659, 7, '202004', '202511', '6843', 250000.00, 'anulada');
insert into tarjeta values(5588490230236186, 8, '201710', '202212', '4930', 75000.00, 'vigente');
insert into tarjeta values(2720849190484829, 9, '201902', '202512', '4983', 400000.00, 'vigente');
insert into tarjeta values(5264727143830907, 10, '202201', '202705', '8469', 250000.00, 'vigente');
insert into tarjeta values(4929597785365045, 11, '201705', '202202', '6235', 75000.00, 'vigente'); --expirada
insert into tarjeta values(5124534106465188, 11, '202102', '202606', '4682', 135000.00, 'vigente');
insert into tarjeta values(4539106380553039, 12, '201805', '202308', '8655', 175000.00, 'vigente');
insert into tarjeta values(4261606383765195, 13, '201904', '202304', '3763', 500000.00, 'vigente');
insert into tarjeta values(4556365787429825, 14, '202005', '202510', '6331', 150000.00, 'suspendida');
insert into tarjeta values(4532474142653407, 15, '202202', '202709', '3018', 120000.00, 'vigente');
insert into tarjeta values(4539722778151788, 16, '202204', '202601', '9123', 150000.00, 'vigente');
insert into tarjeta values(5543040397793513, 17, '202111', '202604', '4172', 300000.00, 'vigente');
insert into tarjeta values(5331682396107249, 18, '201909', '202412', '5390', 100000.00, 'vigente');
insert into tarjeta values(5203094647795928, 19, '202006', '202404', '4529', 150000.00, 'anulada');
insert into tarjeta values(2720409166266195, 20, '202102', '202512', '1252', 1000.00, 'vigente');
--Cliente 6 y 11 tienen dos tarjetas

/*Creamos 5 compras*/
/*insert into compra values(431245, 4756326984155476, 012, '2022-05-08 12:38:05', 6000.00, true); --Compra válida
insert into compra values(466876, 4532969538877007, 010, '2022-05-12 10:32:58', 10000.00, false); --Compra rehazada por tarjeta suspendida
insert into compra values(987896, 5203094647795928, 01, '2022-04-28 17:51:03', 12000.00, false); --Compra rechazada por tarjeta anulada
insert into compra values(273289, 4929597785365045, 011, '2022-06-01 13:04:56', 4000.00, false); --Compra rechazada por tarjeta expirada
insert into compra values(231827, 5588490230236186, 013, '2022-05-23 14:27:06', 76000.00, false); --Compra rechazada por límite de compra de tarjeta excedido
insert into compra values(89034, 2720409166266195, 05, '2022-05-23 14:28:06', 500.00, false);*/

/*Creamos 4 rechazos de compra*/
/*insert into rechazo values(46878767, 4532969538877007, 010, '2022-05-12 10:32:58', 10000.00, 'La tarjeta se encuentra actualmente suspendida.');
insert into rechazo values(32423423, 5203094647795928, 015, '2022-04-28 17:51:03', 12000.00, 'La tarjeta se encuentra actualmente anulada.');
insert into rechazo values(34647372, 4929597785365045, 011, '2022-06-01 13:04:56', 4000.00, 'La tarjeta se encuentra actualmente expirada.');
insert into rechazo values(72637882, 5588490230236186, 013, '2022-05-23 14:27:06', 76000.00, 'El monto supera el límite de compra de la tarjeta.');*/

/*Insertamos datos de los cierres de todas las tarjetas para el año 2022.*/
insert into cierre values(2022,01,3,'2022-01-09','2022-01-30','2022-02-08');
insert into cierre values(2022,06,5, '2022-06-11', '2022-07-11', '2022-07-25'); 
insert into cierre values(2022,06,9, '2022-06-06', '2022-07-06', '2022-07-15');
--las fechas tienen que ser distintas para cada terminación 
--cada terminacion va a tener para cada periodo un cierre distinto

/*Insertamos los datos de cabecera del resumen*/
--insert into cabecera values(234234234, 'Juan', 'Rosas', 'Serano 701', 4756326984155476, '2022-01-09','2022-01-30','2022-02-08', 6000.00);

/*Insertamos los datos de detalle del resumen*/
--insert into detalle values(234234234, 4757990, '2022-05-08 12:38:05', 'Farmacia San luis', 6000.00);

/*Insertamos los datos de alerta a los clientes*/
/*insert into alerta values(354566, 4532969538877007, '2022-05-12 10:32:58', 46878767, 4689, 'Estimado cliente, le informamos que su compra ha sido rechazada ya que su tarjeta se encuentra actualmente suspendida.');
insert into alerta values(242424, 5203094647795928, '2022-04-28 17:51:03', 32423423, 3234, 'Estimado cliente, le informamos que su compra ha sido rechazada ya que su tarjeta se encuentra actualmente anulada.');
insert into alerta values(432543, 4929597785365045, '2022-06-01 13:04:56', 34647372, 5654, 'Estimado cliente, le informamos que su compra ha sido rechazada ya que su tarjeta se encuentra actualmente expirada.');
insert into alerta values(565768, 5588490230236186, '2022-05-23 14:27:06', 72637882, 9866, 'Estimado cliente, le informamos que su compra ha sido rechazada ya que su tarjeta ha superado el límite de compra.');*/

/*Insertamos datos de consumo para realizar pruebas*/
insert into consumo values(4756326984155476, '6713', 015, 2000.00); --Consumo válido

insert into consumo values(5543040397793513, '4172', 017, 100.00); 
insert into consumo values(5543040397793513, '4172', 016, 200.00); 
--consumo con la misma tarjeta en diferentes comercios de = cp en menos de 1 min

insert into consumo values(4823836840552412, '8748', 03, 900.00);
insert into consumo values(4823836840552412, '8748', 020, 500.00); 
--consumo con la misma tarjeta en diferentes comercios de != cp en menos de 5 min

insert into consumo values(2720409166263545, '1252', 015, 2000.00); --Número de tarjeta inexistente
insert into consumo values(5203094647795928, '4529', 08, 1000.00); --Tarjeta no vigente
insert into consumo values(4532969538877007, '6640', 09, 1000.00); --Código de seguridad incorrecto
insert into consumo values(2720849190484829, '4983', 012, 41000.00); --Límite de compra superado
insert into consumo values(4929597785365045, '6235', 011, 500.00); --Tarjeta vencida
insert into consumo values(4556365787429825, '6331', 04, 1000.00); --Tarjeta suspendida

/*Las unicas tablas con datos tienen que ser: cliente, comercio, tarjeta, cierre, consumo*/