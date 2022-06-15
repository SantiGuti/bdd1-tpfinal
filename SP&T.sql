create function autorizar_compras(vnrotarjeta char(16), vcodseguridad char(4), vnrocomercio int, vmonto decimal(7, 2)) returns boolean as $$
declare
    resultado record;
    suma decimal(15,2);
begin
    select * into resultado from tarjeta t where t.nrotarjeta = vnrotarjeta;
    if not found then
        insert into rechazo values (9, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'Tarjeta no valida');
        return false;
    end if;

    select * into resultado from tarjeta t where t.nrotarjeta = vnrotarjeta AND estado = 'suspendida';
    if found then
        insert into rechazo values(5, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'La tarjeta se encuentra suspendida.');
        return false;
    end if;
   
    select * into resultado from tarjeta t where t.nrotarjeta = vnrotarjeta AND estado = 'anulada';
    if found then
        insert into rechazo values(5, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'La tarjeta se encuentra anulada.');
        return false;
    end if;

    select * into resultado from tarjeta t where t.nrotarjeta = vnrotarjeta and estado != 'vigente';
    if found then
        insert into rechazo values(1, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'Tarjeta no vigente.');
        return false;
    end if;

    select * into resultado from tarjeta t where t.nrotarjeta = vnrotarjeta and t.codseguridad != vcodseguridad;
    if found then
        insert into rechazo values(2, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'Código de seguridad inválido.');
        return false;
    end if;
    
    select sum(c.monto) into suma from compra c where c.nrotarjeta = vnrotarjeta and c.pagado = false;
    if suma IS NULL then 
         suma = 0.00;
    end if;
    select * into resultado from tarjeta t where t.nrotarjeta = vnrotarjeta and (suma + vmonto) < t.limitecompra;
    if not found then
        insert into rechazo values(7, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'La compra supera el limite de la tarjeta.');
        return false;
    end if;
        
    select * into resultado from tarjeta t where (t.nrotarjeta = vnrotarjeta) AND TO_DATE(t.validahasta, 'YYYYMM') < CURRENT_DATE;
    if found then
        insert into rechazo values(4, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'Plazo de vigencia expirado.');
        return false;
    end if;

    raise notice 'Compra aceptada.';
    insert into compra values (10, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, true);
    return true;
end;
$$ language plpgsql;



create function generar_resumen(vnrocliente int, aniomes char(6)) returns void as $$
declare
    suma decimal(15, 2);
    fechares date := TO_DATE(aniomes, 'YYYYMM');
    trecord record;
    comprarecord record;
    nroresumencounter int := 0;
    nrolineacounter int := 0;
    datoscliente record;
    vnombrecomercio text;
begin
    select * into datoscliente from cliente cli where cli.nrocliente = vnrocliente;
    for trecord in select t.nrotarjeta from tarjeta t where t.nrocliente = vnrocliente
    LOOP

        FOR comprarecord IN select * from compra c where c.nrotarjeta = trecord.nrotarjeta and (c.fecha - fechares) <= '1month'
        LOOP
            select nombre into vnombrecomercio from comercio co where comprarecord.nrocomercio = co.nrocomercio;
            insert into detalle values(nroresumencounter, nrolineacounter, comprarecord.fecha, vnombrecomercio, comprarecord.monto);
            nrolineacounter = nrolineacounter + 1;
        END LOOP;
        select sum(com.monto) into suma from compra com where com.nrotarjeta = trecord.nrotarjeta and (com.fecha - fechares) <= '1month';
        if suma IS NULL then 
            suma = 0.00;
        end if;
        
        insert into cabecera values(nroresumencounter, datoscliente.nombre, datoscliente.apellido, datoscliente.domicilio, trecord.nrotarjeta, fechares, fechares + interval '1month', fechares + interval '1month' + '1week', suma);
        nroresumencounter = nroresumencounter + 1;
    END LOOP;

end;
$$ language plpgsql;

create function rechazo_alerta() returns trigger as $$
declare
    rechazoInfo record;
begin
    select * into rechazoInfo from rechazo;  
    if found then
        insert into alerta (nrotarjeta, fecha, nrorechazo, codalerta, descripcion)
        values (new.nrotarjeta, new.fecha, new.nrorechazo, 0, new.motivo);
    end if;

    select * into rechazoInfo from rechazo
    where nrotarjeta = new.nrotarjeta
    and   motivo = new.motivo
    and   cast(fecha as date) = cast(new.fecha as date);

    if found then
        insert into alerta(nrotarjeta, fecha, nrorechazo, codalerta, descripcion)
        values (new.nrotarjeta, new.fecha, new.nrorechazo, 32, 'Tarjeta suspendida por exceso del límite de compra en el mismo día.');
        
        update tarjeta set estado = 'suspendida' where nrotarjeta = new.nrotarjeta;
       -- TESTEAR ESTE CASO
    end if;
    return new;
end;
$$ language plpgsql;

create trigger rechazo_alerta
after insert on rechazo
for each row
execute procedure rechazo_alerta();

create function compra_alerta() returns trigger as $$
declare
    compraInfo1min record;
    compraInfo5min record;
begin
    select * into compraInfo1min from compra, comercio
    where compra.nrotarjeta = new.nrotarjeta --se chequea que sea la misma tarjeta
    and comercio.nrocomercio = compra.nrocomercio --se chequea que la compra esté asociada a ese comercio
    and compra.nrocomercio != new.nrocomercio --se chequea que sean != comercios
    and comercio.codigopostal = (select codigopostal from comercio where nrocomercio = new.nrocomercio)
    --se chequea que sea el mismo cp
    and compra.fecha > CURRENT_TIMESTAMP - interval '1 minute';
    -- FIJARSE QUE ESTÉ BIEN LA LOGICA que sea en un lapso menor a 1 min

    if found then
    insert into alerta (nrotarjeta, fecha, codalerta, descripcion)
    values (new.nrotarjeta, new.fecha, 1, 'Dos compras realizadas en menos de 1 min.');
    end if;

    select * into compraInfo5min from compra, comercio --asumimos que son comercios diferentes
    where compra.nrotarjeta = new.nrotarjeta
    and comercio.nrocomercio = compra.nrocomercio
    and compra.nrocomercio != new.nrocomercio
    and comercio.codigopostal != (select codigopostal from comercio where nrocomercio = new.nrocomercio)
    -- diferente cp
    and compra.fecha > CURRENT_TIMESTAMP - interval '5 minute';
    -- FIJARSE QUE ESTÉ BIEN LA LOGICA que sea en un lapso menor a 5 min
    if found then
        insert into alerta (nrotarjeta, fecha, codalerta, descripcion)
        values (new.nrotarjeta, new.fecha, 5, 'Dos compras realizadas en menos de 5 min.');
    end if;
    return new;
end;
$$ language plpgsql;

create trigger compra_alerta
after insert on compra
for each row
execute procedure compra_alerta();

