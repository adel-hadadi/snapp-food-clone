"use client";

import Link from "next/link";
import { useEffect, useState } from "react";
import { IoPersonOutline } from "react-icons/io5";
import Addresses from "../components/addresses";
import Image from "next/image";
import Cart from "./cart";
import useAuthentication from "../hooks/useAuth";
import AxiosInstance from "../utils/axiosInstance";
import { ToastContainer } from "react-toastify";

export default function PanelLayout({ children }) {
  const { user, isLoading } = useAuthentication();
  const [addresses, setAddresses] = useState([]);

  useEffect(() => {
    AxiosInstance.get("http://localhost/api/profile/addresses").then((res) =>
      setAddresses(res.data.data),
    );
  }, [isLoading]);

  return (
    <>
      <header className="flex justify-between items-center h-28 border-b-2 shadow px-10 mb-8">
        <div>
          <Link href="/panel">
            <Image
              src="/images/Snappfood-Logo.png"
              width={100}
              height={100}
              alt="Snapp Food"
            />
          </Link>
        </div>

        <div>
          {addresses.length != 0 && (
            <Addresses
              addresses={addresses}
              defaultAddress={user?.default_address_id}
            />
          )}
        </div>

        <div className="flex">
          <Link href="/profile" className="flex items-center px-5 py-3 text-sm">
            <IoPersonOutline fontWeight="bold" />
            <span className="mx-2">
              {user?.first_name + " " + user?.last_name}
            </span>
          </Link>

          <Cart />
        </div>
      </header>

      <ToastContainer
        position="top-right"
        autoClose={5000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick={false}
        rtl={true}
        pauseOnFocusLoss
        draggable
        pauseOnHover
        theme="light"
        transition={Bounce}
      />

      <div className="container mx-auto px-6">{children}</div>
    </>
  );
}
