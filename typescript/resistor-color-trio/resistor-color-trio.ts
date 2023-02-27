const bandColor = {
  black: 0,
  brown: 1,
  red: 2,
  orange: 3,
  yellow: 4,
  green: 5,
  blue: 6,
  violet: 7,
  grey: 8,
  white: 9,
} as const

type BandColor = keyof typeof bandColor

export function decodedResistorValue([colorBand1, colorBand2, colorBand3]: [BandColor, BandColor, BandColor]): string {
  const trailingZeros = bandColor[colorBand3]
  const numeric =`${bandColor[colorBand1]}${bandColor[colorBand2]}`.padEnd(trailingZeros + 2, '0')

  return numeric.length > 3 ? numeric.replace(/0{3}$/g, '') + ' kiloohms' : numeric  + ' ohms'
}
