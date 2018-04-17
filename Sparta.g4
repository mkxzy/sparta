/*
BSD License

Copyright (c) 2018, Michael Smith
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:

1. Redistributions of source code must retain the above copyright
   notice, this list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright
   notice, this list of conditions and the following disclaimer in the
   documentation and/or other materials provided with the distribution.
3. Neither the NAME of Rainer Schuster nor the NAMEs of its contributors
   may be used to endorse or promote products derived from this software
   without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

grammar Sparta;

program: (stmt)* EOF;

stmt
    : simple_stmt
    ;

simple_stmt
    : expr_stmt
    ;

expr_stmt
    : primary_expr '=' postfix_expr
    | postfix_expr
    ;

primary_expr
    : IDENTIFIER
    ;

//assign_stmt: IDENTIFIER '=' test;

postfix_expr: or_test;

or_test: and_test ('or' and_test)*;

and_test: not_test ('and' not_test)*;

not_test: 'not' not_test | compare_expr;

compare_expr: arith_expr (comp_op arith_expr)?;

comp_op
    : '<'
    |'>'
    |'=='
    |'>='
    |'<='
    |'<>'
    |'!='
//    |'in'
//    |'not' 'in'
//    |'is'
//    |'is' 'not'
    ;

//expr: arith_expr;

//xor_expr: and_expr ('^' and_expr)*;
//
//and_expr: shift_expr ('&' shift_expr)*;
//
//shift_expr: arith_expr (('<<'|'>>') arith_expr)*;

arith_expr: term (('+'|'-') term)*;

term: factor (('*' | '/' | '%') factor)*;

factor
    : ('+' | '-') factor
    | power;

power: atom_expr ('**' factor)?;

atom_expr
    : IDENTIFIER '(' (arg_list)? ')' //函数调用
    | atom
    ;

atom
    : '(' postfix_expr ')'
    | NUMBER_LITERAL
    | IDENTIFIER
    | STRING
    ;

arg_list: argument (',' argument)*;

argument
    : postfix_expr
    ;

STRING
    : '"' StringCharacter* '"'
    ;

fragment
StringCharacter
	: ~["\\\r\n]
	| EscapeSequence
	;

fragment
EscapeSequence
	: '\\' [btnfr"'\\]
	;

//fragment
//UNICODE_CHAR   : ~[\u000A] ;

NUMBER_LITERAL
    : DecimalIntegerLiteral '.' [0-9]*
    | DecimalIntegerLiteral
    ;

IDENTIFIER: Letter LetterOrDigit*;

COMMENT:'/*' .*? '*/' -> skip;

LINE_COMMENT:'//' ~[\r\n]* -> skip;
    
WS: [ \t\u000C\r\n]+ -> skip;

SHEBANG: '#' '!' ~('\n'|'\r')* -> channel(HIDDEN);

fragment LetterOrDigit
    : Letter
    | [0-9]
    ;
fragment Letter
    : [a-zA-Z$_] // these are the "java letters" below 0x7F
    | ~[\u0000-\u007F\uD800-\uDBFF] // covers all characters above 0x7F which are not a surrogate
    | [\uD800-\uDBFF] [\uDC00-\uDFFF] // covers UTF-16 surrogate pairs encodings for U+10000 to U+10FFFF
    ;

fragment DecimalIntegerLiteral
    : '0'
    | [1-9] [0-9]*
    ;
