WELCOME

This is the repository to discover mutants.

INSTALLATION

1. Clone the repository
2. Create a file named .env in the root of the project
3. The .env file must have the following variables:
        - AWS_ACCESS_KEY_ID
        - AWS_SECRET_ACCESS_KEY
Which represent the access data of the AWS user with the respective permissions to use the DynamoDB service
4. To execute the REST API locally, you must start a terminal located in the root folder of the project and execute the following command:
        - go run main.go
If you have Docker installed you can create an image and start it with the following commands:
        - docker build -t [nameImage].
        - docker run -p 8080:8080 -tid [nameImage]
5. Consume the services:
        - POST -> URL: localhost:8080/mutant
                  Body: {"dna": ["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}
        - GET -> URL: localhost:8080/stats


BIENVENIDO

Este es el repositorio para descubrir mutantes.

INSTALACIÓN

1. Clonar el repositorio
2. Crear un archivo con nombre .env en la raíz del proyecto 
3. El archivo .env debe tener las siguientes variables:
        - AWS_ACCESS_KEY_ID
        - AWS_SECRET_ACCESS_KEY
Las cuales representan los datos de acceso del usuario de AWS con los respectivos permisos para utilizar el servicio de DynamoDB
4. Para ejecutar en local el API REST debe iniciar una terminal ubicarse en la carpeta raíz del proyecto y ejecutar el siguiente comando:
        - go run main.go
Si se tiene instalado Docker puede crear una imagen e iniciarla con los siguientes comandos:
        - docker build -t [nameImage] .
        - docker run -p 8080:8080 -tid [nameImage]
5. Consumir los servicios:
        - POST -> URL: localhost:8080/mutant 
                  Body :{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}
        - GET -> URL: localhost:8080/stats
