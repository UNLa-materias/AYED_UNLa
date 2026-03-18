#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#include "persona.h"
#include "carton.h"


///Persona anidada con carton, CON TDA

int main()
{
    srand(time(0));
    printf("------ Clase 1 - AYED - REPASO -------\n");


    PersonaPtr jugador = crearPersonaYCarton(123, "Nico", "Av. Peron 1130" );
    mostrarPersona(jugador);

    //Quiero cambir algo del carton
    setVecPos(getCartonPtr(jugador),2, 99);
    mostrarPersona(jugador);

    int vector[8] = {2,63,96,6,3,2,12,1};

    setVec(getCartonPtr(jugador),vector);
    mostrarPersona(jugador);


    return 0;
}
