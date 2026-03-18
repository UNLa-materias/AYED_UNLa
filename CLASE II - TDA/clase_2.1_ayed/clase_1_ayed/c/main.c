#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#include "persona.h"
#include "carton.h"


///Persona anidada con carton, SIN TDA

int main()
{
    srand(time(0));
    printf("------ Clase 1 - AYED - REPASO -------\n");


    PersonaPtr jugador = crearPersonaYCarton(123, "Nico", "Av. Peron 1130" );
    mostrarPersona(jugador);


    return 0;
}
