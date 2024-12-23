"use client";

import { useRouter } from "next/navigation";
import { useState } from "react";

const Page = () => {
  const [formData, setFormData] = useState({
    phone: "",
  });

  const [showOTPForm, setShowOTPForm] = useState(false);

  const router = useRouter();

  const handleChange = (event) => {
    setFormData({ ...formData, [event.target.name]: event.target.value });
  };

  const handleSendOTP = (event) => {
    event.preventDefault();

    fetch("http://localhost/api/auth/otp", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formData),
    }).then((response) => {
      if (response.ok) {
        setShowOTPForm(true);
      }
    });
  };

  const handleLoginRegister = (event) => {
    event.preventDefault();

    const data = {
      ...formData,
      code: parseInt(formData.code, 10),
    };
    fetch("http://localhost/api/auth/login-register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((res) => {
        localStorage.setItem("access_token", res?.data?.access_token);
        router.push("/panel");
      });
  };

  return (
    <div className="bg-amber-50 h-screen flex justify-center items-center">
      {!showOTPForm ? (
        <form
          className="border-2 shadow-md rounded px-8 pt-6 pb-8 mb-4"
          onSubmit={handleSendOTP}
        >
          <h1 className="font-bold mb-3 text-lg text-cyan-700">
            ورود یا عضویت
          </h1>
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              شماره تلفن همراه
            </label>
            <input
              className="text-center shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              type="text"
              placeholder="۰۹"
              name="phone"
              onChange={handleChange}
            />
          </div>
          <div className="flex items-center justify-between">
            <button
              className="bg-cyan-900 hover:bg-cyan-800 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
              type="submit"
            >
              ارسال کد
            </button>
          </div>
        </form>
      ) : (
        <form
          className="border-2 shadow-md rounded px-8 pt-6 pb-8 mb-4"
          onSubmit={handleLoginRegister}
        >
          <h1 className="font-bold mb-3 text-lg text-cyan-700">
            ورود یا عضویت
          </h1>
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">
              کد پیامک شده
            </label>
            <input
              className="text-center shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              type="number"
              name="code"
              onChange={handleChange}
            />
          </div>
          <div className="flex items-center justify-between">
            <button
              className="bg-cyan-900 hover:bg-cyan-800 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
              type="submit"
            >
              ورود / ثبت نام
            </button>
          </div>
        </form>
      )}
    </div>
  );
};

export default Page;
