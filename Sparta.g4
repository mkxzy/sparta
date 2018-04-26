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

program: stmt* EOF;

stmt
    : assign_stmt
    | fundef_stmt
    | return_stmt
    | funcall_stmt
    | if_stmt
    | for_stmt
    | break_stmt
    | continue_stmt
    ;

assign_stmt: IDENTIFIER '=' test;

fundef_stmt: 'fun' fun_name fun_body;

fun_name: IDENTIFIER;

fun_body: fun_par block;

fun_par: '(' namelist? ')';

namelist: IDENTIFIER (',' IDENTIFIER)*;

return_stmt: 'return' test?;

funcall_stmt: funcall_expr;

funcall_expr: fun_name arg_expr;

// 参数表达式
arg_expr: '(' arg_list? ')';

// 参数列表
arg_list: arg (',' arg)*;

// 参数
arg: test;

if_stmt
    : 'if' test block ('else' 'if' test block)* ('else' block)?;

for_stmt: 'for' IDENTIFIER 'in' test 'to' test for_body;

for_body: block;

block: '{' stmt* '}';

break_stmt: 'break';

continue_stmt: 'continue';

test
    : compare_expr
    ;

//or_test: and_test ('or' and_test)*;
//
//and_test: not_test ('and' not_test)*;
//
//not_test: 'not' not_test | compare_expr;
//
compare_expr: arith_expr (comp_op arith_expr)?;

comp_op
    : '<'
    | '>'
    | '=='
    | '>='
    | '<='
    | '<>'
    | '!='
    ;

//expr: arith_expr;

//xor_expr: and_expr ('^' and_expr)*;
//
//and_expr: shift_expr ('&' shift_expr)*;
//
//shift_expr: arith_expr (('<<'|'>>') arith_expr)*;

// 算数表达式
arith_expr: term (('+'|'-') term)*;

// 乘法
term: factor (('*' | '/' | '%') factor)*;

// 乘法因子
factor: '-'? atom_expr;

//// N次方
//power: atom_expr ('**' factor)?;

//不可分割表达式
atom_expr
    : '(' test ')'  //括号优先表达式
    | funcall_expr          //函数调用
    | IDENTIFIER
    | INTEGER_LITERAL
    | NUMBER_LITERAL
    | STRING
    ;

// 字符串
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

INTEGER_LITERAL: DecimalIntegerLiteral;

NUMBER_LITERAL: DecimalIntegerLiteral '.' [0-9]+ ;

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
