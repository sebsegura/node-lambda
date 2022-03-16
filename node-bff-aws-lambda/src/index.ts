import { Handler, APIGatewayProxyEventV2, APIGatewayProxyResultV2 } from "aws-lambda"

type ProxyHandler = Handler<APIGatewayProxyEventV2, any>

export const handler: ProxyHandler = async (event, context) => {
    const accountId = event.queryStringParameters?.accountId

    const response: APIGatewayProxyResultV2 = {
        statusCode: 200,
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            message: `Hello ${accountId}!`
        })
    }

    return response
}