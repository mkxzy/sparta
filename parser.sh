#!/bin/bash

antlr4 -visitor -Dlanguage=Go -o parser Sparta.g4
