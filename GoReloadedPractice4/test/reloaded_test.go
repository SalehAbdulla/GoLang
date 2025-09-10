package reloaded

import (
	// "strings"
	"reloaded/internal/commands"
	"testing"
)

func TestProcessCommands(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// --- Contractions ---
		{
			name:     "Contraction: don't",
			input:    `' don't '`,
			expected: `'don't'`,
		},

		// --- Text Transformation Commands ---
		{
			name:     "Lowercase command",
			input:    "I should stop SHOUTING (low)",
			expected: "I should stop shouting",
		},
		{
			name:     "Uppercase command",
			input:    "good morning (up)",
			expected: "good MORNING",
		},
		{
			name:     "Capitalize command",
			input:    "this is fun (cap)",
			expected: "this is Fun",
		},
		{
			name:     "Capitalize multiple words",
			input:    "capitalize these words (cap, 2)",
			expected: "capitalize These Words",
		},
		{
			name:     "Uppercase multiple words",
			input:    "This is so exciting (up, 2)",
			expected: "This is SO EXCITING",
		},

		// --- Number Conversion ---
		{
			name:     "Hex to decimal",
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
		{
			name:     "Binary to decimal",
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},
		{
			name:     "Mixed number conversions",
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},

		// --- Article Correction ---
		{
			name:     "A to An before vowel",
			input:    "a apple",
			expected: "an apple",
		},
		{
			name:     "A to An before h",
			input:    "a hours",
			expected: "an hours",
		},
		{
			name:     "Capital A to An",
			input:    "A apple",
			expected: "An apple",
		},
		{
			name:     "In quotes",
			input:    "'a apple'",
			expected: "'an apple'",
		},
		{
			name:     "Multiple corrections",
			input:    "'a apple A orange a umbrella'",
			expected: "'an apple An orange an umbrella'",
		},

		// --- Punctuation Fixes ---
		{
			name:     "Comma spacing",
			input:    "Hello , world",
			expected: "Hello, world",
		},
		{
			name:     "Multiple punctuation",
			input:    "Wait  ! Are you ok ? ",
			expected: "Wait! Are you ok?",
		},
		{
			name:     "Ellipsis preservation",
			input:    "I was thinking ... You were right",
			expected: "I was thinking... You were right",
		},
		{
			name:     "Multiple issues",
			input:    "Hi ,there! How  are you?",
			expected: "Hi, there! How are you?",
		},

		// --- Quotes Handling ---
		{
			name:     "Single word quotes",
			input:    "I am exactly how they describe me: ' awesome '",
			expected: "I am exactly how they describe me: 'awesome'",
		},
		{
			name:     "Multiple word quotes",
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			name:     "Multiple quoted words",
			input:    "'foo'   'bar'",
			expected: "'foo' 'bar'",
		},
		{
			name:     "Mixed quotes",
			input:    `" ' world ' "`,
			expected: `"'world'"`,
		},
		{
			name:     "Complex mixed quotes",
			input:    `they're '"        '   " 'hi     '`,
			expected: `they're '"'" 'hi'`,
		},

		// --- Edge Cases ---
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Whitespace only",
			input:    "     ",
			expected: "",
		},
		{
			name:     "Command with punctuation",
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO!",
		},

		// --- Complex Examples ---
		{
			name:     "Complex transformation sequence",
			input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			name:     "contraction it",
			input:    `it'   s ok`,
			expected: `it's ok`,
		},
		{
			name:     "contraction don't",
			input:    `' don't '`,
			expected: `'don't'`,
		},
		{
			name:     "mixed quotes",
			input:    `they're '"        '   " 'hi     '`,
			expected: `they're '"'" 'hi'`,
		},
		{
			name:     "multiple quoted words",
			input:    `'foo'   'bar'`,
			expected: `'foo' 'bar'`,
		},
		{
			name:     "quoted word with double quotes",
			input:    `" ' world ' "`,
			expected: `"'world'"`,
		},
		// // --- Question cases  ---
		{
			name:     "Question Command",
			input:    "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			expected: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			name:     "Question command",
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:     "Question command",
			input:    "There is no greater agony than bearing a untold story inside you.",
			expected: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:     "Question command",
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			name:     "Question command",
			input:    "There it was. A amazing rock!",
			expected: "There it was. An amazing rock!",
		},
		{
			name:     "Question command",
			input:    "I was thinking ... You were right",
			expected: "I was thinking... You were right",
		},
		{
			name:     "Question command",
			input:    "This is so exciting (up, 2)",
			expected: "This is SO EXCITING",
		},
		{
			name:     "Question command",
			input:    "Welcome to the Brooklyn bridge (cap)",
			expected: "Welcome to the Brooklyn Bridge",
		},
		{
			name:     "Question command",
			input:    "I should stop SHOUTING (low)",
			expected: "I should stop shouting",
		},
		{
			name:     "Question command",
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO!",
		},
		{
			name:     "Question command",
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},
		{
			name:     "Question command",
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO!",
		},
		{
			name:     "Question command",
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
		{
			name:     "Question command",
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO!",
		},
		{
			name:     "Question command",
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO!",
		},

		// --- Commands ---
		{
			name:     "Lowercase command",
			input:    "hello WORLD (low)",
			expected: "hello world",
		},
		{
			name:     "Uppercase command",
			input:    "good morning (up)",
			expected: "good MORNING",
		},
		{
			name:     "Capitalize command",
			input:    "this is fun (cap)",
			expected: "this is Fun",
		},
		{
			name:     "Capitalize 2 words",
			input:    "capitalize these words (cap, 2)",
			expected: "capitalize These Words",
		},

		// --- Vowel fixes ---
		{
			name:     "Simple a -> an",
			input:    "a apple",
			expected: "an apple",
		},
		{
			name:     "Capitalized A",
			input:    "A apple",
			expected: "An apple",
		},
		{
			name:     "Between quotes",
			input:    "'a apple'",
			expected: "'an apple'",
		}, {
			name:     "Multiple between quotes",
			input:    "'a apple A orange a umbrella'",
			expected: "'an apple An orange an umbrella'",
		},

		// --- Punctuation fixes ---
		{
			name:     "Comma spacing",
			input:    "Hello , world",
			expected: "Hello, world",
		},
		{
			name:     "Multiple spaces and punctuation",
			input:    "Wait  ! Are you ok ? ",
			expected: "Wait! Are you ok?",
		},
		{
			name:     "Ellipsis preserved",
			input:    "This is ... really cool",
			expected: "This is... really cool",
		},
		{
			name:     "Multiple issues",
			input:    "Hi ,there! How  are you?",
			expected: "Hi, there! How are you?",
		},

		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Whitespace only",
			input:    "     ",
			expected: "",
		},
		{
			name:     "Multi-Command",
			input:    `they're '"        '   " 'hi     '`,
			expected: `they're '"'" 'hi'`,
		},
		{
			name:     "Multi-Command",
			input:    `a hours`,
			expected: `an hours`,
		},
		{
			name:     "Single Quote Adjusment",
			input:    "' apple                           ' g '",
			expected: "'apple' g'",
		},
		{
			name:     "Multi-Command",
			input:    "(a apple) hussain (up)",
			expected: "(an apple) HUSSAIN",
		},
		{
			name:     "Nested-Command",
			input:    "test test (cap, 10 (bin) )",
			expected: "Test Test",
		},
		{
			name:     "Nested-Command",
			input:    "TEST ((((((((up))))))))",
			expected: "TEST",
		},
		{
			name:     "Nested-Command",
			input:    "1E (((hex)))",
			expected: "30",
		},
		{
			name:     "Nested-Command",
			input:    "test test ((((up,   2   ))))",
			expected: "TEST TEST",
		},
		// --- Quotes handling ---
		{
			name:     "Single Word in Quotes",
			input:    "I am exactly how they describe me: ' awesome '",
			expected: "I am exactly how they describe me: 'awesome'",
		},
		{
			name:     "Multiple Words in Quotes",
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},

		// --- Article correction ---
		{
			name:     "A to An before Vowel",
			input:    "There it was. A amazing rock!",
			expected: "There it was. An amazing rock!",
		},
		{
			name:     "A to An before H",
			input:    "He waited for a hour in the cold.",
			expected: "He waited for an hour in the cold.",
		},
		{
			name:     "Leading and Trailing Quotes",
			input:    `   '   hello   '   `,
			expected: `'hello'`,
		},
		{
			name:     "Mixed Double and Single Quotes",
			input:    ` " ' world '  " `,
			expected: `"'world'"`,
		},
		{
			name:     "Nested Quotes with Spaces",
			input:    `  '"   test   "'  `,
			expected: `'"test"'`,
		},
		{
			name:     "Multiple Words with Quotes",
			input:    ` 'foo'   'bar'   `,
			expected: `'foo' 'bar'`,
		},
		{
			name:     "Quote Inside Word",
			input:    ` it'   s    ok `,
			expected: `it's ok`,
		},
		{
			name:     "Symbols with Quotes",
			input:    `  '   hello!!!   ' `,
			expected: `'hello!!!'`,
		},
		{
			name:     "vowels with Quotes",
			input:    `'a' apple`,
			expected: `'an' apple`,
		},
		{
			name:     "vowels with Quotes",
			input:    `'a, apple`,
			expected: `'an, apple`,
		},
		{
			name: "command in multiple lines",
			input: `hello wo(rld (cap,







2)`,
			expected: `Hello Wo(rld`,
		},
		// --- Contractions ---
		{
			name:     "Contraction: don't",
			input:    `' don't '`,
			expected: `'don't'`,
		},

		// --- Quotes ---
		{
			name:     "Unmatched single quote",
			input:    "It is 'cool",
			expected: "It is 'cool",
		},
		{
			name:     "Unmatched double quote",
			input:    `He said "hello`,
			expected: `He said "hello`,
		},
		{
			name:     "Nested single quotes",
			input:    "He said: 'She said 'hello''",
			expected: "He said: 'She said'hello''",
		},

		// --- Malformed Commands ---
		// {
		// 	name:     "Invalid command syntax",
		// 	input:    "hello (up, ) world",
		// 	expected: "hello (up, ) world", // error message
		// },
		{
			name:     "Unknown command",
			input:    "test (foo)",
			expected: "test (foo)",
		},

		// --- Word Missing Before Command ---
		// {
		// 	name:     "No word before command",
		// 	input:    "(up) hello",
		// 	expected: "(up) hello", // error message
		// },

		// --- Invalid Numbers ---
		// {
		// 	name:     "Invalid number in command",
		// 	input:    "hello (up, x)",
		// 	expected: "hello (up, x)", // error message
		// },
		// {
		// 	name:     "Negative number in command",
		// 	input:    "hello (low, -2)",
		// 	expected: "hello (low, -2)", // error message
		// },

		// --- Hex/Bin Conversion Errors ---
		// {
		// 	name:     "Invalid hex value",
		// 	input:    "G1 (hex) files",
		// 	expected: "G1 (hex) files", // error message
		// },
		// {
		// 	name:     "Invalid binary value",
		// 	input:    "102 (bin) years",
		// 	expected: "102 (bin) years", // error message
		// },

		// --- Punctuation Formatting ---
		{
			name:     "Punctuation without space",
			input:    "Hello ,world!",
			expected: "Hello, world!",
		},
		{
			name:     "Multiple punctuation",
			input:    "Wait !!",
			expected: "Wait!!",
		},
		{
			name:     "Ellipsis",
			input:    "I was thinking ... you know?",
			expected: "I was thinking... you know?",
		},

		// --- Article Adjustment (a → an) ---
		{
			name:     "Basic a→an conversion",
			input:    "a apple",
			expected: "an apple",
		},
		{
			name:     "Capitalized A→An conversion",
			input:    "A orange",
			expected: "An orange",
		},
		{
			name:     "Word starting with consonant",
			input:    "a cat",
			expected: "a cat",
		},
		{
			name:     "Word starting with h",
			input:    "a house",
			expected: "an house",
		},

		// --- Empty or Whitespace Input ---
		{
			name:     "Empty input",
			input:    "",
			expected: "",
		},
		// {
		// 	name:     "Whitespace only input",
		// 	input:    "   ",
		// 	expected: "   ", // error message
		// },

		// --- Extreme Inputs ---
		{
			name:     "Only punctuation",
			input:    "!!!",
			expected: "!!!",
		},
		// {
		// 	name:     "Very long input",
		// 	input:    strings.Repeat("word ", 10000),
		// 	expected: strings.Repeat("word ", 10000), // error message
		// },
		{
			name:     "Mixed pipelines of transformations",
			input:    "' 1e (hex) ' files and ' 10 (bin) ' years (up,2) !?",
			expected: "'30' files and '2' YEARS!?",
		},
		{
			name:     "Mixed pipelines of transformations",
			input:    "hasan ali: (cap, 2) ' hello world one two three. '",
			expected: "Hasan Ali: 'hello world one two three.'",
		},
		// {
		// 	name:     "Empty input",
		// 	input:    "",
		// 	expected: "", // should stay empty, or maybe error if you want strictness
		// },
		// {
		// 	name:     "Whitespace only",
		// 	input:    "     ",
		// 	expected: "", // depending on your rules
		// },
		{
			name:     "Unbalanced parentheses",
			input:    "hello (cap",
			expected: "hello (cap", // should remain unchanged or error
		},
		{
			name:     "Empty command",
			input:    "hello () world",
			expected: "hello () world", // no-op, just preserved
		},
		{
			name:     "Invalid command name",
			input:    "hello (foo)",
			expected: "hello (foo)", // unchanged
		},
		// {
		// 	name:     "Valid command but nonsense count",
		// 	input:    "hello (cap, 0)",
		// 	expected: "hello", // or unchanged, depending on your logic
		// },
		// {
		// 	name:     "Nested commands",
		// 	input:    "hello (up (hex))",
		// 	expected: "hello (up (hex))", // if nesting not supported
		// },
		// {
		// 	name:     "Multiple commands inline",
		// 	input:    "hello (up) world (cap, 2) end (hex)",
		// 	expected: "HELLO World end 656e64", // err message
		// },
		{
			name:     "Only parentheses",
			input:    "()",
			expected: "()", // should be preserved
		},
		// {
		// 	name:     "Non-ASCII characters",
		// 	input:    "héllo (up)",
		// 	expected: "héllo (up)", // should fail or stay unchanged if ASCII enforced
		// },
		// {
		// 	name:     "Large input near limit",
		// 	input:    strings.Repeat("a", 999),
		// 	expected: strings.Repeat("a", 999),
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := commands.ProcessCommands(tt.input)
			if got != tt.expected {
				t.Errorf("Got: %s, Expected.%s", got, tt.expected)
			}
		})
	}
}
