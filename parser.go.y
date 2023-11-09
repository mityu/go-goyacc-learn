%{
package main
%}

%union{
    token Token
    expr  Expr
}

%type<expr> program
%type<expr> expr
%type<expr> primary
%token<token> NUMBER
%token<token> '(' ')'

%left '+' '-'
%left '*' '/'

%%

program
    : expr
    {
        $$ = $1
        yylex.(*Lexer).result = $$
    }

expr
    : expr '+' expr { $$ = ExprBinary{lhs: $1, op: "+", rhs: $3} }
    | expr '-' expr { $$ = ExprBinary{lhs: $1, op: "-", rhs: $3} }
    | expr '*' expr { $$ = ExprBinary{lhs: $1, op: "*", rhs: $3} }
    | expr '/' expr { $$ = ExprBinary{lhs: $1, op: "/", rhs: $3} }
    | primary { $$ = $1 }

primary
    : NUMBER { $$ = ExprNum{literal: $1.literal} }
    | '(' expr ')' { $$ = $2 }

%%
