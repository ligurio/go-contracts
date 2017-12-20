### Introduction

This module provides a draft of Golang package with implementation of
collection of decorators that makes it easy to write software using contracts.

Contracts are a debugging and verification tool. They are declarative
statements about what states a program must be in to be considered "correct" at
runtime. They are similar to assertions, and are verified automatically at
various well-defined points in the program. Contracts can be specified on
functions and on classes.

Contracts consist of two parts: a description and a condition. The description
is simply a human-readable string that describes what the contract is testing,
while the condition is a single function that tests that condition. The
condition is executed automatically and passed certain arguments (which vary
depending on the type of contract), and must return a boolean value: True if
the condition has been met, and False otherwise.

The main idea of using contracts was described in [Applying 'design by
contract'](http://cs.uns.edu.ar/~dcm/tdp/downloads/Material%20de%20Estudio/Design-by-Contract.pdf)
by B. Meyer.
