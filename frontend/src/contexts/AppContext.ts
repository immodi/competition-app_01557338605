import { createContext, useContext } from "react";

interface AppContextInterface {
    header: string;
    isDarkMode: boolean;
    isAuthed: boolean;
    toggleDarkMode: () => void;
    changeHeader: (header: string) => void;
    setIsAuthed: (isAuthed: boolean) => void;
}

export const AppContext = createContext<AppContextInterface>({
    header: "",
    isDarkMode: false,
    isAuthed: false,
    toggleDarkMode: () => {
        console.warn("AppContext not initialized");
    },
    changeHeader: () => {
        console.warn("AppContext not initialized");
    },
    setIsAuthed: () => {
        console.warn("AppContext not initialized");
    },
});

export const useAppContext = () => useContext(AppContext);

export default AppContext;
