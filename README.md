# Flutter GoFiber Demo

Demo to describe the best practices, project structure

## Layered Architecture:

- frontend (presentation)
- app
  - api
    - handlers
    - routes
  - services
    - user-service
    - repo-service
- domain
  - models
    - account
    - repo
- persistence
  - providers
    - account-provider
    - repo-provider
- database
  - dao
    - account
    - repo
  - connection

## References:

https://central.thoughtworks.net/home/catalyst-delivery-accelerator-starter-kits#lang
https://github.com/gofiber/boilerplate
https://github.com/gofiber/recipes
https://en.wikipedia.org/wiki/Multitier_architecture
