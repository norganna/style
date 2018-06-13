// Package style uses strings with "Style tags" embedded in them to generate a styled output.
package style

// A Style tag comprises a string containing:
//    "‹" A single left-pointing angle quotation mark (U+2039) [alt-shift-3 on OS X]
//    The Style word
//    ":" A colon
//    The Style parameter
//    "›" A single right-pointing angle quotation mark (U+203A) [alt-shift-4 on OS X]
//
// These quote characters were chosen because they're relatively uncommonly used in normal text
// and they are easily typed in on the keyboard. However, if you don't want to use the defaults,
// you can change them by using a custom Style config:
//
//    sty := Style.Box()
//    sty.TagSequence("::", "::")
//    sty.Println("Hey ::b:there!::")
//
// A definitive list of the various Style words are contained in the Style() function.
//
// Examples:
//
//    ‹hl:10›    Horizontal line coloured in line colour and consisting of 10 characters.
//    ‹hl:10:>text ›
//               10 char horizontal line coloured in line colour with embedded "text" right aligned.
//               The '>' at the start right aligns it ('|' to center) and the space at the end pushes text back 1 char.
//               e.g.:  "-----text-"
//    ‹hc:Text›  The word "Text" coloured in the header colour.
//    ‹lc:^`\,_,/`^`\,_,/`^`\,_,/`^›
//               A wavy ASCII art line coloured in line colour.
//    ‹li›Text   A line item coloured in the line colour with Text beside it.
//    ‹ll›Text   A last line item coloured in the line colour with Text beside it.
//    ‹li:text›  A line item with text both coloured in the line color.
//    ‹sp:5›     Five characters of space padding.
//    ‹b:Text›   The word "Text" coloured in the bold colour.
//    ‹i:Text›   The word "Text" coloured in the italic colour.
//    ‹e:Text›   The word "Text" coloured in the error colour.
//
// There are also several formatting functions included, which can be combined with Style tags.
//
//    Style.Printf("‹hl:%d› ‹%s:%s›", 3, "b", "Text")
