const earthYears = {
  mercury: 0.2408467,
  venus: 0.61519726,
  earth:  1.0,
  mars:  1.8808158,
  jupiter:  11.862615,
  saturn:  29.447498,
  uranus:  84.016846,
  neptune:  164.79132,
} as const

type Planet = keyof typeof earthYears

function floatToFixed(float: number, toFix: number): number {
  return Number.parseFloat(float.toFixed(toFix))
}

function secondsToYears(seconds: number): number {
  return seconds / (60 * 60 * 24 * 365.245)
}

export function age(planet: Planet, seconds: number): number {
  const factor = earthYears[planet]

  if (!factor) {
    throw new Error('Planet not found')
  }

  const ageInYears = secondsToYears(seconds)
  const result = ageInYears / factor
  console.log('result', result)

  return floatToFixed(result, 2)
}

