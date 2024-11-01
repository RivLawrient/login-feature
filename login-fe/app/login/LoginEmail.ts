import { redirect } from "next/navigation";

export async function LogEmail(email: string, password: string) {
  const url = "https://api.lawrients.my.id/auth/login";
  const response = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email: email,
      password: password,
    }),
    credentials: "include",
  });

  const data = await response.json();

  if (response.ok) {
    redirect("https://lawrients.my.id/home");
  } else {
    alert(data.errors);
  }
}
