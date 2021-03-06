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
        - docker build -t [nameImage] .
        - docker run -p 8080:8080 -tid [nameImage]
5. To deploy the image on AWS, you must first push to hub from the Docker console to upload the image to the cloud.
6. The AWS ECS service was used to deploy this application. The following steps must be followed:
        - First a task definition of type fargate is created in ECS where the container image is the one generated and loaded in Docker. Ports 8080 and 80 are mapped to this container
        - second, a Networking only cluster is created
        - Finally, in the cluster, the task created in the first step is executed
7. Consume the services:
        - POST -> URL: [IP-GenerateTask]:8080/mutant
                  Body: {"dna": ["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}
        - GET -> URL: [IP-GenerateTask]:8080/stats
if run locally consume the services:
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
4. Para ejecutar en local el API REST debe iniciar una terminal ubicada en la carpeta raíz del proyecto y ejecutar el siguiente comando:
        - go run main.go
Si se tiene instalado Docker puede crear una imagen e iniciarla con los siguientes comandos:
        - docker build -t [nameImage] .
        - docker run -p 8080:8080 -tid [nameImage]
5. Para desplegar la imagen en AWS primero se debe hacer push to hub desde la consola de Docker para subir la imagen a la nube.
6. Para desplegar esta aplicación se usó el servicio ECS de AWS. Se debe seguir los siguientes pasos:
        - primero se crea una task definition de tipo fargate en ECS donde la imagen del contenedor es la generada y cargada en Docker. A este contenedor se le mapean los puertos 8080 y 80
        - segundo se crea un cluster de tipo Networking only
        - por ultimo en el cluster se ejecuta la tarea creada en el primer paso
7. Consumir los servicios:
        - POST -> URL: [IP-GenerateTask]:8080/mutant 
                  Body :{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}
        - GET -> URL: [IP-GenerateTask]:8080/stats
si se ejecuta localmente consumir los servicios:
        - POST -> URL: localhost:8080/mutant 
                  Body :{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}
        - GET -> URL: localhost:8080/stats
