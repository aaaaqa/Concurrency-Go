# Concurrency-Go
Repository used to store the source code of the TA1 concurrent and traditional linear regression programs

## Introducción
La concurrencia, introducida por primera vez en “Solution of a problem in concurrent programming control” de Edger Dijkstra, se refiere a la capacidad de un programa para ejecutarse de manera desordenada o parcialmente ordenada, sin que ello afecte los resultados. Por si misma, esta estructura no es inherentemente más eficiente que su contraparte tradicional, pues su propósito radica en la composición de ejecuciones de forma independiente. No obstante, una implementación adecuada de la concurrencia puede potencialmente reducir los tiempos de espera entre algoritmos con recursos compartidos.
En el contexto práctico de este concepto, el curso Programación Concurrente y Distribuida ha propuesto el desarrollo de un algoritmo de regresión lineal utilizando gorutinas en el lenguaje Go. Además, con el objetivo de comparar resultados, también se consideró la implementación del algoritmo siguiendo una estructura a tradicional.
## Propuesta
La implementación de este proyecto consistió en tres etapas. En primer lugar, se inició con la investigación con herramientas virtuales para establecer un plan de acción. Posteriormente, se procedió con el desarrollo y prueba de dos archivos de código fuente en el software Visual Studio Code. Finalmente, se elaboró un informe del proceso que se siguió y los recursos utilizados.
### Primera etapa – Investigación:
Las fuentes utilizadas para la investigación de este proyecto se pueden resumir en tres.
-	Prompts de búsqueda en la aplicación de chatbot ChatGPT:
Esta fuente fue utilizada para recabar información general, previa al desarrollo del código fuente. Los prompts de búsqueda fueron dirigidos para establecer las consideraciones generales del proyecto.

![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/36ba6ffc-0fb6-464d-b462-452d7f9398d1)

![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/ccfb0dfb-a0e8-4ea0-9b9f-cf44f7ffa7ae)

Figura 1. Prompts de búsqueda en agente ChatGPT.

Se puede acceder a la conversación completa mediante el siguiente enlace: https://chat.openai.com/share/a3c3b8bb-aa14-4090-91f8-f2568d5d38c1 
-	Tour de Go:
Posterior a la investigación mediante el agente ChatGPT, se procedió con las herramientas de Tour de Go. El objetivo de esta búsqueda fue el entendimiento de canales en Go.

 ![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/f0eafd8c-d04b-4de0-a278-c57a65869074)

Figura 2. Sección de canales en Tour de Go.

Se puede acceder al recurso utilizado mediante el siguiente enlace: https://go.dev/tour/concurrency/2
-	Stack Overflow:
Esta fuente fue utilizada para profundizar el conocimiento en canales. La principal fuente utilizada en este medio fue la pregunta “What are channels used for?” realizada por el usuario MisterVerleg el 3 de octubre del 2016.

![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/317fd11b-dd66-4a41-867c-0ea9713e3782)

Figura 3. Pregunta en Stack Overflow del usuario MisterVerleg.

Se puede acceder a la pregunta completa mediante el siguiente enlace: https://stackoverflow.com/questions/39826692/what-are-channels-used-for
### Segunda etapa – Desarrollo:
La segunda etapa corresponde al desarrollo de los códigos fuente. Las herramientas utilizadas para este fin consisten en el editor Visual Studio Code y el lenguaje de programación Go. Además, también se creó un repositorio en la plataforma GitHub para almacenar el código.
En primera instancia, se desarrolló el código con estructura tradicional. Entre las consideraciones especiales del procedimiento, se consideró la partición de los arreglos de datos pseudoaleatorios durante el cálculo de la pendiente. El motivo de esta consideración radica en su práctica extendida en otras aplicaciones de aprendizaje automático, como en el entrenamiento redes neuronales, debido a que permite una mejor gestión de recursos computacionales.

![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/dd7e6b8b-ea56-4280-ad39-0b71be85c46c)

Figura 4. Funciones de cálculo de regresión lineal en programación tradicional.

