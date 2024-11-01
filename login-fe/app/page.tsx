"use client";

import axios from "axios";
import Link from "next/link";

export default function Home() {
  const getCookie = async () => {
    try {
      const response = await axios.get("http://127.0.0.1:8080/hai", {
        // withCredentials: true,
      });
      console.log(response);
    } catch (error) {
      console.error("Error getting cookie:", error);
    }
  };
  return (
    <div className="flex-col flex m-1 items-center justify-center  w-full">
      <Link
        href={"/login"}
        className="border border-black w-full max-w-[150px] mb-2 hover:bg-slate-100 flex justify-center items-center"
      >
        LOGIN
      </Link>
      <Link
        href={"/register"}
        className="border border-black w-full max-w-[150px] mb-2 hover:bg-slate-100 flex justify-center items-center"
      >
        REGISTER
      </Link>
      <button
        onClick={getCookie}
        className="border border-black w-full max-w-[150px] mb-2 hover:bg-slate-100 flex justify-center items-center"
      >
        HAI
      </button>
    </div>
  );
}
