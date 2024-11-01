import { redirect } from "next/navigation";

export async function RegGoogle() {
  const url = "https://api.lawrients.my.id/auth/register/google";
  const response = await fetch(url, {});

  const data = await response.json();

  redirect(data.url);
}
