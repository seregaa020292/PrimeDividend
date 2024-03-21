export function deepFreeze(object: any) {
  const propNames = Object.getOwnPropertyNames(object)
  for (const name of propNames) {
    const value = object[name]
    object[name] = value && typeof value === 'object' ? deepFreeze(value) : value
  }
  return Object.freeze(object)
}
