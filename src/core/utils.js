/**
 * Utils
 *
 * Various generic utility functions
 */

import { Translation } from "./translation";

var _t = Translation._t;
var id = -1;

var diacriticsMap = {
  A: "A",
  "\u24B6": "A",
  Ａ: "A",
  À: "A",
  Á: "A",
  Â: "A",
  Ầ: "A",
  Ấ: "A",
  Ẫ: "A",
  Ẩ: "A",
  Ã: "A",
  Ā: "A",
  Ă: "A",
  Ằ: "A",
  Ắ: "A",
  Ẵ: "A",
  Ẳ: "A",
  Ȧ: "A",
  Ǡ: "A",
  Ä: "A",
  Ǟ: "A",
  Ả: "A",
  Å: "A",
  Ǻ: "A",
  Ǎ: "A",
  Ȁ: "A",
  Ȃ: "A",
  Ạ: "A",
  Ậ: "A",
  Ặ: "A",
  Ḁ: "A",
  Ą: "A",
  Ⱥ: "A",
  Ɐ: "A",

  Ꜳ: "AA",
  Æ: "AE",
  Ǽ: "AE",
  Ǣ: "AE",
  Ꜵ: "AO",
  Ꜷ: "AU",
  Ꜹ: "AV",
  Ꜻ: "AV",
  Ꜽ: "AY",
  B: "B",
  "\u24B7": "B",
  Ｂ: "B",
  Ḃ: "B",
  Ḅ: "B",
  Ḇ: "B",
  Ƀ: "B",
  Ƃ: "B",
  Ɓ: "B",

  C: "C",
  "\u24B8": "C",
  Ｃ: "C",
  Ć: "C",
  Ĉ: "C",
  Ċ: "C",
  Č: "C",
  Ç: "C",
  Ḉ: "C",
  Ƈ: "C",
  Ȼ: "C",
  Ꜿ: "C",

  D: "D",
  "\u24B9": "D",
  Ｄ: "D",
  Ḋ: "D",
  Ď: "D",
  Ḍ: "D",
  Ḑ: "D",
  Ḓ: "D",
  Ḏ: "D",
  Đ: "D",
  Ƌ: "D",
  Ɗ: "D",
  Ɖ: "D",
  Ꝺ: "D",

  Ǳ: "DZ",
  Ǆ: "DZ",
  ǲ: "Dz",
  ǅ: "Dz",

  E: "E",
  "\u24BA": "E",
  Ｅ: "E",
  È: "E",
  É: "E",
  Ê: "E",
  Ề: "E",
  Ế: "E",
  Ễ: "E",
  Ể: "E",
  Ẽ: "E",
  Ē: "E",
  Ḕ: "E",
  Ḗ: "E",
  Ĕ: "E",
  Ė: "E",
  Ë: "E",
  Ẻ: "E",
  Ě: "E",
  Ȅ: "E",
  Ȇ: "E",
  Ẹ: "E",
  Ệ: "E",
  Ȩ: "E",
  Ḝ: "E",
  Ę: "E",
  Ḙ: "E",
  Ḛ: "E",
  Ɛ: "E",
  Ǝ: "E",

  F: "F",
  "\u24BB": "F",
  Ｆ: "F",
  Ḟ: "F",
  Ƒ: "F",
  Ꝼ: "F",

  G: "G",
  "\u24BC": "G",
  Ｇ: "G",
  Ǵ: "G",
  Ĝ: "G",
  Ḡ: "G",
  Ğ: "G",
  Ġ: "G",
  Ǧ: "G",
  Ģ: "G",
  Ǥ: "G",
  Ɠ: "G",
  Ꞡ: "G",
  Ᵹ: "G",
  Ꝿ: "G",

  H: "H",
  "\u24BD": "H",
  Ｈ: "H",
  Ĥ: "H",
  Ḣ: "H",
  Ḧ: "H",
  Ȟ: "H",
  Ḥ: "H",
  Ḩ: "H",
  Ḫ: "H",
  Ħ: "H",
  Ⱨ: "H",
  Ⱶ: "H",
  Ɥ: "H",

  I: "I",
  "\u24BE": "I",
  Ｉ: "I",
  Ì: "I",
  Í: "I",
  Î: "I",
  Ĩ: "I",
  Ī: "I",
  Ĭ: "I",
  İ: "I",
  Ï: "I",
  Ḯ: "I",
  Ỉ: "I",
  Ǐ: "I",
  Ȉ: "I",
  Ȋ: "I",
  Ị: "I",
  Į: "I",
  Ḭ: "I",
  Ɨ: "I",

  J: "J",
  "\u24BF": "J",
  Ｊ: "J",
  Ĵ: "J",
  Ɉ: "J",

  K: "K",
  "\u24C0": "K",
  Ｋ: "K",
  Ḱ: "K",
  Ǩ: "K",
  Ḳ: "K",
  Ķ: "K",
  Ḵ: "K",
  Ƙ: "K",
  Ⱪ: "K",
  Ꝁ: "K",
  Ꝃ: "K",
  Ꝅ: "K",
  Ꞣ: "K",

  L: "L",
  "\u24C1": "L",
  Ｌ: "L",
  Ŀ: "L",
  Ĺ: "L",
  Ľ: "L",
  Ḷ: "L",
  Ḹ: "L",
  Ļ: "L",
  Ḽ: "L",
  Ḻ: "L",
  Ł: "L",
  Ƚ: "L",
  Ɫ: "L",
  Ⱡ: "L",
  Ꝉ: "L",
  Ꝇ: "L",
  Ꞁ: "L",

  Ǉ: "LJ",
  ǈ: "Lj",
  M: "M",
  "\u24C2": "M",
  Ｍ: "M",
  Ḿ: "M",
  Ṁ: "M",
  Ṃ: "M",
  Ɱ: "M",
  Ɯ: "M",

  N: "N",
  "\u24C3": "N",
  Ｎ: "N",
  Ǹ: "N",
  Ń: "N",
  Ñ: "N",
  Ṅ: "N",
  Ň: "N",
  Ṇ: "N",
  Ņ: "N",
  Ṋ: "N",
  Ṉ: "N",
  Ƞ: "N",
  Ɲ: "N",
  Ꞑ: "N",
  Ꞥ: "N",

  Ǌ: "NJ",
  ǋ: "Nj",

  O: "O",
  "\u24C4": "O",
  Ｏ: "O",
  Ò: "O",
  Ó: "O",
  Ô: "O",
  Ồ: "O",
  Ố: "O",
  Ỗ: "O",
  Ổ: "O",
  Õ: "O",
  Ṍ: "O",
  Ȭ: "O",
  Ṏ: "O",
  Ō: "O",
  Ṑ: "O",
  Ṓ: "O",
  Ŏ: "O",
  Ȯ: "O",
  Ȱ: "O",
  Ö: "O",
  Ȫ: "O",
  Ỏ: "O",
  Ő: "O",
  Ǒ: "O",
  Ȍ: "O",
  Ȏ: "O",
  Ơ: "O",
  Ờ: "O",
  Ớ: "O",
  Ỡ: "O",
  Ở: "O",
  Ợ: "O",
  Ọ: "O",
  Ộ: "O",
  Ǫ: "O",
  Ǭ: "O",
  Ø: "O",
  Ǿ: "O",
  Ɔ: "O",
  Ɵ: "O",
  Ꝋ: "O",
  Ꝍ: "O",

  Ƣ: "OI",
  Ꝏ: "OO",
  Ȣ: "OU",
  P: "P",
  "\u24C5": "P",
  Ｐ: "P",
  Ṕ: "P",
  Ṗ: "P",
  Ƥ: "P",
  Ᵽ: "P",
  Ꝑ: "P",
  Ꝓ: "P",
  Ꝕ: "P",
  Q: "Q",
  "\u24C6": "Q",
  Ｑ: "Q",
  Ꝗ: "Q",
  Ꝙ: "Q",
  Ɋ: "Q",

  R: "R",
  "\u24C7": "R",
  Ｒ: "R",
  Ŕ: "R",
  Ṙ: "R",
  Ř: "R",
  Ȑ: "R",
  Ȓ: "R",
  Ṛ: "R",
  Ṝ: "R",
  Ŗ: "R",
  Ṟ: "R",
  Ɍ: "R",
  Ɽ: "R",
  Ꝛ: "R",
  Ꞧ: "R",
  Ꞃ: "R",

  S: "S",
  "\u24C8": "S",
  Ｓ: "S",
  ẞ: "S",
  Ś: "S",
  Ṥ: "S",
  Ŝ: "S",
  Ṡ: "S",
  Š: "S",
  Ṧ: "S",
  Ṣ: "S",
  Ṩ: "S",
  Ș: "S",
  Ş: "S",
  Ȿ: "S",
  Ꞩ: "S",
  Ꞅ: "S",

  T: "T",
  "\u24C9": "T",
  Ｔ: "T",
  Ṫ: "T",
  Ť: "T",
  Ṭ: "T",
  Ț: "T",
  Ţ: "T",
  Ṱ: "T",
  Ṯ: "T",
  Ŧ: "T",
  Ƭ: "T",
  Ʈ: "T",
  Ⱦ: "T",
  Ꞇ: "T",

  Ꜩ: "TZ",

  U: "U",
  "\u24CA": "U",
  Ｕ: "U",
  Ù: "U",
  Ú: "U",
  Û: "U",
  Ũ: "U",
  Ṹ: "U",
  Ū: "U",
  Ṻ: "U",
  Ŭ: "U",
  Ü: "U",
  Ǜ: "U",
  Ǘ: "U",
  Ǖ: "U",
  Ǚ: "U",
  Ủ: "U",
  Ů: "U",
  Ű: "U",
  Ǔ: "U",
  Ȕ: "U",
  Ȗ: "U",
  Ư: "U",
  Ừ: "U",
  Ứ: "U",
  Ữ: "U",
  Ử: "U",
  Ự: "U",
  Ụ: "U",
  Ṳ: "U",
  Ų: "U",
  Ṷ: "U",
  Ṵ: "U",
  Ʉ: "U",

  V: "V",
  "\u24CB": "V",
  Ｖ: "V",
  Ṽ: "V",
  Ṿ: "V",
  Ʋ: "V",
  Ꝟ: "V",
  Ʌ: "V",
  Ꝡ: "VY",
  W: "W",
  "\u24CC": "W",
  Ｗ: "W",
  Ẁ: "W",
  Ẃ: "W",
  Ŵ: "W",
  Ẇ: "W",
  Ẅ: "W",
  Ẉ: "W",
  Ⱳ: "W",
  X: "X",
  "\u24CD": "X",
  Ｘ: "X",
  Ẋ: "X",
  Ẍ: "X",

  Y: "Y",
  "\u24CE": "Y",
  Ｙ: "Y",
  Ỳ: "Y",
  Ý: "Y",
  Ŷ: "Y",
  Ỹ: "Y",
  Ȳ: "Y",
  Ẏ: "Y",
  Ÿ: "Y",
  Ỷ: "Y",
  Ỵ: "Y",
  Ƴ: "Y",
  Ɏ: "Y",
  Ỿ: "Y",

  Z: "Z",
  "\u24CF": "Z",
  Ｚ: "Z",
  Ź: "Z",
  Ẑ: "Z",
  Ż: "Z",
  Ž: "Z",
  Ẓ: "Z",
  Ẕ: "Z",
  Ƶ: "Z",
  Ȥ: "Z",
  Ɀ: "Z",
  Ⱬ: "Z",
  Ꝣ: "Z",

  a: "a",
  "\u24D0": "a",
  ａ: "a",
  ẚ: "a",
  à: "a",
  á: "a",
  â: "a",
  ầ: "a",
  ấ: "a",
  ẫ: "a",
  ẩ: "a",
  ã: "a",
  ā: "a",
  ă: "a",
  ằ: "a",
  ắ: "a",
  ẵ: "a",
  ẳ: "a",
  ȧ: "a",
  ǡ: "a",
  ä: "a",
  ǟ: "a",
  ả: "a",
  å: "a",
  ǻ: "a",
  ǎ: "a",
  ȁ: "a",
  ȃ: "a",
  ạ: "a",
  ậ: "a",
  ặ: "a",
  ḁ: "a",
  ą: "a",
  ⱥ: "a",
  ɐ: "a",

  ꜳ: "aa",
  æ: "ae",
  ǽ: "ae",
  ǣ: "ae",
  ꜵ: "ao",
  ꜷ: "au",
  ꜹ: "av",
  ꜻ: "av",
  ꜽ: "ay",
  b: "b",
  "\u24D1": "b",
  ｂ: "b",
  ḃ: "b",
  ḅ: "b",
  ḇ: "b",
  ƀ: "b",
  ƃ: "b",
  ɓ: "b",

  c: "c",
  "\u24D2": "c",
  ｃ: "c",
  ć: "c",
  ĉ: "c",
  ċ: "c",
  č: "c",
  ç: "c",
  ḉ: "c",
  ƈ: "c",
  ȼ: "c",
  ꜿ: "c",
  ↄ: "c",

  d: "d",
  "\u24D3": "d",
  ｄ: "d",
  ḋ: "d",
  ď: "d",
  ḍ: "d",
  ḑ: "d",
  ḓ: "d",
  ḏ: "d",
  đ: "d",
  ƌ: "d",
  ɖ: "d",
  ɗ: "d",
  ꝺ: "d",

  ǳ: "dz",
  ǆ: "dz",

  e: "e",
  "\u24D4": "e",
  ｅ: "e",
  è: "e",
  é: "e",
  ê: "e",
  ề: "e",
  ế: "e",
  ễ: "e",
  ể: "e",
  ẽ: "e",
  ē: "e",
  ḕ: "e",
  ḗ: "e",
  ĕ: "e",
  ė: "e",
  ë: "e",
  ẻ: "e",
  ě: "e",
  ȅ: "e",
  ȇ: "e",
  ẹ: "e",
  ệ: "e",
  ȩ: "e",
  ḝ: "e",
  ę: "e",
  ḙ: "e",
  ḛ: "e",
  ɇ: "e",
  ɛ: "e",
  ǝ: "e",

  f: "f",
  "\u24D5": "f",
  ｆ: "f",
  ḟ: "f",
  ƒ: "f",
  ꝼ: "f",

  g: "g",
  "\u24D6": "g",
  ｇ: "g",
  ǵ: "g",
  ĝ: "g",
  ḡ: "g",
  ğ: "g",
  ġ: "g",
  ǧ: "g",
  ģ: "g",
  ǥ: "g",
  ɠ: "g",
  ꞡ: "g",
  ᵹ: "g",
  ꝿ: "g",

  h: "h",
  "\u24D7": "h",
  ｈ: "h",
  ĥ: "h",
  ḣ: "h",
  ḧ: "h",
  ȟ: "h",
  ḥ: "h",
  ḩ: "h",
  ḫ: "h",
  ẖ: "h",
  ħ: "h",
  ⱨ: "h",
  ⱶ: "h",
  ɥ: "h",

  ƕ: "hv",

  i: "i",
  "\u24D8": "i",
  ｉ: "i",
  ì: "i",
  í: "i",
  î: "i",
  ĩ: "i",
  ī: "i",
  ĭ: "i",
  ï: "i",
  ḯ: "i",
  ỉ: "i",
  ǐ: "i",
  ȉ: "i",
  ȋ: "i",
  ị: "i",
  į: "i",
  ḭ: "i",
  ɨ: "i",
  ı: "i",

  j: "j",
  "\u24D9": "j",
  ｊ: "j",
  ĵ: "j",
  ǰ: "j",
  ɉ: "j",

  k: "k",
  "\u24DA": "k",
  ｋ: "k",
  ḱ: "k",
  ǩ: "k",
  ḳ: "k",
  ķ: "k",
  ḵ: "k",
  ƙ: "k",
  ⱪ: "k",
  ꝁ: "k",
  ꝃ: "k",
  ꝅ: "k",
  ꞣ: "k",

  l: "l",
  "\u24DB": "l",
  ｌ: "l",
  ŀ: "l",
  ĺ: "l",
  ľ: "l",
  ḷ: "l",
  ḹ: "l",
  ļ: "l",
  ḽ: "l",
  ḻ: "l",
  ſ: "l",
  ł: "l",
  ƚ: "l",
  ɫ: "l",
  ⱡ: "l",
  ꝉ: "l",
  ꞁ: "l",
  ꝇ: "l",

  ǉ: "lj",
  m: "m",
  "\u24DC": "m",
  ｍ: "m",
  ḿ: "m",
  ṁ: "m",
  ṃ: "m",
  ɱ: "m",
  ɯ: "m",

  n: "n",
  "\u24DD": "n",
  ｎ: "n",
  ǹ: "n",
  ń: "n",
  ñ: "n",
  ṅ: "n",
  ň: "n",
  ṇ: "n",
  ņ: "n",
  ṋ: "n",
  ṉ: "n",
  ƞ: "n",
  ɲ: "n",
  ŉ: "n",
  ꞑ: "n",
  ꞥ: "n",

  ǌ: "nj",

  o: "o",
  "\u24DE": "o",
  ｏ: "o",
  ò: "o",
  ó: "o",
  ô: "o",
  ồ: "o",
  ố: "o",
  ỗ: "o",
  ổ: "o",
  õ: "o",
  ṍ: "o",
  ȭ: "o",
  ṏ: "o",
  ō: "o",
  ṑ: "o",
  ṓ: "o",
  ŏ: "o",
  ȯ: "o",
  ȱ: "o",
  ö: "o",
  ȫ: "o",
  ỏ: "o",
  ő: "o",
  ǒ: "o",
  ȍ: "o",
  ȏ: "o",
  ơ: "o",
  ờ: "o",
  ớ: "o",
  ỡ: "o",
  ở: "o",
  ợ: "o",
  ọ: "o",
  ộ: "o",
  ǫ: "o",
  ǭ: "o",
  ø: "o",
  ǿ: "o",
  ɔ: "o",
  ꝋ: "o",
  ꝍ: "o",
  ɵ: "o",

  ƣ: "oi",
  ȣ: "ou",
  ꝏ: "oo",
  p: "p",
  "\u24DF": "p",
  ｐ: "p",
  ṕ: "p",
  ṗ: "p",
  ƥ: "p",
  ᵽ: "p",
  ꝑ: "p",
  ꝓ: "p",
  ꝕ: "p",
  q: "q",
  "\u24E0": "q",
  ｑ: "q",
  ɋ: "q",
  ꝗ: "q",
  ꝙ: "q",

  r: "r",
  "\u24E1": "r",
  ｒ: "r",
  ŕ: "r",
  ṙ: "r",
  ř: "r",
  ȑ: "r",
  ȓ: "r",
  ṛ: "r",
  ṝ: "r",
  ŗ: "r",
  ṟ: "r",
  ɍ: "r",
  ɽ: "r",
  ꝛ: "r",
  ꞧ: "r",
  ꞃ: "r",

  s: "s",
  "\u24E2": "s",
  ｓ: "s",
  ß: "s",
  ś: "s",
  ṥ: "s",
  ŝ: "s",
  ṡ: "s",
  š: "s",
  ṧ: "s",
  ṣ: "s",
  ṩ: "s",
  ș: "s",
  ş: "s",
  ȿ: "s",
  ꞩ: "s",
  ꞅ: "s",
  ẛ: "s",

  t: "t",
  "\u24E3": "t",
  ｔ: "t",
  ṫ: "t",
  ẗ: "t",
  ť: "t",
  ṭ: "t",
  ț: "t",
  ţ: "t",
  ṱ: "t",
  ṯ: "t",
  ŧ: "t",
  ƭ: "t",
  ʈ: "t",
  ⱦ: "t",
  ꞇ: "t",

  ꜩ: "tz",

  u: "u",
  "\u24E4": "u",
  ｕ: "u",
  ù: "u",
  ú: "u",
  û: "u",
  ũ: "u",
  ṹ: "u",
  ū: "u",
  ṻ: "u",
  ŭ: "u",
  ü: "u",
  ǜ: "u",
  ǘ: "u",
  ǖ: "u",
  ǚ: "u",
  ủ: "u",
  ů: "u",
  ű: "u",
  ǔ: "u",
  ȕ: "u",
  ȗ: "u",
  ư: "u",
  ừ: "u",
  ứ: "u",
  ữ: "u",
  ử: "u",
  ự: "u",
  ụ: "u",
  ṳ: "u",
  ų: "u",
  ṷ: "u",
  ṵ: "u",
  ʉ: "u",

  v: "v",
  "\u24E5": "v",
  ｖ: "v",
  ṽ: "v",
  ṿ: "v",
  ʋ: "v",
  ꝟ: "v",
  ʌ: "v",
  ꝡ: "vy",
  w: "w",
  "\u24E6": "w",
  ｗ: "w",
  ẁ: "w",
  ẃ: "w",
  ŵ: "w",
  ẇ: "w",
  ẅ: "w",
  ẘ: "w",
  ẉ: "w",
  ⱳ: "w",
  x: "x",
  "\u24E7": "x",
  ｘ: "x",
  ẋ: "x",
  ẍ: "x",

  y: "y",
  "\u24E8": "y",
  ｙ: "y",
  ỳ: "y",
  ý: "y",
  ŷ: "y",
  ỹ: "y",
  ȳ: "y",
  ẏ: "y",
  ÿ: "y",
  ỷ: "y",
  ẙ: "y",
  ỵ: "y",
  ƴ: "y",
  ɏ: "y",
  ỿ: "y",

  z: "z",
  "\u24E9": "z",
  ｚ: "z",
  ź: "z",
  ẑ: "z",
  ż: "z",
  ž: "z",
  ẓ: "z",
  ẕ: "z",
  ƶ: "z",
  ȥ: "z",
  ɀ: "z",
  ⱬ: "z",
  ꝣ: "z",
};

