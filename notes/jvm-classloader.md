---
id: jvm-classloader
aliases:
  - how class loader works?
  - What it does?
tags: []
created: 2025-07-10 15:58:56
modified: 2025-07-14 11:19:00
---

# What it does?
- load the [[dotclass-file|.class]] files from disk to the [[jvm-method-area|method-area]]
- and the corresponding data like..
  * fully qualified name of the class (String -> java.lang.String)
  * Immediate parent (Fullname)
  * what it represents? (class, interface, Enum)
  * method constructors variates
  * modifiers information
  * constant pool information
- then make a isntance of `class classObject {}` to represent the loded class files into [[jvm-heap|heap]] area
# how class loader works?
- ask my parent
- if there in no parent or no resolution from the parent then resolve my self

# Why is it this way??
- Always trust the libs/class that are from the top and decrease trust as we come down 
  * basicaly as you move to parent classloader the security constants decrease
- Depending on the trust level you must conduct security checks
- user level code will go rigorous testing and verification

![[jvm-classloader-ex.excalidraw]]

## Verification Checks
- final classes are not subclassed
- final methods are not overriden
- No illegal method overloading
- bytecode integrity
  * jump intr. (if condition doent send the vm beyond end of method)


# when is it loaded for the first-time (class and interface)
- new instace is created
- invoking a ststic method
- Accessing A static field
  * *exception* when acessing compile time constants
- Subclass is loaded OR *When a Sub Interface Is loaded*
- Run from cmd
- Reflection


> [!note] Class Creation [[java-classobject]]
> JVM uses class-object to create the instance of the class

---

[[jvm-classloader-sybsytem]]
