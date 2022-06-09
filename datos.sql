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
insert into cliente values(13, 'Lourdes', 'Caballero', 'Azcuenaga 3541', '011-58155389'); 
insert into cliente values(14, 'Sol', 'Loza', 'Paraguay 526', '011-53612190'); 
insert into cliente values(15, 'Mercedes', 'Ferreiro', 'Libertad 1221', '011-49197929'); 
insert into cliente values(16, 'Belen', 'Fazzi', 'Salerno 2434', '011-90515713'); 
insert into cliente values(17, 'Federica', 'Funes', 'Las Delicias 1983', '011-20950617'); 
insert into cliente values(18, 'Jimena', 'Venditto', 'Gelly y Obes 941', '011-71630935'); 
insert into cliente values(19, 'Trinidad', 'Precedo', 'Urquiza 1284', '011-62026547'); 
insert into cliente values(20, 'Delfina', 'Pacioni', 'Junin 2764', '011-35627028'); 

insert into comercio values(01, 'Libreria El patito feo', 'Av San Luis 1687', '83645921', '1193155601');
--nrocomercio de 2 digitos para hacerlo distintivo al del cliente?

insert into tarjeta values(4756326984155476, 1, '202201', '202405', '6713', 150000.00, 'vigente');
--Todes les clientes tendrán una tarjeta, excepto dos clientes que tendrán dos tarjetas cada une. Una tarjeta deberá estar expirada en su fecha de vencimiento.

insert into consumo values(4756326984155476, '6713', 15, 2000.00); --consumo válido 

--insert into cierre values();
--La tabla cierre deberá tener los cierres de las tarjetas para todo el año 2022.






