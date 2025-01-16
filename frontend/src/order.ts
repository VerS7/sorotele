export interface OrderDetails {
  fullName: string
  contacts: string
  message: string
}

const api = import.meta.env.API_PATH

export async function sendOrder(details: OrderDetails): Promise<void> {
  const response = await fetch(api + '/request', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(details),
  })

  if (!response.ok) {
    throw { error: new Error(`${response.statusText}`), code: response.status }
  }
}
