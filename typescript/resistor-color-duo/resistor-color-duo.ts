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

export function decodedValue(bandColors: BandColor[]): number {
  const strResult = bandColors
      .slice(0, 2)
      .reduce((acc, curr) => `${acc}${getValue(curr)}`, '')

  return Number.parseInt(strResult, 10)
}

function getValue(key: BandColor): typeof bandColor[BandColor] | never {
  const value = bandColor[key]
  if (value === undefined) {
    throw new Error(`Invalid band color ${key}`)
  }

  return value
}
