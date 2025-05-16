import AppContext from "@/contexts/AppContext";
import { useEffect, useState } from "react";
import { MdDarkMode, MdLightMode } from "react-icons/md";
import { Outlet } from "react-router";

const App: React.FC = () => {
    const [darkMode, setDarkMode] = useState(false);
    const [header, setHeader] = useState("");
    const [isAuthed, setIsAuthed] = useState(false);

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
                isAuthed: isAuthed,
                header: header,
                isDarkMode: darkMode,

                setIsAuthed: (isAuthed: boolean) => setIsAuthed(isAuthed),
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
                        {/* <div className="relative flex items-center">
                            <FaTicketAlt className="w-7 h-7 text-blue-600" />
                            <span className="absolute -top-2 -right-2 bg-red-500 text-white text-xs rounded-full px-1.5">
                                5
                            </span>
                        </div> */}
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
            </div>
        </AppContext.Provider>
    );
};

export default App;
