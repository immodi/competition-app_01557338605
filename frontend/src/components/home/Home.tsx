import SearchBar from "@/components/home/Searchbar";
import Table from "@/components/home/Table";
import { useAppContext } from "@/contexts/AppContext";
import useRequireAuth from "@/hooks/useRequireAuth";
import type { Event } from "@/interfaces/models/event";
import { getEvents, getEventsByCategory, searchEvents } from "@/repo/events";
import { useEffect, useState } from "react";
import toast from "react-hot-toast";
import { useNavigate, useParams } from "react-router";

const ITEMS_PER_PAGE = 8;

const Home: React.FC = () => {
    const { category } = useParams();
    const navigate = useNavigate();
    const { changeHeader, isDarkMode, authData } = useAppContext();
    const isAuthed = useRequireAuth();
    const [events, setEvents] = useState<Event[]>([]);
    const [page, setPage] = useState(1);
    const [totalPages, setTotalPages] = useState(1);
    const [query, setQuery] = useState("");

    function refreshEvents() {
        getEvents(authData.token, page, ITEMS_PER_PAGE)
            .then((response) => {
                setEvents(response.events);
                setTotalPages(
                    Math.max(Math.ceil(response.count / ITEMS_PER_PAGE), 1)
                );
            })
            .catch((error) => {
                console.error(error);
                toast.error(error.message);
            });
    }

    useEffect(() => {
        changeHeader("Events Dashboard");
    }, []);

    useEffect(() => {
        if (!authData.token) return;

        if (!category) {
            refreshEvents();
        } else {
            getEventsByCategory(authData.token, category, page, ITEMS_PER_PAGE)
                .then((response) => {
                    setEvents(response.events);
                    setTotalPages(
                        Math.max(Math.ceil(response.count / ITEMS_PER_PAGE), 1)
                    );
                })
                .catch((error) => {
                    console.error(error);
                    toast.error(error.message);
                });
        }
    }, [authData.token, page, category]);

    if (!isAuthed) {
        return <div className="p-10 text-center text-lg">Unauthorized</div>;
    }

    const buttonBaseClasses = "px-4 py-2 rounded-lg font-semibold transition";

    const buttonClasses = isDarkMode
        ? buttonBaseClasses + " bg-gray-700 hover:bg-gray-600 text-white"
        : buttonBaseClasses + " bg-gray-200 hover:bg-gray-300 text-gray-800";

    return (
        <main className="container mx-auto p-10">
            <SearchBar
                query={query}
                setQuery={setQuery}
                onFocus={() => setPage(1)}
                onSearch={(query) => {
                    if (query === "") {
                        refreshEvents();
                    }
                    if (query && query.length > 0) {
                        searchEvents(
                            authData.token,
                            query,
                            page,
                            ITEMS_PER_PAGE
                        )
                            .then((response) => {
                                setEvents(response.events);
                                setTotalPages(
                                    Math.max(
                                        Math.ceil(
                                            response.count / ITEMS_PER_PAGE
                                        ),
                                        1
                                    )
                                );
                            })
                            .catch((error) => {
                                console.error(error);
                                toast.error(error.message);
                            });
                    }
                }}
            />
            <Table
                events={events}
                refreshEvents={refreshEvents}
                filterCategory={(category) => {
                    setPage(1);
                    setQuery(`?category=${category}`);
                    navigate(`/${category}`);
                }}
            />
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

export default Home;
