#!/bin/bash

antlr4 -visitor -Dlanguage=Go -o parser Sparta.g4

rrd-antlr Sparta.g4
