#!/bin/bash

antlr4 -no-visitor -no-listener -Dlanguage=Go -o parser Sparta.g4

rrd-antlr Sparta.g4
