---
id: java-methods
aliases:
  - Methods
tags: []
created: 2025-07-09 17:00:43
modified: 2025-07-09 18:26:06
---

# Methods
## Types
### Static
- doesn't reqire the Instance of the onject to call
- can be used to facilitate the utilitary oporations
### Instance
- requre the object instance to call
- may modify the onject state

## Singnature
- methodName
- parameter Types in order

> [!note] gotchas!!
> No Return type and Modifiers are taken into account in method signature

## Overloading
- Same name in the same scope of the code
- Diffrent Singnature of the Method to distinguish from each other

> [!note]  type autocasting
> lets say sum is overloded buy not for byte type so it will find the closet type in this case it will be the short datatype 
> this only goes upward and will go for the Object at last
> Call the next big data type
> this happens at compile time also called method binding

> [!note] upcasting in exprs
> when using multiple datatypes java tries to convert every variable type to the biggest one of them

## Overriding
- Same name in the children scope for specialization

## Arguments
### Var args
- only alowed as the last arg and only once
- three dots
```java
void foo(boolean, int... items) {}
foo( true, new int[]{1,2} ); // explicit
foo( true,1,2 ); // same
foo( true ) // also works
```
- how is it diffrent than foo(int[] items) {}
  * flexible
    + dont have to pass null
    + dont have to init a data type
  * example
    + printf(String format, Object... args) {}



[[java--contractor|contractor]]
