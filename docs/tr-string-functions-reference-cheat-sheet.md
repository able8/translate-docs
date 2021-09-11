# Go string handling overview [cheat sheet]

  yourbasic.org/golang

  - [String literals (escape characters)](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/#string-literals-escape-characters)
  - [Concatenate](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/#concatenate)
  - [Equal and compare (ignore case)](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/#equal-and-compare-ignore-case)
  - [Length in bytes or runes](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/#length-in-bytes-or-runes)
  - [Index, substring, iterate](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/#index-substring-iterate)
  - [Search (contains, prefix/suffix, index)](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/#search-contains-prefix-suffix-index)
  - [Replace (uppercase/lowercase, trim)](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/#replace-uppercase-lowercase-trim)
  - [Split by space or comma](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/#split-by-space-or-comma)
  - [Join strings with separator](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/#join-strings-with-separator)
  - [Format and convert](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/#format-and-convert)
  - [Regular expressions](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/#regular-expressions)

  ## String literals (escape characters)

  | Expression                | Result     | Note                                                         |
  | ------------------------- | ---------- | ------------------------------------------------------------ |
  | `""`                      |            | [Default zero value](https://yourbasic.org/golang/default-zero-value/) for type `string` |
  | `"Japan 日本"`            | Japan 日本 | Go code is [Unicode text encoded in UTF‑8](https://yourbasic.org/golang/rune/) |
  | `"\xe6\x97\xa5"`          | 日         | `\xNN` specifies a byte                                      |
  | `"\u65E5"`                | 日         | `\uNNNN` specifies a Unicode value                           |
  | `"\\"`                    | \          | Backslash                                                    |
  | `"\""`                    | "          | Double quote                                                 |
  | `"\n"`                    |            | Newline                                                      |
  | `"\t"`                    |            | Tab                                                          |
  | ``\xe6``                  | \xe6       | Raw string literal*                                          |
  | `html.EscapeString("<>")` | &lt;&gt;   | HTML escape for <, >, &, ' and "                             |
  | `url.PathEscape("A B")`   | A%20B      | URL percent-encoding net/url                                 |

  \* In ```` string literals, text is interpreted literally and backslashes have no special meaning. See [Escapes and multiline strings](https://yourbasic.org/golang/multiline-string/) for more on raw strings, escape characters and string encodings.

  ## Concatenate

  | Expression     | Result | Note          |
  | -------------- | ------ | ------------- |
  | `"Ja" + "pan"` | Japan  | Concatenation |

  > **Performance tips**
  >  See [3 tips for efficient string concatenation](https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/) for how to best use a string builder to concatenate strings without redundant copying.

  ## Equal and compare (ignore case)

  | Expression                            | Result | Note                 |
  | ------------------------------------- | ------ | -------------------- |
  | `"Japan" == "Japan"`                  | true   | Equality             |
  | `strings.EqualFold("Japan", "JAPAN")` | true   | Unicode case folding |
  | `"Japan" < "japan"`                   | true   | Lexicographic order  |

  ## Length in bytes or runes

  | Expression                     | Result | Note                  |
  | ------------------------------ | ------ | --------------------- |
  | `len("日")`                    | 3      | Length in bytes       |
  | `utf8.RuneCountInString("日")` | 1      | in runes unicode/utf8 |
  | `utf8.ValidString("日")`       | true   | UTF-8? unicode/utf8   |

  ## Index, substring, iterate

  | Expression     | Result | Note               |
  | -------------- | ------ | ------------------ |
  | `"Japan"[2]`   | 'p'    | Byte at position 2 |
  | `"Japan"[1:3]` | ap     | Byte indexing      |
  | `"Japan"[:2]`  | Ja     |                    |
  | `"Japan"[2:]`  | pan    |                    |

  A Go [range loop](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) iterates over UTF-8 encoded characters ([runes](https://yourbasic.org/golang/rune/)):

  ```
  for i, ch := range "Japan 日本" {
      fmt.Printf("%d:%q ", i, ch)
  }
  // Output: 0:'J' 1:'a' 2:'p' 3:'a' 4:'n' 5:' ' 6:'日' 9:'本'
  ```

  Iterating over bytes produces nonsense characters for non-ASCII text:

  ```
  s := "Japan 日本"
  for i := 0; i < len(s); i++ {
      fmt.Printf("%q ", s[i])
  }
  // Output: 'J' 'a' 'p' 'a' 'n' ' ' 'æ' '\u0097' '¥' 'æ' '\u009c' '¬'
  ```

  ## Search (contains, prefix/suffix, index)

  | Expression                             | Result | Note                             |
  | -------------------------------------- | ------ | -------------------------------- |
  | `strings.Contains("Japan", "abc")`     | false  | Is abc in Japan?                 |
  | `strings.ContainsAny("Japan", "abc")`  | true   | Is a, b or c in Japan?           |
  | `strings.Count("Banana", "ana")`       | 1      | Non-overlapping instances of ana |
  | `strings.HasPrefix("Japan", "Ja")`     | true   | Does Japan start with Ja?        |
  | `strings.HasSuffix("Japan", "pan")`    | true   | Does Japan end with pan?         |
  | `strings.Index("Japan", "abc")`        | -1     | Index of first abc               |
  | `strings.IndexAny("Japan", "abc")`     | 1      | a, b or c                        |
  | `strings.LastIndex("Japan", "abc")`    | -1     | Index of last abc                |
  | `strings.LastIndexAny("Japan", "abc")` | 3      | a, b or c                        |

  ## Replace (uppercase/lowercase, trim)

  | Expression                                                   | Result | Note                                                 |
  | ------------------------------------------------------------ | ------ | ---------------------------------------------------- |
  | `strings.Replace("foo", "o", ".", 2)`                        | f..    | Replace first two “o” with “.” Use -1 to replace all |
  | `f := func(r rune) rune {  return r + 1}strings.Map(f, "ab")` | bc     | Apply function to each character                     |
  | `strings.ToUpper("Japan")`                                   | JAPAN  | Uppercase                                            |
  | `strings.ToLower("Japan")`                                   | japan  | Lowercase                                            |
  | `strings.Title("ja pan")`                                    | Ja Pan | Initial letters to uppercase                         |
  | `strings.TrimSpace(" foo\n")`                                | foo    | Strip leading and trailing white space               |
  | `strings.Trim("foo", "fo")`                                  |        | Strip *leading and trailing* f:s and o:s             |
  | `strings.TrimLeft("foo", "f")`                               | oo     | *only leading*                                       |
  | `strings.TrimRight("foo", "o")`                              | f      | *only trailing*                                      |
  | `strings.TrimPrefix("foo", "fo")`                            | o      |                                                      |
  | `strings.TrimSuffix("foo", "o")`                             | fo     |                                                      |

  ## Split by space or comma

  | Expression                       | Result       | Note               |
  | -------------------------------- | ------------ | ------------------ |
  | `strings.Fields(" a\t b\n")`     | `["a" "b"]`  | Remove white space |
  | `strings.Split("a,b", ",")`      | `["a" "b"]`  | Remove separator   |
  | `strings.SplitAfter("a,b", ",")` | `["a," "b"]` | Keep separator     |

  ## Join strings with separator

  | Expression                              | Result | Note             |
  | --------------------------------------- | ------ | ---------------- |
  | `strings.Join([]string{"a", "b"}, ":")` | a:b    | Add separator    |
  | `strings.Repeat("da", 2)`               | dada   | 2 copies of “da” |

  ## Format and convert

  | Expression                   | Result  | Note          |
  | ---------------------------- | ------- | ------------- |
  | `strconv.Itoa(-42)`          | `"-42"` | Int to string |
  | `strconv.FormatInt(255, 16)` | `"ff"`  | Base 16       |

  ### Sprintf

  The [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) function is often your best friend when formatting data:

  ```
  s := fmt.Sprintf("%.4f", math.Pi) // s == "3.1416"
  ```

  This [fmt cheat sheet](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) covers the most common formatting flags.

  ## Regular expressions


  For more advanced string handling, see  this [Regular expressions tutorial](https://yourbasic.org/golang/regexp-cheat-sheet/), a gentle introduction to the `regexp` package with cheat sheet and plenty of examples.


  ## Related

  [Runes and character encoding](https://yourbasic.org/golang/rune/)

  A rune is a type meant to represent a Unicode code point. Strings, however, are sequences of bytes (typically containing  Unicode text encoded in UTF-8). 