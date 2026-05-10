async function getVPNStatus() {
  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}/api/v1/vpn/status`,
    {
      cache: "no-store",
    }
  )

  if (!response.ok) {
    throw new Error("Failed to fetch VPN status")
  }

  return response.json()
}

export default async function DashboardPage() {
  const data = await getVPNStatus()

  return (
    <main className="min-h-screen bg-black text-white p-10">
      <div className="max-w-5xl mx-auto">
        <h1 className="text-4xl font-bold mb-8">
          ProtosVPN Dashboard
        </h1>

        <div className="bg-zinc-900 rounded-xl border border-zinc-800 overflow-hidden">
          <table className="w-full">
            <thead className="bg-zinc-800">
              <tr>
                <th className="text-left p-4">User</th>
                <th className="text-left p-4">Real IP</th>
                <th className="text-left p-4">Bytes Received</th>
                <th className="text-left p-4">Bytes Sent</th>
              </tr>
            </thead>

            <tbody>
              {data.clients?.map((client: any) => (
                <tr
                  key={client.name}
                  className="border-t border-zinc-800"
                >
                  <td className="p-4">{client.name}</td>

                  <td className="p-4">
                    {client.real_address}
                  </td>

                  <td className="p-4">
                    {client.bytes_received}
                  </td>

                  <td className="p-4">
                    {client.bytes_sent}
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </main>
  )
}