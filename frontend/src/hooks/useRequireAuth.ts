import { useAppContext } from "@/contexts/AppContext";
import { useEffect } from "react";
import { useNavigate } from "react-router";

export default function useRequireAuth(): boolean {
    const { isAuthed } = useAppContext();
    const navigate = useNavigate();

    useEffect(() => {
        if (!isAuthed) {
            navigate("/auth/login");
        }
    }, [isAuthed, navigate]);

    return isAuthed;
}
