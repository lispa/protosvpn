"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

export default function LoginPage() {
  const router = useRouter();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const [error, setError] = useState("");

  async function handleLogin(
    event: React.FormEvent,
  ) {
    event.preventDefault();

    setError("");

    try {
      const response = await fetch(
        "https://protosvpn.novikoff.org/api/v1/auth/login",
        {
          method: "POST",

          headers: {
            "Content-Type": "application/json",
          },

          body: JSON.stringify({
            username,
            password,
          }),
        },
      );

      if (!response.ok) {
        setError("Invalid credentials");

        return;
      }

      const data = await response.json();

      localStorage.setItem(
        "token",
        data.token,
      );

      router.push("/dashboard");
    } catch {
      setError("Login failed");
    }
  }

  return (
    <main
      className="
        min-h-screen
        flex
        items-center
        justify-center
        bg-gray-100
      "
    >
      <form
        onSubmit={handleLogin}
        className="
          bg-white
          p-8
          rounded-xl
          shadow-lg
          w-full
          max-w-sm
          space-y-4
        "
      >
        <h1 className="text-2xl font-bold">
          ProtosVPN Login
        </h1>

        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) =>
            setUsername(e.target.value)
          }
          className="
            w-full
            border
            rounded
            p-3
          "
        />

        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) =>
            setPassword(e.target.value)
          }
          className="
            w-full
            border
            rounded
            p-3
          "
        />

        {error && (
          <p className="text-red-500">
            {error}
          </p>
        )}

        <button
          type="submit"
          className="
            w-full
            bg-black
            text-white
            p-3
            rounded
          "
        >
          Login
        </button>
      </form>
    </main>
  );
}