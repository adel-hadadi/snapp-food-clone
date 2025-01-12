import { useEffect, useState } from "react";
import Modal from "./modal";
import AxiosInstance from "../utils/axiosInstance";
import useAuthentication from "../hooks/useAuth";
import { redirect } from "next/navigation";

const LoginModal = ({ onToggle }) => {
  const { login, error, isLoading, user } = useAuthentication();
  const [otpFormData, setOTPFormData] = useState({ phone: "" });
  const [loginFormData, setLoginFormData] = useState({
    code: 0,
  });
  const [modalTitle, setModalTitle] = useState("ورود یا عضویت");
  const [showOTPForm, setShowOTPForm] = useState(true);

  const handleOTPFormChange = (e) => {
    setOTPFormData({ ...otpFormData, [e.target.name]: e.target.value });
  };

  const handleLogginFormDataChange = (e) => {
    setLoginFormData({ ...loginFormData, [e.target.name]: e.target.value });
  };

  const handleSubmitForm = (event) => {
    event.preventDefault();

    AxiosInstance.post("/auth/otp", otpFormData).then(() => {
      setModalTitle("تایید شماره");
      setShowOTPForm(false);
    });
  };

  const handleSubmitLogginForm = async (event) => {
    event.preventDefault();

    await login({ code: Number(loginFormData.code), phone: otpFormData.phone });

    if (user) {
      redirect("/panel");
    }
  };

  return (
    <Modal handleToggleModal={onToggle} isOpen={true} title={modalTitle}>
      {showOTPForm ? (
        <form onSubmit={handleSubmitForm}>
          <label className="text-gray-500 text-sm">شماره تلفن</label>
          <input
            type="text"
            name="phone"
            className="w-full border rounded py-3 px-2 mb-2"
            onChange={handleOTPFormChange}
          />
          <button className="w-full py-3 bg-slate-200 rounded">ادامه</button>
        </form>
      ) : (
        <form onSubmit={handleSubmitLogginForm}>
          <label className="text-gray-500 text-sm">
            کد تایید به شماره {otpFormData.phone} فرستاده شد.
          </label>
          <section className="mb-2">
            <input
              type="number"
              name="code"
              className="w-full border rounded py-3 px-2 "
              onChange={handleLogginFormDataChange}
              value={loginFormData.code}
            />
            {error && (
              <p className="text-red-600 text-sm">
                {error.data?.message || error.statusText}
              </p>
            )}
          </section>
          <button className="w-full py-3 bg-slate-200 rounded">
            {isLoading ? "..." : "ورود"}
          </button>
        </form>
      )}
    </Modal>
  );
};
export default LoginModal;
