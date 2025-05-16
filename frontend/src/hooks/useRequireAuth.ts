import { useAppContext } from "@/contexts/AppContext";
import { useEffect } from "react";
import { useNavigate } from "react-router";

export default function useRequireAuth(): boolean {
    const { authData } = useAppContext();
    const navigate = useNavigate();

    useEffect(() => {
        if (!authData.isAuthed) {
            navigate("/auth/login");
        }
    }, [authData.isAuthed, navigate]);

    return authData.isAuthed;
}
