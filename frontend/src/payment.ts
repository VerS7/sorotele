export interface PaymentDetails {
  account: string
  email: string
  sum: string
}

export interface PaymentResponse {
  service: string
  link: string
}

const api = import.meta.env.API_PATH

// Отправляет запрос на оплату
export async function requestPayment(details: PaymentDetails): Promise<PaymentResponse> {
  const response = await fetch(api + '/pay', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(details),
  })
  if (!response.ok) {
    throw { error: new Error(`${response.statusText}`), code: response.status }
  }
  const paymentResponse = await response.json()

  return { service: paymentResponse.service, link: paymentResponse.link } as PaymentResponse
}
