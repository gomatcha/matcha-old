# Text

## Mochi

Option 1.
Each attribute gets a type. And Format.Attributes() returns a slice of attributes that can be type casted. Downside is that it dirtys the package namespace, with a bunch of types.

marker, sentinel, flag, dummy

## Considerations

* Right-to-left text
* Animating font size and color
* Animating individual characters
* Links
* FontTraits (https://developer.apple.com/fonts/TrueType-Reference-Manual/RM09/AppendixF.html#Type34)
* Text spannig across multiple frames
* https://developer.apple.com/reference/foundation/nsattributedstring/character_attributes
https://developer.apple.com/reference/foundation/nsattributedstring/document_attributes
* Tokenizing
* Inline views
* UTF8 normalization
* 

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
## Data structure
Rope
GapBuffer
Zipper
PieceTable Piece Chain
https://www.cs.unm.edu/~crowley/papers/sds.pdf

https://blog.golang.org/normalization

# iOS Fonts

Copperplate
    Copperplate-Light
    Copperplate
    Copperplate-Bold
Heiti SC
Kohinoor Telugu
    KohinoorTelugu-Regular
    KohinoorTelugu-Medium
    KohinoorTelugu-Light
Thonburi
    Thonburi
    Thonburi-Bold
    Thonburi-Light
Heiti TC
Courier New
    CourierNewPS-BoldMT
    CourierNewPS-ItalicMT
    CourierNewPSMT
    CourierNewPS-BoldItalicMT
Gill Sans
    GillSans-Italic
    GillSans-Bold
    GillSans-BoldItalic
    GillSans-LightItalic
    GillSans
    GillSans-Light
    GillSans-SemiBold
    GillSans-SemiBoldItalic
    GillSans-UltraBold
Apple SD Gothic Neo
    AppleSDGothicNeo-Bold
    AppleSDGothicNeo-UltraLight
    AppleSDGothicNeo-Thin
    AppleSDGothicNeo-Regular
    AppleSDGothicNeo-Light
    AppleSDGothicNeo-Medium
    AppleSDGothicNeo-SemiBold
Marker Felt
    MarkerFelt-Thin
    MarkerFelt-Wide
Avenir Next Condensed
    AvenirNextCondensed-BoldItalic
    AvenirNextCondensed-Heavy
    AvenirNextCondensed-Medium
    AvenirNextCondensed-Regular
    AvenirNextCondensed-HeavyItalic
    AvenirNextCondensed-MediumItalic
    AvenirNextCondensed-Italic
    AvenirNextCondensed-UltraLightItalic
    AvenirNextCondensed-UltraLight
    AvenirNextCondensed-DemiBold
    AvenirNextCondensed-Bold
    AvenirNextCondensed-DemiBoldItalic
Tamil Sangam MN
    TamilSangamMN
    TamilSangamMN-Bold
Helvetica Neue
    HelveticaNeue-Italic
    HelveticaNeue-Bold
    HelveticaNeue-UltraLight
    HelveticaNeue-CondensedBlack
    HelveticaNeue-BoldItalic
    HelveticaNeue-CondensedBold
    HelveticaNeue-Medium
    HelveticaNeue-Light
    HelveticaNeue-Thin
    HelveticaNeue-ThinItalic
    HelveticaNeue-LightItalic
    HelveticaNeue-UltraLightItalic
    HelveticaNeue-MediumItalic
    HelveticaNeue
Gurmukhi MN
    GurmukhiMN-Bold
    GurmukhiMN
Times New Roman
    TimesNewRomanPSMT
    TimesNewRomanPS-BoldItalicMT
    TimesNewRomanPS-ItalicMT
    TimesNewRomanPS-BoldMT
Georgia
    Georgia-BoldItalic
    Georgia
    Georgia-Italic
    Georgia-Bold
Apple Color Emoji
    AppleColorEmoji
Arial Rounded MT Bold
    ArialRoundedMTBold
Kailasa
    Kailasa-Bold
    Kailasa
Kohinoor Devanagari
    KohinoorDevanagari-Light
    KohinoorDevanagari-Regular
    KohinoorDevanagari-Semibold
Kohinoor Bangla
    KohinoorBangla-Semibold
    KohinoorBangla-Regular
    KohinoorBangla-Light
Chalkboard SE
    ChalkboardSE-Bold
    ChalkboardSE-Light
    ChalkboardSE-Regular
Sinhala Sangam MN
    SinhalaSangamMN-Bold
    SinhalaSangamMN
PingFang TC
    PingFangTC-Medium
    PingFangTC-Regular
    PingFangTC-Light
    PingFangTC-Ultralight
    PingFangTC-Semibold
    PingFangTC-Thin
Gujarati Sangam MN
    GujaratiSangamMN-Bold
    GujaratiSangamMN
Damascus
    DamascusLight
    DamascusBold
    DamascusSemiBold
    DamascusMedium
    Damascus
Noteworthy
    Noteworthy-Light
    Noteworthy-Bold
Geeza Pro
    GeezaPro
    GeezaPro-Bold
Avenir
    Avenir-Medium
    Avenir-HeavyOblique
    Avenir-Book
    Avenir-Light
    Avenir-Roman
    Avenir-BookOblique
    Avenir-MediumOblique
    Avenir-Black
    Avenir-BlackOblique
    Avenir-Heavy
    Avenir-LightOblique
    Avenir-Oblique
Academy Engraved LET
    AcademyEngravedLetPlain
Mishafi
    DiwanMishafi
Futura
    Futura-CondensedMedium
    Futura-CondensedExtraBold
    Futura-Medium
    Futura-MediumItalic
    Futura-Bold
Farah
    Farah
Kannada Sangam MN
    KannadaSangamMN
    KannadaSangamMN-Bold
Arial Hebrew
    ArialHebrew-Bold
    ArialHebrew-Light
    ArialHebrew
Arial
    ArialMT
    Arial-BoldItalicMT
    Arial-BoldMT
    Arial-ItalicMT
Party LET
    PartyLetPlain
Chalkduster
    Chalkduster
Hoefler Text
    HoeflerText-Italic
    HoeflerText-Regular
    HoeflerText-Black
    HoeflerText-BlackItalic
Optima
    Optima-Regular
    Optima-ExtraBlack
    Optima-BoldItalic
    Optima-Italic
    Optima-Bold
Palatino
    Palatino-Bold
    Palatino-Roman
    Palatino-BoldItalic
    Palatino-Italic
Lao Sangam MN
    LaoSangamMN
Malayalam Sangam MN
    MalayalamSangamMN-Bold
    MalayalamSangamMN
Al Nile
    AlNile-Bold
    AlNile
Bradley Hand
    BradleyHandITCTT-Bold
PingFang HK
    PingFangHK-Ultralight
    PingFangHK-Semibold
    PingFangHK-Thin
    PingFangHK-Light
    PingFangHK-Regular
    PingFangHK-Medium
Trebuchet MS
    Trebuchet-BoldItalic
    TrebuchetMS
    TrebuchetMS-Bold
    TrebuchetMS-Italic
Helvetica
    Helvetica-Bold
    Helvetica
    Helvetica-LightOblique
    Helvetica-Oblique
    Helvetica-BoldOblique
    Helvetica-Light
Courier
    Courier-BoldOblique
    Courier
    Courier-Bold
    Courier-Oblique
Cochin
    Cochin-Bold
    Cochin
    Cochin-Italic
    Cochin-BoldItalic
Hiragino Mincho ProN
    HiraMinProN-W6
    HiraMinProN-W3
Devanagari Sangam MN
    DevanagariSangamMN
    DevanagariSangamMN-Bold
Oriya Sangam MN
    OriyaSangamMN
    OriyaSangamMN-Bold
Snell Roundhand
    SnellRoundhand-Bold
    SnellRoundhand
    SnellRoundhand-Black
Zapf Dingbats
    ZapfDingbatsITC
Bodoni 72
    BodoniSvtyTwoITCTT-Bold
    BodoniSvtyTwoITCTT-Book
    BodoniSvtyTwoITCTT-BookIta
Verdana
    Verdana-Italic
    Verdana-BoldItalic
    Verdana
    Verdana-Bold
American Typewriter
    AmericanTypewriter-CondensedLight
    AmericanTypewriter
    AmericanTypewriter-CondensedBold
    AmericanTypewriter-Light
    AmericanTypewriter-Semibold
    AmericanTypewriter-Bold
    AmericanTypewriter-Condensed
Avenir Next
    AvenirNext-UltraLight
    AvenirNext-UltraLightItalic
    AvenirNext-Bold
    AvenirNext-BoldItalic
    AvenirNext-DemiBold
    AvenirNext-DemiBoldItalic
    AvenirNext-Medium
    AvenirNext-HeavyItalic
    AvenirNext-Heavy
    AvenirNext-Italic
    AvenirNext-Regular
    AvenirNext-MediumItalic
Baskerville
    Baskerville-Italic
    Baskerville-SemiBold
    Baskerville-BoldItalic
    Baskerville-SemiBoldItalic
    Baskerville-Bold
    Baskerville
Khmer Sangam MN
    KhmerSangamMN
Didot
    Didot-Italic
    Didot-Bold
    Didot
Savoye LET
    SavoyeLetPlain
Bodoni Ornaments
    BodoniOrnamentsITCTT
Symbol
    Symbol
Menlo
    Menlo-Italic
    Menlo-Bold
    Menlo-Regular
    Menlo-BoldItalic
Bodoni 72 Smallcaps
    BodoniSvtyTwoSCITCTT-Book
Papyrus
    Papyrus
    Papyrus-Condensed
Hiragino Sans
    HiraginoSans-W3
    HiraginoSans-W6
PingFang SC
    PingFangSC-Ultralight
    PingFangSC-Regular
    PingFangSC-Semibold
    PingFangSC-Thin
    PingFangSC-Light
    PingFangSC-Medium
Myanmar Sangam MN
    MyanmarSangamMN-Bold
    MyanmarSangamMN
Euphemia UCAS
    EuphemiaUCAS-Italic
    EuphemiaUCAS
    EuphemiaUCAS-Bold
Telugu Sangam MN
Bangla Sangam MN
Zapfino
    Zapfino
Bodoni 72 Oldstyle
    BodoniSvtyTwoOSITCTT-Book
    BodoniSvtyTwoOSITCTT-Bold
    BodoniSvtyTwoOSITCTT-BookIt
