import { useAppContext } from "@/contexts/AppContext";
import { registerUser } from "@/repo/auth";
import React, { useEffect, useState } from "react";
import toast from "react-hot-toast";
import { NavLink, useNavigate } from "react-router";
import AdminDisclaimer from "./AdminDisclamer";

const Register: React.FC = () => {
    const { isDarkMode, changeHeader, authData } = useAppContext();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [confirm, setConfirm] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        if (authData.isAuthed) {
            navigate("/");
        }
    }, [authData.isAuthed, navigate]);

    useEffect(() => {
        changeHeader("Register");
    }, [changeHeader]);

    if (authData.isAuthed) {
        return <div className="p-10 text-center text-lg">Unauthorized</div>;
    }

    function handleRegisterCallback(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();

        console.log(username);

        if (password === confirm) {
            registerUser(username, password)
                .then((res) => {
                    if (res.token) {
                        authData.setToken(res.token);
                        navigate("/");
                        toast.success("Registered successfully");
                    }
                })
                .catch((err) => {
                    toast.error(err.message);
                    console.warn(err);
                });
        } else {
            toast.error("Passwords do not match");
        }
    }

    return (
        <div
            className={`${
                isDarkMode ? "dark" : ""
            } min-h-screen flex items-center justify-center p-6 bg-gray-50 dark:bg-gray-900 transition-colors duration-300`}
        >
            <div className="w-full max-w-md space-y-4">
                <AdminDisclaimer />

                <form
                    action="#"
                    onSubmit={handleRegisterCallback}
                    className="w-full max-w-md bg-white dark:bg-gray-800 rounded-2xl shadow p-8 space-y-6"
                >
                    <h2 className="text-2xl font-bold text-gray-800 dark:text-gray-100 text-center">
                        Register
                    </h2>
                    <div className="space-y-4">
                        <label className="block">
                            <span className="text-gray-700 dark:text-gray-200">
                                Username
                            </span>
                            <input
                                type="text"
                                className="mt-1 block w-full rounded-lg bg-gray-100 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
                                value={username}
                                onChange={(e) => setUsername(e.target.value)}
                                required
                            />
                        </label>
                        <label className="block">
                            <span className="text-gray-700 dark:text-gray-200">
                                Password
                            </span>
                            <input
                                type="password"
                                className="mt-1 block w-full rounded-lg bg-gray-100 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                                required
                            />
                        </label>
                        <label className="block">
                            <span className="text-gray-700 dark:text-gray-200">
                                Confirm Password
                            </span>
                            <input
                                type="password"
                                className="mt-1 block w-full rounded-lg bg-gray-100 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
                                value={confirm}
                                onChange={(e) => setConfirm(e.target.value)}
                                required
                            />
                        </label>
                    </div>
                    <button
                        type="submit"
                        className="w-full py-2 cursor-pointer bg-blue-500 dark:bg-blue-600 text-white font-semibold rounded-lg hover:bg-blue-600 dark:hover:bg-blue-700 transition"
                    >
                        Create Account
                    </button>
                    <p className="text-center text-sm text-gray-600 dark:text-gray-400">
                        Already have an account?{" "}
                        <NavLink
                            to="/auth/login"
                            className="text-blue-500 hover:underline"
                            end
                        >
                            Log in
                        </NavLink>
                    </p>
                </form>
            </div>
        </div>
    );
};

export default Register;
