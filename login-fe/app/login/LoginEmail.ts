import { redirect } from "next/navigation";

export async function LogEmail(email: string, password: string) {
  const url = "http://127.0.0.1:8080/auth/login";
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
    redirect("http://127.0.0.1:3000/home");
  } else {
    alert(data.errors);
  }
}
