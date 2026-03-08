#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#include "persona.h"
#include "carton.h"


///Docentes, Nico, Javi Vescio y Diego Cañete
//Usamos mucho el SLACK para las preguntas rápidas
//Todo estará siempre disponible en el material digital
///Tienen una guia de TPs, con los ejercicios que les iremos pidiendo
//Y en caso que no esten cancheros u olvidaron C, les dejamos un repaso de programación
///Hay un total de 8 clases antes del parcial, en cada clase habrá 1 tarea o 2,para entregar
///Además tienen un TP final individual.




///repaso - Unidad 0
//Tp de nivelación, quiniela Plus
//Un jugador (nombre y dni) compra un cartón (8 numeros del 0 al 99 sin repetir y ordenados)
//Se sortean 20 numeros, del 0 al 99, acá si se pueden repetir, si adivinaste los 8 tenes el primer premio
//con 7 o 6 aciertos premios inferiores.


int main()
{
    srand(time(0));
    printf("------ Clase 1 - AYED - REPASO -------\n");


    PersonaPtr jugador = crearPersona(123, "Nico");
    mostrarPersona(jugador);


    CartonPtr carton = crearCarton1("Av. Peron 1130");
    mostrarCarton(carton);


    return 0;
}
