#ifndef CARTON_H_INCLUDED
#define CARTON_H_INCLUDED

struct CartonE{

    int vec[8];
    char direccion[15];
    //fecha tiene que ser domingo!!!

};

typedef struct CartonE * CartonPtr;


CartonPtr crearCartonTeclado(); //Comprador seleccione los numeros

CartonPtr crearCarton2(int v[8], char d[15]);

CartonPtr crearCartonAutomatico(char d[15]);

void destruirCarton(CartonPtr c);

void mostrarCarton(CartonPtr c);

//ordenar
//buscar
//buscarAcierto
//verificar triunfo!!!


#endif // CARTON_H_INCLUDED
