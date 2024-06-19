# Teamcore service

Este proyecto es un servicio para consumir una url del servidor y devolver un JSON modificado para cumplir una tarea específica

### Primeros pasos
- Instalar docker - https://docs.docker.com/engine/install/
- Instalar Google Cloud SDK - https://cloud.google.com/sdk/docs/install-sdk?hl=es-419

### Obtener id de proyecto Google cloud
- Ve a google coud console y copia el `PROJECT_ID` (Nota: copiar el id del proyecto, no el numero, ejemplo: `glowing-harmony-123456-v9`, no el numero `123456789012`; es importante que la cuenta de Google se encuentre activa con metodo de pago, asi sea en modo gratuito y que si esta en modo de prueba no haya caducado, ya que de lo contrario aunque permita usar casi todo no dejara subir la imagen docker y no se podra completar los pasos de forma exitosa)

### Contruir imagen docker para usar en Google cloud
- Abrir la terminal en la carpeta raíz del proyecto
- Ejecutar el comando `docker build -t gcr.io/PROJECT_ID/go .`

### Usar en Google Cloud Run
- Autenticarse con Google Cloud
    `gcloud auth login`
- Configurar el proyecto
    `gcloud config set project PROJECT_ID`
- Activa Container Registry API
    `gcloud services enable containerregistry.googleapis.com`
- Activa Enable Artifact Registry API
    `gcloud services enable artifactregistry.googleapis.com --project=PROJECT_ID`
- Autenticar Docker con Google Container Registry
    `gcloud auth configure-docker`
- Subir imagen docker en Google Container Registry
    `docker push gcr.io/PROJECT_ID/go`
- Desplegar en Google Cloud Run
    `gcloud run deploy go --image gcr.io/PROJECT_ID/go --platform managed --region us-central1`
    Confirmar con y cuando pregunte

Una vez realizado todos los pasos Google Cloud Run te proporcionará una URL. Puedes probar tu servicio haciendo una solicitud HTTP a esa URL.

### Prueba del resultado subido a cuenta personal
    URL: https://go-qfmcustb5a-uc.a.run.app