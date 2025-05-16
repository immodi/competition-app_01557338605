import type { AuthInterface } from "@/interfaces/auth";
import type { User } from "@/interfaces/models/user";
import { getUserData } from "@/repo/user";
import { useState, useEffect } from "react";

function useAuthed(): AuthInterface {
    const [isAuthed, setIsAuthed] = useState(!!localStorage.getItem("token"));
    const [userData, setUserData] = useState<User>({
        userId: 0,
        createdAt: new Date(),
        role: "user",
        tickets: 0,
        username: "",
    });
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
