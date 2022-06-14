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
    
    -- TESTEAR ESTE
    -- NOT NULL VALUE A VER SI
    -- NOS AHORRAMOS UN IF
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
    insert into compra values (6, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, true);
    return true;
end;
$$ language plpgsql;



create function generar_resumen(vnrocliente int, aniomes char(6)) returns void as $$
declare
    suma decimal(15, 2);
    fechares date := TO_DATE(aniomes, 'YYYYMM');
    trecord record;
    comprarecord record;
    nroresumencounter int := 15;
    nrolineacounter int := 15;
    datoscliente record;
    nombrecomercio text;
begin
    select * into datoscliente from cliente cli where cli.nrocliente = vnrocliente;
    for trecord in select t.nrotarjeta from tarjeta t where t.nrocliente = vnrocliente
    LOOP

        FOR comprarecord IN select * from compra c where c.nrotarjeta = trecord.nrotarjeta and date_trunc('month', c.fecha) = date_trunc('month', fechares) and date_trunc('year', c.fecha) = date_trunc('year', fechares)
        LOOP
            select nombrecomercio into nombrecomercio from comercio co where comprarecord.nrocomercio = co.nrocomercio;
            insert into detalle values(nroresumencounter, nrolineacounter, nombrecomercio, comprarecord.monto);
            nrolineacounter = nrolineacounter + 1;
        END LOOP;
        select sum(com.monto) into suma from compra com where com.nrotarjeta = trecord.nrotarjeta and date_trunc('month', com.fecha) = date_trunc('month', fechares) and date_trunc('year', com.fecha) = date_trunc('year', fechares);
        if suma IS NULL then 
            suma = 0.00;
        end if;
        
        insert into cabecera values(nroresumencounter, datoscliente.nombre, datoscliente.apellido, datoscliente.domicilio, trecord.nrotarjeta, fechares, fechares, fechares, suma);
        nroresumencounter = nroresumencounter + 1;
    END LOOP;

end;
$$ language plpgsql;
