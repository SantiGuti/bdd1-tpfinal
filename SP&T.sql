create function aut_compras(vnrotarjeta char(16), vcodseguridad char(4), vnrocomercio int, vmonto decimal(7, 2)) returns boolean as $$
declare
    numtarjeta char(16);

begin
   
    select t.nrotarjeta into numtarjeta from tarjeta t where t.nrotarjeta = vnrotarjeta and estado = 'vigente';
    if not found then
        insert into rechazo values(1, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'Tarjeta no válida o no vigente');
        return false;

        
    end if;

    if t.nrotarjeta from tarjeta t where (t.codseguridad != vcodseguridad) AND (t.nrotarjeta = vnrotarjeta) then
        insert into rechazo values(2, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'Código de seguridad inválido.');
        return false;
    end if;

    /*select t.nrotarjeta into numtarjeta from compra where pagado = false and nrotarjeta = vnrotarjeta SUM(monto) + vmonto > (select limitecompra from tarjeta where nrotarjeta = vnrotarjeta);
    if not found then
        insert into rechazo values(8456, vnrotarjeta, vnrocomercio, '2021-08-24 13:32:58', vmonto, "supera limite de compra");
        return false;
    end if;*/

    if t.nrotarjeta from tarjeta t where CAST(t.validahasta AS DATE) < CURRENT_DATE then 
        insert into rechazo values(4, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'Plazo de vigencia expirado.');
        return false;
    end if;

    if t.nrotarjeta from tarjeta t where t.estado = 'suspendida' then
        insert into rechazo values(5, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'La tarjeta se encuentra suspendida.');
        return false;
    end if;

    raise notice 'Compra aceptada.';
    insert into compra values (6, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, true);
    return true;
end;
$$ language plpgsql;
