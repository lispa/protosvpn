export async function GET() {
  const response = await fetch(
    "http://api:8080/api/v1/vpn/status",
    {
      cache: "no-store",
    }
  )

  const data = await response.json()

  return Response.json(data)
}