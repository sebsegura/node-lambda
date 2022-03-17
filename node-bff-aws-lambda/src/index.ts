import {
  Handler,
  APIGatewayProxyEventV2,
  APIGatewayProxyResultV2,
} from "aws-lambda"
import { Transaction, createMetric, createTrx } from "./services"

type ProxyHandler = Handler<APIGatewayProxyEventV2, any>

export const handler: ProxyHandler = async (event, context) => {
  const body: Transaction = JSON.parse(event.body)

  const promises: Promise<any>[] = [createTrx(body), createMetric(body)]

  let response: APIGatewayProxyResultV2
  try {
    const [resTrx, resMetrics] = await Promise.all(promises)
    response = {
      statusCode: 200,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        ...resMetrics,
        ...resTrx,
      }),
    }
  } catch (err) {
    response = {
      statusCode: 400,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        message: "error",
      }),
    }
  }

  return response
}
