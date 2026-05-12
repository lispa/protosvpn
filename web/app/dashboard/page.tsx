"use client"

import { useEffect, useState } from "react"
import { useRouter } from "next/navigation"

type Client = {
  name: string
  real_address: string
  bytes_received: string
  bytes_sent: string
}

type VPNClient = {
  name: string
  status: string
}

export default function DashboardPage() {
  const router = useRouter()

  const [clients, setClients] = useState<Client[]>([])
  const [vpnClients, setVpnClients] = useState<VPNClient[]>([])
  const [loading, setLoading] = useState(true)
  const [clientName, setClientName] = useState("")
  const [creatingClient, setCreatingClient] = useState(false)

  async function fetchVPNClients() {
  try {
    const token = localStorage.getItem("token")

    const response = await fetch(
      "/api/v1/vpn/clients",
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    )

    const data = await response.json()

    setVpnClients(
      (data.clients || []).filter(
        (client: VPNClient) =>
          client.name !==
          "protosvpn.novikoff.org"
      )
    )
  } catch (error) {
    console.error(error)
  }
}

  async function fetchStatus() {
    try {
      const token = localStorage.getItem("token")

      const response = await fetch(
        "/api/v1/vpn/status",
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      )

      if (response.status === 401) {
        localStorage.removeItem("token")

        router.push("/login")

        return
      }

      const data = await response.json()

      setClients(data.clients || [])
    } catch (error) {
      console.error(
        "Failed to fetch VPN status",
        error
      )
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    const token = localStorage.getItem("token")

    if (!token) {
      router.push("/login")

      return
    }

    fetchStatus()
    fetchVPNClients()

    const interval = setInterval(() => {
      fetchStatus()
      fetchVPNClients()
    }, 5000)

    return () => clearInterval(interval)
  }, [router])

  function handleLogout() {
    localStorage.removeItem("token")

    router.push("/login")
  }

  async function handleDownloadClient(
  clientName: string
) {
  try {
    const token = localStorage.getItem("token")

    const response = await fetch(
      `/api/v1/vpn/download-client?name=${clientName}`,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    )

    const blob = await response.blob()

    const url =
      window.URL.createObjectURL(blob)

    const link =
      document.createElement("a")

    link.href = url

    link.download = `${clientName}.ovpn`

    document.body.appendChild(link)

    link.click()

    link.remove()
  } catch (error) {
    console.error(error)

    alert("Failed to download config")
  }
}

  async function handleRevokeClient(
  clientName: string
) {
  try {
    const token = localStorage.getItem("token")

    const response = await fetch(
      "/api/v1/vpn/revoke-client",
      {
        method: "POST",

        headers: {
          "Content-Type": "application/json",

          Authorization: `Bearer ${token}`,
        },

        body: JSON.stringify({
          name: clientName,
        }),
      }
    )

    if (!response.ok) {
      alert("Failed to revoke client")

      return
    }

    await fetchVPNClients()

    alert("Client revoked")
  } catch (error) {
    console.error(error)

    alert("Failed to revoke client")
  }
}

  async function handleCreateClient() {
  if (!clientName) {
    return
  }

  try {
    setCreatingClient(true)

    const token = localStorage.getItem("token")

    const createResponse = await fetch(
      "/api/v1/vpn/create-client",
      {
        method: "POST",

        headers: {
          "Content-Type": "application/json",

          Authorization: `Bearer ${token}`,
        },

        body: JSON.stringify({
          name: clientName,
        }),
      }
    )

    if (!createResponse.ok) {
      alert("Failed to create VPN client")

      return
    }

    const downloadResponse = await fetch(
      `/api/v1/vpn/download-client?name=${clientName}`,
      {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    )

    const blob = await downloadResponse.blob()

    const url = window.URL.createObjectURL(blob)

    const link = document.createElement("a")

    link.href = url

    link.download = `${clientName}.ovpn`

    document.body.appendChild(link)

    link.click()

    link.remove()

    setClientName("")

    alert("VPN client created")
  } catch (error) {
    console.error(error)

    alert("Failed to create VPN client")
  } finally {
    setCreatingClient(false)
  }
}

  return (
    <main className="min-h-screen bg-black text-white p-10">
      <div className="max-w-5xl mx-auto">
        <div className="flex items-center justify-between mb-8">
          <h1 className="text-4xl font-bold">
            ProtosVPN Dashboard
          </h1>

          <div className="flex items-center gap-6">
            <div className="text-zinc-400">
              Online Users: {clients.length}
            </div>

            <button onClick={handleLogout} className="bg-red-600 hover:bg-red-700 px-4 py-2 rounded">
              Logout
            </button>
          </div>
        </div>

        <div className="bg-zinc-900 rounded-xl border border-zinc-800 p-6 mb-6">
  <h2 className="text-2xl font-bold mb-4">
    Create VPN Client
  </h2>

  <div className="flex gap-4">
    <input type="text" placeholder="Client name" value={clientName} onChange={(e) =>setClientName(e.target.value)} className="flex-1 bg-zinc-800 border border-zinc-700 rounded p-3"/>

    <button
      onClick={handleCreateClient}
      disabled={creatingClient}
      className="bg-green-600 hover:bg-green-700 px-6 rounded">
      {creatingClient
        ? "Creating..."
        : "Create"}
    </button>
  </div>
</div>

        <div className="bg-zinc-900 rounded-xl border border-zinc-800 overflow-hidden">
          <table className="w-full">
            <thead className="bg-zinc-800">
              <tr>
                <th className="text-left p-4">
                  User
                </th>

                <th className="text-left p-4">
                  Real IP
                </th>

                <th className="text-left p-4">
                  Bytes Received
                </th>

                <th className="text-left p-4">
                  Bytes Sent
                </th>
              </tr>
            </thead>

            <tbody>
              {loading ? (
                <tr>
                  <td
                    colSpan={4}
                    className="
                      p-6
                      text-center
                      text-zinc-400
                    "
                  >
                    Loading...
                  </td>
                </tr>
              ) : clients.length === 0 ? (
                <tr>
                  <td
                    colSpan={4}
                    className="
                      p-6
                      text-center
                      text-zinc-400
                    "
                  >
                    No active VPN clients
                  </td>
                </tr>
              ) : (
                clients.map((client) => (
                  <tr
                    key={client.name}
                    className="
                      border-t
                      border-zinc-800
                    "
                  >
                    <td className="p-4">
                      <div className="flex items-center gap-2">
                        <div
                          className="
                            w-2
                            h-2
                            rounded-full
                            bg-green-500
                          "
                        />

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
          <div className="mt-10">
  <h2 className="text-3xl font-bold mb-6">
    VPN Clients
  </h2>

  <div className="bg-zinc-900 rounded-xl border border-zinc-800 overflow-hidden">
    <table className="w-full">
      <thead className="bg-zinc-800">
        <tr>
          <th className="text-left p-4">
            Client
          </th>

          <th className="text-left p-4">
            Status
          </th>

          <th className="text-left p-4">
            Actions
          </th>
        </tr>
      </thead>

      <tbody>
        {vpnClients.map((client) => (
          <tr
            key={client.name}
            className="
              border-t
              border-zinc-800
            "
          >
            <td className="p-4">
              {client.name}
            </td>

            <td className="p-4">
              <span
                className={
                  client.status ===
                  "active"
                    ? "text-green-500"
                    : "text-red-500"
                }
              >
                {client.status}
              </span>
            </td>

            <td className="p-4">
              <div className="flex gap-3">
                <button
                  onClick={() =>
                    handleDownloadClient(
                      client.name
                    )
                  }
                  className="
                    bg-blue-600
                    hover:bg-blue-700
                    px-4
                    py-2
                    rounded
                  "
                >
                  Download
                </button>

                {client.status ===
                  "active" && (
                  <button
                    onClick={() =>
                      handleRevokeClient(
                        client.name
                      )
                    }
                    className="
                      bg-red-600
                      hover:bg-red-700
                      px-4
                      py-2
                      rounded
                    "
                  >
                    Revoke
                  </button>
                )}
              </div>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  </div>
</div>
        </div>
      </div>
    </main>
  )
}