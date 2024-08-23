# TOMLANG
## Un proyecto para aprender como funciona un interprete
### UPDATE => Se está realizando el compilador de la segunda parte.

**Intérprete en Go**

Este repositorio contiene un intérprete escrito en Go basado en el libro "Writing an Interpreter in Go" de Thorsten Ball. El intérprete está escrito de manera similar al del libro, con algunas modificaciones menores, como la traducción de los mensajes de error al español.

**Características**

* Interpreta un lenguaje de scripting simple
* Soporta variables, tipos de datos básicos y operadores
* Permite la definición de funciones y estructuras de control
* Incluye una biblioteca estándar con funciones para la entrada/salida

**Uso**

Para ejecutar el intérprete, simplemente ejecute el comando `go run main.go`. Se le presentará un intérprete interactivo donde puede ingresar comandos y expresiones.

**Ejemplos**

```go
# Sumar dos números
2 + 3

# Definir una variable
let x = 5;

# Imprimir el valor de una variable
puts(x)

# Función para calcular el factorial
fn factorial(x) {
  if (x == 0) {
    return 1;
  }
  return x * factorial(x-1);
}

# Calcular el factorial de 5
factorial(x)
```

**Modificaciones**

Las siguientes son algunas de las modificaciones que se han realizado al código del libro:

* Los mensajes de error se han traducido al español.

**Próximos pasos**

Este intérprete es un proyecto en curso. Se están realizando mejoras continuamente, como la adición de nuevas características y la mejora del rendimiento.

**Pendiente**

* Soporte para el manejo de archivos.
* Loops

**Enlaces**

* Libro "Writing an Interpreter in Go": [https://interpreterbook.com/](https://interpreterbook.com/)
* Libro "Writing a Compiler in Go": [https://compilerbook.com/](https://compilerbook.com/)
