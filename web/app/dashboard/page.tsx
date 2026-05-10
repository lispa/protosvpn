"use client"

import { useEffect, useState } from "react"

type Client = {
  name: string
  real_address: string
  bytes_received: string
  bytes_sent: string
}

export default function DashboardPage() {
  const [clients, setClients] = useState<Client[]>([])
  const [loading, setLoading] = useState(true)

  async function fetchStatus() {
    try {
      const response = await fetch(
        "/api/v1/vpn/status"
      )

      const data = await response.json()

      setClients(data.clients || [])
    } catch (error) {
      console.error("Failed to fetch VPN status", error)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchStatus()

    const interval = setInterval(() => {
      fetchStatus()
    }, 5000)

    return () => clearInterval(interval)
  }, [])

  return (
    <main className="min-h-screen bg-black text-white p-10">
      <div className="max-w-5xl mx-auto">
        <div className="flex items-center justify-between mb-8">
          <h1 className="text-4xl font-bold">
            ProtosVPN Dashboard
          </h1>

          <div className="text-zinc-400">
            Online Users: {clients.length}
          </div>
        </div>

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
              {loading ? (
                <tr>
                  <td
                    colSpan={4}
                    className="p-6 text-center text-zinc-400"
                  >
                    Loading...
                  </td>
                </tr>
              ) : clients.length === 0 ? (
                <tr>
                  <td
                    colSpan={4}
                    className="p-6 text-center text-zinc-400"
                  >
                    No active VPN clients
                  </td>
                </tr>
              ) : (
                clients.map((client) => (
                  <tr
                    key={client.name}
                    className="border-t border-zinc-800"
                  >
                    <td className="p-4">
                      <div className="flex items-center gap-2">
                        <div className="w-2 h-2 rounded-full bg-green-500" />

                        {client.name}
                      </div>
                    </td>

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
                ))
              )}
            </tbody>
          </table>
        </div>
      </div>
    </main>
  )
}