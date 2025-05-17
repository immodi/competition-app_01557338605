import { useAppContext } from "@/contexts/AppContext";
import type { TableProps } from "@/interfaces/tableData";
import { assignUserToEvent, deleteEvent } from "@/repo/events"; // ⬅️ added deleteEvent here
import { getUserEventIds } from "@/repo/user";
import { useEffect, useState } from "react";
import type React from "react";
import toast from "react-hot-toast";
import { FaPencil, FaTicket, FaTrash } from "react-icons/fa6"; // ⬅️ added FaTrash icon
import { IoEyeSharp } from "react-icons/io5";
import { useNavigate } from "react-router";

const Table: React.FC<TableProps> = ({
    events,
    filterCategory,
    refreshEvents,
}) => {
    const { authData } = useAppContext();
    const [userEventIds, setUserEventIds] = useState<number[]>([]);
    const navigate = useNavigate();

    useEffect(() => {
        if (authData.token && authData.userData.userId !== 0) {
            getUserEventIds(authData.token, authData.userData.userId)
                .then(setUserEventIds)
                .catch((error) => {
                    console.error(error);
                    toast.error(error.message);
                });
        }
    }, [authData.token, authData.userData.userId, authData.userData.tickets]);

    function deleteEventCallback(eventId: number) {
        if (window.confirm("Are you sure you want to delete this event?")) {
            // console.log(eventId);

            deleteEvent(authData.token, eventId)
                .then(() => {
                    toast.success("Event deleted successfully");
                    refreshEvents();
                })
                .catch((error) => {
                    console.error(error);
                    toast.error(error.message);
                });
        }
    }

    return (
        <div className="overflow-x-auto">
            <table className="min-w-full border border-gray-200 bg-white dark:bg-gray-800 dark:border-gray-700 rounded-lg">
                <thead className="bg-gray-100 dark:bg-gray-700">
                    <tr>
                        {[
                            "Event Id",
                            "Event Name",
                            "Description",
                            "Category",
                            "Date",
                            "Venue",
                            "Price",
                        ].map((heading) => (
                            <th
                                key={heading}
                                className="px-4 py-3 text-left text-gray-700 dark:text-gray-200 font-semibold border-b dark:border-gray-600"
                            >
                                {heading}
                            </th>
                        ))}
                        <th className="px-4 py-3 text-center text-gray-700 dark:text-gray-200 font-semibold border-b dark:border-gray-600">
                            Actions
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {events.map((event) => (
                        <tr
                            key={event.id}
                            className="hover:bg-gray-50 border-b dark:border-gray-600 dark:hover:bg-gray-700"
                        >
                            <td className="px-4 py-3 dark:text-gray-100">
                                {event.id}
                            </td>
                            <td className="px-4 py-3 dark:text-gray-100">
                                {event.name}
                            </td>
                            <td className="px-4 py-3 dark:text-gray-100">
                                {event.description}
                            </td>
                            <td className="px-4 py-3">
                                <button
                                    onClick={() =>
                                        filterCategory(event.category)
                                    }
                                    className="px-3 py-1 cursor-pointer rounded-full text-sm font-medium bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200 hover:opacity-90 transition"
                                >
                                    {event.category}
                                </button>
                            </td>
                            <td className="px-4 py-3 dark:text-gray-100">
                                {new Date(event.date).toLocaleString()}
                            </td>
                            <td className="px-4 py-3 dark:text-gray-100">
                                {event.venue}
                            </td>
                            <td className="px-4 py-3 dark:text-gray-100">
                                ${event.price}
                            </td>
                            <td className="px-4 py-3 items-center">
                                <div className="flex justify-center gap-2">
                                    {!userEventIds.includes(event.id) ? (
                                        <button
                                            onClick={() => {
                                                assignUserToEvent(
                                                    authData.token,
                                                    event.id,
                                                    authData.userData.userId
                                                )
                                                    .then(() => {
                                                        authData.refreshUserData();
                                                        toast.success(
                                                            "Event booked successfully"
                                                        );
                                                        navigate(
                                                            "/booking-success",
                                                            {
                                                                state: {
                                                                    name: event.name,
                                                                },
                                                            }
                                                        );
                                                    })
                                                    .catch((error) => {
                                                        console.error(error);
                                                        toast.error(
                                                            error.message
                                                        );
                                                    });
                                            }}
                                            title="Book"
                                            className="px-3 py-1 bg-blue-500 cursor-pointer text-white rounded hover:bg-blue-600 transition"
                                        >
                                            <FaTicket />
                                        </button>
                                    ) : (
                                        <button
                                            disabled
                                            title="Booked"
                                            className="px-3 py-1 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 dark:bg-gray-600 dark:text-gray-100 dark:hover:bg-gray-500 transition"
                                        >
                                            <FaTicket />
                                        </button>
                                    )}
                                    <button
                                        onClick={() =>
                                            navigate(`/event/${event.id}`)
                                        }
                                        title="View Details"
                                        className="px-3 py-1 bg-blue-600 text-white cursor-pointer rounded hover:bg-blue-700 dark:bg-blue-500 dark:hover:bg-blue-600 transition"
                                    >
                                        <IoEyeSharp />
                                    </button>
                                    {authData.userData.role === "admin" && (
                                        <>
                                            <button
                                                onClick={() =>
                                                    navigate(
                                                        `/event/edit/${event.id}`
                                                    )
                                                }
                                                title="Edit Event"
                                                className="px-3 py-1 bg-green-600 text-white cursor-pointer rounded hover:bg-green-700 dark:bg-green-500 dark:hover:bg-green-600 transition"
                                            >
                                                <FaPencil />
                                            </button>
                                            <button
                                                onClick={() =>
                                                    deleteEventCallback(
                                                        event.id
                                                    )
                                                }
                                                title="Delete Event"
                                                className="px-3 py-1 bg-red-600 text-white cursor-pointer rounded hover:bg-red-700 dark:bg-red-500 dark:hover:bg-red-600 transition"
                                            >
                                                <FaTrash />
                                            </button>
                                        </>
                                    )}
                                </div>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default Table;
