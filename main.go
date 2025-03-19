package main

import (
	"genesis/pos/reportes_pos/comunes/contenedor"
	"genesis/pos/reportes_pos/dominio/constantes"
	adaptador_servidor "genesis/pos/reportes_pos/presentacion/api/gin/config/adaptador"
	"log"
	"os"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.SetOutput(os.Stdout)
	log.Println(`***** ***                                                                                         ***** **          * ***          *******    
			  ******  * **                                                   *                                  ******  ****       *  ****        *       ***  
			 **   *  *  **                                                  **                                 **   *  *  ***     *  *  ***      *         **  
			*    *  *   **                                                  **                                *    *  *    ***   *  **   ***     **        *   
			    *  *    *                ****       ****    ***  ****     ********              ****              *  *      **  *  ***    ***     ***          
			   ** **   *       ***      * ***  *   * ***  *  **** **** * ********     ***      * **** *          ** **      ** **   **     **    ** ***        
			   ** **  *       * ***    *   ****   *   ****    **   ****     **       * ***    **  ****           ** **      ** **   **     **     *** ***      
			   ** ****       *   ***  **    **   **    **     **            **      *   ***  ****              **** **      *  **   **     **       *** ***    
			   ** **  ***   **    *** **    **   **    **     **            **     **    ***   ***            * *** **     *   **   **     **         *** ***  
			   ** **    **  ********  **    **   **    **     **            **     ********      ***             ** *******    **   **     **           ** *** 
			   *  **    **  *******   **    **   **    **     **            **     *******         ***           ** ******      **  **     **            ** ** 
			      *     **  **        **    **   **    **     **            **     **         ****  **           ** **           ** *      *              * *  
			  ****      *** ****    * *******     ******      ***           **     ****    * * **** *            ** **            ***     *     ***        *   
			 *  ****    **   *******  ******       ****        ***           **     *******     ****             ** **             *******     *  *********    
			*    **     *     *****   **                                             *****                  **   ** **               ***      *     *****      
			*                         **                                                                   ***   *  *                         *                
			 **                       **                                                                    ***    *                           **              
			                           **                                                                    ******                                            
			                                                                                                   ***`)

	if err := contenedor.InicializarContenedor(); err != nil {
		log.Fatal(err)
		panic(err)
	}

	log.Println("INICIANDO SERVICIO REPORTES POS: ", constantes.HOST_IP, ":", constantes.HOST_PORT)
	if err := adaptador_servidor.Start(); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
