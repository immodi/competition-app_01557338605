import type { AuthInterface } from "@/interfaces/auth";
import type { User } from "@/interfaces/models/user";
import { getUserData } from "@/repo/user";
import { useState, useEffect } from "react";

const BLANK_USER: User = {
    userId: 0,
    username: "",
    createdAt: new Date(),
    tickets: 0,
    role: "user",
};

function useAuthed(): AuthInterface {
    const [isAuthed, setIsAuthed] = useState(!!localStorage.getItem("token"));
    const [userData, setUserData] = useState<User>(BLANK_USER);
    const [token, setToken] = useState(
        () => localStorage.getItem("token") ?? ""
    );

    useEffect(() => {
        if (token) {
            localStorage.setItem("token", token);
            refreshUserData();
            setIsAuthed(true);
        } else {
            localStorage.removeItem("token");
            setIsAuthed(false);
        }
    }, [token]);

    function logout() {
        setToken("");
        setUserData(BLANK_USER);
    }

    function refreshUserData() {
        getUserData(token).then(setUserData);
    }

    return {
        isAuthed,
        token,
        userData,
        setToken,
        logout,
        refreshUserData,
    };
}

export default useAuthed;
