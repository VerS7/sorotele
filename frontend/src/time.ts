export function formatDateTime(datetime: Date): string {
  const date = new Date(datetime)
  return (
    date.toLocaleDateString('ru-RU', { dateStyle: 'medium' }) +
    ' в ' +
    date.toLocaleTimeString('ru-RU', { timeStyle: 'short' })
  )
}

export function getUnixTime(datetime: Date): number {
  const date = new Date(datetime)
  return date.getTime()
}
