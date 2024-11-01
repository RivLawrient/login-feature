import { redirect } from "next/navigation";

export async function LogGoogle() {
  const url = "https://api.lawrients.my.id/auth/login/google";
  const response = await fetch(url);

  const data = await response.json();

  redirect(data.url);
}
