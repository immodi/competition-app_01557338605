import Table from "@/components/home/Table";
import { useAppContext } from "@/contexts/AppContext";
import useRequireAuth from "@/hooks/useRequireAuth";
import type { Event } from "@/interfaces/models/event";
import { getEvents } from "@/repo/events";
import { useEffect, useState } from "react";
import toast from "react-hot-toast";

const ITEMS_PER_PAGE = 8;

const HomeRoute: React.FC = () => {
    const { changeHeader, isDarkMode, authData } = useAppContext();
    const isAuthed = useRequireAuth();
    const [events, setEvents] = useState<Event[]>([]);
    const [page, setPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);

    useEffect(() => {
        changeHeader("Events Dashboard");
    }, []);

    useEffect(() => {
        if (!authData.token) return;

        getEvents(authData.token, page, ITEMS_PER_PAGE)
            .then((response) => {
                setEvents(response.events);
                setTotalPages(Math.ceil(response.count / ITEMS_PER_PAGE));
            })
            .catch((error) => {
                console.error(error);
                toast.error(error.message);
            });
    }, [authData.token, page]);

    if (!isAuthed) {
        return <div className="p-10 text-center text-lg">Unauthorized</div>;
    }

    const buttonBaseClasses = "px-4 py-2 rounded-lg font-semibold transition";

    const buttonClasses = isDarkMode
        ? buttonBaseClasses + " bg-gray-700 hover:bg-gray-600 text-white"
        : buttonBaseClasses + " bg-gray-200 hover:bg-gray-300 text-gray-800";

    return (
        <main className="container mx-auto p-10">
            <Table events={events} />
            <div className="flex justify-center gap-4 mt-6">
                <button
                    disabled={page === 1}
                    onClick={() => setPage((p) => Math.max(1, p - 1))}
                    className={`${buttonClasses} ${
                        page === 1 ? "opacity-50 cursor-not-allowed" : ""
                    }`}
                >
                    Prev
                </button>

                <span className="flex items-center select-none">
                    Page {page} / {totalPages}
                </span>

                <button
                    disabled={page === totalPages}
                    onClick={() => setPage((p) => Math.min(totalPages, p + 1))}
                    className={`${buttonClasses} ${
                        page === totalPages
                            ? "opacity-50 cursor-not-allowed"
                            : ""
                    }`}
                >
                    Next
                </button>
            </div>
        </main>
    );
};

export default HomeRoute;
