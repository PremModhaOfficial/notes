---
id: jvm-classloader-sybsytem
aliases:
  - jvm-classloader-sybsytem
  - Load
tags:
  - #concept/java/jvm
  - #syntheses/jvm-classloader-sybsytem
created: 2025-07-11 10:49:10
modified: 2025-07-11 18:44:33
---


# Load
- invoke classloader and try to load classes by hierarchy and security
# Link
- resole symbolic ref to actual ref / direct refrences
- alocate space for the static variables and assigne default values

> [!note] Alocate The space for every single static variale
> even the user defined values are set to default for the time being

# Init
- itialize the static variables with user-defined values and execute static blocks inside resolved classes
