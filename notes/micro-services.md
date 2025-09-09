---
id: micro-services
aliases: []
tags: []
backlinks: []
created: 2025-12-31 12:10:59
links: []
modified: 2026-01-01 12:26:22
status: reviewed
---

# Micro-Services

## Spill the Monolith into a micro-services architecure
### intro
- Traditional e-com

### Attempt 1 -> spilt by service layer
- components
    * fronted
    * backend
    * database
- prolems
    * in order to cange the frontend we need to also cordinate with the backend team for the new stuff and apis

- require / solution
    * we must group the dependent entities together
    * in other words related elements that change togather should *stay* together also called **COHESION** #cohesion

### Attempt 2 -> split by tech (programming language)
- components
    * fronted
        + a mobile-app team (ie. dart)
        + a web-app team (ie. react-ts)
    * backend
        + recommendations-A (c++)
        + recommendations-B (python for ML)
        + core app (java)
    * database
- prolems
    * the management is hard to get right in order to change the recommendation part of the product we dont know wich team will be responsible because it is not defined purely based on tech

- require / solution
    * the components should follow [[single-responsibility]] principle, and do one thing and only one thing but do it exceptionaly well

### Attempt 3 -> split everything (per package or per class)
- components
    * fronted x 3
    * backend x 15
    * database x 1
- prolems
    * all the servises need to comunicate to eachother in order to progress
    * this is bad and the management of such service is too complex and not feasible from a business Perspective

- require / solution
    * we need **Loose Coupling**

WE need to check if the servises follow
- cohesion
- single-responsibility
- Loose Coupling


## HOW to atualy spilt the architecure
### By Business Capabilities
- Product Service
- Orders Service
- Web-app Service
- Shipping Service
- Inventory Service
- Reviews Service

### By Domain / SubDomain (By view of devs)
- Core 
    -> key differentiator (main selling point of the product)
    -> without this the product cannot provide value
- Support
    -> integral to the delevry of the core Capabilities
    -> donesnt differente us from the competitors
- Generic
    -> not specific to anything / can be used from libs / shelf

#### example
- Core
    * Product catalog
- Support
    * Orders
    * Inventory
    * Shipping
- Generic
    * Reviews
    * Payments
    * Search
    * Images
    * Image Compressions
    * Security
    * web ui

