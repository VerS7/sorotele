import Cookies from 'js-cookie'

const api = import.meta.env.API_PATH

export type Role = 'Admin' | 'User'

export interface UserCredentials {
  login: string
  password: string
}

export interface UserData {
  account: string
  name: string
  surname: string
  rateName: string
  ratePrice: number
  role: Role
}

export interface NewUserData {
  name: string
  surname: string
  password: string
  role: Role
  rateID: number
}

export interface NewUserResponse {
  account: string
  name: string
  surname: string
  token: string
}

export interface History {
  amount: number
  datetime: Date
}

export interface UserDynamicData {
  balance: string
  history: History[]
}

// Загружает токен доступа из Cookies
export function loadAuthToken(): string {
  const token = Cookies.get('st-auth-token')
  if (
    token === undefined ||
    typeof token === undefined ||
    typeof token?.length === undefined ||
    token?.length === 0
  ) {
    throw new Error('Token not found in Cookies')
  }
  return token
}

// Записывает токен доступа в Cookies
export function setAuthToken(token: string) {
  Cookies.set('st-auth-token', token, { expires: 30 })
}

// Удаляет токен доступа из Cookies
export function removeAuthToken() {
  Cookies.remove('st-auth-token')
}

// Загружает данные пользователя из SessionStorage
export function loadSessionUserData(): UserData {
  const data = sessionStorage.getItem('UserData')
  if (data === null) {
    throw new Error(`Session user data not found`)
  }
  return JSON.parse(data) as UserData
}

// Записывает данные пользователя в SessionStorage
export function setSessionUserData(data: UserData) {
  sessionStorage.setItem('UserData', JSON.stringify(data))
}

// Удаляет данные пользоватяля в SessionStorage
export function removeSessionUserData() {
  sessionStorage.removeItem('UserData')
}

// Запрашивает токен доступа
export async function requestAccessToken(creds: UserCredentials): Promise<string> {
  const response = await fetch(api + '/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(creds),
  })

  if (!response.ok) {
    throw { error: new Error(`${response.statusText}`), code: response.status }
  }

  const data = await response.json()
  return data.token
}

// Запрашивает данные пользователя
export async function getUserData(token: string): Promise<UserData> {
  const response = await fetch(api + '/user', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authentication-Token': token,
    },
  })

  if (!response.ok) {
    throw { error: new Error(`${response.statusText}`), code: response.status }
  }

  return await response.json()
}

// Запрашивает динамическую информацию пользователя (баланс и т.п.)
export async function getUserDynamicData(token: string): Promise<UserDynamicData> {
  const response = await fetch(api + '/user/data', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authentication-Token': token,
    },
  })

  if (!response.ok) {
    throw { error: new Error(`${response.statusText}`), code: response.status }
  }

  return await response.json()
}

// Создать пользователя
export async function CreateNewUser(token: string, user: NewUserData): Promise<NewUserResponse> {
  const response = await fetch(api + '/user/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authentication-Token': token,
    },
    body: JSON.stringify(user),
  })
  if (!response.ok) {
    throw { error: new Error(`${response.statusText}`), code: response.status }
  }
  return response.json()
}
