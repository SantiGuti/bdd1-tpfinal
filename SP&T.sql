create function aut_compras(vnrotarjeta char(16), vcodseguridad char(4), vnrocomercio int, vmonto decimal(7, 2)) returns boolean as $$
declare
    resultado record;
    suma decimal(15,2);
    

begin

    select * into resultado from tarjeta t where t.nrotarjeta = vnrotarjeta AND estado = 'suspendida';
    if found then
        insert into rechazo values(5, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'La tarjeta se encuentra suspendida.');
        return false;
    end if;
   
    select * into resultado from tarjeta t where t.nrotarjeta = vnrotarjeta and estado != 'vigente';
    if found then
        insert into rechazo values(1, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'Tarjeta no válida o no vigente');
        return false;
    end if;

    select * into resultado from tarjeta t where t.nrotarjeta = vnrotarjeta and t.codseguridad != vcodseguridad;
    if found then
        insert into rechazo values(2, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'Código de seguridad inválido.');
        return false;
    end if;

    select sum(c.monto) into suma from compra c where c.nrotarjeta = vnrotarjeta and c.pagado = false;
    if suma IS NULL then 
         suma = 00.00;
    end if;
    select * into resultado from tarjeta t where t.nrotarjeta = vnrotarjeta and (suma + vmonto) < t.limitecompra;
    if not found then
        insert into rechazo values(7, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto,'supera limite de compra');
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

create function generarResumen(numclient int, periodo date) returns void as $$
declare
    -- nombreCliente text
    -- apellidoCliente text
    -- direccionCliente text
    nrotarjeta char(16);
    -- fechaVencimiento date
    -- totalAPagar decimal (8,2)
    result record;
    suma decimal(8, 2);
    nombrecomercio text;
begin
    select c.nrotarjeta into nrotarjeta from cliente cl where cl.numclient = numclient;
    select sum(c.monto) into suma from compra co where co.nrotarjeta = nrotarjeta and co.fecha = periodo and co.pagado = true;
    select * into result from cliente c where c.numclient = numclient;
    if found then
        insert into cabecera(54, result.nombre, result.apellido, result.domicilio, nrotarjeta, date_trunc('month', periodo), date_trunc('month', periodo), date_trunc('month', periodo), suma);
    end if;
    for row in select * from compra com where com.nrotarjeta = nrotarjeta and com.fecha = periodo and com.pagado = true LOOP;
        select comercio.nombrecomercio into nombrecomercio from comercio where comercio.nrocomercio = com.nrocomercio;
        insert into detalle(54, row, c.fecha, nombrecomercio, c.monto);
    end loop;

end;
$$ language plpgsql;
