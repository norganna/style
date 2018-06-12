# Style

Style will style up your output text from within `(s)print(ln)(f)`'s using style tags.

```go
package main

import "github.com/norganna/style"

func main() {
    style.Printlnf("‹bold:%s›", "Hello world")
}
```

## Tags

A Style tag comprises a string containing:

 *  "‹" A single left-pointing angle quotation mark (U+2039) [alt-shift-3 on OS X]
 *  The Style word
 *  ":" A colon
 *  The Style parameter
 *  "›" A single right-pointing angle quotation mark (U+203A) [alt-shift-4 on OS X]

These quote characters were chosen because they're relatively uncommonly used in normal text
and they are easily typed in on the keyboard. However, if you don't want to use the defaults,
you can change them by using a custom Style config:

```go
sty := Style.Box()
sty.TagSequence("::", "::")
sty.Println("Hey ::b:there!::")
```

A definitive list of the various Style words are contained in the Style() function.

Examples:

 *  `‹hl:10›`    Horizontal line coloured in line colour and consisting of 10 characters.
 *  `‹hc:Text›`  The word "Text" coloured in the header colour.
 *  `‹lc:^\,_,/^\,_,/^\,_,/^›`
              A wavy ASCII art line coloured in line colour.
 *  `‹li›Text`   A line item coloured in the line colour with Text beside it.
 *  `‹ll›Text`   A last line item coloured in the line colour with Text beside it.
 *  `‹b:Text›`   The word "Text" coloured in the bold colour.
 *  `‹i:Text›`   The word "Text" coloured in the italic colour.
 *  `‹e:Text›`   The word "Text" coloured in the error colour.

There are also several formatting functions included, which can be combined with Style tags.

```go
Style.Printf("‹hl:%d› ‹%s:%s›", 3, "b", "Text")
 ```