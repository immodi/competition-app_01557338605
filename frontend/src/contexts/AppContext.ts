import type { AuthInterface } from "@/interfaces/auth";
import { createContext, useContext } from "react";

interface AppContextInterface {
    header: string;
    isDarkMode: boolean;
    authData: AuthInterface;

    toggleDarkMode: () => void;
    changeHeader: (header: string) => void;
}

export const AppContext = createContext<AppContextInterface>({
    header: "",
    isDarkMode: false,
    authData: {
        isAuthed: false,
        userData: {
            userId: 0,
            createdAt: new Date(),
            role: "user",
            tickets: 0,
            username: "",
        },

        setToken: () => {
            console.warn("AppContext not initialized");
        },
        token: "",
        logout: () => {
            console.warn("AppContext not initialized");
        },
        refreshUserData: () => {
            console.warn("AppContext not initialized");
        },
    },

    toggleDarkMode: () => {
        console.warn("AppContext not initialized");
    },
    changeHeader: () => {
        console.warn("AppContext not initialized");
    },
});

export const useAppContext = () => useContext(AppContext);

export default AppContext;
