export interface Rate {
  id: number
  name: string
  price: number
}

export interface NewRate {
  name: string
  price: number
}

const api = import.meta.env.API_PATH

// Отправляет запрос на оплату
export async function createNewRate(token: string, rate: NewRate): Promise<void> {
  const response = await fetch(api + '/rate/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authentication-Token': token,
    },
    body: JSON.stringify(rate),
  })
  if (!response.ok) {
    throw { error: new Error(`${response.statusText}`), code: response.status }
  }
}

// Получить все тарифы
export async function getAllRates(token: string): Promise<Rate[]> {
  const response = await fetch(api + '/rates', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authentication-Token': token,
    },
  })
  if (!response.ok) {
    throw { error: new Error(`${response.statusText}`), code: response.status }
  }
  return response.json()
}
