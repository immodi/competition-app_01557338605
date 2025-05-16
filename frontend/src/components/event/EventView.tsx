import { useParams, useNavigate } from "react-router";
import { useAppContext } from "@/contexts/AppContext";
import { useEffect, useState } from "react";
import { IoArrowBack } from "react-icons/io5";
import toast from "react-hot-toast";
import { getEventById } from "@/repo/events";
import type { Event, EventTranslation } from "@/interfaces/models/event";
import useRequireAuth from "@/hooks/useRequireAuth";

const EventView: React.FC = () => {
    const { id } = useParams();
    const isAuthed = useRequireAuth();
    const navigate = useNavigate();
    const { authData } = useAppContext();
    const [event, setEvent] = useState<Event | null>(null);
    const [selectedTranslation, setSelectedTranslation] =
        useState<EventTranslation | null>(null);

    useEffect(() => {
        if (!authData.token) return;

        if (authData.token && id) {
            getEventById(authData.token, parseInt(id))
                .then(setEvent)
                .catch((error) => {
                    console.error(error);
                    toast.error(error.message);
                });
        }
    }, [id, authData.token]);

    if (!isAuthed) {
        return <div className="p-10 text-center text-lg">Unauthorized</div>;
    }

    if (!event) {
        return (
            <div className="p-10 text-center text-gray-800 dark:text-gray-100">
                Loading event details...
            </div>
        );
    }

    const displayedName = selectedTranslation?.name || event.name;
    const displayedDescription =
        selectedTranslation?.description || event.description;
    const displayedVenue = selectedTranslation?.venue || event.venue;

    return (
        <main className="container mx-auto p-8">
            <div className="w-full flex justify-between items-center">
                <h1 className="text-3xl font-bold mb-6 text-gray-800 dark:text-gray-100">
                    View Event Details
                </h1>

                <button
                    onClick={() => navigate(-1)}
                    className="flex cursor-pointer items-center gap-2 mb-6 px-4 py-2 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 dark:bg-gray-600 dark:text-gray-100 dark:hover:bg-gray-500 transition"
                >
                    <IoArrowBack />
                    Back
                </button>
            </div>

            <div className="bg-white dark:bg-gray-800 rounded-lg shadow p-6 space-y-4">
                <div className="flex flex-col md:flex-row md:items-center justify-between">
                    <h1 className="text-3xl font-bold text-gray-800 dark:text-gray-100">
                        {displayedName}
                    </h1>

                    {event.translations && event.translations.length > 0 && (
                        <select
                            onChange={(e) => {
                                const lang = e.target.value;
                                if (lang === "") {
                                    setSelectedTranslation(null);
                                } else {
                                    const translation = event.translations.find(
                                        (t) => t.language === lang
                                    );
                                    setSelectedTranslation(translation || null);
                                }
                            }}
                            className="mt-4 md:mt-0 px-3 py-2 border border-gray-300 dark:border-gray-600 rounded bg-white dark:bg-gray-700 text-gray-700 dark:text-gray-100"
                        >
                            <option value="">EN</option>
                            {event.translations.map((t) => (
                                <option key={t.language} value={t.language}>
                                    {t.language.toUpperCase()}
                                </option>
                            ))}
                        </select>
                    )}
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                    <div>
                        <table className="w-full text-left border-collapse">
                            <tbody>
                                <tr className="border-b border-gray-300 dark:border-gray-600">
                                    <th className="py-3 pr-4 text-gray-800 dark:text-gray-100 font-semibold w-32">
                                        Description:
                                    </th>
                                    <td className="py-3 text-gray-700 dark:text-gray-300">
                                        {displayedDescription}
                                    </td>
                                </tr>
                                <tr className="border-b border-gray-300 dark:border-gray-600">
                                    <th className="py-3 pr-4 text-gray-800 dark:text-gray-100 font-semibold">
                                        Category:
                                    </th>
                                    <td className="py-3 text-gray-700 dark:text-gray-300">
                                        {event.category}
                                    </td>
                                </tr>
                                <tr className="border-b border-gray-300 dark:border-gray-600">
                                    <th className="py-3 pr-4 text-gray-800 dark:text-gray-100 font-semibold">
                                        Date:
                                    </th>
                                    <td className="py-3 text-gray-700 dark:text-gray-300">
                                        {new Date(event.date).toLocaleString()}
                                    </td>
                                </tr>
                                <tr className="border-b border-gray-300 dark:border-gray-600">
                                    <th className="py-3 pr-4 text-gray-800 dark:text-gray-100 font-semibold">
                                        Venue:
                                    </th>
                                    <td className="py-3 text-gray-700 dark:text-gray-300">
                                        {displayedVenue}
                                    </td>
                                </tr>
                                <tr>
                                    <th className="py-3 pr-4 text-gray-800 dark:text-gray-100 font-semibold">
                                        Price:
                                    </th>
                                    <td className="py-3 text-gray-700 dark:text-gray-300">
                                        ${event.price}
                                    </td>
                                </tr>
                            </tbody>
                        </table>
                    </div>

                    <div className="flex items-center justify-center">
                        {event.image ? (
                            <img
                                src={`data:image/png;base64,${event.image}`}
                                alt={event.name}
                                className="w-full h-64 md:h-80 object-cover rounded-lg shadow-md"
                            />
                        ) : (
                            <div className="w-full h-64 md:h-80 bg-gray-200 dark:bg-gray-700 rounded-lg flex items-center justify-center text-gray-500 dark:text-gray-400">
                                No Image Available
                            </div>
                        )}
                    </div>
                </div>
            </div>
        </main>
    );
};

export default EventView;
