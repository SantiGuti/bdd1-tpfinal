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
insert into comercio values(016, 'Pintureria Ruiseñor', 'Catamarca 3023', 'B1636DKI', '011-11486352');
insert into comercio values(017, 'Kiosco Los amigos', 'Italia 473', 'B1663NXJ', '011-70546601');
insert into comercio values(018, 'Rotiseria La dorada', 'Avellaneda 3712', 'B1708GHY', '011-51365309');
insert into comercio values(019, 'Electrodomesticos Fraverino', 'El Churrinche 3084', 'B1834COJ', '011-74885170');
insert into comercio values(020, 'Casa de ropa Mimo', 'Derqui 946', 'B1804EMT', '011-74777097');
--nrocomercio de 2 digitos para hacerlo distintivo al del cliente?

insert into tarjeta values(4756326984155476, 1, '202201', '202405', '6713', 150000.00, 'vigente');
--Todes les clientes tendrán una tarjeta, excepto dos clientes que tendrán dos tarjetas cada une. Una tarjeta deberá estar expirada en su fecha de vencimiento.

insert into consumo values(4756326984155476, '6713', 15, 2000.00); --consumo válido 

--insert into cierre values();
--La tabla cierre deberá tener los cierres de las tarjetas para todo el año 2022.






