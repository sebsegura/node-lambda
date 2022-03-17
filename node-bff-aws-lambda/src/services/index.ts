interface RequestOptions {
  method: "GET" | "POST" | "PUT" | "DELETE"
  headers?: Record<string, string>
  body?: string
}

const METRICS_API_URL = process.env.METRICS_API_URL
const TRX_API_URL = process.env.TRX_API_URL

const apiCall = async <T>(url: string, options: RequestOptions) => {
  const { default: fetch } = await import("node-fetch")
  const headers = {
    "Content-Type": "application/json",
    ...options.headers,
  }

  try {
    const res = await fetch(url, {
      headers,
      ...options,
    })

    let payload: any

    try {
      payload = await res.json()
    } catch {
      payload = null
    }

    if (!res.ok) {
      throw new Error("error")
    }

    return payload as T
  } catch (err) {
    throw err
  }
}

const apiPost = async <T>(url: string, body: string) => {
  return apiCall<T>(url, {
    method: "POST",
    body,
  })
}

export interface Metric {
  accountId: string
  date: string
}

export interface Transaction {
  accountId: string
  amount: number
}

export interface Status {
  status: string
}

export const createMetric = async (trx: Transaction) => {
  return apiPost<Metric>(METRICS_API_URL, JSON.stringify(trx))
}

export const createTrx = async (trx: Transaction) => {
  return apiPost<Status>(TRX_API_URL, JSON.stringify(trx))
}
