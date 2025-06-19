---
id: open-close
aliases:
  - open-close
  - open-close-principal
tags: []
created: 2025-06-19 15:01:20
modified: 2025-06-19 18:13:42
---

# open-close

- new functionalities should be able to be added in the code with minimal chages to the existing code
- adding new behaviour should not cause or reqire modification of existing code

> [!faq] Definition
> software-entities should be open for Extention but closr for modification


**violation of [[single-responsibility]] will case the violation of [[open-close]] and vice versha**

## TO Avoid violation
- extend be extended explicitly A common interface should be defined instead


## Example Problems
### Progress Dialogs
- the progress bar on the ui should be displayed for various reasons like downloding the images or loading music 
- this should be determided at **runtime** and may change over time whan a new feature gets added

## helpful designe patterns
[[strategy-pattern]]
[[decorator-pattern]]
