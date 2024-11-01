import { redirect } from "next/navigation";

export async function RegGoogle() {
  const url = "http://127.0.0.1:8080/auth/register/google";
  const response = await fetch(url, {});

  const data = await response.json();

  redirect(data.url);
}
