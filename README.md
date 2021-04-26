# go-patron.creac-abstract.factory

## Abstract Factory en Go

Abstract Factory es un patrón de diseño creacional que resuelve el problema de crear familias enteras de productos sin especificar sus clases concretas.

El patrón Abstract Factory define una interfaz para crear todos los productos, pero deja la propia creación de productos para las clases de fábrica concretas. Cada tipo de fábrica se corresponde con cierta variedad de producto.

El código cliente invoca los métodos de creación de un objeto de fábrica en lugar de crear los productos directamente con una llamada al constructor (operador new). Como una fábrica se corresponde con una única variante de producto, todos sus productos serán compatibles.

El código cliente trabaja con fábricas y productos únicamente a través de sus interfaces abstractas. Esto permite al mismo código cliente trabajar con productos diferentes. Simplemente, creas una nueva clase fábrica concreta y la pasas al código cliente.

# Ejemplo conceptual

Digamos que tienes que comprar dos productos diferentes de equipamiento deportivo: un par de zapatillas y una camiseta. Te gustaría comprar todo el equipamiento de la misma marca, para que los artículos combinen.

Si intentamos traducir esto en código, la fábrica abstracta nos ayudará a crear grupos de productos para que siempre coincidan entre sí.
