---
id: java-polymorphism
aliases:
  - polymorphism
tags:
  - #lang/java
created: 2025-07-15 14:29:51
modified: 2025-07-15 15:13:36
---

# polymorphism

## Method Overriding
### Rules
- Same Params and Return Type
- has same or more access level
## Method OverLoading

## Method Resolution
### CompileTime
- Check the Refrence Type if it has the method implemented or not
- this ensures that at runtime we have atleast something to call
### RunTime
- Check the childern / most specific implementattion

### WHY ??
- becouse when we Refrence a variable in a general manner we want tou use it in a general manner
- But then we shuld not be thinking of specifics if we did then we must use the more specific or the actual refrence name itself

## Method Binding
### Method Sinature Binding | Compile-time Binding | Early Binding
- complier checks if the ref type has this method signature
- if it has it then write the details to the bytecode OTHERWISE throw compile time error
- Applies to the Static and Instance Methods

### Method Implementation Binding | Runtime Binding | Late Binding
- Static -> At compile time [!note] still called Compiltime binding
- Instance -> At Runtime