export var utils = {
  /**
   * Throws an error if the given condition is not true
   *
   * @param {any} bool
   */
  assert: function (bool) {
    if (!bool) {
      throw new Error("AssertionError");
    }
  },
  /**
   * Check if the value is a bin_size or not.
   * If not, compute an approximate size out of the base64 encoded string.
   *
   * @param  {string} value original format
   * @return {string} bin_size (human-readable)
   */
  binaryToBinsize: function (value) {
    if (!this.is_bin_size(value)) {
      // Computing approximate size out of base64 encoded string
      // http://en.wikipedia.org/wiki/Base64#MIME
      return this.human_size(value.length / 1.37);
    }
    // already bin_size
    return value;
  },
  /**
   * Confines a value inside an interval
   *
   * @param {number} [val] the value to confine
   * @param {number} [min] the minimum of the interval
   * @param {number} [max] the maximum of the interval
   * @return {number} val if val is in [min, max], min if val < min and max
   *   otherwise
   */
  confine: function (val, min, max) {
    return Math.max(min, Math.min(max, val));
  },
  /**
   * @param {number} value
   * @param {integer} decimals
   * @returns {boolean}
   */
  float_is_zero: function (value, decimals) {
    var epsilon = Math.pow(10, -decimals);
    return Math.abs(utils.round_precision(value, epsilon)) < epsilon;
  },
  /**
   * Generate a unique numerical ID
   *
   * @returns {integer}
   */
  generateID: function () {
    return ++id;
  },
  /**
   * Read the cookie described by c_name
   *
   * @param {string} c_name
   * @returns {string}
   */
  get_cookie: function (c_name) {
    var cookies = document.cookie ? document.cookie.split("; ") : [];
    for (var i = 0, l = cookies.length; i < l; i++) {
      var parts = cookies[i].split("=");
      var name = parts.shift();
      var cookie = parts.join("=");

      if (c_name && c_name === name) {
        return cookie;
      }
    }
    return "";
  },
  /**
   * Gets dataURL (base64 data) from the given file or blob.
   * Technically wraps FileReader.readAsDataURL in Promise.
   *
   * @param {Blob|File} file
   * @returns {Promise} resolved with the dataURL, or rejected if the file is
   *  empty or if an error occurs.
   */
  getDataURLFromFile: function (file) {
    if (!file) {
      return Promise.reject();
    }
    return new Promise(function (resolve, reject) {
      var reader = new FileReader();
      reader.addEventListener("load", function () {
        resolve(reader.result);
      });
      reader.addEventListener("abort", reject);
      reader.addEventListener("error", reject);
      reader.readAsDataURL(file);
    });
  },
  /**
   * Returns a human readable number (e.g. 34000 -> 34k).
   *
   * @param {number} number
   * @param {integer} [decimals=0]
   *        maximum number of decimals to use in human readable representation
   * @param {integer} [minDigits=1]
   *        the minimum number of digits to preserve when switching to another
   *        level of thousands (e.g. with a value of '2', 4321 will still be
   *        represented as 4321 otherwise it will be down to one digit (4k))
   * @param {function} [formatterCallback]
   *        a callback to transform the final number before adding the
   *        thousands symbol (default to adding thousands separators (useful
   *        if minDigits > 1))
   * @returns {string}
   */
  human_number: function (number, decimals, minDigits, formatterCallback) {
    number = Math.round(number);
    decimals = decimals | 0;
    minDigits = minDigits || 1;
    formatterCallback = formatterCallback || utils.insert_thousand_seps;

    var d2 = Math.pow(10, decimals);
    var val = _t("kMGTPE");
    var symbol = "";
    var numberMagnitude = number.toExponential().split("e")[1];
    // the case numberMagnitude >= 21 corresponds to a number
    // better expressed in the scientific format.
    if (numberMagnitude >= 21) {
      // we do not use number.toExponential(decimals) because we want to
      // avoid the possible useless O decimals: 1e.+24 preferred to 1.0e+24
      number = Math.round(number * Math.pow(10, decimals - numberMagnitude)) / d2;
      // formatterCallback seems useless here.
      return number + "e" + numberMagnitude;
    }
    var sign = Math.sign(number);
    number = Math.abs(number);
    for (var i = val.length; i > 0; i--) {
      var s = Math.pow(10, i * 3);
      if (s <= number / Math.pow(10, minDigits - 1)) {
        number = Math.round((number * d2) / s) / d2;
        symbol = val[i - 1];
        break;
      }
    }
    number = sign * number;
    return formatterCallback("" + number) + symbol;
  },
  /**
   * Returns a human readable size
   *
   * @param {Number} size number of bytes
   */
  human_size: function (size) {
    var units = _t("Bytes|Kb|Mb|Gb|Tb|Pb|Eb|Zb|Yb").split("|");
    var i = 0;
    while (size >= 1024) {
      size /= 1024;
      ++i;
    }
    return size.toFixed(2) + " " + units[i].trim();
  },
  /**
   * Insert "thousands" separators in the provided number (which is actually
   * a string)
   *
   * @param {String} num
   * @returns {String}
   */
  insert_thousand_seps: function (num) {
    var negative = num[0] === "-";
    num = negative ? num.slice(1) : num;
    return (negative ? "-" : "") + utils.intersperse(num, _t.database.parameters.grouping, _t.database.parameters.thousands_sep);
  },
  /**
   * Intersperses ``separator`` in ``str`` at the positions indicated by
   * ``indices``.
   *
   * ``indices`` is an array of relative offsets (from the previous insertion
   * position, starting from the end of the string) at which to insert
   * ``separator``.
   *
   * There are two special values:
   *
   * ``-1``
   *   indicates the insertion should end now
   * ``0``
   *   indicates that the previous section pattern should be repeated (until all
   *   of ``str`` is consumed)
   *
   * @param {String} str
   * @param {Array<Number>} indices
   * @param {String} separator
   * @returns {String}
   */
  intersperse: function (str, indices, separator) {
    separator = separator || "";
    var result = [],
      last = str.length;

    for (var i = 0; i < indices.length; ++i) {
      var section = indices[i];
      if (section === -1 || last <= 0) {
        // Done with string, or -1 (stops formatting string)
        break;
      } else if (section === 0 && i === 0) {
        // repeats previous section, which there is none => stop
        break;
      } else if (section === 0) {
        // repeat previous section forever
        //noinspection AssignmentToForLoopParameterJS
        section = indices[--i];
      }
      result.push(str.substring(last - section, last));
      last -= section;
    }

    var s = str.substring(0, last);
    if (s) {
      result.push(s);
    }
    return result.reverse().join(separator);
  },
  /**
   * @param {any} object
   * @param {any} path
   * @returns
   */
  into: function (object, path) {
    if (!_(path).isArray()) {
      path = path.split(".");
    }
    for (var i = 0; i < path.length; i++) {
      object = object[path[i]];
    }
    return object;
  },
  /**
   * @param {string} v
   * @returns {boolean}
   */
  is_bin_size: function (v) {
    return /^\d+(\.\d*)? [^0-9]+$/.test(v);
  },
  /**
   * Returns whether the given anchor is valid.
   *
   * This test is useful to prevent a crash that would happen if using an invalid
   * anchor as a selector.
   *
   * @param {string} anchor
   * @returns {boolean}
   */
  isValidAnchor: function (anchor) {
    return /^#[\w-]+$/.test(anchor);
  },
  /**
   * @param {any} node
   * @param {any} human_readable
   * @param {any} indent
   * @returns {string}
   */
  json_node_to_xml: function (node, human_readable, indent) {
    // For debugging purpose, this function will convert a json node back to xml
    indent = indent || 0;
    var sindent = human_readable ? new Array(indent + 1).join("\t") : "",
      r = sindent + "<" + node.tag,
      cr = human_readable ? "\n" : "";

    if (typeof node === "string") {
      return sindent + node.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;");
    } else if (typeof node.tag !== "string" || !node.children instanceof Array || !node.attrs instanceof Object) {
      throw new Error(_.str.sprintf(_t("Node [%s] is not a JSONified XML node"), JSON.stringify(node)));
    }
    for (var attr in node.attrs) {
      var vattr = node.attrs[attr];
      if (typeof vattr !== "string") {
        // domains, ...
        vattr = JSON.stringify(vattr);
      }
      vattr = vattr.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;");
      if (human_readable) {
        vattr = vattr.replace(/&quot;/g, "'");
      }
      r += " " + attr + '="' + vattr + '"';
    }
    if (node.children && node.children.length) {
      r += ">" + cr;
      var childs = [];
      for (var i = 0, ii = node.children.length; i < ii; i++) {
        childs.push(utils.json_node_to_xml(node.children[i], human_readable, indent + 1));
      }
      r += childs.join(cr);
      r += cr + sindent + "</" + node.tag + ">";
      return r;
    } else {
      return r + "/>";
    }
  },
  /**
   * Left-pad provided arg 1 with zeroes until reaching size provided by second
   * argument.
   *
   * @see rpad
   *
   * @param {number|string} str value to pad
   * @param {number} size size to reach on the final padded value
   * @returns {string} padded string
   */
  lpad: function (str, size) {
    str = "" + str;
    return new Array(size - str.length + 1).join("0") + str;
  },
  /**
   * performs a half up rounding with a fixed amount of decimals, correcting for float loss of precision
   * See the corresponding float_round() in server/tools/float_utils.py for more info
   * @param {Number} value the value to be rounded
   * @param {Number} decimals the number of decimals. eg: round_decimals(3.141592,2) -> 3.14
   */
  round_decimals: function (value, decimals) {
    /**
     * The following decimals introduce numerical errors:
     * Math.pow(10, -4) = 0.00009999999999999999
     * Math.pow(10, -5) = 0.000009999999999999999
     *
     * Such errors will propagate in round_precision and lead to inconsistencies between Python
     * and JavaScript. To avoid this, we parse the scientific notation.
     */
    return utils.round_precision(value, parseFloat("1e" + -decimals));
  },
  /**
   * performs a half up rounding with arbitrary precision, correcting for float loss of precision
   * See the corresponding float_round() in server/tools/float_utils.py for more info
   *
   * @param {number} value the value to be rounded
   * @param {number} precision a precision parameter. eg: 0.01 rounds to two digits.
   */
  round_precision: function (value, precision) {
    if (!value) {
      return 0;
    } else if (!precision || precision < 0) {
      precision = 1;
    }
    var normalized_value = value / precision;
    var epsilon_magnitude = Math.log(Math.abs(normalized_value)) / Math.log(2);
    var epsilon = Math.pow(2, epsilon_magnitude - 52);
    normalized_value += normalized_value >= 0 ? epsilon : -epsilon;

    /**
     * Javascript performs strictly the round half up method, which is asymmetric. However, in
     * Python, the method is symmetric. For example:
     * - In JS, Math.round(-0.5) is equal to -0.
     * - In Python, round(-0.5) is equal to -1.
     * We want to keep the Python behavior for consistency.
     */
    var sign = normalized_value < 0 ? -1.0 : 1.0;
    var rounded_value = sign * Math.round(Math.abs(normalized_value));
    return rounded_value * precision;
  },
  /**
   * @see lpad
   *
   * @param {string} str
   * @param {number} size
   * @returns {string}
   */
  rpad: function (str, size) {
    str = "" + str;
    return str + new Array(size - str.length + 1).join("0");
  },
  /**
   * Create a cookie
   * @param {String} name the name of the cookie
   * @param {String} value the value stored in the cookie
   * @param {Integer} ttl time to live of the cookie in millis. -1 to erase the cookie.
   */
  set_cookie: function (name, value, ttl) {
    ttl = ttl || 24 * 60 * 60 * 365;
    document.cookie = [name + "=" + value, "path=/", "max-age=" + ttl, "expires=" + new Date(new Date().getTime() + ttl * 1000).toGMTString()].join(";");
  },
  /**
   * Sort an array in place, keeping the initial order for identical values.
   *
   * @param {Array} array
   * @param {function} iteratee
   */
  stableSort: function (array, iteratee) {
    var stable = array.slice();
    return array.sort(function stableCompare(a, b) {
      var order = iteratee(a, b);
      if (order !== 0) {
        return order;
      } else {
        return stable.indexOf(a) - stable.indexOf(b);
      }
    });
  },
  /**
   * @param {any} array
   * @param {any} elem1
   * @param {any} elem2
   */
  swap: function (array, elem1, elem2) {
    var i1 = array.indexOf(elem1);
    var i2 = array.indexOf(elem2);
    array[i2] = elem1;
    array[i1] = elem2;
  },

  /**
   * @param {string} value
   * @param {boolean} allow_mailto
   * @returns boolean
   */
  is_email: function (value, allow_mailto) {
    // http://stackoverflow.com/questions/46155/validate-email-address-in-javascript
    var re;
    if (allow_mailto) {
      re = /^(mailto:)?(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
    } else {
      re = /^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
    }
    return re.test(value);
  },

  /**
   * @param {any} str
   * @param {any} elseValues
   * @param {any} trueValues
   * @param {any} falseValues
   * @returns
   */
  toBoolElse: function (str, elseValues, trueValues, falseValues) {
    var ret = _.str.toBool(str, trueValues, falseValues);
    if (_.isUndefined(ret)) {
      return elseValues;
    }
    return ret;
  },
  /**
   * @todo: is this really the correct place?
   *
   * @param {any} data
   * @param {any} f
   */
  traverse_records: function (data, f) {
    if (data.type === "record") {
      f(data);
    } else if (data.data) {
      for (var i = 0; i < data.data.length; i++) {
        utils.traverse_records(data.data[i], f);
      }
    }
  },
  /**
   * Replace diacritics character with ASCII character
   *
   * @param {string} str diacritics string
   * @param {boolean} casesensetive
   * @returns {string} ASCII string
   */
  unaccent: function (str, casesensetive) {
    str = str.replace(/[^\u0000-\u007E]/g, function (accented) {
      return diacriticsMap[accented] || accented;
    });
    return casesensetive ? str : str.toLowerCase();
  },
  /**
   * @param {any} node
   * @param {any} strip_whitespace
   * @returns
   */
  xml_to_json: function (node, strip_whitespace) {
    switch (node.nodeType) {
      case 9:
        return utils.xml_to_json(node.documentElement, strip_whitespace);
      case 3:
      case 4:
        return strip_whitespace && node.data.trim() === "" ? undefined : node.data;
      case 1:
        var attrs = node.attributes; // $(node).getAttributes();
        return {
          tag: node.tagName.toLowerCase(),
          attrs: attrs,
          children: _.compact(
            _.map(node.childNodes, function (node) {
              return utils.xml_to_json(node, strip_whitespace);
            })
          ),
        };
    }
  },
  /**
   * @param {any} node
   * @returns {string}
   */
  xml_to_str: function (node) {
    var str = "";
    if (window.XMLSerializer) {
      str = new XMLSerializer().serializeToString(node);
    } else if (window.ActiveXObject) {
      str = node.xml;
    } else {
      throw new Error(_t("Could not serialize XML"));
    }
    // Browsers won't deal with self closing tags except void elements:
    // http://www.w3.org/TR/html-markup/syntax.html
    var void_elements = "area base br col command embed hr img input keygen link meta param source track wbr".split(" ");

    // The following regex is a bit naive but it's ok for the xmlserializer output
    str = str.replace(/<([a-z]+)([^<>]*)\s*\/\s*>/g, function (match, tag, attrs) {
      if (void_elements.indexOf(tag) < 0) {
        return "<" + tag + attrs + "></" + tag + ">";
      } else {
        return match;
      }
    });
    return str;
  },
  /**
   * Visit a tree of objects, where each children are in an attribute 'children'.
   * For each children, we call the callback function given in arguments.
   *
   * @param {Object} tree an object describing a tree structure
   * @param {function} f a callback
   */
  traverse: function (tree, f) {
    if (f(tree)) {
      _.each(tree.children, function (c) {
        utils.traverse(c, f);
      });
    }
  },
  /**
   * Enhanced traverse function with 'path' building on traverse.
   *
   * @param {Object} tree an object describing a tree structure
   * @param {function} f a callback
   * @param {Object} path the path to the current 'tree' object
   */
  traversePath: function (tree, f, path) {
    path = path || [];
    f(tree, path);
    _.each(tree.children, function (node) {
      utils.traversePath(node, f, path.concat(tree));
    });
  },
  /**
   * Visit a tree of objects and freeze all
   *
   * @param {Object} obj
   */
  deepFreeze: function (obj) {
    var propNames = Object.getOwnPropertyNames(obj);
    propNames.forEach(function (name) {
      var prop = obj[name];
      if (typeof prop == "object" && prop !== null) utils.deepFreeze(prop);
    });
    return Object.freeze(obj);
  },

  /**
   * Find the closest value of the given one in the provided array
   *
   * @param {Number} num
   * @param {Array} arr
   * @returns {Number|undefined}
   */
  closestNumber: function (num, arr) {
    var curr = arr[0];
    var diff = Math.abs(num - curr);
    for (var val = 0; val < arr.length; val++) {
      var newdiff = Math.abs(num - arr[val]);
      if (newdiff < diff) {
        diff = newdiff;
        curr = arr[val];
      }
    }
    return curr;
  },
  /**
   * Returns the domain targeting assets files.
   *
   * @returns {Array} Domain of assets files
   */
  assetsDomain: function () {
    return [
      "&",
      ["res_model", "=", "ir.ui.view"],
      "|",
      "|",
      "|",
      ["name", "=like", "%.assets_%.css"],
      ["name", "=like", "%.assets_%.js"],
      ["name", "=", "web_editor.summernote.css"],
      ["name", "=", "web_editor.summernote.js"],
    ];
  },
};

var vectors = (window.vectors = window.vectors || {});
vectors.utils = vectors.utils || {};
// All **ECMAScript 5** native function implementations that we hope to use
// are declared here.
var nativeIsArray = Array.isArray,
  nativeKeys = Object.keys,
  nativeCreate = Object.create;

// An internal function for creating assigner functions.
var createAssigner = function (keysFunc, undefinedOnly) {
  return function (obj) {
    var length = arguments.length;
    if (length < 2 || obj == null) return obj;
    for (var index = 1; index < length; index++) {
      var source = arguments[index],
        keys = keysFunc(source),
        l = keys.length;
      for (var i = 0; i < l; i++) {
        var key = keys[i];
        if (!undefinedOnly || obj[key] === void 0) obj[key] = source[key];
      }
    }
    return obj;
  };
};

// Internal function that returns an efficient (for current engines) version
// of the passed-in callback, to be repeatedly applied in other Underscore
// functions.
var optimizeCb = function (func, context, argCount) {
  if (context === void 0) return func;
  switch (argCount == null ? 3 : argCount) {
    case 1:
      return function (value) {
        return func.call(context, value);
      };
    case 2:
      return function (value, other) {
        return func.call(context, value, other);
      };
    case 3:
      return function (value, index, collection) {
        return func.call(context, value, index, collection);
      };
    case 4:
      return function (accumulator, value, index, collection) {
        return func.call(context, accumulator, value, index, collection);
      };
  }
  return function () {
    return func.apply(context, arguments);
  };
};

// A mostly-internal function to generate callbacks that can be applied
// to each element in a collection, returning the desired result — either
// identity, an arbitrary callback, a property matcher, or a property accessor.
var cb = function (value, context, argCount) {
  if (value == null) return vectors.utils.identity;
  if (vectors.utils.isFunction(value)) return optimizeCb(value, context, argCount);
  if (vectors.utils.isObject(value)) return vectors.utils.matcher(value);
  return vectors.utils.property(value);
};
vectors.utils.iteratee = function (value, context) {
  return cb(value, context, Infinity);
};

// Keys in IE < 9 that won't be iterated by `for key in ...` and thus missed.
var hasEnumBug = !{ toString: null }.propertyIsEnumerable("toString");
var nonEnumerableProps = ["valueOf", "isPrototypeOf", "toString", "propertyIsEnumerable", "hasOwnProperty", "toLocaleString"];

function collectNonEnumProps(obj, keys) {
  var nonEnumIdx = nonEnumerableProps.length;
  var constructor = obj.constructor;
  var proto = (vectors.utils.isFunction(constructor) && constructor.prototype) || ObjProto;

  // Constructor is a special case.
  var prop = "constructor";
  if (vectors.utils.has(obj, prop) && !vectors.utils.contains(keys, prop)) keys.push(prop);

  while (nonEnumIdx--) {
    prop = nonEnumerableProps[nonEnumIdx];
    if (prop in obj && obj[prop] !== proto[prop] && !vectors.utils.contains(keys, prop)) {
      keys.push(prop);
    }
  }
}
// Keep the identity function around for default iteratees.
vectors.utils.identity = function (value) {
  return value;
};

vectors.utils.property = function (key) {
  return function (obj) {
    return obj == null ? void 0 : obj[key];
  };
};

// Use a comparator function to figure out the smallest index at which
// an object should be inserted so as to maintain order. Uses binary search.
vectors.utils.sortedIndex = function (array, obj, iteratee, context) {
  iteratee = cb(iteratee, context, 1);
  var value = iteratee(obj);
  var low = 0,
    high = array.length;
  while (low < high) {
    var mid = Math.floor((low + high) / 2);
    if (iteratee(array[mid]) < value) low = mid + 1;
    else high = mid;
  }
  return low;
};

// Generator function to create the findIndex and findLastIndex functions
function createIndexFinder(dir) {
  return function (array, predicate, context) {
    predicate = cb(predicate, context);
    var length = array != null && array.length;
    var index = dir > 0 ? 0 : length - 1;
    for (; index >= 0 && index < length; index += dir) {
      if (predicate(array[index], index, array)) return index;
    }
    return -1;
  };
}

// Returns the first index on an array-like that passes a predicate test
vectors.utils.findIndex = createIndexFinder(1);

vectors.utils.findLastIndex = createIndexFinder(-1);

// Return the first value which passes a truth test. Aliased as `detect`.
vectors.utils.find = vectors.utils.detect = function (obj, predicate, context) {
  var key;
  if (isArrayLike(obj)) {
    key = vectors.utils.findIndex(obj, predicate, context);
  } else {
    key = vectors.utils.findKey(obj, predicate, context);
  }
  if (key !== void 0 && key !== -1) return obj[key];
};

// The cornerstone, an `each` implementation, aka `forEach`.
// Handles raw objects in addition to array-likes. Treats all
// sparse array-likes as if they were dense.
vectors.utils.each = vectors.utils.forEach = function (obj, iteratee, context) {
  iteratee = optimizeCb(iteratee, context);
  var i, length;
  if (isArrayLike(obj)) {
    for (i = 0, length = obj.length; i < length; i++) {
      iteratee(obj[i], i, obj);
    }
  } else {
    var keys = vectors.utils.keys(obj);
    for (i = 0, length = keys.length; i < length; i++) {
      iteratee(obj[keys[i]], keys[i], obj);
    }
  }
  return obj;
};

// Return all the elements that pass a truth test.
// Aliased as `select`.
vectors.utils.filter = vectors.utils.select = function (obj, predicate, context) {
  var results = [];
  predicate = cb(predicate, context);
  vectors.utils.each(obj, function (value, index, list) {
    if (predicate(value, index, list)) results.push(value);
  });
  return results;
};

// Return the position of the first occurrence of an item in an array,
// or -1 if the item is not included in the array.
// If the array is large and already in sort order, pass `true`
// for **isSorted** to use binary search.
vectors.utils.indexOf = function (array, item, isSorted) {
  var i = 0,
    length = array && array.length;
  if (typeof isSorted == "number") {
    i = isSorted < 0 ? Math.max(0, length + isSorted) : isSorted;
  } else if (isSorted && length) {
    i = vectors.utils.sortedIndex(array, item);
    return array[i] === item ? i : -1;
  }
  if (item !== item) {
    return vectors.utils.findIndex(slice.call(array, i), vectors.utils.isNaN);
  }
  for (; i < length; i++) if (array[i] === item) return i;
  return -1;
};

// Produce a duplicate-free version of the array. If the array has already
// been sorted, you have the option of using a faster algorithm.
// Aliased as `unique`.
vectors.utils.uniq = vectors.utils.unique = function (array, isSorted, iteratee, context) {
  if (array == null) return [];
  if (!vectors.utils.isBoolean(isSorted)) {
    context = iteratee;
    iteratee = isSorted;
    isSorted = false;
  }
  if (iteratee != null) iteratee = cb(iteratee, context);
  var result = [];
  var seen = [];
  for (var i = 0, length = array.length; i < length; i++) {
    var value = array[i],
      computed = iteratee ? iteratee(value, i, array) : value;
    if (isSorted) {
      if (!i || seen !== computed) result.push(value);
      seen = computed;
    } else if (iteratee) {
      if (!vectors.utils.contains(seen, computed)) {
        seen.push(computed);
        result.push(value);
      }
    } else if (!vectors.utils.contains(result, value)) {
      result.push(value);
    }
  }
  return result;
};
// Take the difference between one array and a number of other arrays.
// Only the elements present in just the first array will remain.
vectors.utils.difference = function (array) {
  var rest = flatten(arguments, true, true, 1);
  return vectors.utils.filter(array, function (value) {
    return !vectors.utils.contains(rest, value);
  });
};
// Convenience version of a common use case of `filter`: selecting only objects
// containing specific `key:value` pairs.
vectors.utils.where = function (obj, attrs) {
  return vectors.utils.filter(obj, vectors.utils.matcher(attrs));
};

// Return the results of applying the iteratee to each element.
vectors.utils.map = vectors.utils.collect = function (obj, iteratee, context) {
  iteratee = cb(iteratee, context);
  var keys = !isArrayLike(obj) && vectors.utils.keys(obj),
    length = (keys || obj).length,
    results = Array(length);
  for (var index = 0; index < length; index++) {
    var currentKey = keys ? keys[index] : index;
    results[index] = iteratee(obj[currentKey], currentKey, obj);
  }
  return results;
};
// Convenience version of a common use case of `map`: fetching a property.
vectors.utils.pluck = function (obj, key) {
  return vectors.utils.map(obj, vectors.utils.property(key));
};
// Determine if the array or object contains a given value (using `===`).
// Aliased as `includes` and `include`.
vectors.utils.contains = vectors.utils.includes = vectors.utils.include = function (obj, target, fromIndex) {
  if (!isArrayLike(obj)) obj = vectors.utils.values(obj);
  return vectors.utils.indexOf(obj, target, typeof fromIndex == "number" && fromIndex) >= 0;
};

// Returns whether an object has a given set of `key:value` pairs.
vectors.utils.isMatch = function (object, attrs) {
  var keys = vectors.utils.keys(attrs),
    length = keys.length;
  if (object == null) return !length;
  var obj = Object(object);
  for (var i = 0; i < length; i++) {
    var key = keys[i];
    if (attrs[key] !== obj[key] || !(key in obj)) return false;
  }
  return true;
};

// Is a given array, string, or object empty?
// An "empty" object has no enumerable own-properties.
vectors.utils.isEmpty = function (obj) {
  if (obj == null) return true;
  if (isArrayLike(obj) && (vectors.utils.isArray(obj) || vectors.utils.isString(obj) || vectors.utils.isArguments(obj))) return obj.length === 0;
  return vectors.utils.keys(obj).length === 0;
};

// Is a given value a DOM element?
vectors.utils.isElement = function (obj) {
  return !!(obj && obj.nodeType === 1);
};

// Is a given value an array?
// Delegates to ECMA5's native Array.isArray
vectors.utils.isArray =
  nativeIsArray ||
  function (obj) {
    return toString.call(obj) === "[object Array]";
  };

// Is a given variable an object?
vectors.utils.isObject = function (obj) {
  var type = typeof obj;
  return type === "function" || (type === "object" && !!obj);
};

// Is the given value `NaN`? (NaN is the only number which does not equal itself).
vectors.utils.isNaN = function (obj) {
  return vectors.utils.isNumber(obj) && obj !== +obj;
};

// Add some isType methods: isArguments, isFunction, isString, isNumber, isDate, isRegExp, isError.
["Arguments", "Function", "String", "Number", "Date", "RegExp", "Error"].forEach(function (name) {
  vectors.utils["is" + name] = function (obj) {
    return toString.call(obj) === "[object " + name + "]";
  };
});

// Shortcut function for checking if an object has a given property directly
// on itself (in other words, not on a prototype).
vectors.utils.has = function (obj, key) {
  return obj != null && hasOwnProperty.call(obj, key);
};

// Generate a unique integer id (unique within the entire client session).
// Useful for temporary DOM ids.
var idCounter = 0;
vectors.utils.uniqueId = function (prefix) {
  var id = ++idCounter + "";
  return prefix ? prefix + id : id;
};

// Retrieve the names of an object's own properties.
// Delegates to **ECMAScript 5**'s native `Object.keys`
vectors.utils.keys = function (obj) {
  if (!vectors.utils.isObject(obj)) return [];
  if (nativeKeys) return nativeKeys(obj);
  var keys = [];
  for (var key in obj) if (vectors.utils.has(obj, key)) keys.push(key);
  // Ahem, IE < 9.
  if (hasEnumBug) collectNonEnumProps(obj, keys);
  return keys;
};

// Retrieve all the property names of an object.
vectors.utils.allKeys = function (obj) {
  if (!vectors.utils.isObject(obj)) return [];
  var keys = [];
  for (var key in obj) keys.push(key);
  // Ahem, IE < 9.
  if (hasEnumBug) collectNonEnumProps(obj, keys);
  return keys;
};

// Retrieve the values of an object's properties.
vectors.utils.values = function (obj) {
  var keys = vectors.utils.keys(obj);
  var length = keys.length;
  var values = Array(length);
  for (var i = 0; i < length; i++) {
    values[i] = obj[keys[i]];
  }
  return values;
};

// Helper for collection methods to determine whether a collection
// should be iterated as an array or as an object
// Related: http://people.mozilla.org/~jorendorff/es6-draft.html#sec-tolength
var MAX_ARRAY_INDEX = Math.pow(2, 53) - 1;
var isArrayLike = function (collection) {
  var length = collection && collection.length;
  return typeof length == "number" && length >= 0 && length <= MAX_ARRAY_INDEX;
};

// Determine whether all of the elements match a truth test.
// Aliased as `all`.
vectors.utils.every = vectors.utils.all = function (obj, predicate, context) {
  predicate = cb(predicate, context);
  var keys = !isArrayLike(obj) && vectors.utils.keys(obj),
    length = (keys || obj).length;
  for (var index = 0; index < length; index++) {
    var currentKey = keys ? keys[index] : index;
    if (!predicate(obj[currentKey], currentKey, obj)) return false;
  }
  return true;
};

// Return a copy of the object only containing the whitelisted properties.
vectors.utils.pick = function (object, oiteratee, context) {
  var result = {},
    obj = object,
    iteratee,
    keys;
  if (obj == null) return result;
  if (vectors.utils.isFunction(oiteratee)) {
    keys = vectors.utils.allKeys(obj);
    iteratee = optimizeCb(oiteratee, context);
  } else {
    keys = flatten(arguments, false, false, 1);
    iteratee = function (value, key, obj) {
      return key in obj;
    };
    obj = Object(obj);
  }
  for (var i = 0, length = keys.length; i < length; i++) {
    var key = keys[i];
    var value = obj[key];
    if (iteratee(value, key, obj)) result[key] = value;
  }
  return result;
};

/**
 * computes (Math.floor(a/b), a%b and passes that to the callback.
 *
 * returns the callback's result
 */
vectors.utils.divmod = function divmod(a, b, fn) {
  var mod = a % b;
  // in python, sign(a % b) === sign(b). Not in JS. If wrong side, add a
  // round of b
  if ((mod > 0 && b < 0) || (mod < 0 && b > 0)) {
    mod += b;
  }
  return fn(Math.floor(a / b), mod);
};

/**
 * Passes the fractional and integer parts of x to the callback, returns
 * the callback's result
 */
vectors.utils.modf = function modf(x, fn) {
  var mod = x % 1;
  if (mod < 0) {
    mod += 1;
  }
  return fn(mod, Math.floor(x));
}; /*
vectors.utils.extend = function (defaults, options) {
var extended = {};
var prop;
for (prop in defaults) {
  if (Object.prototype.hasOwnProperty.call(defaults, prop)) {
    extended[prop] = defaults[prop];
  }
}
for (prop in options) {
  if (Object.prototype.hasOwnProperty.call(options, prop)) {
    extended[prop] = options[prop];
  }
}
return extended;
};
*/ // Extend a given object with all the properties in passed-in object(s).

/**
 * Merge defaults with user options
 * @private
 * @param {Object} defaults Default settings
 * @param {Object} options User options
 * @returns {Object} Merged values of defaults and options
 */ vectors.utils.extend = createAssigner(vectors.utils.allKeys);

// Assigns a given object with all the own properties in the passed-in object(s)
// (https://developer.mozilla.org/docs/Web/JavaScript/Reference/Global_Objects/Object/assign)
vectors.utils.extendOwn = vectors.utils.assign = createAssigner(vectors.utils.keys);

// Fill in a given object with default properties.
vectors.utils.defaults = createAssigner(vectors.utils.allKeys, true);
