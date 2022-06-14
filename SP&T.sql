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
    insert into compra values (6, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, true);
    return true;
end;
$$ language plpgsql;



create function generar_resumen(vnrocliente int, aniomes int) returns void as $$
declare
    vnrotarjeta char(16);
    suma decimal(15, 2);
    fecha date := TO_DATE(aniomes, 'YYYYMM');
    nroresumencounter int := 0;
    nrolineacounter int := 0;
    datoscliente record;
    datoscompra record;
    nombrecomercio text;
begin
    select * into datoscliente from cliente where cliente.nrocliente = vnrocliente;
    FOR nrotarjeta IN select tarjeta.nrotarjeta into vnrotarjeta FROM tarjeta where cliente.nrocliente = vnrocliente;
    LOOP

        FOR monto IN select * into datoscompra from compra c where c.nrotarjeta = vnrotarjeta and c.periodo = fecha;
        LOOP
            select comercio.nombrecomercio into nombrecomercio from comercio co where datoscompra.nrocomercio = co.nrocomercio;
            insert into detalle(nroresumencounter, nrolineacounter, nombrecomercio, datoscompra.monto);
            UPDATE nrolineacounter SET nrolineacounter = nrolineacounter + 1;
        END LOOP;

        select sum(com.monto) into suma from compra com where com.nrotarjeta = vnrotarjeta and com.periodo = fecha;
        -- if suma IS NULL then 
        --     suma = 0.00;
        -- end if;
        insert into cabecera(nroresumencounter, datoscliente.nombre, datoscliente.apellido, datoscliente.domicilio, vnrotarjeta, fecha, fecha, fecha, suma);
        UPDATE nroresumencounter SET nroresumencounter = nroresumencounter + 1;
    END LOOP;
    
end;
$$ language plpgsql;
