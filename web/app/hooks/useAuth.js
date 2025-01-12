import AxiosInstance from "../utils/axiosInstance";

const { useState, useEffect } = require("react");

const useAuthentication = () => {
    const [user, setUser] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);

    const login = async (credentials) => {
        try {
            setIsLoading(true);
            const response = await AxiosInstance.post(
                "/auth/login-register",
                credentials,
            );

            const { access_token, refresh_token } = response.data.data;
            localStorage.setItem("access_token", access_token);
            localStorage.setItem("refresh_token", refresh_token);

            profile()

            setError(null)
            setIsLoading(false);
        } catch (err) {
            setError(err.response);
            setIsLoading(false);
        }
    };

    const logout = () => {
        localStorage.removeItem("access_token");
        setUser(null);
    };

    const profile = async () => {
            const token = localStorage.getItem("access_token");
            if (token) {
                try {
                    const response = await AxiosInstance.get("/profile/personal-info");
                    setUser(response.data.data);
                } catch (err) {
                    setError("invalid token");
                    const response = await AxiosInstance.post("/auth/refresh", {refresh_token: localStorage.getItem("refresh_token")})
                    localStorage.setItem("access_token", response.data.data.access_token)
                }
            }
            setIsLoading(false);
    }

    useEffect(() => {
        profile();
    }, []);

    return { user, isLoading, error, login, logout };
};

export default useAuthentication;
