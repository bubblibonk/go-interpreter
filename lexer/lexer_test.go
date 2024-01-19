package lexer
import(
	"testing"
	"monkey/token"
)
func TestNextToken(t *testing.T){
	input:=`=+(){},a;`
	tests:= []struct{
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN,"="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l:= New(input)

	for i,tt:=range tests{
		tok:=l.NextToken()
		
		if tok.Type!=tt.expectedType{
			t.Fatalf("tests[%d], wrong type expected: %q, got %q.",i,tt.expectedType,tok.Type)

		}
		if tok.Literal!=tt.expectedLiteral{
			t.Fatalf("tests[%d],wrong literal expected:%q got: %q",i,tt.expectedLiteral,tok.Literal)
		}
	}
}
