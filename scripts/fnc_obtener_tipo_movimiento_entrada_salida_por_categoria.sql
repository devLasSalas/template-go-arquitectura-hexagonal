CREATE OR REPLACE FUNCTION tienda.fnc_obtener_tipo_movimiento_entrada_salida_por_categoria(i_id_categoria int)
RETURNS json
LANGUAGE plpgsql
AS $$
----------------------------------------------------------------------------------
-- Autor: Carlos de las salas                                                   --
-- Sprint: 16                                                                   --
-- Fecha Creacion: 09/07/2025                                                   --
-- Descripcion: Obtener los tipos de movimiento por categoria id                --
----------------------------------------------------------------------------------
DECLARE
    v_resultado json;
BEGIN
    SELECT COALESCE(json_agg(row_to_json(ces)), '[]'::json)
    INTO v_resultado
    FROM tienda.tbl_tipo_movimiento_entrada_salida ces
	WHERE ces.estado_id = 1 AND ces.categoria_entrada_salida_id = i_id_categoria;

    RETURN v_resultado;
END;
$$;
