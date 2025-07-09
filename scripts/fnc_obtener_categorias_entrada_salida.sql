CREATE OR REPLACE FUNCTION tienda.fnc_obtener_categorias_entrada_salida()
RETURNS json
LANGUAGE plpgsql
AS $$
----------------------------------------------------------------------------------
-- Autor: Carlos de las salas                                                   --
-- Sprint: 16                                                                   --
-- Fecha Creacion: 09/07/2025                                                   --
-- Descripcion: Obtener las categorias de entradas y salidas                    --
----------------------------------------------------------------------------------

DECLARE
    v_resultado json;
BEGIN
    SELECT COALESCE(json_agg(row_to_json(ces)), '[]'::json)
    INTO v_resultado
    FROM tienda.tbl_categoria_entrada_salida ces;

    RETURN v_resultado;
END;
$$;