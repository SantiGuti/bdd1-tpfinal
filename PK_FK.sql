/*PRIMARY KEYS*/
alter table cliente	add constraint cliente_pk 	primary key (nrocliente);
alter table tarjeta	add constraint tarjeta_pk 	primary key (nrotarjeta);
alter table comercio add constraint comercio_pk primary key (nrocomercio);
alter table compra	add constraint compra_pk	primary key (nrooperacion);
alter table rechazo	add constraint rechazo_pk	primary key (nrorechazo);
alter table cierre	add constraint cierre_pk	primary key (año, mes, terminacion);
alter table cabecera add constraint cabecera_pk	primary key (nroresumen);
alter table detalle	add constraint detalle_pk	primary key (nroresumen, nrolinea);
alter table alerta	add constraint alerta_pk	primary key (nroalerta);

/*FOREIGN KEYS*/
alter table tarjeta add constraint tarjeta_fk foreign key (nrocliente) references cliente (nrocliente);
alter table compra add constraint compra_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
alter table compra add constraint compra_fk foreign key (nrocomercio) references comercio (nrocomercio);
alter table rechazo add constraint rechazo_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
alter table rechazo add constraint rechazo_fk foreign key (nrocomercio) references comercio (nrocomercio);
alter table cabecera add constraint cabecera_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
alter table detalle add constraint detalle_fk foreign key (nombrecomercio) references comercio (nombre);
alter table alerta add constraint alerta_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
alter table alerta add constraint alerta_fk foreign key (nrorechazo) references rechazo (nrorechazo);
alter table consumo add constraint consumo_fk foreign key (nrotarjeta) references tarjeta (nrotarjeta);
alter table consumo add constraint consumo_fk foreign key (codseguridad) references tarjeta (codseguridad);
alter table consumo add constraint consumo_fk foreign key (nrocomercio) references comercio (nrocomercio);
alter table consumo add constraint consumo_fk foreign key (monto) references compra (monto);