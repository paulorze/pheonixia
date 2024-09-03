# PhoenixAI
PhoenixAI es una API RESTful para el manejo y análisis de documentación utilizando inteligencia artificial.

## Características
- **Autenticación**: Registro de usuarios, inicio de sesión y manejo de sesiones mediante tokens JWT.
- **Gestión de Documentación**: Ingestión de Documentación y almacenamiento de documentación vectorizada en Base de Datos.
- **Protección de Rutas**: Autenticación y autorización de usuarios mediante JWT.

## Requisitos
- Go 1.18 o superior
- PostgreSQL
- Librerías:
  - `github.com/gofiber/fiber/v2`
  - `github.com/golang-jwt/jwt/v4`
  - `github.com/golang-jwt/jwt/v5`
  - `github.com/google/uuid`
  - `github.com/jackc/pgx/v5`
  - `github.com/joho/godotenv`
  - `github.com/jung-kurt/gofpdf`
  - `github.com/stretchr/testify`
  - `github.com/tmc/langchaingo`
  - `golang.org/x/crypto`
  - `gorm.io/driver/postgres`
  - `gorm.io/gorm`

## Instalación

1. Clona este repositorio:

    ```sh
    git clone https://github.com/Florencia-Harmath/phoenixia.git
    cd phoenixia
    ```

2. Crea un archivo `.env` en la raíz del proyecto con las siguientes variables de entorno:

    ```
        DB_HOST=localhost
        DB_PORT=5432
        DB_PASS=postgres
        DB_USER=postgres
        DB_SSLMODE=disable
        DB_NAME=pgvector
        JWT_SECRET=supersecret
    ```

3. Instala las dependencias:

    ```sh
    go mod tidy
    ```

4. Asegurate que tu versión de PostgreSQL es compatible con vectores:
### Windows

Ensure [C++ support in Visual Studio](https://learn.microsoft.com/en-us/cpp/build/building-on-the-command-line?view=msvc-170#download-and-install-the-tools) is installed, and run:

```cmd
call "C:\Program Files\Microsoft Visual Studio\2022\Community\VC\Auxiliary\Build\vcvars64.bat"
```

Note: The exact path will vary depending on your Visual Studio version and edition

Then use `nmake` to build:

```cmd
set "PGROOT=C:\Program Files\PostgreSQL\16"
cd %TEMP%
git clone --branch v0.7.4 https://github.com/pgvector/pgvector.git
cd pgvector
nmake /F Makefile.win
nmake /F Makefile.win install
```

5. Inicia el servidor del backend:

    ```sh
    go run cmd/main.go
    ```

## Endpoints

### Usuarios

- **Listado de usuarios**
  - `GET /api/users/`

- **Obtener usuario**
  - `GET /api/users/:id`

- **Actualización de usuario**
  - `PUT /api/users/:id`
  - Cuerpo de la solicitud:
    ```json
    {
      "ID": "123asd123asd",
      "email": "john.doe@example.com",
      "firstName": "John",
      "lastName": "Doe",
    }
    ```

- **Eliminación de usuario**
  - `DELETE /api/users/:id`

- **Registro de usuario**
  - `POST /api/users/register`
  - Cuerpo de la solicitud:
    ```json
    {
      "username": "johndoe",
      "password": "your-password",
      "email": "john.doe@example.com",
      "firstName": "John",
      "lastName": "Doe"
    }
    ```

### Engine

- **Subida de archivo PDF**
  - `POST /api/engine/upload`

- **Consulta a la AI**
  - `POST /api/engine/ask`

### Próximos Pasos
- **Incorporar funcionalidad JWT.**
- **Incorporar funcionalidad de autenticación.**
- **Incorporar módulos de testing.**
- **Revisar arquitectura. Ver ubicación /database, /engine, /errors, /storage, /utils. Reveer si no hay que separar las utils en otro lugar**
- **Revisar coherencia de nombres de archivos y variables.**
- **Actualizar tipos de errores en Engine.**
- **Agregar historial en /api/engine/ask.**
- **Agregar todo al main.**