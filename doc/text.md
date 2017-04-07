# Text

## Mochi

Option 1.
Each attribute gets a type. And Format.Attributes() returns a slice of attributes that can be type casted. Downside is that it dirtys the package namespace, with a bunch of types.

## Considerations

* Right-to-left text
* Animating font size and color
* Animating individual characters
* Links
* FontTraits (https://developer.apple.com/fonts/TrueType-Reference-Manual/RM09/AppendixF.html#Type34)
* Text spannig across multiple frames
* https://developer.apple.com/reference/foundation/nsattributedstring/character_attributes
https://developer.apple.com/reference/foundation/nsattributedstring/document_attributes

## Flutter

RichText
* TextSpan
    * String
    * TextStyle
    * Children
* MaxLines
* SoftWrap
* TextScaleFactor // font size multiplier?
* TextAlign

Text
* String
* MaxLines
* Overflow
    * clip
    * ellipses
    * fade
* SoftWrap
* TextScaleFactor
* TextAlign
    * center
    * justify
    * left
    * right
* TextStyle
    * Color
    * DecorationColor
    * Decoration
        * linethrough
        * none
        * overline
        * underline
    * DecorationStyle
        * dashed
        * dotted
        * double
        * solid
        * wavy
    * FontFamily
    * FontSize
    * FontStyle
    * FontWeight
    * Height // LineHeight
    * LetterSpacing
    * TextBaseline
    * WordSpacing

ParagraphStyle (Used for drawing)
* TextAlign
* FontWeight
* FontStyle
* MaxLines
* FontFamily
* FontSize
* LineHeight
* Ellipsis (String)?

## ReactNative

## CoreText

## UIKit

NSAttributedString
* NSTextAttachment (Inserting Images into text)
* BackgroundColor
* BaselineOffset
* Font
    * Size
    * Style
    * Family
* ForegroundColor
* Kerning
* Ligature
* Links
* ParagraphStyle
    * Alignment
    * HeadIndent
    * TailIndent
    * LineHeightMultiple // The natural line height of the receiver is multiplied by this factor (if positive) before being constrained by minimum and maximum line height. The default value of this property is 0.0.
    * MinimumLineHeight
    * MaximumLineHeight
    * LineSpacing // Additional spacing between lines
    * ParagraphSpacing
    * ParagraphSpacingBefore
    * LineBreakMode
        * WordWrapping
        * CharWrapping
        * Clipping
        * TruncatingHead
        * TruncatingTail
        * TruncatingMiddle
    * HyphenationFactor
    * WritingDirection
* Superscript
* Underline

## HTML

## Android - Spannable, Editable

## TTTAttributedLabel

## YetiCharacterLabelExample

## Draft.js
## Rope (Data structure)