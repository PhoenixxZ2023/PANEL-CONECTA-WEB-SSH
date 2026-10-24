# PANEL CONECTA WEB SSH

Panel web empresarial para administrar servidores SSH, usuarios, accesos y múltiples nodos desde una plataforma centralizada.

## Estado
Base inicial del proyecto. La arquitectura está preparada para crecer por módulos sin depender del panel legacy.

## Objetivos
- Administración web centralizada.
- Gestión de múltiples servidores/nodos.
- Gestión de usuarios SSH y accesos.
- Health checks y monitoreo.
- Autenticación y roles.
- API interna para agentes.
- CLI de recuperación independiente del dominio.
- Despliegue reproducible para Ubuntu amd64/arm64.

## Estructura
- `cmd/panel`: entrada del servidor web.
- `internal`: lógica privada del backend.
- `web`: interfaz web y recursos estáticos.
- `agent`: componentes del agente de nodo.
- `configs`: configuraciones.
- `migrations`: migraciones de base de datos.
- `systemd`: unidades de servicio.
- `nginx`: reverse proxy.
- `scripts`: validaciones y mantenimiento.
- `tests`: pruebas.

## Desarrollo
Requiere Go 1.24+.

```bash
go test ./...
go build ./cmd/panel
```

Ejecutar:
```bash
PANEL_ADDR=:8080 go run ./cmd/panel
```

Health:
```bash
curl http://127.0.0.1:8080/healthz
```

## Regla del proyecto
Este repositorio es independiente de cualquier panel legacy. No se deben borrar ni modificar sistemas existentes para desarrollar este proyecto.
