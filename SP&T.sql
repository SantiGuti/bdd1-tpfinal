create function aut_compras(vnrotarjeta char(16), vcodseguridad char(4), vnrocomercio int, vmonto decimal(7, 2)) returns boolean as $$
declare
    resultado record;
    suma decimal(8,2);

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

    WITH suma AS (select sum(c.monto) from compra c where c.nrotarjeta = vnrotarjeta and c.pagado = false)
    select * into resultado from tarjeta t where t.nrotarjeta = vnrotarjeta and suma < t.limitecompra;
    if not found then
        insert into rechazo values(7, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'supera limite de compra');
        return false;
    end if;
    
    select * into resultado from tarjeta t where (t.codseguridad != vcodseguridad) AND CAST(t.validahasta AS DATE) < CURRENT_DATE;
    if not found then
        insert into rechazo values(4, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, 'Plazo de vigencia expirado.');
        return false;
    end if;*/

    raise notice 'Compra aceptada.';
    insert into compra values (6, vnrotarjeta, vnrocomercio, CURRENT_TIMESTAMP, vmonto, true);
    return true;
end;
$$ language plpgsql;
