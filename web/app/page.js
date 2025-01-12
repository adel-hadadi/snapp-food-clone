"use client";

import useAuthentication from "./hooks/useAuth";
import { redirect } from "next/navigation";
import Image from "next/image";
import { FaStore } from "react-icons/fa";
import { useState } from "react";
import LoginModal from "./components/login_modal";

export default function Home() {
  const [showLogginModal, setShowLogginModal] = useState(false);
  const handleShowLoginModal = () => {
    setShowLogginModal((p) => !p);
  };

  const { user } = useAuthentication();
  if (user) {
    redirect("/panel");
  }

  return (
    <div>
      <section
        className="bg-slate-100 w-100 rounded-2xl relative overflow-hidden px-3"
        style={{ height: "44rem" }}
      >
        <section className="flex justify-between items-center">
          <Image
            src="/images/Snappfood-Logo.png"
            height={150}
            width={150}
            alt="Snapp Food"
          />
          <section className="z-10">
            <button className="inline-flex items-center mx-2 py-3">
              <FaStore className="ml-1" />
              ثبت نام فروشندگان
            </button>
            <button
              className="py-3 px-1 bg-pink-500 text-white rounded mx-2 font-bold"
              onClick={handleShowLoginModal}
            >
              ورود یا عضویت
            </button>
            {showLogginModal && <LoginModal onToggle={handleShowLoginModal} />}
          </section>
        </section>

        <section
          className="absolute left-0 bottom-0 w-1/2 z-0"
          style={{ height: "85%" }}
        >
          <section className="h-full w-full relative">
            <Image src="/images/hero-image.png" alt="hero" fill={true} />
          </section>
        </section>
      </section>
    </div>
  );
}
