create function aut_compras(vnrotarjeta char(16), vcodseguridad char(4), vnrocomercio int, vmonto decimal(7, 2)) returns boolean as $$
declare
    resultado record;
begin
    
    select * into resultado from tarjeta where nrotarjeta = vnrotarjeta and estado = 'vigente';
    if not found then
        insert into rechazo values(8456, vnrotarjeta, vnrocomercio, '2021-08-24 13:32:58', vmonto, "tarjeta no valida o no vigente");
        return false;
    end if;

    select * into resultado from tarjeta where codseguridad = vcodseguridad;
    if not found then
        insert into rechazo values(8456, vnrotarjeta, vnrocomercio, '2021-08-24 13:32:58', vmonto, "codigo de seguridad invalido");
    end if;

    select * into resultado from compra where pagado = false;
    if not found then
        insert into rechazo values(8456, vnrotarjeta, vnrocomercio, '2021-08-24 13:32:58', vmonto, "codigo de seguridad invalido");
    end if;

    insert into compra values (84561, vnrotarjeta, vnrocomercio, '2021-08-24 13:32:58', vmonto, true);
    return true;
end;
$$ language plpgsql;
