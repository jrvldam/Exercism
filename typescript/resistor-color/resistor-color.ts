type Color = keyof typeof colors

const colors = {
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

export const colorCode = (color: Color) => {
  return colors[color]
}

export const COLORS = Object.keys(colors)



