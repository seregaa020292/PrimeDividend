export function array<T extends Array<T>>(arr: T): T {
  return arr[Math.floor(Math.random() * arr.length)]
}

export function integer(min: number, max: number): number {
  return Math.floor(min + Math.random() * (max + 1 - min))
}

export function boolean(): boolean {
  return Math.random() < 0.5
}
