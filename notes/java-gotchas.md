---
id: java-gotchas
aliases:
  - JAVA-Gotchas
tags: []
created: 2025-07-18 18:18:22
modified: 2025-07-23 10:15:29
---

# JAVA-Gotchas

- [x] removeIf(Lambda) removes all of the matches
- [x] $ is valid IDENTIFIER!!!
- [x] equals(Object o) is overriden? ((double check the overridden mathod may be not overriden))
- Other than premitive final objects that are final will not be compiletime redolved
- instanceof will throw when checking for siblings
- LocalDate has no constructors
- String Builder is Mutable  in java immutable actualy refers to it returning new obj instade of calculated new obj
- compile time checks are only for the checked exeptions
- try without actual risky code and that becomes the risky code
- javac -d input.java -o output.java
- StrinBuilder has no equals(Object o) method overriden always return false unless it is the self
- arraylistObj.add(not-a-index) -->> theow an error
- interface method public abstract implecitly
- you can upcast or cast to primetive >> because Integer and Byte are sibs they dont conert to eachouther implecitly
- AutoBoxing and U pacastong are done Exculsivelily not at the same time
- Interface will not inherit and implement the method unless it is default method
- Casting is runtime Problem
- Exception null Ref is type bound but compiletime throw null is valid but also a runtime error
- string null != String\[0\]
- null_String.toString().length() = 4;
- null_String.toString().equals("null"); /true
- Dad said that i will call my method called Other than child must cll list own Other method which is overridden
- Handel of declare iven if you do the try-finally <- which is not the exeption handeling in pure form
- both try catch has the return block making the outsid eblock unreachable
- checked exception cannot be handelled if not thrown
