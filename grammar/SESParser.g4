grammar SESParser; // sequenced event set pattern matching language

options {
    language=Go;
}

////////////////////////// PARSER ///////////////////////////////////////////////////////////

parse
    :   ses_window? ses+ group? EOF
    ;

ses_window
    :   WINDOW_TEXT
    ;

ses
    :   set_window? event+
    ;

event
    : EVENT name=ID qty=event_qty? (WHERE event_expression (AND event_expression)*)?
    ;

event_qty
    :   PLUS # qty_plus
    |   ASTERISK # qty_asterisc
    |   (L_CURLY_BRACKET from=NUMBER? COMMA to=NUMBER? R_CURLY_BRACKET) # qty_precise
    |   (L_CURLY_BRACKET exact=NUMBER R_CURLY_BRACKET) # qty_precise_alt
    ;

set_window
    :   AND? THEN
    |   AND? THEN? SKIP_ skip=dateInterval (AND? WITHIN within=dateInterval)?
    |   AND? THEN? WITHIN within=dateInterval
    ;

date
    :   absoluteDate
    |   relativeDate
    ;

absoluteDate
    :   STRING
    ;

relativeDate
    :   LAST last=dateInterval
    ;

dateInterval
    :   num=NUMBER unit=DATE_UNIT (AND extra=dateInterval)*
    ;

event_expression
    :   left=expr_operand op=OP_LOGICAL right=expr_operand
    ;

expr_operand
    :   modifier=ID L_BRACKET eventAttr R_BRACKET # attrModified
    |   eventAttr # attr
    |   literal   # lit
    ;

eventAttr
    :   (eventName=ID DOT)? attrName=ID //qualified or short attribute name
    ;

literal
    :   NUMBER # literal_number
    |   STRING # literal_string
    ;

group
    :   GROUP eventAttr
    ;

////////////////////////////// LEXER ///////////////////////////////////////////////////////

WS : [ \r\t\n]+ -> skip ;
LINE_COMMENT: '//' ~[\r\n]* ('\r'? '\n' | EOF) -> skip;
SPACE: [ \t\r\n]+;
WINDOW_TEXT : WINDOW ANY+? (NL|SEMI) ; // delimiters are important, as the window specification is arbitrary text

// KEYWORDS
EVENT: E V E N T;
WINDOW: W I N D O W;
FROM: F R O M;
TO: T O;
LAST: L A S T;
SKIP_: S K I P;
WITHIN: W I T H I N;
THEN: T H E N;
WHERE: W H E R E;
AND: A N D;
OR: O R;
GROUP: G R O U P SPACE B Y;

// Interval type Keywords
DATE_UNIT
    : QUARTER
    | YEAR
    | MONTH
    | WEEK
    | DAY
    | HOUR
    | MINUTE
    | SECOND
    | MICROSECOND
    ;

// Operators. Comparation
OP_LOGICAL
    : EQ
    | NEQ
    | RE_EQ
    | GT
    | GTE
    | LT
    | LTE
    ;

// Constructors symbols
DOT: '.';
L_BRACKET: '(';
R_BRACKET: ')';
L_CURLY_BRACKET: '{';
R_CURLY_BRACKET: '}';
COMMA: ',';
SEMI: ';';
PLUS: '+';
ASTERISK: '*';

// LITERALS
NUMBER: DIGIT+ (DOT DIGIT+)?;
STRING: DQUOTA_STRING | SQUOTA_STRING;
ID: [a-zA-Z_] [0-9a-zA-Z_]*;

// FRAGMENTS
fragment EQ: '=';
fragment RE_EQ: '~=';
fragment NEQ: '!=';
fragment GT: '>';
fragment GTE: '>=';
fragment LT: '<';
fragment LTE: '<=';

fragment QUARTER: Q U A R T E R S?;
fragment YEAR: Y E A R S?;
fragment MONTH: M O N T H S?;
fragment WEEK: W E E K S?;
fragment DAY: D A Y S?;
fragment HOUR: H O U R S?;
fragment MINUTE: M I N U T E S?;
fragment SECOND: S E C O N D S?;
fragment MICROSECOND: M I C R O S E C O N D S?;

fragment DQUOTA_STRING: '"' ( '\\'. | '""' | ~('"'| '\\') )* '"';
fragment SQUOTA_STRING: '\'' ('\\'. | '\'\'' | ~('\'' | '\\'))* '\'';

fragment DIGIT: [0-9];
fragment SIGN:  [+-];

fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];

fragment NL : [\n];
fragment ANY : .;