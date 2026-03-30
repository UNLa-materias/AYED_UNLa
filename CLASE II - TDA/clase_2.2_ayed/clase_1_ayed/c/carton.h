#ifndef CARTON_H_INCLUDED
#define CARTON_H_INCLUDED

struct CartonE;

typedef struct CartonE * CartonPtr;


char* getDireccion(CartonPtr c);
int* getVec(CartonPtr c);

void setDireccion(CartonPtr c, char nuevaDireccion[15]);
void setVec(CartonPtr c, int nuevoVec[8]);

//Cuando se trabaja con arreglos se suelen hacer los get y set por posición
int getVecPos(CartonPtr c, int pos);
void setVecPos(CartonPtr c, int pos, int nuevoElementoVec);




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
