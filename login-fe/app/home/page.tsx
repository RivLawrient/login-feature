"use client";

import axios from "axios";
import { useRouter } from "next/navigation";

export default function Homes() {
  const getCookie = async () => {
    try {
      const response = await axios.get("https://api.lawrients.my.id/api/user", {
        withCredentials: true,
      });
      console.log(response);
    } catch (error) {
      console.error("Error getting cookie:", error);
    }
  };
  const router = useRouter();

  const logout = async () => {
    try {
      const response = await axios.get("https://api.lawrients.my.id/bye", {
        withCredentials: true,
      });
      router.push("/");
      console.log(response);
    } catch (error) {
      console.error("Error getting cookie:", error);
    }
  };
  return (
    <div className="flex flex-col items-center justify-center w-full">
      <button onClick={getCookie}>klik aku</button>
      <button onClick={logout}>logout</button>
      <video controls>
        <source src="/v1.mov" type="video/mov" />
      </video>
    </div>
  );
}
