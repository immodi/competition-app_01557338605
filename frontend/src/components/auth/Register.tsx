import { useAppContext } from "@/contexts/AppContext";
import React, { useEffect, useState } from "react";
import { NavLink } from "react-router";

const Register: React.FC = () => {
    const { isDarkMode, changeHeader } = useAppContext();
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [confirm, setConfirm] = useState("");

    useEffect(() => {
        changeHeader("Register");
    }, []);

    return (
        <div
            className={`${
                isDarkMode ? "dark" : ""
            } min-h-screen flex items-center justify-center p-6 bg-gray-50 dark:bg-gray-900 transition-colors duration-300`}
        >
            <form className="w-full max-w-md bg-white dark:bg-gray-800 rounded-2xl shadow p-8 space-y-6">
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
                            Email
                        </span>
                        <input
                            type="email"
                            className="mt-1 block w-full rounded-lg bg-gray-100 dark:bg-gray-700 border border-gray-300 dark:border-gray-600 px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-400"
                            value={email}
                            onChange={(e) => setEmail(e.target.value)}
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
                    className="w-full py-2 bg-blue-500 dark:bg-blue-600 text-white font-semibold rounded-lg hover:bg-blue-600 dark:hover:bg-blue-700 transition"
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
    );
};

export default Register;
