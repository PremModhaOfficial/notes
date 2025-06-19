---
id: liskov-substitution
aliases:
  - liskov-substitution
tags: []
created: 2025-06-19 18:01:35
modified: 2025-06-19 18:14:04
---

# liskov-substitution

## Motivation
- we usually create class hierachies during the application development
- we extend some classes 
- but some child classes have exceptions on it behaviours
- this is not good software designe

## Do not break the functionalities of superclass in subclass
- use should be able to substitute a child class instace inplace of super class and the behaviours doesnt break the systems
```java
class Vehicle {}
class Car extends Vehicle {}

psvm() {
  run(new Car());
  run(new Vehicle());
}
```
this should result in same behaviour or should not break the functionalities

violation of liskov-substitution is also violation of [[open-close|open-close-principal]]

## designe patters that helsp
[[strategy-pattern]]
[[template-pattern]]