En segunda instancia, se desarrolló el código con estructura concurrente. Dado que el cálculo parcial de la pendiente puede ser ejecutada de manera desordenada, el punto de mejora entre ambos códigos radica en la implementación de gorutinas para la computación de dicho valor, pues al compartir la referencia de las variables xMean, yMean y slope, se puede aplicar eficientemente esta estructura. Por otro lado, es importante precisar que se utilizan una cantidad de canales equivalente al número de segmentos en los que el arreglo es divido para evitar la condición de carrera entre procesos.

![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/eb0e7d7b-444e-4dba-a508-bf1708264081)

Figura 5. Funciones de cálculo de regresión lineal en programación concurrente.

### Tercera etapa – Informe:
Finalmente, se realizó un informe del proceso que se llevó a cabo para el desarrollo de la tarea académica, precisando la información encontrada en este documento.
## Resultados
### Primera prueba:
En esta prueba se comprobó el funcionamiento del código a escala pequeña. Los arreglos utilizados como datos fueron definidos de la manera que se visualiza en la figura 6.

![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/03c1ab90-7910-4c37-a37a-25c6f096118a)

Figura 6. Definición de arreglos con datos generados manualmente.

![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/27212f28-f9d8-4641-948e-a5220a64fdab)

Figura 7. Resultados de los códigos fuente – Prueba de funcionamiento.

Como se evidencia en la figura 7, el resultado de ambos códigos es correcto. Otro aspecto a resaltar es el tiempo de ejecución, el cual es menor en la programación tradicional.
### Segunda prueba:
En esta prueba se comprobó la velocidad de ejecución del código en una escala mayor, tomando 1 millón de datos pseudoaleatorios para cada arreglo entre los valores 0 y 1,000. La definición de los valores corresponde a la figura 8.

![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/a64583e8-0d72-4266-b7f7-02122418e82f)

Figura 8. Definición de arreglos con 1 millón de datos pseudoaleatorios.

![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/99a9a4e5-7161-4f34-bc88-118d9a3021c8)

Figura 9. Resultados de códigos fuente con 10 segmentos de arreglo – Prueba de velocidad.

![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/c5cbf908-2737-4f51-8565-5cc9add44dcd)

Figura 10. Resultados de códigos fuente con 10 segmentos de arreglo – Prueba de velocidad. 

Como se evidencia en las figuras 9 y 10, el tiempo de ejecución del código con programación concurrente es menor. Aunque la diferencia es despreciable a esta escala, supone una disminución de aproximadamente 22% en el tiempo total.
### Tercera prueba:
En esta prueba se comprobó la existencia de condición de carrera en el código concurrente. Para ello, se utilizó el indicador -race durante la ejecución del programa. Esta herramienta integrada de Go permite detectar si existe condición de carrera mediante la comprobación del cuándo y como se ha accedido a la memoria, así como el acceso desincronizado a variables compartidas.

![image](https://github.com/aaaaqa/Concurrency-Go/assets/91223158/7c74c346-23a5-41f6-9f91-df5d66966f09)

Figura 11. Resultado de códigos fuente - Prueba de condición de carrera.

Aunque la programación tradicional no es susceptible a la aparición de condiciones de carrera, el uso de esta herramienta supone una ejecución más lenta, por lo que también se consideró el uso del indicador -race para tener un punto de comparación bajo condiciones similares.
Como se evidencia en la figura 10, la herramienta de detección de código fuente no imprimió alguna alerta en el terminal, por lo que se puede asegurar que no existe dicho error.
## Conclusiones
Tras el análisis de los resultados, se puede concluir lo siguiente:
-	El uso de concurrencia en el cálculo de regresiones lineales es despreciable en conjuntos de datos reducidos. Sin embargo, los beneficios de su aplicación pueden ser cruciales en implementaciones que abarquen una mayor cantidad de información.
-	El desarrollo de este trabajo académico cumple con los requerimientos solicitados por el curso Programación Concurrente y Distribuida, tomando en consideración la implementación de programación concurrente y tradicional, así como la prevención de la condición de carrera y el uso exclusivo de librerías nativas en Go.
-	El uso de canales para la comunicación entre gorutinas evitó exitosamente la condición de carrera.
