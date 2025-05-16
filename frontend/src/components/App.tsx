import AppContext from "@/contexts/AppContext";
import useAuthed from "@/hooks/useAuthed";
import { useEffect, useState } from "react";
import { Toaster } from "react-hot-toast";
import { FaTicket } from "react-icons/fa6";
import { MdDarkMode, MdLightMode } from "react-icons/md";
import { Outlet, useNavigate } from "react-router";

const App: React.FC = () => {
    const [darkMode, setDarkMode] = useState(true);
    const [header, setHeader] = useState("");
    const authData = useAuthed();
    const navigate = useNavigate();

    useEffect(() => {
        if (darkMode) {
            document.documentElement.classList.add("dark");
        } else {
            document.documentElement.classList.remove("dark");
        }
    }, [darkMode]);

    return (
        <AppContext.Provider
            value={{
                header: header,
                isDarkMode: darkMode,
                authData: authData,

                toggleDarkMode: () => setDarkMode(!darkMode),
                changeHeader: (header: string) => setHeader(header),
            }}
        >
            <div className="p-6 bg-gray-50 dark:bg-gray-900 min-h-screen transition-colors duration-300 text-gray-800 dark:text-white">
                <header className="flex items-center justify-between mb-8">
                    <h1 className="text-2xl font-bold text-gray-800 dark:text-white">
                        {header}
                    </h1>
                    <div className="flex items-center gap-6">
                        {authData.isAuthed && (
                            <>
                                {authData.userData.role === "admin" && (
                                    <button
                                        onClick={() =>
                                            navigate("/event/create")
                                        }
                                        className={`ml-3 px-4 py-2 cursor-pointer rounded-lg transition text-sm relative bottom-0.5 ${
                                            darkMode
                                                ? "bg-blue-500 hover:bg-blue-600 text-white"
                                                : "bg-blue-600 hover:bg-blue-700 text-white"
                                        }`}
                                    >
                                        Create Event
                                    </button>
                                )}
                                <button
                                    onClick={authData.logout}
                                    className={`px-4 py-2 cursor-pointer rounded-lg transition text-sm relative bottom-0.5 ${
                                        darkMode
                                            ? "bg-red-500 hover:bg-red-600 text-white"
                                            : "bg-red-600 hover:bg-red-700 text-white"
                                    }`}
                                >
                                    Logout
                                </button>
                                <div className="relative flex items-center">
                                    <FaTicket className="w-7 h-7 text-gray-800 dark:text-white" />
                                    <span className="absolute -top-2 -right-2 bg-red-500 text-white text-xs rounded-full px-1.5">
                                        {authData.userData.tickets}
                                    </span>
                                </div>
                            </>
                        )}

                        {darkMode ? (
                            <MdLightMode
                                onClick={() => setDarkMode(false)}
                                className="w-8 h-8 cursor-pointer dark:text-white hover:text-gray-800 transition"
                                title="Switch to Light Mode"
                            />
                        ) : (
                            <MdDarkMode
                                onClick={() => setDarkMode(true)}
                                className="w-8 h-8 cursor-pointer text-gray-600 dark:text-white hover:text-gray-800 transition"
                                title="Enable Dark Mode"
                            />
                        )}
                    </div>
                </header>
                <Outlet />
                <Toaster />
            </div>
        </AppContext.Provider>
    );
};

export default App;
