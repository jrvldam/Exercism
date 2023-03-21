const dnaToRna: { [key: string]: string } = {
  G: 'C',
  C: 'G',
  T: 'A',
  A: 'U',
} as const

export function toRna(nucleotides: string): string {
  return nucleotides.split('').reduce((acc, curr) => { 
    if (!dnaToRna[curr]) {
      throw new Error('Invalid input DNA.')
    }

    return acc + dnaToRna[curr]
  }, '')
}
