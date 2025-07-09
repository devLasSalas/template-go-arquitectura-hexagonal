CREATE OR REPLACE FUNCTION tienda.fnc_obtener_bodegas_por_empresa(i_empresa_id int)
RETURNS json
LANGUAGE plpgsql
AS $$
----------------------------------------------------------------------------------
-- Autor: Carlos de las salas                                                   --
-- Sprint: 16                                                                   --
-- Fecha Creacion: 09/07/2025                                                   --
-- Descripcion: Obtener las bodegas por empresa id                   			--
----------------------------------------------------------------------------------
DECLARE
    v_resultado json;
BEGIN
    SELECT COALESCE(json_agg(row_to_json(ces)), '[]'::json)
    INTO v_resultado
    FROM tienda.tbl_bodegas ces
	WHERE ces.estado_id = 1 AND empresas_id = i_empresa_id;

    RETURN v_resultado;
END;
$$;