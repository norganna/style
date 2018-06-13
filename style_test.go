package style

import "testing"

func fn(name string) func(string) string {
	return func(in string) string {
		return name + "(" + in + ")"
	}
}

var testConfig = &Config{
	ConfigGenerators: ConfigGenerators{
		GenDH: Repeater("D"),
		GenHL: Repeater("S"),
		GenSP: Repeater("_"),
		GenLI: Literal("I"),
		GenLL: Literal("L"),
	},
	ConfigColours: ConfigColours{
		HC: fn("HC"),
		LC: fn("LC"),
		BC: fn("BC"),
		IC: fn("IC"),
		EC: fn("EC"),
	},
}

func Test_Style(t *testing.T) {
	testConfig.TagSequence("‹", "›")
	runTests(t, testConfig, [][2]string{
		{"‹ test ‹dh:3››", "‹ test LC(DDD)›"},
		{"‹ test ‹bc:test››", "‹ test BC(test)›"},
		{"‹lc:‹bc:test››", "LC(BC(test))"},
		{
			"‹sp:10:<test› ‹sp:10:|test› ‹sp:10:>test› ‹sp:10:test›",
			"test______ ___test___ ______test test______",
		},
		{"‹sp:10:|tests›", "__tests___"},
		{
			"‹DH:1›‹HL:1›‹LI›‹LL›‹HC:X›‹LC:X›‹BC:X›‹IC:X›‹EC:X›",
			"LC(D)LC(S)LC(I)LC(L)HC(X)LC(X)BC(X)IC(X)EC(X)",
		},
	})
}

func Test_TagChars(t *testing.T) {
	testConfig.TagSequence("X", "Y")
	runTests(t, testConfig, [][2]string{
		{"X test Xdh:3YY", "X test LC(DDD)Y"},
		{"X test XBC:How is this?Y Y", "X test BC(How is this?) Y"},
	})
}

func Test_TagMultiChars(t *testing.T) {
	testConfig.TagSequence("{{", "}}")
	runTests(t, testConfig, [][2]string{
		{"{{ test {{BC:How {} is this?}} } }}", "{{ test BC(How {} is this?) } }}"},
	})
}

func Test_TagSameChars(t *testing.T) {
	testConfig.TagSequence("::", "::")
	runTests(t, testConfig, [][2]string{
		{"::BC:test::", "BC(test)"},
	})
}

func Test_Repeater(t *testing.T) {
	r := Repeater("∫")

	tests := []struct {
		len    int
		text   string
		expect string
	}{
		{10, "hello", "hello∫∫∫∫∫"},
		{10, "|hello", "∫∫hello∫∫∫"},
		{10, ">hello", "∫∫∫∫∫hello"},
		{10, " hello", "∫hello∫∫∫∫"},
		{10, "|hello", "∫∫hello∫∫∫"},
		{10, ">hello ", "∫∫∫∫hello∫"},
		{10, ">hello world", "hello∫wor…"},
	}

	for _, test := range tests {
		got := r(test.len, test.text)
		if got != test.expect {
			t.Errorf("Repeated text failed with %dx “%s”, expected “%s”, got “%s”",
				test.len, test.text, test.expect, got)
		}
	}
}

func runTests(t *testing.T, c *Config, tests [][2]string) {
	for _, test := range tests {
		v := c.Style(test[0])
		if v != test[1] {
			t.Errorf("Error testing “%s”, expected “%s”, got “%s”.\n",
				test[0], test[1], v,
			)
		}
	}
}
